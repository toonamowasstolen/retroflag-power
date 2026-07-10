#!/bin/sh
set -eu

SCRIPT_NAME="$(basename "$0")"
DEFAULT_DURATION_SECONDS=120
DEFAULT_INTERVAL_SECONDS=5
DURATION_SECONDS="$DEFAULT_DURATION_SECONDS"
INTERVAL_SECONDS="$DEFAULT_INTERVAL_SECONDS"
PLAIN=0
OUTPUT_FILE=''
TOTAL_STEPS=5
START_EPOCH=0
START_UTC=''
START_LOCAL=''
START_PROC_UPTIME=''
SAMPLES=0
CHECKPOINTS=''
WARNINGS=''
MISSING_EVIDENCE=''
THROTTLED_VALUES=''
MIN_TEMP=''
MAX_TEMP=''
FRONTEND_EVER_DETECTED='no'
FRONTEND_FIRST_SAMPLE='unavailable'
DISPLAY_SUMMARY='unavailable'
FIRST_DISPLAY_HINT_SAMPLE='unavailable'
FIRST_DISPLAY_HINT_SUMMARY='unavailable'
FIRST_SYSTEMD_TIMING_SAMPLE='unavailable'
FIRST_SYSTEMD_TIMING_SUMMARY='unavailable'
FIRST_JOURNAL_HINT_SAMPLE='unavailable'
FIRST_JOURNAL_HINT_SUMMARY='unavailable'
FIRST_DMESG_HINT_SAMPLE='unavailable'
FIRST_DMESG_HINT_SUMMARY='unavailable'
SYSTEMD_TIMING_STATUS='not_checked'
JOURNAL_STATUS='not_checked'
DMESG_STATUS='not_checked'
LOAD_LAST='unavailable'
MEM_LAST='unavailable'
DISK_LAST='unavailable'

usage() {
	cat <<'EOF'
GPi Case 2 True Boot Trace Lantern

Read-only foreground startup trace for a responsive GPi Case 2 handheld.

What it gathers:
  - Host, user, kernel, uptime, and /proc uptime metadata.
  - Bounded boot-trace checkpoints for systemd timing, selected journal/dmesg
    boot hints, frontend clues, framebuffer/display hints, vcgencmd values,
    memory, disk, and load.
  - One final Boot Trace Ledger artifact file that can be retrieved with scp.

Safety boundaries:
  - Read-only observation only.
  - No boot config, service, GPIO, power, display, shutdown, sleep/resume,
    RetroPie, or EmulationStation config changes.
  - The side switch remains the normal stock shutdown path while responsive.
  - The top sleep/resume button remains suspect unless a procedure explicitly
    says otherwise.
  - SSH to retropi@gpi is optional support, not the primary handheld UX.

Default artifact path:
  /home/retropi/gpi-case2-true-boot-trace-lantern-YYYYMMDD-HHMMSS.txt

Options:
  --duration SECONDS  Bounded capture length. Default: 120.
  --interval SECONDS  Seconds between checkpoints. Default: 5.
  --output FILE       Final Boot Trace Ledger path. Default: $HOME timestamp file.
  --plain             Disable color, glyphs, and live terminal control.
  --help              Show this help.

Examples:
  scp scripts/gpi-case2-true-boot-trace-lantern.sh retropi@gpi:/home/retropi/
  ssh retropi@gpi
  sh /home/retropi/gpi-case2-true-boot-trace-lantern.sh --plain --duration 120
  scp retropi@gpi:/home/retropi/gpi-case2-true-boot-trace-lantern-YYYYMMDD-HHMMSS.txt .
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
		--interval)
			if [ "$#" -lt 2 ]; then
				printf 'ERROR: --interval needs a positive integer number of seconds.\n' >&2
				exit 2
			fi
			INTERVAL_SECONDS="$2"
			shift 2
			;;
		--interval=*)
			INTERVAL_SECONDS="${1#--interval=}"
			shift
			;;
		--output)
			if [ "$#" -lt 2 ]; then
				printf 'ERROR: --output needs a file path.\n' >&2
				exit 2
			fi
			OUTPUT_FILE="$2"
			shift 2
			;;
		--output=*)
			OUTPUT_FILE="${1#--output=}"
			shift
			;;
		--*)
			printf 'ERROR: Unknown option: %s\n' "$1" >&2
			printf 'Action: Run %s --help for supported options.\n' "$SCRIPT_NAME" >&2
			exit 2
			;;
		*)
			printf 'ERROR: Unexpected argument: %s\n' "$1" >&2
			printf 'Action: Run %s --help for supported options.\n' "$SCRIPT_NAME" >&2
			exit 2
			;;
	esac
