#!/bin/sh

curl -H "Accept: application/json" -H "Content-type: application/json" -v -X POST "http://47.100.227.175:8000/judge/result" -d '{
  "problem_id": 1,
  "submission_id": 1,
  "submission_code": "#include <iostream>\n int main() { \n std::cout << \"11hello world\" << std::endl; \n return 0; \n }",
  "resolution_path": "",
  "cpu_time_limit": 1000,
  "memory_limit": 64,
  "output_limit": 0,
  "language": "C++",
  "judge_preference": 0,
  "spj":  false
}'