package dto

/**
 * 解决方案数据传输对象
 * 用户需要传入一个或多个解决方案，每一个解决方案包括：
 * 1. 输入文件 stdin --- stdIn
 * 2. 输出文件 stdout --- expectedStdOut
 * 例如：
 * {
 * "stdIn": "your_download_path",
 * "expectedStdOut":"your_download_path"
 * }
 */

type SolutionDTO struct {
	StdIn          string
	ExpectedStdOut string
}
