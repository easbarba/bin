#!/usr/bin/env bash
#
# Bin is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# Bin is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with Bin. If not, see <https://www.gnu.org/licenses/>.
#

# Debug Options
set -euo pipefail

STEP=3
STATE="$1"

# TRANSLATE HUMAN-FRIENDLY STATE TO VOLUME MANAGERS COMMAND
case $STATE in
    up)
        STATE='+'
        ;;
    down)
        STATE='-'
        ;;
    toggle) ;;

esac

usage() {
    printf %s "
Usage:
        up:     raise volume
      down:     decrease volume
    toggle:     mute or power on volume
"
    exit 0
}

# VOLUME MANAGERS

_pactl() {
    case $STATE in
        toggle)
            pactl set-sink-mute @DEFAULT_SINK@ toggle
            ;;
        + | -)
            pactl set-sink-volume @DEFAULT_SINK@ "$STATE$STEP%"
            ;;
        *) exit 1 ;;
    esac
}

# RUN
[[ $# -eq 0 ]] && usage

[[ -x $(command -v pactl) ]] && _pactl
