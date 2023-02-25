#!/usr/bin/env python3

# Packer is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# Packer is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with Packer. If not, see <https://www.gnu.org/licenses/>.

# TODO: log commands

import sys
import os
import json
import argparse
from pathlib import Path

gota_home = Path.home().joinpath(".config", "gota")

# TODO: check deps
def pack(data: dict) -> None:
    print("\n", data["lang"], "\n")

    # install all packages
    for key, value in data["packages"].items():
        for i in value:
            print("\n", i, "\n")
            os.system(f"{data['command']} i")

    # run post command
    if "post" in data.keys():
        os.system(data['post'])
    else:
        print("post command not provided")

def run(where: str, func) -> None:
    for it in gota_home.joinpath(where).iterdir():
        with open(it) as json_data:
            data = json.load(json_data)

        func(data)

# PARSE CLI
parser = argparse.ArgumentParser(
                    prog = 'Gota',
                    description = 'Install [every/any]thing',
                    epilog = 'Praise the sun!')
parser.add_argument('-p', '--pack',  action='store_true')
args = parser.parse_args()

if args.pack:
    run("pack", pack)

if args.distro:
    run("distro", distro)

exit()
