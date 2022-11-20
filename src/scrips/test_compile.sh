#!/bin/sh


SUBMISSION_PATH=$1;

mkdir "$SUBMISSION_PATH"

cd "$SUBMISSION_PATH"

touch compile.out
touch compile.err
touch test.out
touch exp.out

echo "hello world" > test.out
echo "hello world" > exp.out

echo "{
          "realTimeCost": 0,
          "cpuTimeCost": 0,
          "memoryCost": 0,
          "condition": 1,
          "stdinPath": "/home/cu1/judgeEnvironment/resolutions/1/hello.in",
          "stdoutPath": "$SUBMISSION_PATH/test.out",
          "stderrPath": "",
          "loggerPath": ""
      } " > compile.err