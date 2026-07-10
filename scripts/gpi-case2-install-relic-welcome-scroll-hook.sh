#!/bin/sh
set -eu

SCRIPT_NAME="$(basename "$0")"
ACTION=''
PLAIN=0
DRY_RUN=0

HOOK_FILE_NAME=".gpi-relic-welcome-scroll-hook.sh"
DISABLED_FILE_NAME=".gpi-relic-welcome-scroll.disabled"
SCROLL_SCRIPT_NAME="gpi-case2-relic-welcome-scroll.sh"
BASHRC_NAME=".bashrc"
BLOCK_BEGIN="# >>> gpi-case2 relic welcome scroll hook >>>"
BLOCK_END="# <<< gpi-case2 relic welcome scroll hook <<<"

usage() {
	cat <<'EOF'
GPi Case 2 Relic Welcome Scroll Login Hook Installer

Recovery-first installer for wiring the Relic Welcome Scroll into interactive
SSH login. It is disabled by default: nothing is installed unless --install is
explicitly provided by the operator.

Safety boundaries:
  - User-level hook only, installed under $HOME.
  - Adds one small guarded source block to $HOME/.bashrc.
  - Creates backups before modifying $HOME/.bashrc.
  - Does not edit sshd config, PAM, MOTD, services, boot config, GPIO,
    display, shutdown, sleep, or resume behavior.
  - The hook stays quiet for non-interactive SSH, scp, and automation.

Options:
  --install    Install or refresh the guarded login hook.
  --uninstall  Remove the .bashrc source block and move the hook aside.
  --status     Print current hook status without changing files.
  --dry-run    Show intended changes without changing files.
  --plain      Disable color and decorative output.
  --help       Show this help.

Disable switches after install:
  touch $HOME/.gpi-relic-welcome-scroll.disabled
  GPI_RELIC_WELCOME_SCROLL_DISABLED=1 ssh retropi@gpi

scp-first field trail:
  scp scripts/gpi-case2-relic-welcome-scroll.sh retropi@gpi:/home/retropi/
  scp scripts/gpi-case2-install-relic-welcome-scroll-hook.sh retropi@gpi:/home/retropi/
  ssh retropi@gpi 'sh /home/retropi/gpi-case2-install-relic-welcome-scroll-hook.sh --status'
  ssh retropi@gpi 'sh /home/retropi/gpi-case2-install-relic-welcome-scroll-hook.sh --install'
EOF
}

while [ "$#" -gt 0 ]; do
	case "$1" in
		--help|-h)
			usage
			exit 0
			;;
		--install|--uninstall|--status)
			if [ -n "$ACTION" ]; then
				printf 'ERROR: Choose only one action: --install, --uninstall, or --status.\n' >&2
				exit 2
			fi
			ACTION="${1#--}"
			shift
			;;
		--dry-run)
			DRY_RUN=1
			shift
			;;
		--plain)
			PLAIN=1
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

if [ -z "$ACTION" ]; then
	printf 'ERROR: No action selected. Use --status, --install, or --uninstall.\n' >&2
	printf 'Action: Run %s --help for the recovery map.\n' "$SCRIPT_NAME" >&2
	exit 2
fi

USE_COLOR=0
if [ "$PLAIN" -eq 0 ] && [ -t 1 ]; then
	USE_COLOR=1
fi

if [ "$USE_COLOR" -eq 1 ]; then
	C_RESET="$(printf '\033[0m')"
	C_BOLD="$(printf '\033[1m')"
	C_GREEN="$(printf '\033[32m')"
	C_YELLOW="$(printf '\033[33m')"
else
	C_RESET=''
	C_BOLD=''
	C_GREEN=''
	C_YELLOW=''
fi

