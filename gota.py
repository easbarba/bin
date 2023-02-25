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

json_file = sys.argv[1]

with open(json_file) as json_data:
    data = json.load(json_data)

print("\n", data["lang"], "\n")

# TODO: check deps

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
