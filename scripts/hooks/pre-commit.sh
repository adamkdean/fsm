#!/bin/bash
#
#  _____ ____  __  __
# |  ___/ ___||  \/  |
# | |_  \___ \| |\/| |
# |  _|  ___) | |  | |
# |_|   |____/|_|  |_|
#
# Finite State Machine
# (c) 2018 Adam K Dean

#
# First we want to format any modified file
#
for FILE in $(git diff --cached --name-only --diff-filter=d | grep "\.go")
do
  SRC=$FILE make fmt
  git add "$FILE"
done

#
# Next we want to lint
#
make lint
