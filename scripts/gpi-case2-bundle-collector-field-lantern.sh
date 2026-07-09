#!/bin/sh
set -eu

SCRIPT_NAME="$(basename "$0")"
DEFAULT_DURATION_SECONDS=90
DURATION_SECONDS="$DEFAULT_DURATION_SECONDS"
PLAIN=0
OUTPUT_PARENT="${HOME:-.}"
TOTAL_STEPS=6
DMESG_PATTERN="under-voltage|undervoltage|voltage|thrott|vc4|v3d|drm|dpi|kms|framebuffer|panel|mailbox|usb|hid|input|xpad|audio|snd|mmc|ext4|filesystem|I/O error|rcu|stall|hung|blocked|watchdog"
FRONTEND_PATTERN="emulationstation|emulationstation.sh|EmulationStation|ES-DE|retroarch|runcommand"

usage() {
	cat <<'EOF'
GPi Case 2 Bundle Collector Field Lantern

Manual post-boot evidence collector for a responsive GPi Case 2.

What it collects:
  - Remembered boot and current-session log excerpts for power, display, USB,
    controller, RCU, watchdog, MMC, and filesystem clues.
  - vcgencmd get_throttled near the top of power evidence.
  - vcgencmd measure_volts clearly labeled as internal/core rail evidence,
    not the GPi Case 2 5V input rail.
  - Temperature, uptime, command availability, selected boot context files,
    ps snapshots, and best-effort frontend detection.
  - One sample per second for the requested duration.

What it does not prove:
  - It is not a true early boot-time recorder.
  - It cannot determine the exact second of early boot undervoltage unless a
    boot-time recorder was already active.
  - It cannot report watts, amps, TDP, power draw, or actual 5V rail voltage.
  - Frontend/process detection is a clue reported as detected, not detected,
    or uncertain. It is not proof that EmulationStation is or is not visible.

Expected duration:
  - About 90 seconds by default, plus a short bundling stage.
  - Use --duration SECONDS for a shorter field retest or local smoke test.

Safety boundaries:
  - No GPIO reads or writes.
  - No shutdown, reboot, halt, suspend, or resume.
  - No systemd activation or service changes.
  - No rc.local replacement.
  - No SafeShutdown.py replacement.
  - No firmware flashing, installer execution, automatic fixes, telemetry, or
    hardware modification instructions.

Output artifact pattern:
  /home/retropi/gpi-case2-bundle-collector-field-lantern-YYYYMMDD-HHMMSS.tar.gz

Options:
  --duration SECONDS  Capture sample duration. Default: 90.
  --output-dir DIR    Parent directory for the timestamped satchel. Default: $HOME.
  --plain             Disable color, glyphs, and live terminal control.
  --help              Show this help.

Examples:
  scp scripts/gpi-case2-bundle-collector-field-lantern.sh retropi@gpi:/home/retropi/
  ssh retropi@gpi
  sh /home/retropi/gpi-case2-bundle-collector-field-lantern.sh --duration 90
  scp retropi@gpi:/home/retropi/gpi-case2-bundle-collector-field-lantern-YYYYMMDD-HHMMSS.tar.gz .
EOF
}

positive_integer() {
	case "$1" in
		''|*[!0-9]*)
			return 1
			;;
	esac
	[ "$1" -gt 0 ]
}