HOME_DIR="${HOME:-/home/retropi}"
HOOK_FILE="$HOME_DIR/$HOOK_FILE_NAME"
DISABLED_FILE="$HOME_DIR/$DISABLED_FILE_NAME"
SCROLL_SCRIPT="$HOME_DIR/$SCROLL_SCRIPT_NAME"
BASHRC_FILE="$HOME_DIR/$BASHRC_NAME"

say() {
	printf '%s\n' "$1"
}

status_line() {
	printf '  %-16s %s\n' "$1:" "$2"
}

timestamp() {
	date +%Y%m%d-%H%M%S 2>/dev/null || date -u +%Y%m%d-%H%M%S 2>/dev/null || printf 'unknown-time'
}

contains_hook_block() {
	[ -f "$BASHRC_FILE" ] || return 1
	grep -F "$BLOCK_BEGIN" "$BASHRC_FILE" >/dev/null 2>&1
}

backup_bashrc() {
	[ -f "$BASHRC_FILE" ] || return 0
	backup="$BASHRC_FILE.relic-welcome-scroll.bak-$(timestamp)"
	if [ -e "$backup" ]; then
		n=1
		while [ -e "$backup.$n" ]; do
			n=$((n + 1))
		done
		backup="$backup.$n"
	fi
	if [ "$DRY_RUN" -eq 1 ]; then
		say "DRY-RUN: would back up $BASHRC_FILE to $backup"
	else
		cp "$BASHRC_FILE" "$backup"
		say "Backup written: $backup"
	fi
}

write_hook_file() {
	tmp="$HOOK_FILE.tmp.$$"
	if [ "$DRY_RUN" -eq 1 ]; then
		say "DRY-RUN: would write guarded hook $HOOK_FILE"
		return 0
	fi
	cat >"$tmp" <<'HOOK_EOF'
# GPi Case 2 Relic Welcome Scroll guarded SSH login hook.
# Installed by gpi-case2-install-relic-welcome-scroll-hook.sh.

__gpi_relic_welcome_scroll() {
	case "$-" in
		*i*) ;;
		*) return 0 ;;
	esac

	[ -t 1 ] || return 0
	[ -z "${SSH_ORIGINAL_COMMAND:-}" ] || return 0
	[ -z "${GPI_RELIC_WELCOME_SCROLL_DISABLED:-}" ] || return 0
	[ ! -e "$HOME/.gpi-relic-welcome-scroll.disabled" ] || return 0
	[ -x "$HOME/gpi-case2-relic-welcome-scroll.sh" ] || return 0

	if command -v timeout >/dev/null 2>&1; then
		timeout 2s "$HOME/gpi-case2-relic-welcome-scroll.sh" 2>/dev/null || true
	else
		"$HOME/gpi-case2-relic-welcome-scroll.sh" 2>/dev/null || true
	fi
}

__gpi_relic_welcome_scroll
unset -f __gpi_relic_welcome_scroll 2>/dev/null || true
HOOK_EOF
	chmod 0644 "$tmp"
	mv "$tmp" "$HOOK_FILE"
	say "Hook written: $HOOK_FILE"
}

append_bashrc_block() {
	if contains_hook_block; then
		say "Status: .bashrc already has the Relic Welcome Scroll hook block."
		return 0
	fi
	backup_bashrc
	if [ "$DRY_RUN" -eq 1 ]; then
		say "DRY-RUN: would append guarded source block to $BASHRC_FILE"
		return 0
	fi
	{
		printf '\n%s\n' "$BLOCK_BEGIN"
		printf '[ -r "$HOME/%s" ] && . "$HOME/%s"\n' "$HOOK_FILE_NAME" "$HOOK_FILE_NAME"
		printf '%s\n' "$BLOCK_END"
	} >>"$BASHRC_FILE"
	say "Source block added: $BASHRC_FILE"
}

