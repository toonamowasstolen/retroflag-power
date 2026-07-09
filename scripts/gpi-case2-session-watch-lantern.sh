#!/bin/sh
set -eu

SCRIPT_NAME="$(basename "$0")"
DEFAULT_DURATION_SECONDS=300
DEFAULT_INTERVAL_SECONDS=15
DURATION_SECONDS="$DEFAULT_DURATION_SECONDS"
INTERVAL_SECONDS="$DEFAULT_INTERVAL_SECONDS"
PLAIN=0
OUTPUT_FILE=''
TOTAL_STEPS=4
START_EPOCH=0
START_UTC=''
START_LOCAL=''
SAMPLES=0
CHECKPOINTS=''
WARNINGS=''
MIN_TEMP=''
MAX_TEMP=''
FRONTEND_EVER_DETECTED='no'
FRONTEND_FIRST_SAMPLE='unavailable'
THROTTLED_VALUES=''

usage() {
	cat <<'EOF'
GPi Case 2 Session Watch Lantern

Read-only foreground session watcher for a responsive GPi Case 2 handheld.

What it watches:
  - Host, user, kernel, and uptime metadata where safely available.
  - Bounded runtime checkpoints for throttling, temperature, load, memory,
    disk space, frontend/emulator process clues, and narrow kernel warnings.
  - One final Ledger artifact file that can be retrieved or shared after
    review instead of pasting long terminal output.

Safety boundaries:
  - Read-only observation only.
  - No GPIO reads or writes.
  - No shutdown, reboot, halt, suspend, sleep, resume, power, display, config,
    service, installer, firmware, or SafeShutdown behavior changes.
  - The side switch remains the normal stock shutdown path while responsive.
  - The top sleep/resume button remains suspect unless a procedure explicitly
    says otherwise.

Default artifact path:
  /home/retropi/gpi-case2-session-watch-lantern-YYYYMMDD-HHMMSS.txt

Options:
  --duration SECONDS  Bounded watch length. Default: 300.
  --interval SECONDS  Seconds between checkpoints. Default: 15.
  --output FILE       Final Ledger artifact path. Default: $HOME timestamp file.
  --plain             Disable color, glyphs, and live terminal control.
  --help              Show this help.

Examples:
  scp scripts/gpi-case2-session-watch-lantern.sh retropi@gpi:/home/retropi/
  ssh retropi@gpi
  sh /home/retropi/gpi-case2-session-watch-lantern.sh --duration 600
  scp retropi@gpi:/home/retropi/gpi-case2-session-watch-lantern-YYYYMMDD-HHMMSS.txt .
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

loadavg() {
	read_first_line /proc/loadavg
}

uptime_seconds() {
	if [ -r /proc/uptime ]; then
		awk '{print $1}' /proc/uptime 2>/dev/null || printf 'unavailable'
	else
		printf 'unavailable'
	fi
}

mem_available_kb() {
	if [ -r /proc/meminfo ]; then
		awk '/^MemAvailable:/ {print $2 " kB"; found=1} END {if (!found) print "unavailable"}' /proc/meminfo 2>/dev/null || printf 'unavailable'
	else
		printf 'unavailable'
	fi
}

disk_root_available() {
	if command -v df >/dev/null 2>&1; then
		df -k / 2>/dev/null | awk 'NR==2 {print $4 " kB available on /"; found=1} END {if (!found) print "unavailable"}'
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

dmesg_hint() {
	pattern='under-voltage|undervoltage|voltage|thrott|vc4|v3d|drm|dpi|kms|framebuffer|panel|usb|hid|input|audio|snd|mmc|ext4|filesystem|I/O error|rcu|stall|hung|watchdog|suspend|resume|power-save'
	if command -v dmesg >/dev/null 2>&1; then
		dmesg 2>/dev/null | grep -Ei "$pattern" | tail -n 1 | one_line || printf 'no_recent_matching_kernel_hint'
	else
		printf 'command_unavailable:dmesg'
	fi
}

step_line() {
	current="$1"
	label="$2"
	text="$3"
	printf '%s[[%s/%s]]%s %s %s\n' "$C_CYAN" "$current" "$TOTAL_STEPS" "$C_RESET" "$(glyph "$label")" "$text"
}

add_warning() {
	if [ -z "$WARNINGS" ]; then
		WARNINGS="$1"
	else
		WARNINGS="${WARNINGS}
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

warning_count() {
	if [ -n "$WARNINGS" ]; then
		printf '%s\n' "$WARNINGS" | awk 'END {print NR}'
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

sample_once() {
	elapsed="$1"
	SAMPLES=$((SAMPLES + 1))
	captured_at="$(now_utc)"
	uptime_value="$(uptime_seconds)"
	throttled="$(command_one_line vcgencmd get_throttled)"
	temp="$(command_one_line vcgencmd measure_temp)"
	volts="$(command_one_line vcgencmd measure_volts)"
	load_value="$(loadavg)"
	mem_value="$(mem_available_kb)"
	disk_value="$(disk_root_available)"
	frontend_value="$(frontend_hint)"
	kernel_hint="$(dmesg_hint)"

	track_throttled_value "$throttled"
	track_temp_value "$temp"
	track_frontend_value "$frontend_value"

	case "$throttled$temp$volts$kernel_hint" in
		*command_unavailable*)
			add_warning "sample ${SAMPLES}: one or more optional commands were unavailable"
			;;
	esac

	append_checkpoint "${SAMPLES}|${captured_at}|elapsed=${elapsed}s|uptime=${uptime_value}|throttled=${throttled}|temp=${temp}|volts_internal_core=${volts}|load=${load_value}|mem_available=${mem_value}|disk=${disk_value}|frontend=${frontend_value}|kernel_hint=${kernel_hint}"
}

write_artifact() {
	status="$1"
	end_epoch="$(now_epoch)"
	end_utc="$(now_utc)"
	end_local="$(now_local)"
	total_elapsed=$((end_epoch - START_EPOCH))
	if [ "$total_elapsed" -lt 0 ]; then
		total_elapsed=0
	fi
	ssh_watch_completed='no'
	if [ "$status" = "completed" ]; then
		ssh_watch_completed='yes'
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
	warnings_total="$(warning_count)"

	{
		printf 'GPi Case 2 Session Watch Lantern Ledger\n'
		printf '\n'
		printf 'status: %s\n' "$status"
		printf 'read_only_marker: READ-ONLY / NO CHANGES MADE\n'
		printf 'artifact_path: %s\n' "$OUTPUT_FILE"
		printf '\n'
		printf 'Session Metadata\n'
		printf 'started_utc: %s\n' "$START_UTC"
		printf 'started_local: %s\n' "$START_LOCAL"
		printf 'ended_utc: %s\n' "$end_utc"
		printf 'ended_local: %s\n' "$end_local"
		printf 'requested_duration_seconds: %s\n' "$DURATION_SECONDS"
		printf 'sample_interval_seconds: %s\n' "$INTERVAL_SECONDS"
		printf 'observed_duration_seconds: %s\n' "$total_elapsed"
		printf 'observed_duration: %s\n' "$(elapsed_hms "$total_elapsed")"
		printf 'sample_count: %s\n' "$SAMPLES"
		printf '\n'
		printf 'Host Metadata\n'
		printf 'hostname: %s\n' "$(command_one_line hostname)"
		printf 'user: %s\n' "$(command_one_line id -un)"
		printf 'kernel: %s\n' "$(command_one_line uname -a)"
		printf 'uptime_now: %s\n' "$(command_one_line uptime)"
		printf 'proc_uptime_now: %s\n' "$(uptime_seconds)"
		printf '\n'
		printf 'Safety Boundary\n'
		printf '%s\n' '- This script writes only this final artifact file.'
		printf '%s\n' '- It does not change power, display, GPIO, config, services, shutdown, sleep, or resume behavior.'
		printf '%s\n' '- The side switch remains the normal stock shutdown path while responsive.'
		printf '%s\n' '- The top sleep/resume button remains suspect unless a procedure explicitly says otherwise.'
		printf '%s\n' '- SSH to retropi@gpi is optional support, not the primary handheld UX.'
		printf '\n'
		printf 'Observed Checkpoints\n'
		printf 'sample|captured_at|details\n'
		if [ -n "$CHECKPOINTS" ]; then
			printf '%s\n' "$CHECKPOINTS"
		else
			printf 'none\n'
		fi
		printf '\n'
		printf 'Warnings And Missing Evidence\n'
		if [ -n "$WARNINGS" ]; then
			printf '%s\n' "$WARNINGS"
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
		printf 'temperature_min: %s\n' "$temp_min_summary"
		printf 'temperature_max: %s\n' "$temp_max_summary"
		printf 'frontend_detected_ever: %s\n' "$FRONTEND_EVER_DETECTED"
		printf 'frontend_first_detected_sample: %s\n' "$FRONTEND_FIRST_SAMPLE"
		printf 'ssh_side_watch_completed_normally: %s\n' "$ssh_watch_completed"
		printf 'throttled_raw_values_observed: %s\n' "$throttled_summary"
		printf 'throttling_note: Raw vcgencmd get_throttled values are preserved for later interpretation; do not overclaim cause, power state, battery state, charger state, or emulator performance from this Ledger alone.\n'
		printf 'warnings_missing_evidence_count: %s\n' "$warnings_total"
		printf '\n'
		printf 'Final Artifact Path\n'
		printf '%s\n' "$OUTPUT_FILE"
	} > "$OUTPUT_FILE"
}

finish_interrupted() {
	printf '\n'
	step_line 4 ledger "Watch interrupted; sealing the Ledger with completed checkpoints."
	write_artifact "interrupted"
	printf 'Status:   interrupted\n'
	printf 'Artifact: %s\n' "$OUTPUT_FILE"
	exit 130
}

case "$OUTPUT_FILE" in
	'')
		OUTPUT_FILE="${HOME:-.}/gpi-case2-session-watch-lantern-$(stamp_local).txt"
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

trap finish_interrupted INT TERM

printf '%sGPi Case 2 Session Watch Lantern%s\n' "$C_BOLD" "$C_RESET"
printf '%s %s\n' "$(glyph lantern)" "Read-only handheld session watch. No changes made."
printf 'Duration: %s (%s seconds)\n' "$(elapsed_hms "$DURATION_SECONDS")" "$DURATION_SECONDS"
printf 'Interval: %s seconds\n' "$INTERVAL_SECONDS"
printf 'Artifact: %s\n' "$OUTPUT_FILE"
printf '\n'

step_line 1 map "Preparing the session map."
printf 'Relic:    GPi Case 2 handheld-first watch\n'
printf 'Started:  %s\n' "$START_LOCAL"
printf 'Promise:  read-only / no power, display, GPIO, config, service, or shutdown changes\n'
printf '\n'

step_line 2 rune "Watching checkpoints."
elapsed=0
while [ "$elapsed" -lt "$DURATION_SECONDS" ]; do
	sample_once "$elapsed"
	remaining=$((DURATION_SECONDS - elapsed))
	printf 'Progress: elapsed=%s remaining=%s samples=%s status=watching\n' "$(elapsed_hms "$elapsed")" "$(elapsed_hms "$remaining")" "$SAMPLES"

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
done

if [ "$elapsed" -lt "$DURATION_SECONDS" ]; then
	elapsed="$DURATION_SECONDS"
fi
sample_once "$elapsed"
printf 'Progress: elapsed=%s remaining=00:00 samples=%s status=sampled\n' "$(elapsed_hms "$elapsed")" "$SAMPLES"
printf '\n'

step_line 3 ledger "Writing the final Ledger artifact."
write_artifact "completed"

step_line 4 done "Lantern sealed."
total=$(( $(now_epoch) - START_EPOCH ))
if [ "$total" -lt 0 ]; then
	total=0
fi
printf 'Status:   completed\n'
printf 'Duration: %s\n' "$(elapsed_hms "$total")"
printf 'Artifact: %s\n' "$OUTPUT_FILE"
printf 'Reminder: inspect and redact the Ledger before sharing.\n'