while [ "$#" -gt 0 ]; do
	case "$1" in
		--help|-h)
			usage
			exit 0
			;;
		--plain)
			PLAIN=1
			shift
			;;
		--duration)
			if [ "$#" -lt 2 ]; then
				printf 'ERROR: --duration needs a positive integer number of seconds.\n' >&2
				exit 2
			fi
			DURATION_SECONDS="$2"
			shift 2
			;;
		--duration=*)
			DURATION_SECONDS="${1#--duration=}"
			shift
			;;
		--output-dir)
			if [ "$#" -lt 2 ]; then
				printf 'ERROR: --output-dir needs a directory path.\n' >&2
				exit 2
			fi
			OUTPUT_PARENT="$2"
			shift 2
			;;
		--output-dir=*)
			OUTPUT_PARENT="${1#--output-dir=}"
			shift
			;;
		--*)
			printf 'ERROR: Unknown option: %s\n' "$1" >&2
			printf 'Action: Run %s --help for supported options.\n' "$SCRIPT_NAME" >&2
			exit 2
			;;
		*)
			if positive_integer "$1" && [ "$DURATION_SECONDS" = "$DEFAULT_DURATION_SECONDS" ]; then
				DURATION_SECONDS="$1"
				shift
			else
				printf 'ERROR: Unexpected argument: %s\n' "$1" >&2
				printf 'Action: Use --duration SECONDS, or run %s --help.\n' "$SCRIPT_NAME" >&2
				exit 2
			fi
			;;
	esac
done

if ! positive_integer "$DURATION_SECONDS"; then
	printf 'ERROR: duration must be a positive integer number of seconds.\n' >&2
	exit 2
fi

IS_INTERACTIVE=0
if [ -t 1 ] && [ "$PLAIN" -eq 0 ]; then
	IS_INTERACTIVE=1
fi

USE_GLYPHS=0
case "${LC_ALL:-}${LC_CTYPE:-}${LANG:-}" in
	*UTF-8*|*utf8*|*UTF8*)
		if [ "$IS_INTERACTIVE" -eq 1 ]; then
			USE_GLYPHS=1
		fi
		;;
esac

USE_COLOR=0
if [ "$IS_INTERACTIVE" -eq 1 ] && [ -z "${NO_COLOR:-}" ]; then
	USE_COLOR=1
fi

if [ "$USE_COLOR" -eq 1 ]; then
	C_RESET="$(printf '\033[0m')"
	C_BOLD="$(printf '\033[1m')"
	C_DIM="$(printf '\033[2m')"
	C_GREEN="$(printf '\033[32m')"
	C_YELLOW="$(printf '\033[33m')"
	C_CYAN="$(printf '\033[36m')"
else
	C_RESET=''
	C_BOLD=''
	C_DIM=''
	C_GREEN=''
	C_YELLOW=''
	C_CYAN=''
fi

glyph() {
	label="$1"
	if [ "$USE_GLYPHS" -eq 1 ]; then
		case "$label" in
			lantern) printf '🏮' ;;
			scroll) printf '📖' ;;
			spark) printf '✨' ;;
			map) printf '🧭' ;;
			disk) printf '💾' ;;
			done) printf '✅' ;;
			*) printf '•' ;;
		esac
	else
		printf '[%s]' "$label"
	fi
}

elapsed_hms() {
	total="$1"
	h=$((total / 3600))
	m=$(((total % 3600) / 60))
	s=$((total % 60))
	if [ "$h" -gt 0 ]; then
		printf '%02d:%02d:%02d' "$h" "$m" "$s"
	else
		printf '%02d:%02d' "$m" "$s"
	fi
}

now_epoch() {
	date +%s 2>/dev/null || printf '0'
}

now_utc() {
	date -u +%Y-%m-%dT%H:%M:%SZ 2>/dev/null || date 2>/dev/null || printf 'unknown-time'
}

step_line() {
	current="$1"
	label="$2"
	text="$3"
	printf '%s[[%s/%s]]%s %s %s\n' "$C_CYAN" "$current" "$TOTAL_STEPS" "$C_RESET" "$(glyph "$label")" "$text"
}

stage_begin() {
	STAGE_STEP="$1"
	STAGE_LABEL="$2"
	STAGE_TEXT="$3"
	STAGE_KEY="$4"
	STAGE_START_EPOCH="$(now_epoch)"
	step_line "$STAGE_STEP" "$STAGE_LABEL" "$STAGE_TEXT"
}

