# Gopssh

[TOC]

Gopssh (go parallel ssh tool) 是一款由 Go 语言编写，高效且易上手的 ssh 工具，它可以通过 ssh 连接来批量地执行命令、执行脚本、上传文件。
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
    hosts:
      mysql_master:
        ips: 
        - 172.16.8.83
        # username: root
        password: cm9vdCQxMjM= # echo -n 'root$123' | base64
      mysql_slaves:
        ips: 
        - 172.16.8.84
        - 172.16.8.85
        - 172.16.8.86
        username: mysql
        password: bXlzcWwkMTIz # gopssh base64 -c 'mysql$123'
    ```

- 完整的配置文件

    ```yaml
    # The global setting will take effect for all hosts,
    #   if he does not set the corresponding value
    # The 'global' is not necessary
    global:
      # If 'global.port' is not set, the default port is 22
      port: 23
      # If 'global.port' is not set, the default username is root
      username: root
      # The 'global.password' have not default value, 
      #   please set 'global.password' or set 'password' for each host
      password: cm9vdCQxMjM=
      # The 'global.labels' will add for all hosts
      # All host always have label all=all,
      #   and you can override all=all by setting 'global.labels'
      labels: 
        # all: all
        app: my_host
    hosts:
      # Execute command 'ls -l /' for mysql_master 
      #   gopssh execute -h mysql_master 'ls -l /'
      mysql_master:
        # The 'hosts.<host_alias>.ip' is necessary
        ips: 
        - 172.16.8.82
        # Overwrite 'global.port' 23
        port: 30
        # * Overwrite 'global.username' root
        username: mysql
        # Overwrite 'global.password' cm9vdCQxMjM=
        # Please enter user password with base64
        #   Encode password with base64 with command:
        #     echo -n 'mysql$123' | base64
        #     gopssh base64 'mysql$123'
        #   Decode base64 password with command:
        #     echo -n 'bXlzcWwkMTIz' | base64 -d
        #     gopssh base64 -d 'bXlzcWwkMTIz'
        password: bXlzcWwkMTIz
        labels:
          # Execute command 'ls -l /' for all host that have label app=mysql
          #   gopssh execute -l app=mysql 'ls -l /'
          app: mysql
          mysql: master
          # app: my_host
      # The sample without comment
      mysql_slaves:
        ips: 
        - 172.16.8.84
        - 172.16.8.85
        port: 22
        username: mysql
        password: bXlzcWwkMTIz
        labels:
          app: mysql
          mysql: slave
      # The sample with 'global'
      nginx:
        ip: 172.16.8.87
        # port: 23
        # username: root
        # password: cm9vdCQxMjM=
        # labels:
        #   all: all
    ```