package xoj

import "fmt"

func compileError(id int, problem, msg string) {
	fmt.Println(id, " ", problem, " ", msg)
}

func doneProblem(id int, problem, message string, memory, mytime int, username string, contest, result int, testcase string) {
	fmt.Println(id, " ", problem, " ", message, " ", mytime, " ", username, " ", contest, " ", result, " ", testcase)
}