stage_end() {
	end_epoch="$(now_epoch)"
	duration=$((end_epoch - STAGE_START_EPOCH))
	{
		printf '%s\t%s\t%s\t%s\t%s\n' "$STAGE_STEP" "$STAGE_KEY" "$STAGE_START_EPOCH" "$end_epoch" "$duration"
	} >> "${ROOT}/timing.tsv"
}

csv_escape() {
	printf '%s' "$1" | sed 's/"/""/g; s/^/"/; s/$/"/'
}

one_line() {
	tr '\n' ' ' | sed 's/[[:space:]]*$//'
}

command_status() {
	if command -v "$1" >/dev/null 2>&1; then
		printf 'available'
	else
		printf 'unavailable'
	fi
}

run_command() {
	label="$1"
	shift
	output="${ROOT}/${label}.txt"
	{
		printf 'command:'
		for arg in "$@"; do
			printf ' %s' "$arg"
		done
		printf '\n'

		if command -v "$1" >/dev/null 2>&1; then
			"$@" 2>&1 || printf '\ncommand-exit: non-zero\n'
		else
			printf 'command-unavailable: %s\n' "$1"
		fi
	} > "${output}"
}

command_value() {
	if command -v "$1" >/dev/null 2>&1; then
		"$@" 2>&1 || true
	else
		printf 'command-unavailable:%s' "$1"
	fi
}

copy_if_readable() {
	source_path="$1"
	safe_name="$(printf '%s' "$source_path" | sed 's#^/##; s#[^A-Za-z0-9._-]#_#g')"
	output="${ROOT}/file-${safe_name}.txt"

	if [ -r "$source_path" ]; then
		{
			printf 'source: %s\n' "$source_path"
			printf 'status: readable\n\n'
			cat "$source_path" 2>&1 || printf '\nread-error\n'
		} > "$output"
	else
		{
			printf 'source: %s\n' "$source_path"
			printf 'status: unavailable-or-not-readable\n'
		} > "$output"
	fi
}

latest_dmesg_match() {
	if command -v dmesg >/dev/null 2>&1; then
		if dmesg_output="$(dmesg 2>&1)"; then
			printf '%s\n' "$dmesg_output" | grep -Ei "${DMESG_PATTERN}" | tail -n 1 || \
				printf 'no-matching-dmesg-lines'
		else
			printf 'dmesg-read-unavailable:%s' "$(printf '%s' "$dmesg_output" | one_line)"
		fi
	else
		printf 'command-unavailable:dmesg'
	fi
}

write_matching_dmesg_excerpt() {
	output="$1"
	{
		printf 'pattern: %s\n\n' "${DMESG_PATTERN}"
		if command -v dmesg >/dev/null 2>&1; then
			if dmesg_output="$(dmesg 2>&1)"; then
				printf '%s\n' "$dmesg_output" | grep -Ei "${DMESG_PATTERN}" | tail -n 120 || \
					printf 'no-matching-dmesg-lines\n'
			else
				printf 'dmesg-read-unavailable:\n%s\n' "$dmesg_output"
			fi
		else
			printf 'command-unavailable: dmesg\n'
		fi
	} > "$output"
}

write_matching_journal_excerpt() {
	output="$1"
	{
		printf 'pattern: %s\n\n' "${DMESG_PATTERN}"
		if command -v journalctl >/dev/null 2>&1; then
			if journal_output="$(journalctl -b --no-pager 2>&1)"; then
				printf '%s\n' "$journal_output" | grep -Ei "${DMESG_PATTERN}" | tail -n 160 || \
					printf 'no-matching-journal-lines\n'
			else
				printf 'journal-read-unavailable:\n%s\n' "$journal_output"
			fi
		else
			printf 'command-unavailable: journalctl\n'
		fi
	} > "$output"
}

write_ps_snapshot() {
	name="$1"
	output="${ROOT}/ps-${name}.txt"
	{
		printf 'snapshot: %s\n' "$name"
		printf 'captured_at: %s\n' "$(now_utc)"
		printf 'frontend_pattern: %s\n\n' "$FRONTEND_PATTERN"
		if command -v ps >/dev/null 2>&1; then
			ps axww -o pid=,ppid=,stat=,comm=,args= 2>&1 || ps aux 2>&1 || ps 2>&1 || true
		else
			printf 'command-unavailable: ps\n'
		fi
	} > "$output"
}

