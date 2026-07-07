---
title: GPi Case 2 Hardware Findings and KMS Power Notes
status: field-notes
created: 2026-07-07
project: RetroFlag Power
device: RetroFlag GPi Case 2 + Raspberry Pi Compute Module 4
tone: warm retro/adventurer field notes
---

# GPi Case 2 Hardware Findings and KMS Power Notes

These notes preserve the field findings from the GPi Case 2 KMS migration and power-behavior investigation.

The short version: the LCD has now been proven to work under modern KMS DPI with Broadcom V3D hardware acceleration. The remaining dragons are power-policy cleanup, power-quality headroom, and a few case-specific buttons/behaviors that are not ordinary EmulationStation inputs.

## Confirmed Victory: KMS DPI + V3D Works

The GPi Case 2 LCD was successfully brought up on the modern KMS path.

Observed proof:

```text
/proc/fb:
0 vc4drmfb

kmsprint:
DPI-1 (connected)
640x480@59.99
FB 640x480

EmulationStation log:
GL vendor:   Broadcom
GL renderer: V3D 4.2
GL version:  2.1 Mesa 20.3.5
```

This replaced the earlier failure mode where EmulationStation used Mesa llvmpipe software rendering, pegged CPU, generated heat, and did not have a valid KMS display path.

## Known-Good KMS Display Configuration

Working `/boot/config.txt` ingredients:

```ini
dtparam=audio=off
arm_64bit=1
disable_fw_kms_setup=1

# Legacy RetroFlag DPI/FKMS path disabled
#display_rotate=0
#dtoverlay=dpi24
#overscan_left=0
#overscan_right=0
#overscan_top=0
#overscan_bottom=0
#framebuffer_width=640
#framebuffer_height=480
#framebuffer_depth=8
#enable_dpi_lcd=1
#display_default_lcd=1
#dpi_group=2
#dpi_mode=87
#dpi_output_format=0x00016
#hdmi_timings=640 0 41 40 41 480 0 18 9 18 0 0 0 60 0 24000000 1

[pi4]
#dtoverlay=vc4-fkms-v3d
#max_framebuffers=2

[all]
#dtoverlay=vc4-fkms-v3d
#overscan_scale=1
dtoverlay=dwc2,dr_mode=host
dtoverlay=disable-bt

# Disabled because audremap claims GPIO12/13 and conflicts with KMS DPI.
#dtoverlay=audremap

# GPi Case 2 LCD - modern KMS DPI path
dtoverlay=vc4-kms-v3d,noaudio
max_framebuffers=0

dtoverlay=vc4-kms-dpi-generic,clock-frequency=24000000,rgb666-padhi
dtparam=hactive=640,hfp=41,hsync=40,hbp=41
dtparam=vactive=480,vfp=18,vsync=9,vbp=18
dtparam=rotate=0
```

The two critical changes were:

1. Disable `audremap`.
2. Use `rgb666-padhi` instead of `rgb888`.

## Failed State: Why KMS Initially Did Not Bind DPI

Earlier KMS attempts failed with this kernel clue:

```text
pin gpio12 already requested by fe00b840.mailbox; cannot claim for fe208000.dpi
vc4_dpi fe208000.dpi: Error applying setting, reverse things back
```

The resource map showed:

```text
audremap default pins: GPIO12 & GPIO13
dpi24: GPIO0-27
vc4-kms-dpi-generic rgb888: GPIO0-27
dpi18cpadhi: GPIO0-9, GPIO12-17, GPIO20-25
vc4-kms-dpi-generic rgb666-padhi: GPIO0-9, GPIO12-17, GPIO20-25
```

Interpretation:

- `audremap` claimed GPIO12/13 for PWM audio.
- KMS DPI needed GPIO12 and failed to bind.
- `rgb888` was also too broad for this hardware, claiming GPIO0-27.
- `rgb666-padhi` matched the safer padded 18-bit DPI layout and avoided GPIO18/19 and GPIO26/27.

## GPIO / Resource Map

Known or inferred case resources:

| Resource | GPIO / Device | Notes |
|---|---:|---|
| KMS DPI LCD | GPIO0-9, GPIO12-17, GPIO20-25 | Working via `vc4-kms-dpi-generic,rgb666-padhi`. |
| PWM audio via `audremap` | GPIO12/13 by default | Conflicts with KMS DPI. Leave disabled for now. |
| RetroFlag dock / HDMI detect script | GPIO18 | Old `lcdfirst.sh`/`lcdnext.sh` use GPIO18 as `HDMI_HPD`. |
| Safe shutdown switch input | GPIO26 | `SafeShutdown.py` uses BCM GPIO26 as `powerPin`. |
| Power enable hold | GPIO27 | `SafeShutdown.py` uses BCM GPIO27 as `powerenPin` and drives it high. |
| Built-in controller | USB HID / Xbox 360 style | Appears as Microsoft Xbox 360 pad / GBA Pi Case+ / Nuvoton. |
| USB audio | GeneralPlus USB Audio Device | Detected over internal USB. Potential future audio path if PWM audio remains disabled. |
| LCD/backlight/power-save behavior | Case power board / top button | Screen can turn off while Linux and SSH stay alive. Top button wakes it. |

