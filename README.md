# About Go-Notes

Go-Notes 项目为golang源码阅读笔记。目的是在go开源项目源码研读过程中根据作者理解做些记录，同时方便其他人研读著名golang开源项目源代码。
另外欢迎大家参与该记录和解读工作，如果有意向参与，可以发邮件至sioomy@qq.com或者直接搜索笔者微信，微信号:sioomy

# How to Debug（调试方法）

笔者通过goland进行调试，步骤如下：
```
1.配置环境变量，goland(2020.1)中打开Goland->Preferences->Go->GOROOT
配置新的goroot环境变量地址为项目根目录地址如:/Users/sioomy/work/golang/go

2.配置debug命令，右键单机项目目录中./src/cmd/go/main.go
选择Create 'go build main.go'
在弹出窗口中配置参数：
Run kind:Directory
Directory:/Users/sioomy/work/golang/go/src/cmd/go
Working directory:/Users/sioomy/work/golang/go/src
Program arguments:build ../test/varinit.go

3.运行，点击右上角的小虫子按钮，或右键单机main.go，选择Debug 'go build main.go'
```
根据以上步骤即可调试go主程序，上述事例中为通过build命令编译一个简单的golang代码../test/varinit.go

# The Go Programming Language

Go is an open source programming language that makes it easy to build simple,
reliable, and efficient software.

![Gopher image](doc/gopher/fiveyears.jpg)
*Gopher image by [Renee French][rf], licensed under [Creative Commons 3.0 Attributions license][cc3-by].*

Our canonical Git repository is located at https://go.googlesource.com/go.
There is a mirror of the repository at https://github.com/golang/go.

Unless otherwise noted, the Go source files are distributed under the
BSD-style license found in the LICENSE file.

### Download and Install

#### Binary Distributions

Official binary distributions are available at https://golang.org/dl/.

After downloading a binary release, visit https://golang.org/doc/install
or load [doc/install.html](./doc/install.html) in your web browser for installation
instructions.

#### Install From Source

If a binary distribution is not available for your combination of
operating system and architecture, visit
https://golang.org/doc/install/source or load [doc/install-source.html](./doc/install-source.html)
in your web browser for source installation instructions.

### Contributing

Go is the work of thousands of contributors. We appreciate your help!

To contribute, please read the contribution guidelines:
	https://golang.org/doc/contribute.html

Note that the Go project uses the issue tracker for bug reports and
proposals only. See https://golang.org/wiki/Questions for a list of
places to ask questions about the Go language.

[rf]: https://reneefrench.blogspot.com/
[cc3-by]: https://creativecommons.org/licenses/by/3.0/