frontend_detection() {
	if command -v pgrep >/dev/null 2>&1; then
		if pgrep -af "$FRONTEND_PATTERN" >/dev/null 2>&1; then
			printf 'detected'
			return
		fi
		printf 'not_detected'
		return
	fi

	if command -v ps >/dev/null 2>&1; then
		if ps axww 2>/dev/null | grep -Ei "$FRONTEND_PATTERN" | grep -v grep >/dev/null 2>&1; then
			printf 'detected'
		else
			printf 'not_detected'
		fi
		return
	fi

	printf 'uncertain'
}

write_frontend_detection_details() {
	output="${ROOT}/frontend-detection.txt"
	{
		printf 'frontend_detection_policy: detected / not_detected / uncertain\n'
		printf 'important_note: detection is best-effort evidence, not proof. The first field run missed visibly running EmulationStation.\n'
		printf 'patterns: %s\n\n' "$FRONTEND_PATTERN"
		printf 'pgrep_status: %s\n' "$(command_status pgrep)"
		printf 'ps_status: %s\n\n' "$(command_status ps)"
		if command -v pgrep >/dev/null 2>&1; then
			printf 'pgrep_matches:\n'
			pgrep -af "$FRONTEND_PATTERN" 2>&1 || printf 'no-pgrep-matches\n'
			printf '\n'
		fi
		if command -v ps >/dev/null 2>&1; then
			printf 'ps_matching_lines:\n'
			ps axww 2>/dev/null | grep -Ei "$FRONTEND_PATTERN" | grep -v grep 2>&1 || printf 'no-ps-matches\n'
		fi
	} > "$output"
}

print_progress() {
	sample="$1"
	total="$2"
	elapsed="$3"
	remaining=$((total - sample))
	if [ "$remaining" -lt 0 ]; then
		remaining=0
	fi

	if [ "$IS_INTERACTIVE" -eq 1 ]; then
		width=20
		filled=$((sample * width / total))
		bar=''
		j=0
		while [ "$j" -lt "$width" ]; do
			if [ "$j" -lt "$filled" ]; then
				bar="${bar}#"
			else
				bar="${bar}-"
			fi
			j=$((j + 1))
		done
		printf '\rSamples [%s] %s/%s elapsed %s ETA %s' "$bar" "$sample" "$total" "$(elapsed_hms "$elapsed")" "$(elapsed_hms "$remaining")"
	else
		if [ "$sample" -eq 1 ] || [ "$sample" -eq "$total" ] || [ $((sample % 5)) -eq 0 ]; then
			printf '[[3/%s]] %s Sample %s/%s, elapsed %s, remaining about %s.\n' "$TOTAL_STEPS" "$(glyph spark)" "$sample" "$total" "$(elapsed_hms "$elapsed")" "$(elapsed_hms "$remaining")"
		fi
	fi
}

STAMP="$(date +%Y%m%d-%H%M%S 2>/dev/null || printf 'unknown-time')"
ROOT="${OUTPUT_PARENT%/}/gpi-case2-bundle-collector-field-lantern-${STAMP}"
BUNDLE="${ROOT}.tar.gz"
CAPTURE_START_EPOCH="$(now_epoch)"
CAPTURE_START_UTC="$(now_utc)"

if [ "$PLAIN" -eq 0 ] && [ "$IS_INTERACTIVE" -eq 1 ]; then
	printf '%s' "$C_YELLOW"
	printf '        .-.\n'
	printf '       (   )\n'
	printf '        | |\n'
	printf '      __| |__\n'
	printf '     /  | |  \\\n'
	printf '    /___| |___\\\n'
	printf '        |_|\n'
	printf '%s' "$C_RESET"
