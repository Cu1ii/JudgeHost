#!/bin/sh

SUBMISSION_PATH=$1
ID=$2

cd "$SUBMISSION_PATH"

rm -rf "$ID".out "$ID"ce.txt

set -e

timeout 10 go build -o "$ID".o Main.go 2>"$ID"ce.txt