## Old RetroFlag Scripts

Current legacy scripts are not KMS-aware.

### GPi Case 2 Power Latch Trap

Treat the GPi Case 2 power path like a discovered dungeon hazard: the side
switch looks like a familiar hard power switch, but the verified behavior says
it is part of a latched power-control ritual.

Verified behavior:

- The GPi Case 2 side power switch does not directly cut battery power.
- The side switch appears to signal a GPIO watched by
  `/opt/RetroFlag/SafeShutdown.py`.
- If the RetroFlag script is disabled, the side switch may no longer trigger
  shutdown.
- The top power-save/resume button toggles sleep/wake behavior and is not a
  hard power or reset button.
- In sleep/power-save, Wi-Fi may go down. If that happens, the SSH recovery path
  can vanish while the handheld is still alive in the dark.

Observed process shape from the stock script:

```text
sudo python /opt/RetroFlag/SafeShutdown.py
└─ python /opt/RetroFlag/SafeShutdown.py
   ├─ python /opt/RetroFlag/SafeShutdown.py
   └─ python /opt/RetroFlag/SafeShutdown.py
```

The process tree matches the script's use of `multiprocessing.Process` to spawn
separate workers: one for `poweroff()` and one for `lcdrun()`.

Key verified script signs:

```python
import RPi.GPIO as GPIO
from multiprocessing import Process

GPIO.setup(powerPin, GPIO.IN, pull_up_down=GPIO.PUD_UP)
GPIO.wait_for_edge(powerPin, GPIO.FALLING)
# shutdown -h now

GPIO.setup(powerenPin, GPIO.OUT)
GPIO.output(powerenPin, GPIO.HIGH)

powerProcess = Process(target = poweroff)
lcdrunProcess = Process(target = lcdrun)
```

Interpretation:

- `powerPin` is the side-switch shutdown signal path, observed as BCM GPIO26.
- `powerenPin` is the power-enable latch path, observed as BCM GPIO27.
- The script drives the power-enable GPIO HIGH at startup, likely keeping the
  case power latch enabled.
- `poweroff()` watches the side switch and eventually calls `shutdown -h now`.
- `lcdrun()` participates in LCD/power-save/resume behavior and currently calls
  the old RetroFlag LCD scripts.

Warning:

Do not disable the RetroFlag `SafeShutdown.py` script on real GPi Case 2
hardware until `retroflag-powerd` has a verified replacement for the power
latch, shutdown switch, and power-save/resume behavior.

The replacement boundary is mapped in the
[SafeShutdown replacement boundary map](../03-operations/safeshutdown-replacement-boundary-map.md).
Use that compass before planning any service install, `rc.local` change, GPIO
write, shutdown execution, or `SafeShutdown.py` replacement.

Safe replacement requires `retroflag-powerd` to own all three parts of the
trapdoor mechanism:

1. Power-enable latch behavior.
2. Side-switch shutdown detection.
3. Power-save/resume behavior currently handled by `lcdrun()`.

### `/etc/rc.local`

Previously started:

```sh
sh /opt/RetroFlag/lcdfirst.sh&
sudo python /opt/RetroFlag/SafeShutdown.py&
```

`lcdfirst.sh` was disabled during KMS investigation. `SafeShutdown.py` may still restart after boot unless commented out in `/etc/rc.local`.

### `/opt/RetroFlag/SafeShutdown.py`

Observed script behavior:

```python
powerPin = 26
powerenPin = 27

GPIO.setup(powerPin, GPIO.IN, pull_up_down=GPIO.PUD_UP)
GPIO.setup(powerenPin, GPIO.OUT)
GPIO.output(powerenPin, GPIO.HIGH)
```

Power-off loop:

```python
GPIO.wait_for_edge(powerPin, GPIO.FALLING)
sudo killall emulationstation
sudo sleep 5s
sudo shutdown -h now
```

LCD loop:

```python
while True:
    os.system("sh /opt/RetroFlag/lcdnext.sh")
    time.sleep(1)
```

The LCD loop is dangerous for the KMS migration because `lcdnext.sh` rewrites `/boot/config.txt` between LCD and HDMI modes. It also expected GPIO18 to be exported and produced:

```text
/sys/class/gpio/gpio18/value: No such file or directory
/opt/RetroFlag/lcdnext.sh: 17
```

Recommended future direction: replace the old script with a small systemd service that only handles power enable and shutdown. Do not keep LCD/HDMI config switching in the safe-shutdown path.

## Power-Save Behavior

Observed behavior:

- Screen turns off.
- Power indicator flashes.
- SSH remains connected.
- EmulationStation continues running.
- Face buttons do not wake the screen.
- The top button wakes the screen.
- On wake, a fresh undervoltage event may appear.

This strongly suggests a GPi Case 2 case-level power-save mode rather than Linux suspend.

Observed wake-time undervoltage:

