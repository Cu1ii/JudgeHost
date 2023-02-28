# JudgeHost


[![](https://img.shields.io/badge/Version-0.2.0-blue)](https://github.com/Cu1ii/JudgeHost) ![](https://img.shields.io/badge/go-1.19.3-brightgreen?logo=go)

这是基于 Go 的 Online Judge 平台的**判题服务器模块**, , 在 Linux 下运行

本程序基于调用 QingdaoU / OnlineJudge 平台提供的**判题核心**来对题目结果进行判断, 即该判题服务器（JudgeHost）负责接收用户的提交并将代码**编译、运行、比较，并返回判断情况**其中，代码运行的核心被单独分离在这个仓库  [Judger Sandbox(Seccomp)](https://github.com/QingdaoU/Judger)

提供 **Docker 环境**作为运行环境

### 快速上手

```shell
git clone https://github.com/Cu1ii/JudgeHost.git
```
**使用 docker**

使用 `grpc` 服务来实现对该判题服务器的调用


**手动构建镜像**
```shell
docker build . -t your-name
docker run -d -p port:8000 -v your-volumn:/home/cu1/XOJ --name your-name your-image-name
```
**使用默认提供镜像**
```shell
docker run -d -p port:8000 -v your-volumn:/home/cu1/XOJ --name your-name cu1ii/judge-host:0.2.0
```
### 测试

```go
package test

import (
"context"
"fmt"
"github.com/sirupsen/logrus"
"google.golang.org/grpc"
"testing"
)

func TestJudgeService(t *testing.T) {

	logrus.Info("the judge server begin to run")

	conn, err := grpc.Dial("your-host:your-port", grpc.WithInsecure())
	if err != nil {
		logrus.Fatal(err)
	}
	defer conn.Close()
	client := NewJudgeServiceClient(conn)
	reply, err := client.Judge(context.Background(), &JudgeRequest{ProblemId: 1,
		SubmissionId:    1,
		SubmissionCode:  "#include <iostream>\n int main() { \n std::cout << \"hello world\" << std::endl; \n return 0; \n }",
		ResolutionPath:  "",
		TimeLimit:       1000,
		MemoryLimit:     64,
		OutputLimit:     0,
		Language:        "C++",
		JudgePreference: 0,
		Spj:             false,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Println(*reply)
}
```

### 版本日志

最新版本 `v0.2.0`

### 其他

关于判题核心请移步 [Judger Sandbox(Seccomp)](https://github.com/QingdaoU/Judger)


### 待实现
- [x] 实现 SPJ
- [ ] 实现远程拉取判题数据

### Contributions

欢迎大家提 Issue, PR

### *License*

```
MIT License

Copyright (c) 2022 cu1

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
