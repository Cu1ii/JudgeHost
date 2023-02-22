#!/bin/sh

SUBMISSION_PATH=$1
ID=$2

set -e
cd "$SUBMISSION_PATH"
rm -rf "$ID"ce.txt

javac Main.java -d "$ID" 2>"$ID"ce.txt
