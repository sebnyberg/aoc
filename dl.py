#!/usr/bin/env python3

from aocd import get_data

import os
import re
import sys

path = os.getcwd()

yearm = re.search('20\\d{2}', path)
if not yearm:
    print("could not find '20xx' in current path")
    sys.exit(1)
year = yearm.group()

daym = re.search('day(\\d{2})', path)
if not daym:
    print("could not find 'dayXX' in current path")
    sys.exit(1)
day = int(daym.group(1))

with open("input", "w") as f:
    f.write(get_data(day=day, year=year))