done

if ! positive_integer "$DURATION_SECONDS"; then
	printf 'ERROR: duration must be a positive integer number of seconds.\n' >&2
	exit 2
fi

if ! positive_integer "$INTERVAL_SECONDS"; then
	printf 'ERROR: interval must be a positive integer number of seconds.\n' >&2
	exit 2
fi

IS_INTERACTIVE=0
if [ -t 1 ] && [ "$PLAIN" -eq 0 ]; then
	IS_INTERACTIVE=1
fi

USE_COLOR=0
if [ "$IS_INTERACTIVE" -eq 1 ] && [ -z "${NO_COLOR:-}" ]; then
	USE_COLOR=1
fi

USE_GLYPHS=0
case "${LC_ALL:-}${LC_CTYPE:-}${LANG:-}" in
	*UTF-8*|*utf8*|*UTF8*)
		if [ "$IS_INTERACTIVE" -eq 1 ]; then
			USE_GLYPHS=1
		fi
		;;
esac

if [ "$USE_COLOR" -eq 1 ]; then
	C_RESET="$(printf '\033[0m')"
	C_BOLD="$(printf '\033[1m')"
	C_DIM="$(printf '\033[2m')"
	C_GREEN="$(printf '\033[32m')"
	C_CYAN="$(printf '\033[36m')"
else
	C_RESET=''
	C_BOLD=''
	C_DIM=''
	C_GREEN=''
	C_CYAN=''
fi

glyph() {
	label="$1"
	if [ "$USE_GLYPHS" -eq 1 ]; then
		case "$label" in
			lantern) printf '🏮' ;;
			map) printf '🧭' ;;
			rune) printf '✦' ;;
			ledger) printf '📖' ;;
			done) printf '✅' ;;
			*) printf '•' ;;
		esac
	else
		printf '[%s]' "$label"
	fi
}

now_epoch() {
	date +%s 2>/dev/null || printf '0'
}

now_utc() {
	date -u +%Y-%m-%dT%H:%M:%SZ 2>/dev/null || date 2>/dev/null || printf 'unknown-time'
}

now_local() {
	date +%Y-%m-%dT%H:%M:%S%z 2>/dev/null || date 2>/dev/null || printf 'unknown-time'
}

