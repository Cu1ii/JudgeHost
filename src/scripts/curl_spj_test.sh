#!/bin/sh

curl -H "Accept: application/json" -H "Content-type: application/json" -v -X POST "http://localhost:8000/judge/result" -d '{
  "submission_code": "#include <iostream>\n#include <stdio.h>\nint a[1009][1009];\nint main()\n{\nusing namespace std;\nint m, n, s;\nlong s1 = 0,s2 = 0;\nscanf(\"%d %d\", &m, &n);\nint b[2];\nint i, j;\nfor (i = 0; i < m; i++)\nfor (j = 0; j < n; j++)\nscanf(\"%d\", &a[i][j]), s1 += a[i][j];\nfor (i = 0; i < m; i++) {\ns = 0;\nfor (j = 0; j < n; j++) {\nif (a[i][j] % 2 != 0) {\nb[s] = j;\ns++;\n}\nif (s == 2) {\na[i][b[0]]--;\na[i][b[1]]++;\ns = 0;\n}\n}\nif (s == 1 && i != m - 1) {\na[i][b[0]]--;\na[i + 1][n - 1]++;\n}\n}\nint sum = 0;\nfor (i = 0; i < m; i++)\nfor (j = 0; j < n; j++){\nif (a[i][j] % 2 == 0) {\nsum++;\n}\ns2 += a[i][j];\n}\nprintf(\"%d\\n\", sum);\nfor (i = 0; i < m; i++) {\nfor (j = 0; j < n - 1; j++)\nprintf(\"%d \", a[i][j]);\nprintf(\"%d\", a[i][n - 1]); \nprintf(\"\\n\"); \n } \n}",
  "real_time_limit": 1000,
  "cpu_time_limit": 3000,
  "memory_limit": 65536,
  "output_limit": 3000,
  "language": "C_PLUS_PLUS",
  "judge_preference": "ACM",
  "spj": true,
  "solutions": [
    {
      "std_in": "/home/cu1/judgeEnvironment/resolutions/2/test1.in",
      "expected_std_out": "/home/cu1/judgeEnvironment/resolutions/2/test1.out"
    },
    {
      "std_in": "/home/cu1/judgeEnvironment/resolutions/2/test2.in",
      "expected_std_out": "/home/cu1/judgeEnvironment/resolutions/2/test2.out"
    },
    {
      "std_in": "/home/cu1/judgeEnvironment/resolutions/2/test3.in",
      "expected_std_out": "/home/cu1/judgeEnvironment/resolutions/2/test3.out"
    },
    {
      "std_in": "/home/cu1/judgeEnvironment/resolutions/2/test4.in",
      "expected_std_out": "/home/cu1/judgeEnvironment/resolutions/2/test4.out"
    },
    {
      "std_in": "/home/cu1/judgeEnvironment/resolutions/2/test5.in",
      "expected_std_out": "/home/cu1/judgeEnvironment/resolutions/2/test5.out"
    },
    {
      "std_in": "/home/cu1/judgeEnvironment/resolutions/2/test6.in",
      "expected_std_out": "/home/cu1/judgeEnvironment/resolutions/2/test6.out"
    },
    {
      "std_in": "/home/cu1/judgeEnvironment/resolutions/2/test7.in",
      "expected_std_out": "/home/cu1/judgeEnvironment/resolutions/2/test7.out"
    },
    {
      "std_in": "/home/cu1/judgeEnvironment/resolutions/2/test8.in",
      "expected_std_out": "/home/cu1/judgeEnvironment/resolutions/2/test8.out"
    },
    {
      "std_in": "/home/cu1/judgeEnvironment/resolutions/2/test9.in",
      "expected_std_out": "/home/cu1/judgeEnvironment/resolutions/2/test9.out"
    },
    {
      "std_in": "/home/cu1/judgeEnvironment/resolutions/2/test10.in",
      "expected_std_out": "/home/cu1/judgeEnvironment/resolutions/2/test10.out"
    }
  ]
}'

