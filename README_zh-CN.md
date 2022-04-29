# Gopssh

[TOC]

Gopssh (go parallel ssh tool) 是一款由 Go 语言编写，高效且易上手的 ssh 工具，它可以通过 ssh 连接来批量地执行命令、检查连通性、上传文件。
该项目灵感来自于 [pypssh](https://github.com/souloss/pypssh)，在这里感谢它的作者 [souloss](https://github.com/souloss)。

- [中文](./README_zh-CN.md)
- [English](./README.md)

## 安装

- 下载预构置的二进制文件

    可以在 GitHub 上该仓库的 [Release](TODO) 找到最新的预构置二进制文件

- 手动编译安装

    在拉取仓库之后可以通过执行仓库目录下的脚本文件 `./hack/build.sh` 来构建二进制文件

    ```bash
    git clone https://github.com/XIniris/gopssh
    cd gopssh
    ./build.sh
    ```

    也可以手动通过 `go build` 来构建二进制文件

    ```bash
    git clone https://github.com/XIniris/gopssh
    cd gopssh
    go build -o gopssh
    ```

- 检查

    在获取到二进制文件之后，可以通过 `version` 来检查程序是否能够正常运行

    ```bash
    gopssh version
    ```

## 用法


## 配置

- 最简单的配置文件

    ```yaml
   
    ```

- 完整的配置文件

    ```yaml
   
    ```