stamp_local() {
	date +%Y%m%d-%H%M%S 2>/dev/null || date -u +%Y%m%d-%H%M%S 2>/dev/null || printf 'unknown-time'
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

one_line() {
	tr '\n' ' ' | sed 's/[[:space:]]*$//'
}

read_first_line() {
	path="$1"
	if [ -r "$path" ]; then
		sed -n '1p' "$path" 2>/dev/null | one_line
	else
		printf 'unavailable'
	fi
}

command_one_line() {
	if command -v "$1" >/dev/null 2>&1; then
		"$@" 2>&1 | one_line
	else
		printf 'command_unavailable:%s' "$1"
	fi
}

command_status_text() {
	name="$1"
	shift
	if ! command -v "$name" >/dev/null 2>&1; then
		printf 'command_unavailable:%s' "$name"
		return
	fi
	output="$("$name" "$@" 2>&1 || true)"
	if [ -z "$output" ]; then
		printf 'empty_output:%s' "$name"
	else
		printf '%s' "$output" | one_line
	fi
}

uptime_seconds() {
	if [ -r /proc/uptime ]; then
		awk '{print $1}' /proc/uptime 2>/dev/null || printf 'unavailable'
	else
		printf 'unavailable'
	fi
}

loadavg() {
	read_first_line /proc/loadavg
}

mem_summary() {
	if [ -r /proc/meminfo ]; then
		awk '/^MemTotal:/ {total=$2 " kB"} /^MemAvailable:/ {avail=$2 " kB"} END {if (total || avail) printf "total=%s available=%s", total, avail; else print "unavailable"}' /proc/meminfo 2>/dev/null || printf 'unavailable'
	elif command -v free >/dev/null 2>&1; then
		free -h 2>&1 | one_line
	else
		printf 'unavailable'
	fi
}

disk_summary() {
	if command -v df >/dev/null 2>&1; then
		df -k / "${HOME:-/home/retropi}" 2>/dev/null | awk 'NR==1 {next} {printf "%s:%s_kB_available ", $6, $4; found=1} END {if (!found) print "unavailable"}' | one_line
	else
		printf 'command_unavailable:df'
	fi
}

frontend_hint() {
	if command -v ps >/dev/null 2>&1; then
		if ps -eo comm 2>/dev/null | grep -Eiq '^(emulationstation|retroarch|runcommand|pegasus-fe|attract|es-de)$'; then
			printf 'detected'
		else
			printf 'not_detected'
		fi
	else
		printf 'uncertain:ps_unavailable'
	fi
}

display_hint() {
	graphics='unavailable'
	drm='unavailable'
	fbdev='unavailable'
	if [ -d /sys/class/graphics ]; then
		graphics="$(find /sys/class/graphics -maxdepth 1 -type l -o -type d 2>/dev/null | sed 's#.*/##' | sort | one_line)"
	fi
	if [ -d /sys/class/drm ]; then
		drm="$(find /sys/class/drm -maxdepth 1 -type l -o -type d 2>/dev/null | sed 's#.*/##' | sort | one_line)"
	fi
	if ls /dev/fb* >/dev/null 2>&1; then
		fbdev="$(ls /dev/fb* 2>/dev/null | one_line)"
	fi
	printf 'graphics=%s drm=%s fbdev=%s' "$graphics" "$drm" "$fbdev"
}

systemd_boot_timing() {
	if command -v systemd-analyze >/dev/null 2>&1; then
		output="$(systemd-analyze time 2>&1 || true)"
		if [ -z "$output" ]; then
			SYSTEMD_TIMING_STATUS='empty_output'
			printf 'empty_output:systemd-analyze'
		elif printf '%s\n' "$output" | grep -Eiq 'permission|denied|failed|not booted|offline'; then
			SYSTEMD_TIMING_STATUS='unavailable_or_restricted'
			printf '%s' "$output" | one_line
		else
			SYSTEMD_TIMING_STATUS='captured'
			printf '%s' "$output" | one_line
		fi
	else
		SYSTEMD_TIMING_STATUS='command_unavailable'
		printf 'command_unavailable:systemd-analyze'
	fi
}

journal_boot_hints() {
	pattern='under-voltage|undervoltage|voltage|thrott|vc4|v3d|drm|dpi|kms|framebuffer|panel|gpio12|usb|hid|input|xpad|audio|snd|mmc|ext4|I/O error|rcu|stall|hung|watchdog|emulationstation|retroarch'
	if command -v journalctl >/dev/null 2>&1; then
		output="$(journalctl -b --no-pager -n 300 2>&1 | grep -Ei "$pattern" | tail -n 12 || true)"
		if [ -z "$output" ]; then
			JOURNAL_STATUS='no_matching_boot_hints'
			printf 'no_matching_boot_hints'
		elif printf '%s\n' "$output" | grep -Eiq 'permission|denied|failed|No journal files'; then
			JOURNAL_STATUS='unavailable_or_restricted'
			printf '%s' "$output" | one_line
		else
			JOURNAL_STATUS='captured'
			printf '%s' "$output" | one_line
		fi
	else
		JOURNAL_STATUS='command_unavailable'
		printf 'command_unavailable:journalctl'
	fi
}

dmesg_boot_hints() {
	pattern='under-voltage|undervoltage|voltage|thrott|vc4|v3d|drm|dpi|kms|framebuffer|panel|gpio12|usb|hid|input|xpad|audio|snd|mmc|ext4|I/O error|rcu|stall|hung|watchdog|emulationstation|retroarch'
	if command -v dmesg >/dev/null 2>&1; then
		output="$(dmesg 2>&1 | grep -Ei "$pattern" | tail -n 12 || true)"
		if [ -z "$output" ]; then
			DMESG_STATUS='no_matching_boot_hints'
			printf 'no_matching_boot_hints'
		elif printf '%s\n' "$output" | grep -Eiq 'permission|denied|failed|Operation not permitted'; then
			DMESG_STATUS='unavailable_or_restricted'
			printf '%s' "$output" | one_line
		else
			DMESG_STATUS='captured'
			printf '%s' "$output" | one_line
		fi
	else
		DMESG_STATUS='command_unavailable'
		printf 'command_unavailable:dmesg'
	fi
}

ssh_context() {
	if [ -n "${SSH_CONNECTION:-}" ] || [ -n "${SSH_CLIENT:-}" ] || [ -n "${SSH_TTY:-}" ]; then
		printf 'likely_ssh_session'
	elif [ -t 0 ]; then
		printf 'local_or_unknown_terminal'
	else
		printf 'unknown_noninteractive'
	fi
}

add_warning() {
	if [ -z "$WARNINGS" ]; then
		WARNINGS="$1"
	else
		WARNINGS="${WARNINGS}
$1"
	fi
}

add_missing() {
	if [ -n "$MISSING_EVIDENCE" ] && printf '%s\n' "$MISSING_EVIDENCE" | grep -Fx "$1" >/dev/null 2>&1; then
		return
	fi
	if [ -z "$MISSING_EVIDENCE" ]; then
		MISSING_EVIDENCE="$1"
	else
		MISSING_EVIDENCE="${MISSING_EVIDENCE}
$1"
	fi
}

append_checkpoint() {
	line="$1"
	if [ -z "$CHECKPOINTS" ]; then
		CHECKPOINTS="$line"
	else
		CHECKPOINTS="${CHECKPOINTS}
$line"
	fi
}

line_count_or_zero() {
	if [ -n "$1" ]; then
		printf '%s\n' "$1" | awk 'END {print NR}'
	else
		printf '0'
	fi
}

track_throttled_value() {
	value="$1"
	if [ -z "$THROTTLED_VALUES" ]; then
		THROTTLED_VALUES="$value"
		return
	fi
	if ! printf '%s\n' "$THROTTLED_VALUES" | grep -Fx "$value" >/dev/null 2>&1; then
		THROTTLED_VALUES="${THROTTLED_VALUES}
$value"
	fi
}

track_temp_value() {
	value="$1"
	temp_number="$(printf '%s\n' "$value" | sed -n "s/^temp=\([-0-9.][0-9.]*\)'C$/\1/p")"
	if [ -z "$temp_number" ]; then
		return
	fi
	if [ -z "$MIN_TEMP" ]; then
		MIN_TEMP="$temp_number"
		MAX_TEMP="$temp_number"
		return
	fi
	if awk "BEGIN {exit !($temp_number < $MIN_TEMP)}"; then
		MIN_TEMP="$temp_number"
	fi
	if awk "BEGIN {exit !($temp_number > $MAX_TEMP)}"; then
		MAX_TEMP="$temp_number"
	fi
}

track_frontend_value() {
	value="$1"
	if [ "$value" = "detected" ] && [ "$FRONTEND_EVER_DETECTED" = "no" ]; then
		FRONTEND_EVER_DETECTED='yes'
		FRONTEND_FIRST_SAMPLE="$SAMPLES"
	fi
}

track_first_summary_value() {
	kind="$1"
	value="$2"
	sample="$3"
	case "$value" in
		''|unavailable|command_unavailable:*|empty_output:*|no_matching_boot_hints|*permission*|*Permission*|*denied*|*Operation\ not\ permitted*|*restricted*|*not\ booted*|*No\ journal\ files*)
			return
			;;
	esac
	if [ "$kind" = "display" ] && [ "$value" = "graphics=unavailable drm=unavailable fbdev=unavailable" ]; then
		return
	fi
	case "$kind" in
		display)
			if [ "$FIRST_DISPLAY_HINT_SAMPLE" = "unavailable" ]; then
				FIRST_DISPLAY_HINT_SAMPLE="$sample"
				FIRST_DISPLAY_HINT_SUMMARY="$value"
			fi
			;;
		systemd)
			if [ "$FIRST_SYSTEMD_TIMING_SAMPLE" = "unavailable" ]; then
				FIRST_SYSTEMD_TIMING_SAMPLE="$sample"
				FIRST_SYSTEMD_TIMING_SUMMARY="$value"
			fi
			;;
		journal)
			if [ "$FIRST_JOURNAL_HINT_SAMPLE" = "unavailable" ]; then
				FIRST_JOURNAL_HINT_SAMPLE="$sample"
				FIRST_JOURNAL_HINT_SUMMARY="$value"
			fi
			;;
		dmesg)
			if [ "$FIRST_DMESG_HINT_SAMPLE" = "unavailable" ]; then
				FIRST_DMESG_HINT_SAMPLE="$sample"
				FIRST_DMESG_HINT_SUMMARY="$value"
			fi
			;;
	esac
}

