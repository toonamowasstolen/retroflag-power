#!/bin/sh
set -eu

SCRIPT_NAME="$(basename "$0")"
PLAIN=0
SHOW_ART=1
COMPACT=0

usage() {
	cat <<'EOF'
GPi Case 2 Relic Welcome Scroll Preview

Fast read-only SSH greeting preview for the GPi Case 2 Relic.

Safety boundaries:
  - Manual preview only. Nothing is installed into SSH login, MOTD, shell
    startup, services, boot config, GPIO, display, shutdown, sleep, or resume.
  - Read-only local system-info clues only.
  - Missing tools show as unavailable; the scroll keeps going.

Options:
  --plain      Disable color and decorative banner output.
  --no-art     Hide the ASCII banner in normal mode.
  --compact    Print a shorter field set.
  --help       Show this help.

Examples:
  scp scripts/gpi-case2-relic-welcome-scroll.sh retropi@gpi:/home/retropi/
  ssh retropi@gpi 'chmod +x /home/retropi/gpi-case2-relic-welcome-scroll.sh'
  ssh retropi@gpi '/home/retropi/gpi-case2-relic-welcome-scroll.sh'
  ssh retropi@gpi '/home/retropi/gpi-case2-relic-welcome-scroll.sh --plain'
EOF
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
		--no-art)
			SHOW_ART=0
			shift
			;;
		--compact)
			COMPACT=1
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

if [ -n "${NO_COLOR:-}" ]; then
	PLAIN=1
fi

USE_COLOR=0
if [ "$PLAIN" -eq 0 ] && [ -t 1 ]; then
	USE_COLOR=1
fi

if [ "$USE_COLOR" -eq 1 ]; then
	C_RESET="$(printf '\033[0m')"
	C_BOLD="$(printf '\033[1m')"
	C_DIM="$(printf '\033[2m')"
	C_GREEN="$(printf '\033[32m')"
	C_CYAN="$(printf '\033[36m')"
	C_YELLOW="$(printf '\033[33m')"
else
	C_RESET=''
	C_BOLD=''
	C_DIM=''
	C_GREEN=''
	C_CYAN=''
	C_YELLOW=''
fi

one_line() {
	tr '\n' ' ' | sed 's/^[[:space:]]*//; s/[[:space:]]*$//; s/[[:space:]][[:space:]]*/ /g'
}

command_one_line() {
	if command -v "$1" >/dev/null 2>&1; then
		"$@" 2>/dev/null | one_line
	else
		printf 'unavailable'
	fi
}

field() {
	label="$1"
	value="$2"
	printf '  %-14s %s\n' "$label:" "$value"
}

hostname_value() {
	value="$(command_one_line hostname)"
	if [ "$value" != "unavailable" ] && [ -n "$value" ]; then
		printf '%s\n' "$value"
		return
	fi
	if [ -r /etc/hostname ]; then
		sed -n '1p' /etc/hostname 2>/dev/null | one_line
	else
		printf 'unknown-relic'
	fi
}

user_value() {
	value="$(command_one_line id -un)"
	if [ "$value" != "unavailable" ] && [ -n "$value" ]; then
		printf '%s\n' "$value"
	elif [ -n "${USER:-}" ]; then
		printf '%s\n' "$USER"
	elif [ -n "${LOGNAME:-}" ]; then
		printf '%s\n' "$LOGNAME"
	else
		printf 'unknown-user'
	fi
}

uptime_value() {
	if [ -r /proc/uptime ]; then
		awk '
			{
				total = int($1)
				days = int(total / 86400)
				hours = int((total % 86400) / 3600)
				mins = int((total % 3600) / 60)
				if (days > 0) {
					printf "%dd %dh %dm\n", days, hours, mins
				} else if (hours > 0) {
					printf "%dh %dm\n", hours, mins
				} else {
					printf "%dm\n", mins
				}
				found = 1
			}
			END { if (!found) print "unavailable" }
		' /proc/uptime 2>/dev/null || printf 'unavailable\n'
	else
		value="$(command_one_line uptime -p)"
		if [ "$value" = "unavailable" ] || [ -z "$value" ]; then
			command_one_line uptime
		else
			printf '%s\n' "$value"
		fi
	fi
}

kernel_value() {
	value="$(command_one_line uname -srmo)"
	if [ "$value" = "unavailable" ] || [ -z "$value" ]; then
		command_one_line uname -sr
	else
		printf '%s\n' "$value"
	fi
}

