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

PACK=$1

readarray -t packed < <(jq --compact-output '.[]' "$PACK")

for item in "${packed[@]}"; do
    command=$(jq '.command' <<<"$item")
    readarray -t packages < <(jq --compact-output '.packages' <<<"$item")

    for package in "${packages[@]}"; do
        for pack in "${package[@]}"; do
            echo $package
            echo $pack
        done
    done
done