record_optional_evidence() {
	field="$1"
	value="$2"
	case "$value" in
		*command_unavailable*|*permission*|*Permission*|*denied*|*Operation\ not\ permitted*|*empty_output*|unavailable*)
			add_missing "${field}: ${value}"
			;;
	esac
}

evidence_status() {
	value="$1"
	case "$value" in
		command_unavailable:*)
			printf 'command_unavailable'
			;;
		empty_output:*)
			printf 'empty_output'
			;;
		no_matching_boot_hints)
			printf 'no_matching_boot_hints'
			;;
		*permission*|*Permission*|*denied*|*Operation\ not\ permitted*|*restricted*|*not\ booted*|*No\ journal\ files*)
			printf 'unavailable_or_restricted'
			;;
		*)
			printf 'captured'
			;;
	esac
}

sample_once() {
	elapsed="$1"
	SAMPLES=$((SAMPLES + 1))
	captured_at="$(now_utc)"
	proc_uptime="$(uptime_seconds)"
	systemd_timing="$(systemd_boot_timing)"
	journal_hints="$(journal_boot_hints)"
	dmesg_hints="$(dmesg_boot_hints)"
	SYSTEMD_TIMING_STATUS="$(evidence_status "$systemd_timing")"
	JOURNAL_STATUS="$(evidence_status "$journal_hints")"
	DMESG_STATUS="$(evidence_status "$dmesg_hints")"
	throttled="$(command_one_line vcgencmd get_throttled)"
	temp="$(command_one_line vcgencmd measure_temp)"
	volts="$(command_one_line vcgencmd measure_volts)"
	load_value="$(loadavg)"
	mem_value="$(mem_summary)"
	disk_value="$(disk_summary)"
	frontend_value="$(frontend_hint)"
	display_value="$(display_hint)"

	LOAD_LAST="$load_value"
	MEM_LAST="$mem_value"
	DISK_LAST="$disk_value"
	DISPLAY_SUMMARY="$display_value"
	track_throttled_value "$throttled"
	track_temp_value "$temp"
	track_frontend_value "$frontend_value"
	track_first_summary_value display "$display_value" "$SAMPLES"
	track_first_summary_value systemd "$systemd_timing" "$SAMPLES"
	track_first_summary_value journal "$journal_hints" "$SAMPLES"
	track_first_summary_value dmesg "$dmesg_hints" "$SAMPLES"

	record_optional_evidence "systemd_boot_timing" "$systemd_timing"
	record_optional_evidence "journal_boot_hints" "$journal_hints"
	record_optional_evidence "dmesg_boot_hints" "$dmesg_hints"
	record_optional_evidence "vcgencmd_get_throttled" "$throttled"
	record_optional_evidence "vcgencmd_measure_temp" "$temp"
	record_optional_evidence "vcgencmd_measure_volts" "$volts"

	case "$throttled$dmesg_hints$journal_hints" in
		*under-voltage*|*undervoltage*|*thrott*|*voltage*)
			add_warning "sample ${SAMPLES}: boot power or throttling hint present; preserve raw value and avoid overclaiming cause"
			;;
	esac

	append_checkpoint "${SAMPLES}|${captured_at}|elapsed=${elapsed}s|proc_uptime=${proc_uptime}|systemd=${systemd_timing}|journal=${journal_hints}|dmesg=${dmesg_hints}|throttled=${throttled}|temp=${temp}|volts_internal_core=${volts}|load=${load_value}|mem=${mem_value}|disk=${disk_value}|frontend=${frontend_value}|display=${display_value}"
}

