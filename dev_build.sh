#!/usr/bin/env bash
set -x -e
if test -f "dead-drop"; then
    rm dead-drop
fi

go build -ldflags "-w -s" -o dead-drop .
./dead-drop