fi

printf '%sGPi Case 2 Bundle Collector Field Lantern%s\n' "$C_BOLD" "$C_RESET"
printf 'Save Room Tech: local, read-only, human-carried.\n'
printf 'SAFETY: no GPIO reads, no GPIO writes, no shutdown, no reboot, no service changes.\n'
printf 'Output satchel: %s\n' "$ROOT"
printf 'Expected duration: about %s seconds plus bundling.\n\n' "$DURATION_SECONDS"

stage_begin 1 lantern "Preparing capture satchel..." "prepare"
if ! mkdir -p "${ROOT}"; then
	printf 'ERROR: Could not create output directory: %s\n' "$ROOT" >&2
	printf 'Action: Check free space and permissions, then rerun the command.\n' >&2
	exit 1
fi
{
	printf 'step\tstage\tstart_epoch\tend_epoch\tduration_seconds\n'
} > "${ROOT}/timing.tsv"
{
	printf 'GPi Case 2 Bundle Collector Field Lantern timing\n'
	printf 'capture_start_utc: %s\n' "$CAPTURE_START_UTC"
	printf 'capture_start_epoch: %s\n' "$CAPTURE_START_EPOCH"
	printf 'note: unusually slow capture stages may be diagnostic evidence.\n\n'
} > "${ROOT}/timing.txt"

{
	printf 'GPi Case 2 Bundle Collector Field Lantern\n'
	printf 'captured_local_stamp: %s\n' "${STAMP}"
	printf 'capture_start_utc: %s\n' "$CAPTURE_START_UTC"
	printf 'duration_seconds_requested: %s\n' "${DURATION_SECONDS}"
	printf 'output_folder: %s\n' "${ROOT}"
	printf 'bundle: %s\n\n' "${BUNDLE}"
	printf 'safety: local read-only capture; no GPIO reads, no GPIO writes, no shutdown, no reboot, no systemd activation, no installer, no telemetry, no automatic fixes.\n'
	printf 'note: manual post-boot evidence collector; not a true early boot recorder.\n'
	printf 'note: keep the handheld active and stop before idle auto power-save can trigger. Avoid the top sleep/resume button unless a procedure explicitly says otherwise.\n'
	printf 'note: the side power switch is the normal stock shutdown control while the system is responsive.\n'
	printf 'note: get_throttled reports firmware throttling flags, not watts, TDP, amps, power draw, or actual 5V rail voltage.\n'
	printf 'note: measure_volts reports an internal/core rail, not the GPi Case 2 5V input rail.\n'
	printf 'note: frontend process detection is a clue reported as detected, not_detected, or uncertain; it is not authoritative proof.\n\n'
	printf 'power_evidence:\n'
	printf -- '- vcgencmd_get_throttled: %s\n' "$(command_value vcgencmd get_throttled | one_line)"
	printf -- '- vcgencmd_measure_volts_internal_core_rail_not_5v_input: %s\n' "$(command_value vcgencmd measure_volts | one_line)"
	printf -- '- vcgencmd_measure_temp: %s\n\n' "$(command_value vcgencmd measure_temp | one_line)"
	printf 'command_availability:\n'
	for cmd in date uname uptime cat awk sed grep tail pgrep ps tar dmesg journalctl systemd-analyze lsusb mount df free vcgencmd; do
		printf -- '- %s: %s\n' "$cmd" "$(command_status "$cmd")"
	done
} > "${ROOT}/report.txt"

{
	printf 'GPi Case 2 Bundle Collector Field Lantern manifest\n'
	printf 'captured_local_stamp: %s\n' "${STAMP}"
	printf 'capture_start_utc: %s\n' "$CAPTURE_START_UTC"
	printf 'script: %s\n' "$SCRIPT_NAME"
	printf 'portable: yes\n'
	printf 'capture_kind: manual-post-boot-evidence-collector\n'
	printf 'true_boot_time_recorder: no\n'
	printf 'requires_repo_checkout: no\n'
	printf 'requires_git: no\n'
	printf 'requires_go: no\n'
	printf 'requires_root_writes: no\n'
	printf 'network_contact: no\n'
} > "${ROOT}/manifest.txt"
stage_end