temperature_value() {
	if command -v vcgencmd >/dev/null 2>&1; then
		value="$(vcgencmd measure_temp 2>/dev/null | one_line || printf 'unavailable')"
		if [ -n "$value" ]; then
			printf '%s\n' "$value"
		else
			printf 'unavailable\n'
		fi
	else
		printf 'unavailable\n'
	fi
}

throttled_value() {
	if command -v vcgencmd >/dev/null 2>&1; then
		value="$(vcgencmd get_throttled 2>/dev/null | one_line || printf 'unavailable')"
		if [ -n "$value" ]; then
			printf '%s\n' "$value"
		else
			printf 'unavailable\n'
		fi
	else
		printf 'unavailable\n'
	fi
}

disk_free_value() {
	if command -v df >/dev/null 2>&1; then
		df -hP / 2>/dev/null | awk 'NR == 2 {print $4 " free on " $6; found=1} END {if (!found) print "unavailable"}'
	else
		printf 'unavailable\n'
	fi
}

load_value() {
	if [ -r /proc/loadavg ]; then
		awk '{print $1 " " $2 " " $3}' /proc/loadavg 2>/dev/null || printf 'unavailable\n'
	else
		value="$(command_one_line uptime)"
		case "$value" in
			*load\ average:*|*load\ averages:*)
				printf '%s\n' "$value"
				;;
			*)
				printf 'unavailable\n'
				;;
		esac
	fi
}

memory_available_value() {
	if [ -r /proc/meminfo ]; then
		awk '
			/^MemAvailable:/ {
				mb = int($2 / 1024)
				printf "%d MB\n", mb
				found = 1
			}
			END { if (!found) print "unavailable" }
		' /proc/meminfo 2>/dev/null || printf 'unavailable\n'
	else
		printf 'unavailable\n'
	fi
}

address_value() {
	if command -v hostname >/dev/null 2>&1; then
		value="$(hostname -I 2>/dev/null | one_line || printf '')"
		if [ -n "$value" ]; then
			printf '%s\n' "$value"
			return
		fi
	fi
	if command -v ip >/dev/null 2>&1; then
		value="$(ip -o -4 addr show scope global 2>/dev/null | awk '{print $4}' | one_line || printf '')"
		if [ -n "$value" ]; then
			printf '%s\n' "$value"
			return
		fi
	fi
	printf 'unavailable\n'
}

print_art() {
	cat <<'EOF'
       _.-._
    .-'  |  '-.        Relic Welcome Scroll
   /  .--+--.  \       GPi Case 2 Field Lantern
   |  |  *  |  |
   \  '--+--'  /       SSH support map
    '-.  |  .-'
       '-'
EOF
}

HOSTNAME_VALUE="$(hostname_value)"
USER_VALUE="$(user_value)"
UPTIME_VALUE="$(uptime_value)"
KERNEL_VALUE="$(kernel_value)"
TEMP_VALUE="$(temperature_value)"
THROTTLED_VALUE="$(throttled_value)"
DISK_VALUE="$(disk_free_value)"
LOAD_VALUE="$(load_value)"
MEMORY_VALUE="$(memory_available_value)"
ADDRESS_VALUE="$(address_value)"

if [ "$PLAIN" -eq 0 ] && [ "$SHOW_ART" -eq 1 ]; then
	printf '%s%s%s\n' "$C_CYAN" "$(print_art)" "$C_RESET"
else
	printf 'Relic Welcome Scroll Preview\n'
fi

printf '%s%s%s\n' "$C_BOLD" "GPi Case 2 SSH field map" "$C_RESET"
printf '%s\n' "${C_DIM}Manual preview only; not installed into SSH login.${C_RESET}"
printf '\n'

field "Relic" "$HOSTNAME_VALUE"
field "SSH target" "retropi@gpi"
field "User" "$USER_VALUE"
field "Uptime" "$UPTIME_VALUE"
field "Kernel" "$KERNEL_VALUE"
field "Disk free" "$DISK_VALUE"
field "Load avg" "$LOAD_VALUE"
field "Mem avail" "$MEMORY_VALUE"

if [ "$COMPACT" -eq 0 ]; then
	field "Temp" "$TEMP_VALUE"
	field "Throttled" "$THROTTLED_VALUE"
	field "Address" "$ADDRESS_VALUE"
fi

printf '\n'
printf '%s%s%s\n' "$C_GREEN" "Field Lantern reminder: this scroll is read-only and changes nothing." "$C_RESET"
printf '%s%s%s\n' "$C_YELLOW" "Keep it manually invoked until a later recovery-first wiring quest." "$C_RESET"