timing_bucket() {
	if [ "$FRONTEND_EVER_DETECTED" = "yes" ]; then
		printf 'frontend-start-or-responsive-session'
	elif printf '%s\n%s\n' "$CHECKPOINTS" "$DISPLAY_SUMMARY" | grep -Eiq 'vc4|v3d|drm|kms|dpi|framebuffer|fbdev=/dev/fb'; then
		printf 'display-handoff-or-framebuffer-clues'
	elif [ "$SYSTEMD_TIMING_STATUS" = "captured" ]; then
		printf 'early-userspace'
	else
		printf 'inconclusive'
	fi
}

write_artifact() {
	status="$1"
	end_epoch="$(now_epoch)"
	end_utc="$(now_utc)"
	end_local="$(now_local)"
	end_proc_uptime="$(uptime_seconds)"
	total_elapsed=$((end_epoch - START_EPOCH))
	if [ "$total_elapsed" -lt 0 ]; then
		total_elapsed=0
	fi
	temp_min_summary='unavailable'
	temp_max_summary='unavailable'
	if [ -n "$MIN_TEMP" ]; then
		temp_min_summary="${MIN_TEMP}'C"
		temp_max_summary="${MAX_TEMP}'C"
	fi
	throttled_summary='unavailable'
	if [ -n "$THROTTLED_VALUES" ]; then
		throttled_summary="$(printf '%s' "$THROTTLED_VALUES" | one_line)"
	fi
	warnings_total="$(line_count_or_zero "$WARNINGS")"
	missing_total="$(line_count_or_zero "$MISSING_EVIDENCE")"

	{
		printf 'GPi Case 2 True Boot Trace Lantern Ledger\n'
		printf '\n'
		printf 'status: %s\n' "$status"
		printf 'read_only_marker: READ-ONLY / NO CHANGES MADE\n'
		printf 'artifact_path: %s\n' "$OUTPUT_FILE"
		printf '\n'
		printf 'Boot Trace Metadata\n'
		printf 'started_utc: %s\n' "$START_UTC"
		printf 'started_local: %s\n' "$START_LOCAL"
		printf 'ended_utc: %s\n' "$end_utc"
		printf 'ended_local: %s\n' "$end_local"
		printf 'requested_duration_seconds: %s\n' "$DURATION_SECONDS"
		printf 'sample_interval_seconds: %s\n' "$INTERVAL_SECONDS"
		printf 'observed_duration_seconds: %s\n' "$total_elapsed"
		printf 'observed_duration: %s\n' "$(elapsed_hms "$total_elapsed")"
		printf 'sample_count: %s\n' "$SAMPLES"
		printf 'start_proc_uptime_seconds: %s\n' "$START_PROC_UPTIME"
		printf 'end_proc_uptime_seconds: %s\n' "$end_proc_uptime"
		printf '\n'
		printf 'Host Metadata\n'
		printf 'hostname: %s\n' "$(command_one_line hostname)"
		printf 'user: %s\n' "$(command_one_line id -un)"
		printf 'kernel: %s\n' "$(command_one_line uname -a)"
		printf 'uptime_now: %s\n' "$(command_one_line uptime)"
		printf 'proc_uptime_now: %s\n' "$end_proc_uptime"
		printf 'boot_id: %s\n' "$(read_first_line /proc/sys/kernel/random/boot_id)"
		printf 'device_model: %s\n' "$(read_first_line /proc/device-tree/model)"
		printf 'ssh_context: %s\n' "$(ssh_context)"
		printf '\n'
		printf 'Safety Boundary\n'
		printf '%s\n' '- This script writes only this final Boot Trace Ledger artifact.'
		printf '%s\n' '- It does not change boot config, services, GPIO, power, display, shutdown, sleep/resume, RetroPie, or EmulationStation config.'
		printf '%s\n' '- The side switch remains the normal stock shutdown path while responsive.'
		printf '%s\n' '- The top sleep/resume button remains suspect unless a procedure explicitly says otherwise.'
		printf '%s\n' '- SSH to retropi@gpi is optional support, not the primary handheld UX.'
		printf '\n'
		printf 'Observed Boot Trace Checkpoints\n'
		printf 'sample|captured_at|details\n'
		if [ -n "$CHECKPOINTS" ]; then
			printf '%s\n' "$CHECKPOINTS"
		else
			printf 'none\n'
		fi
		printf '\n'
		printf 'Warnings\n'
		if [ -n "$WARNINGS" ]; then
			printf '%s\n' "$WARNINGS"
		else
			printf 'none_recorded\n'
		fi
		printf '\n'
		printf 'Missing Evidence\n'
		if [ -n "$MISSING_EVIDENCE" ]; then
			printf '%s\n' "$MISSING_EVIDENCE"
		else
			printf 'none_recorded\n'
		fi
		printf '\n'
		printf 'Artifact Summary\n'
		printf 'completion_status: %s\n' "$status"
		printf 'requested_duration_seconds: %s\n' "$DURATION_SECONDS"
		printf 'observed_duration_seconds: %s\n' "$total_elapsed"
		printf 'observed_duration: %s\n' "$(elapsed_hms "$total_elapsed")"
		printf 'sample_count: %s\n' "$SAMPLES"
		printf 'sample_interval_seconds: %s\n' "$INTERVAL_SECONDS"
		printf 'started_utc: %s\n' "$START_UTC"
		printf 'ended_utc: %s\n' "$end_utc"
		printf 'start_proc_uptime_seconds: %s\n' "$START_PROC_UPTIME"
		printf 'end_proc_uptime_seconds: %s\n' "$end_proc_uptime"
		printf 'kernel: %s\n' "$(command_one_line uname -a)"
		printf 'systemd_boot_timing_status: %s\n' "$SYSTEMD_TIMING_STATUS"
		printf 'journal_boot_hints_status: %s\n' "$JOURNAL_STATUS"
		printf 'dmesg_boot_hints_status: %s\n' "$DMESG_STATUS"
		printf 'frontend_detected_ever: %s\n' "$FRONTEND_EVER_DETECTED"
		printf 'frontend_first_detected_sample: %s\n' "$FRONTEND_FIRST_SAMPLE"
		printf 'first_display_hint_sample: %s\n' "$FIRST_DISPLAY_HINT_SAMPLE"
		printf 'first_display_hint_summary: %s\n' "$FIRST_DISPLAY_HINT_SUMMARY"
		printf 'first_systemd_timing_sample: %s\n' "$FIRST_SYSTEMD_TIMING_SAMPLE"
		printf 'first_systemd_timing_summary: %s\n' "$FIRST_SYSTEMD_TIMING_SUMMARY"
		printf 'first_journal_hint_sample: %s\n' "$FIRST_JOURNAL_HINT_SAMPLE"
		printf 'first_journal_hint_summary: %s\n' "$FIRST_JOURNAL_HINT_SUMMARY"
		printf 'first_dmesg_hint_sample: %s\n' "$FIRST_DMESG_HINT_SAMPLE"
		printf 'first_dmesg_hint_summary: %s\n' "$FIRST_DMESG_HINT_SUMMARY"
		printf 'display_hint_summary: %s\n' "$DISPLAY_SUMMARY"
		printf 'throttled_raw_values_observed: %s\n' "$throttled_summary"
		printf 'temperature_min: %s\n' "$temp_min_summary"
		printf 'temperature_max: %s\n' "$temp_max_summary"
		printf 'load_last: %s\n' "$LOAD_LAST"
		printf 'memory_last: %s\n' "$MEM_LAST"
		printf 'disk_last: %s\n' "$DISK_LAST"
		printf 'ssh_context: %s\n' "$(ssh_context)"
		printf 'warnings_count: %s\n' "$warnings_total"
		printf 'missing_evidence_count: %s\n' "$missing_total"
		printf 'cautious_timing_bucket: %s\n' "$(timing_bucket)"
		printf 'first_visible_screen_note: human_observed_not_script_observed\n'
		printf 'interpretation_note: This Boot Trace Ledger captures a bounded post-boot startup trail. It does not prove sleep, resume, shutdown, battery health, charger quality, or a single root cause.\n'
		printf '\n'
		printf 'Final Artifact Path\n'
		printf '%s\n' "$OUTPUT_FILE"
	} > "$OUTPUT_FILE"
}