stage_begin 2 scroll "Reading boot scrolls..." "boot-scrolls"
copy_if_readable /proc/cmdline
copy_if_readable /proc/uptime
copy_if_readable /proc/device-tree/model
copy_if_readable /boot/config.txt
copy_if_readable /boot/cmdline.txt
copy_if_readable /boot/firmware/config.txt
copy_if_readable /boot/firmware/cmdline.txt

run_command uname uname -a
run_command uptime uptime
run_command vcgencmd-get-throttled vcgencmd get_throttled
run_command vcgencmd-measure-volts-internal-core-rail-not-5v-input vcgencmd measure_volts
run_command vcgencmd-measure-temp vcgencmd measure_temp
run_command mount mount
run_command df df -h
run_command free free -h
run_command lsusb lsusb
run_command systemd-analyze-blame systemd-analyze blame
run_command systemd-analyze-critical-chain systemd-analyze critical-chain

write_matching_dmesg_excerpt "${ROOT}/dmesg-power-display-usb-xpad-rcu-watchdog-mmc-ext4.txt"
write_matching_journal_excerpt "${ROOT}/journal-power-display-usb-xpad-rcu-watchdog-mmc-ext4.txt"
write_ps_snapshot "before-sampling"
write_frontend_detection_details
stage_end

stage_begin 3 spark "Sampling power runes..." "sampling"
printf '%s\n' \
	'captured_at,uptime_seconds,throttled,core_volts_internal_not_5v_input,temp,frontend_detection,latest_dmesg_match' \
	> "${ROOT}/trace.csv"

{
	printf 'GPi Case 2 Bundle Collector process milestones\n'
	printf 'captured_local_stamp: %s\n\n' "${STAMP}"
	printf 'note: frontend/process detection is best-effort and may miss a visible frontend.\n'
	printf 'states: detected, not_detected, uncertain\n\n'
} > "${ROOT}/process-milestones.txt"

FRONTEND_WAS_DETECTED=0
i=0
while [ "$i" -lt "$DURATION_SECONDS" ]; do
	sample=$((i + 1))
	NOW="$(now_utc)"
	UPTIME="$(awk '{print $1}' /proc/uptime 2>/dev/null || printf 'unknown')"
	THROTTLED="$(command_value vcgencmd get_throttled | one_line)"
	VOLTS="$(command_value vcgencmd measure_volts | one_line)"
	TEMP="$(command_value vcgencmd measure_temp | one_line)"
	FRONTEND_STATE="$(frontend_detection)"

	if [ "$FRONTEND_STATE" = "detected" ] && [ "$FRONTEND_WAS_DETECTED" -eq 0 ]; then
		printf '%s uptime=%s frontend first observed detected by process search\n' "$NOW" "$UPTIME" \
			>> "${ROOT}/process-milestones.txt"
		FRONTEND_WAS_DETECTED=1
	fi

	LATEST_DMESG="$(latest_dmesg_match | one_line)"

	{
		csv_escape "$NOW"; printf ','
		csv_escape "$UPTIME"; printf ','
		csv_escape "$THROTTLED"; printf ','
		csv_escape "$VOLTS"; printf ','
		csv_escape "$TEMP"; printf ','
		csv_escape "$FRONTEND_STATE"; printf ','
		csv_escape "$LATEST_DMESG"; printf '\n'
	} >> "${ROOT}/trace.csv"

	print_progress "$sample" "$DURATION_SECONDS" "$sample"

	i=$sample
	sleep 1
done
if [ "$IS_INTERACTIVE" -eq 1 ]; then
	printf '\n'
fi
write_ps_snapshot "after-sampling"
write_frontend_detection_details
stage_end

