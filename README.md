# JudgeHost 用户文档

这是 Online Judge 平台的**判题服务器模块**, 基于 YuJudge-Online-Judge JudgeHost的 Go 的重构版本, 在 Linux 下运行

本程序基于调用 YuJudge-Online-Judge 平台提供的**判题核心**来对题目结果进行判断, 即该判题服务器（JudgeHost）负责接收用户的提交并将代码**编译、运行、比较，并返回判断情况**其中，代码运行的核心被单独分离在这个仓库  [Yu-Judge-Core](https://github.com/yuzhanglong/YuJudge-Core), 考虑到判题的速度、短时间内可能需要大量判题，JudgeHost可能需要考虑多线程、集群相关

后续会提供 **Docker 环境**作为运行环境

### 快速上手

```shell
git clone https://github.com/Cu1ii/JudgeHost.git
```

### 使用了

- [Gin](https://github.com/gin-gonic/gin)
- [logrus](https://github.com/sirupsen/logrus)
- [ants](https://github.com/panjf2000/ants)
- [viper](https://github.com/spf13/viper)
- [Yu-Judge-Core](https://github.com/yuzhanglong/YuJudge-Core)

### 版本日志

最新版本 `0.1`

### 其他

关于判题核心请移步 [Yu-Judge-Core](https://github.com/yuzhanglong/YuJudge-Core)

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