```text
hwmon hwmon1: Undervoltage detected!
hwmon hwmon1: Voltage normalised
```

Working theory:

```text
Case enters its own LCD/backlight/power-save mode.
Top button wakes the LCD/backlight/power rail.
Wake causes a short current spike.
If charger/cable/headroom is marginal, the CM4 sees undervoltage.
```

The old SafeShutdown script may not be required for the display-off event itself, but it may participate in later auto-shutdown behavior or power-state handling.

## Power Quality Notes

Observed throttling/power state:

```text
vcgencmd get_throttled = 0x50000
```

Observed dmesg:

```text
Undervoltage detected
Voltage normalised
```

Likely causes:

1. USB charger/cable cannot maintain enough voltage under load.
2. GPi Case 2 charge/load-sharing board adds voltage drop or current limits.
3. LCD/backlight wake creates a short current surge.
4. The system is running CM4 + LCD + USB hub + controller + USB audio + Wi-Fi + battery/power board.

RetroFlag lists the case input requirement as 5V/3A. A random USB-A 5V/2A supply may be marginal.

Future hardware investigation should identify:

- USB-C/DC input path.
- Battery charger IC.
- Load-sharing / power-path IC.
- Boost/buck regulator feeding CM4.
- Whether the board expects only 5V input.
- Whether a USB-C 5V/3A sink or PD trigger + 5.1V buck retrofit is practical.

Do not feed 9V/12V into the RetroFlag board unless the board is proven to accept it. Assume the board expects 5V.

## Button and Input Notes

During EmulationStation button mapping:

- Main controls were detected as an Xbox 360-style gamepad.
- The built-in controller appeared in logs as a Nuvoton `GBA Pi Case+` / Microsoft Xbox 360 pad.
- The button above Select and left of the RetroFlag logo button was not detected by EmulationStation like the other face/control buttons.
- The top button wakes the case from power-save mode.
- The top button appears to be case/power-board behavior rather than a normal gamepad input.

Future work should identify all event devices:

```bash
cat /proc/bus/input/devices
ls -la /dev/input/by-id /dev/input/by-path 2>/dev/null
evtest
```

The undetected button may be wired to the case power-management board, a separate GPIO, or a HID event that EmulationStation ignores.

## Known-Good Proof Commands

Display/KMS proof:

```bash
cat /proc/fb
kmsprint 2>&1 | sed -n '1,120p'
dmesg | grep -iE "vc4|v3d|drm|dpi|gpio12|gpio13|framebuffer|kms|mailbox|pinctrl|panel" | tail -240
```

EmulationStation renderer proof:

```bash
tail -80 /opt/retropie/configs/all/emulationstation/es_log.txt \
  | grep -iE "GL vendor|GL renderer|GL version|llvmpipe|V3D|Broadcom"
```

Power watch:

```bash
watch -n 2 'date; vcgencmd measure_temp; vcgencmd get_throttled; ps -C emulationstation -o pid,pcpu,pmem,etime,cmd; dmesg | grep -i "undervoltage\|voltage normalised" | tail -8'
```

Script/process watch:

```bash
ps -ef | grep -E "SafeShutdown|lcdnext|lcdfirst|python|emulationstation" | grep -v grep
```

## Recommended Follow-Up Quests

### Quest: Build a KMS-Safe Power Keeper

Goal:

- Replace legacy `SafeShutdown.py` with a systemd service.
- Hold GPIO27 high if required by the power board.
- Watch GPIO26 for switch-off.
- Preserve the top-button power-save/resume behavior currently guarded by the
  old `lcdrun()` path.
- Cleanly shut down EmulationStation and Linux.
- Never run `lcdnext.sh`.
- Never rewrite `/boot/config.txt`.
- Log power events clearly.

### Quest: Add an EmulationStation GPU Guard

Goal:

- Start EmulationStation only if KMS/DRM is healthy.
- Refuse startup if `DPI-1` is missing.
- Refuse startup if Mesa renderer is `llvmpipe`.
- Log a clear diagnostic to `/tmp` or `/var/log/retroflag-power`.

### Quest: Map GPi Case 2 Buttons

Goal:

- Identify all `/dev/input` devices.
- Run `evtest` against each device.
- Press every physical control.
- Document which controls are gamepad HID inputs, GPIO inputs, power-board inputs, or undetected.
- Pay special attention to the button above Select and left of the RetroFlag logo.

### Quest: Power Port Cartography

Goal:

- Open case and inspect power board.
- Identify charger/power-path/regulator chips.
- Confirm whether input is strictly 5V.
- Decide whether a USB-C 5V/3A sink retrofit is enough or whether USB-C PD + buck-to-5.1V is worth exploring.
- Preserve battery safety and original charging behavior.

## Current Standing

The screen dragon is slain:

```text
KMS DPI works.
EmulationStation renders on Broadcom V3D.
The LCD is no longer trapped behind FKMS or llvmpipe.
```

The next dungeon is power behavior:

```text
Understand GPi Case 2 power-save mode.
Replace old RetroFlag scripts.
Improve power-supply headroom.
Map the extra buttons.
```