stage_begin 4 map "Writing Ledger notes..." "ledger"
CAPTURE_END_PRE_BUNDLE_EPOCH="$(now_epoch)"
{
	printf '\ncompleted_sampling_at: %s\n' "$(now_utc)"
	printf 'final_throttled: %s\n' "$(command_value vcgencmd get_throttled | one_line)"
	printf 'final_core_volts_internal_not_5v_input: %s\n' "$(command_value vcgencmd measure_volts | one_line)"
	printf 'final_temp: %s\n' "$(command_value vcgencmd measure_temp | one_line)"
	printf 'final_frontend_detection: %s\n\n' "$(frontend_detection)"
	printf 'collector_limits:\n'
	printf -- '- gathers remembered boot logs after the handheld is responsive\n'
	printf -- '- samples current and sticky get_throttled state\n'
	printf -- '- cannot determine the exact second of early boot undervoltage unless a boot-time recorder was already running\n'
	printf -- '- cannot report watts, TDP, amps, power draw, or actual 5V rail voltage\n'
	printf -- '- cannot make frontend/process detection authoritative\n'
} >> "${ROOT}/report.txt"
{
	printf 'capture_sampling_end_utc: %s\n' "$(now_utc)"
	printf 'capture_sampling_elapsed_seconds: %s\n' "$((CAPTURE_END_PRE_BUNDLE_EPOCH - CAPTURE_START_EPOCH))"
} >> "${ROOT}/manifest.txt"
stage_end

stage_begin 5 disk "Sealing archive..." "archive"
if command -v tar >/dev/null 2>&1; then
	tar -czf "${BUNDLE}" -C "$(dirname "${ROOT}")" "$(basename "${ROOT}")"
	ARTIFACT="${BUNDLE}"
else
	printf 'command-unavailable: tar\n' > "${ROOT}/bundle-unavailable.txt"
	ARTIFACT="${ROOT}"
fi
stage_end

CAPTURE_END_EPOCH="$(now_epoch)"
CAPTURE_END_UTC="$(now_utc)"
TOTAL_DURATION=$((CAPTURE_END_EPOCH - CAPTURE_START_EPOCH))
{
	printf '\nper_stage_durations:\n'
	awk 'NR > 1 { printf "- step_%s_%s: %s seconds\n", $1, $2, $5 }' "${ROOT}/timing.tsv" 2>/dev/null || true
	printf '\ncapture_end_utc: %s\n' "$CAPTURE_END_UTC"
	printf 'capture_end_epoch: %s\n' "$CAPTURE_END_EPOCH"
	printf 'total_duration_seconds: %s\n' "$TOTAL_DURATION"
	printf 'total_duration_human: %s\n' "$(elapsed_hms "$TOTAL_DURATION")"
} >> "${ROOT}/timing.txt"
{
	printf 'capture_end_utc: %s\n' "$CAPTURE_END_UTC"
	printf 'total_duration_seconds: %s\n' "$TOTAL_DURATION"
	printf 'total_duration_human: %s\n' "$(elapsed_hms "$TOTAL_DURATION")"
	printf 'timing_note: unusually slow capture stages may be diagnostic evidence.\n'
} >> "${ROOT}/report.txt"

if [ "$ARTIFACT" = "$BUNDLE" ]; then
	tar -czf "${BUNDLE}" -C "$(dirname "${ROOT}")" "$(basename "${ROOT}")"
fi

step_line 6 done "Lantern sealed."
printf '\n'
printf '%sArtifact:%s %s\n' "$C_GREEN" "$C_RESET" "$ARTIFACT"
printf 'Folder:   %s\n' "$ROOT"
printf 'Duration: %s\n\n' "$(elapsed_hms "$TOTAL_DURATION")"
printf 'Retrieve from your workstation:\n'
if [ "$ARTIFACT" = "$BUNDLE" ]; then
	printf 'scp retropi@gpi:%s .\n' "$BUNDLE"
else
	printf 'scp -r retropi@gpi:%s .\n' "$ROOT"
fi
