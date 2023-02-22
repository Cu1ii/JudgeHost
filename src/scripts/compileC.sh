#!/bin/sh

SUBMISSION_PATH=$1
ID=$2

cd "$SUBMISSION_PATH"

rm -rf "$ID".out "$ID"ce.txt

set -e

timeout 10 gcc Main.c -fmax-errors=3 -o "$ID".o -O2 -std=c11 2>"$ID"ce.txt