step_line() {
	current="$1"
	label="$2"
	text="$3"
	printf '%s[[%s/%s]]%s %s %s\n' "$C_CYAN" "$current" "$TOTAL_STEPS" "$C_RESET" "$(glyph "$label")" "$text"
}

finish_interrupted() {
	printf '\n'
	step_line 5 ledger "Trace interrupted; sealing the Boot Trace Ledger with completed checkpoints."
	write_artifact "interrupted"
	printf 'Status:   interrupted\n'
	printf 'Artifact: %s\n' "$OUTPUT_FILE"
	exit 130
}

case "$OUTPUT_FILE" in
	'')
		OUTPUT_FILE="${HOME:-.}/gpi-case2-true-boot-trace-lantern-$(stamp_local).txt"
		;;
	*/*)
		output_parent="$(dirname "$OUTPUT_FILE")"
		if [ ! -d "$output_parent" ]; then
			printf 'ERROR: output parent directory does not exist: %s\n' "$output_parent" >&2
			exit 2
		fi
		;;
esac

START_EPOCH="$(now_epoch)"
START_UTC="$(now_utc)"
START_LOCAL="$(now_local)"
START_PROC_UPTIME="$(uptime_seconds)"

trap finish_interrupted INT TERM

printf '%sGPi Case 2 True Boot Trace Lantern%s\n' "$C_BOLD" "$C_RESET"
printf '%s %s\n' "$(glyph lantern)" "Read-only startup trace. No changes made."
printf 'Duration: %s (%s seconds)\n' "$(elapsed_hms "$DURATION_SECONDS")" "$DURATION_SECONDS"
printf 'Interval: %s seconds\n' "$INTERVAL_SECONDS"
printf 'Artifact: %s\n' "$OUTPUT_FILE"
printf '\n'

step_line 1 map "Preparing the boot map."
printf 'Relic:    GPi Case 2 handheld-first boot trace\n'
printf 'Started:  %s\n' "$START_LOCAL"
printf 'Promise:  read-only / no boot, service, GPIO, power, display, shutdown, sleep, or resume changes\n'
printf '\n'

step_line 2 rune "Taking the first startup snapshot."
sample_once 0
printf 'Progress: elapsed=00:00 remaining=%s samples=%s status=first-snapshot frontend=%s warnings=%s\n' "$(elapsed_hms "$DURATION_SECONDS")" "$SAMPLES" "$FRONTEND_EVER_DETECTED" "$(line_count_or_zero "$WARNINGS")"
printf '\n'

step_line 3 rune "Following bounded boot-trace checkpoints."
elapsed=0
while [ "$elapsed" -lt "$DURATION_SECONDS" ]; do
	next_sleep="$INTERVAL_SECONDS"
	next_elapsed=$((elapsed + next_sleep))
	if [ "$next_elapsed" -gt "$DURATION_SECONDS" ]; then
		next_sleep=$((DURATION_SECONDS - elapsed))
	fi
	if [ "$next_sleep" -le 0 ]; then
		break
	fi
	sleep "$next_sleep"
	elapsed=$((elapsed + next_sleep))
	sample_once "$elapsed"
	remaining=$((DURATION_SECONDS - elapsed))
	printf 'Progress: elapsed=%s remaining=%s samples=%s status=tracing frontend=%s warnings=%s missing=%s\n' "$(elapsed_hms "$elapsed")" "$(elapsed_hms "$remaining")" "$SAMPLES" "$FRONTEND_EVER_DETECTED" "$(line_count_or_zero "$WARNINGS")" "$(line_count_or_zero "$MISSING_EVIDENCE")"
done
printf '\n'

step_line 4 ledger "Writing the final Boot Trace Ledger artifact."
write_artifact "completed"

step_line 5 done "Lantern sealed."
total=$(( $(now_epoch) - START_EPOCH ))
if [ "$total" -lt 0 ]; then
	total=0
fi
printf 'Status:   completed\n'
printf 'Duration: %s\n' "$(elapsed_hms "$total")"
printf 'Artifact: %s\n' "$OUTPUT_FILE"
printf 'Reminder: inspect and redact the Ledger before sharing.\n'
