#!/bin/sh
set -eu

DURATION_SECONDS="${1:-90}"
case "$DURATION_SECONDS" in
	''|*[!0-9]*)
		printf 'error: duration must be a positive integer number of seconds\n' >&2
		exit 2
		;;
esac
if [ "$DURATION_SECONDS" -eq 0 ]; then
	printf 'error: duration must be greater than zero seconds\n' >&2
	exit 2
fi

STAMP="$(date +%Y%m%d-%H%M%S 2>/dev/null || printf 'unknown-time')"
ROOT="${HOME:-.}/gpi-case2-bundle-collector-field-lantern-${STAMP}"
BUNDLE="${ROOT}.tar.gz"
DMESG_PATTERN="under-voltage|undervoltage|voltage|thrott|vc4|v3d|drm|dpi|kms|framebuffer|panel|mailbox|usb|hid|input|xpad|audio|snd|mmc|ext4|filesystem|I/O error|rcu|stall|hung|blocked|watchdog"

mkdir -p "${ROOT}"

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

{
	printf 'GPi Case 2 Bundle Collector Field Lantern\n'
	printf 'captured_local_stamp: %s\n' "${STAMP}"
	printf 'duration_seconds: %s\n' "${DURATION_SECONDS}"
	printf 'output_folder: %s\n' "${ROOT}"
	printf 'bundle: %s\n\n' "${BUNDLE}"
	printf 'safety: local read-only capture; no GPIO reads, no GPIO writes, no shutdown, no reboot, no systemd activation, no installer, no telemetry, no automatic fixes.\n'
	printf 'note: manual post-boot bundle collector; not a true early boot recorder.\n'
	printf 'note: keep the handheld active and stop before idle auto power-save can trigger. Avoid the top sleep/resume button unless a procedure explicitly says otherwise.\n'
	printf 'note: get_throttled reports firmware throttling flags, not watts, TDP, amps, power draw, or actual 5V rail voltage.\n'
	printf 'note: measure_volts reports an internal/core rail, not the GPi Case 2 5V input rail.\n'
	printf 'note: emulationstation process detection is a clue, not authoritative proof.\n\n'
	printf 'command_availability:\n'
	for cmd in date uname uptime cat awk sed grep tail pgrep tar dmesg journalctl systemd-analyze lsusb mount df free vcgencmd; do
		printf -- '- %s: %s\n' "$cmd" "$(command_status "$cmd")"
	done
} > "${ROOT}/report.txt"

{
	printf 'GPi Case 2 Bundle Collector Field Lantern manifest\n'
	printf 'captured_local_stamp: %s\n' "${STAMP}"
	printf 'script: gpi-case2-bundle-collector-field-lantern.sh\n'
	printf 'portable: yes\n'
	printf 'capture_kind: manual-post-boot-bundle-collector\n'
	printf 'true_boot_time_recorder: no\n'
	printf 'requires_repo_checkout: no\n'
	printf 'requires_git: no\n'
	printf 'requires_go: no\n'
	printf 'requires_root_writes: no\n'
	printf 'network_contact: no\n'
} > "${ROOT}/manifest.txt"

copy_if_readable /proc/cmdline
copy_if_readable /proc/uptime
copy_if_readable /proc/device-tree/model
copy_if_readable /boot/config.txt
copy_if_readable /boot/cmdline.txt
copy_if_readable /boot/firmware/config.txt
copy_if_readable /boot/firmware/cmdline.txt

run_command uname uname -a
run_command uptime uptime
run_command mount mount
run_command df df -h
run_command free free -h
run_command lsusb lsusb
run_command vcgencmd-get-throttled vcgencmd get_throttled
run_command vcgencmd-measure-volts vcgencmd measure_volts
run_command vcgencmd-measure-temp vcgencmd measure_temp
run_command systemd-analyze-blame systemd-analyze blame
run_command systemd-analyze-critical-chain systemd-analyze critical-chain

write_matching_dmesg_excerpt "${ROOT}/dmesg-power-display-usb-xpad-rcu-watchdog-mmc-ext4.txt"
write_matching_journal_excerpt "${ROOT}/journal-power-display-usb-xpad-rcu-watchdog-mmc-ext4.txt"

printf '%s\n' \
	'captured_at,uptime_seconds,throttled,core_volts_not_5v_input,temp,emulationstation_running_unreliable,latest_dmesg_match' \
	> "${ROOT}/trace.csv"

{
	printf 'GPi Case 2 Bundle Collector process milestones\n'
	printf 'captured_local_stamp: %s\n\n' "${STAMP}"
	printf 'note: process detection is best-effort and may miss a visible frontend.\n\n'
} > "${ROOT}/process-milestones.txt"

ES_WAS_RUNNING=0
i=0
while [ "$i" -lt "$DURATION_SECONDS" ]; do
	NOW="$(date -u +%Y-%m-%dT%H:%M:%SZ 2>/dev/null || date 2>/dev/null || printf 'unknown-time')"
	UPTIME="$(awk '{print $1}' /proc/uptime 2>/dev/null || printf 'unknown')"
	THROTTLED="$(command_value vcgencmd get_throttled | one_line)"
	VOLTS="$(command_value vcgencmd measure_volts | one_line)"
	TEMP="$(command_value vcgencmd measure_temp | one_line)"

	if command -v pgrep >/dev/null 2>&1 && pgrep -x emulationstation >/dev/null 2>&1; then
		ES_RUNNING="yes"
		if [ "$ES_WAS_RUNNING" -eq 0 ]; then
			printf '%s uptime=%s emulationstation first observed running\n' "$NOW" "$UPTIME" \
				>> "${ROOT}/process-milestones.txt"
			ES_WAS_RUNNING=1
		fi
	else
		ES_RUNNING="no"
	fi

	LATEST_DMESG="$(latest_dmesg_match | one_line)"

	{
		csv_escape "$NOW"; printf ','
		csv_escape "$UPTIME"; printf ','
		csv_escape "$THROTTLED"; printf ','
		csv_escape "$VOLTS"; printf ','
		csv_escape "$TEMP"; printf ','
		csv_escape "$ES_RUNNING"; printf ','
		csv_escape "$LATEST_DMESG"; printf '\n'
	} >> "${ROOT}/trace.csv"

	i=$((i + 1))
	sleep 1
done

{
	printf '\ncompleted_at: %s\n' "$(date -u +%Y-%m-%dT%H:%M:%SZ 2>/dev/null || date 2>/dev/null || printf 'unknown-time')"
	printf 'final_throttled: %s\n' "$(command_value vcgencmd get_throttled | one_line)"
	printf 'final_core_volts_not_5v_input: %s\n' "$(command_value vcgencmd measure_volts | one_line)"
	printf 'final_temp: %s\n\n' "$(command_value vcgencmd measure_temp | one_line)"
	printf 'collector_limits:\n'
	printf -- '- gathers remembered boot logs after the handheld is responsive\n'
	printf -- '- samples current and sticky get_throttled state\n'
	printf -- '- cannot determine the exact second of early boot undervoltage unless a boot-time recorder was already running\n'
	printf -- '- cannot report watts, TDP, amps, power draw, or actual 5V rail voltage\n'
} >> "${ROOT}/report.txt"

if command -v tar >/dev/null 2>&1; then
	tar -czf "${BUNDLE}" -C "$(dirname "${ROOT}")" "$(basename "${ROOT}")"
	printf '%s\n' "${BUNDLE}"
else
	printf 'command-unavailable: tar\n' > "${ROOT}/bundle-unavailable.txt"
	printf '%s\n' "${ROOT}"
fi