remove_bashrc_block() {
	if ! contains_hook_block; then
		say "Status: no Relic Welcome Scroll hook block found in $BASHRC_FILE."
		return 0
	fi
	backup_bashrc
	if [ "$DRY_RUN" -eq 1 ]; then
		say "DRY-RUN: would remove guarded source block from $BASHRC_FILE"
		return 0
	fi
	tmp="$BASHRC_FILE.tmp.$$"
	awk -v begin="$BLOCK_BEGIN" -v end="$BLOCK_END" '
		$0 == begin {skip = 1; next}
		$0 == end {skip = 0; next}
		skip != 1 {print}
	' "$BASHRC_FILE" >"$tmp"
	mv "$tmp" "$BASHRC_FILE"
	say "Source block removed: $BASHRC_FILE"
}

ensure_scroll_ready() {
	if [ ! -f "$SCROLL_SCRIPT" ]; then
		printf 'ERROR: Welcome Scroll script missing: %s\n' "$SCROLL_SCRIPT" >&2
		printf 'Action: Copy %s to %s before --install.\n' "$SCROLL_SCRIPT_NAME" "$HOME_DIR" >&2
		exit 1
	fi
	if [ ! -x "$SCROLL_SCRIPT" ]; then
		if [ "$DRY_RUN" -eq 1 ]; then
			say "DRY-RUN: would chmod +x $SCROLL_SCRIPT"
		else
			chmod +x "$SCROLL_SCRIPT"
			say "Scroll made executable: $SCROLL_SCRIPT"
		fi
	fi
}

print_status() {
	say "${C_BOLD}Relic Welcome Scroll hook status${C_RESET}"
	status_line "Home" "$HOME_DIR"
	if [ -f "$SCROLL_SCRIPT" ]; then
		if [ -x "$SCROLL_SCRIPT" ]; then
			status_line "Scroll" "ready at $SCROLL_SCRIPT"
		else
			status_line "Scroll" "present but not executable at $SCROLL_SCRIPT"
		fi
	else
		status_line "Scroll" "missing at $SCROLL_SCRIPT"
	fi
	if [ -f "$HOOK_FILE" ]; then
		status_line "Hook" "present at $HOOK_FILE"
	else
		status_line "Hook" "not installed"
	fi
	if contains_hook_block; then
		status_line ".bashrc" "source block present"
	else
		status_line ".bashrc" "source block absent"
	fi
	if [ -e "$DISABLED_FILE" ]; then
		status_line "Disable rune" "present at $DISABLED_FILE"
	else
		status_line "Disable rune" "absent"
	fi
	status_line "Boundaries" "no sshd, PAM, service, boot, GPIO, display, power, sleep, or resume changes"
	if [ "$DRY_RUN" -eq 1 ]; then
		status_line "Mode" "dry-run"
	fi
}

install_hook() {
	ensure_scroll_ready
	write_hook_file
	append_bashrc_block
	say "${C_GREEN}Installed: interactive SSH logins may show the Relic Welcome Scroll.${C_RESET}"
	say "Recovery: touch $DISABLED_FILE or run $SCRIPT_NAME --uninstall."
}

uninstall_hook() {
	remove_bashrc_block
	if [ -f "$HOOK_FILE" ]; then
		disabled="$HOOK_FILE.uninstalled-$(timestamp)"
		if [ -e "$disabled" ]; then
			n=1
			while [ -e "$disabled.$n" ]; do
				n=$((n + 1))
			done
			disabled="$disabled.$n"
		fi
		if [ "$DRY_RUN" -eq 1 ]; then
			say "DRY-RUN: would move $HOOK_FILE to $disabled"
		else
			mv "$HOOK_FILE" "$disabled"
			say "Hook moved aside: $disabled"
		fi
	else
		say "Status: no hook file found at $HOOK_FILE."
	fi
	say "${C_YELLOW}Uninstalled: the Relic Welcome Scroll hook is no longer sourced by .bashrc.${C_RESET}"
}

case "$ACTION" in
	status)
		print_status
		;;
	install)
		install_hook
		;;
	uninstall)
		uninstall_hook
		;;
esac
