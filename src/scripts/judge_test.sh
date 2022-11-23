#!/bin/sh

curl -H "Accept: application/json" -H "Content-type: application/json" -v -X POST "http://localhost:8000/judge/result" -d '{
  "submission_code": "#include <iostream>\n int main() { \n std::cout << \"11hello world\" << std::endl; \n return 0; \n }",
  "real_time_limit": 1000,
  "cpu_time_limit": 3000,
  "memory_limit": 65536,
  "output_limit": 3000,
  "language": "C_PLUS_PLUS",
  "judge_preference": "ACM",
  "solutions": [
    {
      "std_in": "/home/cu1/judgeEnvironment/resolutions/1/hello.in",
      "expected_std_out": "/home/cu1/judgeEnvironment/resolutions/1/hello.out"
    }
  ]
}'