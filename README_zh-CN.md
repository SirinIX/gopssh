# Gopssh

[TOC]

Gopssh (go parallel ssh tool) 是一款由 Go 语言编写，高效且易上手的 ssh 工具，它可以通过 ssh 连接来批量地执行命令、检查连通性、上传文件。
该项目灵感来自于 [pypssh](https://github.com/souloss/pypssh)，在这里感谢它的作者 [souloss](https://github.com/souloss)。

- [中文](./README_zh-CN.md)
- [English](./README.md)

## 简介

Gopssh 可以批量地为不同标签的主机执行命令、上传文件。

你可以将主机信息添加到默认的配置文件 `~/.gopssh/inventory.yaml` 中，并为它们添加各种标签，通过标签来指定不同的主机，批量地为这些主机执行各种操作。

一个简单的配置文件如下，可以在 labels 中指定各种标签，需要注意的是其中 password 需要通过 base64 编码。

```yaml
groups:
  - ips:
    - 192.168.8.83
    username: root
    password: QWJjIUAjMTM1
    labels:
      mysql: "master"
      middleware: "mysql"
  - ips:
    - 192.168.8.84
    - 192.168.8.85
    username: root
    password: QWJjIUAjMTM1
    labels:
      mysql: "slave"
      middleware: "mysql"
```

以上面的配置文件为例，如果要为所有有标签 `mysql=master` 的主机执行命令，可以执行

```bash
gopssh execute -l mysql=master -c 'ls -l /'
```

通过 `-l` 指定标签，支持 `<key>=<val>` 和 `<key>!=<val>` 多个标签通过 `,` 分割。

通过 `-c` 指定需要执行的命令，如果需要查看更多地使用方法可以通过 `gopssh --help`。


## 安装

- 下载预构置的二进制文件

    可以在 GitHub 上该仓库的 [Release](TODO) 找到最新的预构置二进制文件。

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

    在获取到二进制文件之后，可以通过 `version` 来检查程序是否能够正常运行。

    ```bash
    gopssh version
    ```

## 用法

### execute

`execute` 用于执行命令，并返回执行结果（stdout 和 stderr）。

目前支持选项有：

- `-c`
  
    指定需要执行的命令

- `-f`

    指定使用的配置文件，默认为 `$HOME/.gopssh/inventory.yaml`。

- `-l`

    指定需要执行命令的主机带有的标签，支持 `=` 和 `!=`，多个标签通过 `,` 分割， 如果不指定则选定所有主机。

- `-n`

    是否使用缓存，默认使用，添加 `-n` 重新解析配置文件。

通过 `gopssh execute -h` 可以查看详细的使用方法。

```bash
Execute command and return result

Usage:
  root execute [flags]

Examples:
  Simple:                 gopssh execute -c 'ls -l'
  Specify config:         gopssh execute -c 'ls -l' -f /sample.yaml
  Select host to execute: gopssh execute -c 'ls -l' -l app=mysql
  Execute without cache:  gopssh execute -c 'ls -l' -n

Flags:
  -c, --command string       command to execute
  -f, --config-file string   config file path
  -h, --help                 help for execute
  -l, --labels string        label to filter on, default select all host, supports '=', and '!=' (e.g. -l key1=value1,key2!=value2
  -n, --without-cache        not use cache, default use cache
```

### upload

`upload` 用于上传本地文件到指定的主机。

主要关心选项，其余选项和 `execute` 中的类似：

- `-i`

    需要上传的本地文件路径。

- `-o`

    文件需要上传到远程的路径，需要该文件不存在于远程主机。

通过 `gopssh upload -h` 可以查看详细的使用方法。

```bash
Upload file to remote

Usage:
  root upload [flags]

Examples:
 Simple:                 gopssh upload -i sample.txt -o /tmp/upload.txt
  Specify config:         gopssh execute -i sample.txt -o /tmp/upload.txt -f /sample.yaml
  Select host to execute: gopssh execute -i sample.txt -o /tmp/upload.txt -l app=mysql
  Execute without cache:  gopssh execute -i sample.txt -o /tmp/upload.txt -n

Flags:
  -f, --config-file string   config file path
  -h, --help                 help for upload
  -l, --labels string        label to filter on, default select all host, supports '=', and '!=' (e.g. -l key1=value1,key2!=value2
  -o, --output-path string   upload file download path
  -i, --upload-file string   the file to upload
  -n, --without-cache        not use cache, default use cache
```

### check

`check` 用于检查主机的联通性。

通过 `gopssh check -h` 可以查看详细的使用方法。

```bash
Check all IP ports in the configuration file for connectivity

Usage:
  root check [flags]

Examples:
  Simple:               gopssh check
  Specify config:       gopssh check -f config.yaml
  Select host to check: gopssh check -l app=mysql
  Check without cache:  gopssh check -f config.yaml -n

Flags:
  -f, --config-file string   config file path
  -h, --help                 help for check
  -l, --labels string        label to filter on, default select all host, supports '=', and '!=' (e.g. -l key1=value1,key2!=value2
  -n, --without-cache        not use cache, default use cache
```

### get

`get` 用于获取主机的详细信息，包括 Ip、端口、用户名、密码（base64 编码）和 标签。

通过 `gopssh get -h` 可以查看详细的使用方法。

```bash
Get connection instance of config file

Usage:
  root get [flags]

Examples:
  Simple:                 gopssh get
  Specify config:         gopssh get -f /sample.yaml
  Select host to execute: gopssh get -l app=mysql
  Execute without cache:  gopssh get -n

Flags:
  -f, --config-file string   config file path
  -h, --help                 help for get
  -l, --labels string        label to filter on, default select all host, supports '=', and '!=' (e.g. -l key1=value1,key2!=value2
  -n, --without-cache        not use cache, default use cache
```

### base64

`base64` 可以将字符串经过 base64 编码，或者从 base64 解码。

支持选项：

- `-c`

  指定需要编码 / 解码的字符串。

- `-d`

  是否为解码模式，默认为编码模式。

通过 `gopssh base64 -h` 可以查看详细的使用方法。

```bash
Encode or decode content with base64

Usage:
  root base64 [flags]

Examples:
  Encode: gopssh base64 -c 'root$123'
  Decode: gopssh base64 -d -c 'cm9vdA=='

Flags:
  -c, --content string   decode / encode data content (required)
  -d, --decode           decode or encode, default is encode
  -h, --help             help for base64
```

### template

`template` 可以获取默认的配置文件模板。

支持选项：

- `-o`

  默认输出到控制台，也可以通过 `-o` 指定保存到文件。

- `-t`

  指定配置文件模板的格式，支持 `yaml` 和 `json`

通过 `gopssh template -h` 可以查看详细的使用方法。

```bash
Dump config template, yaml or json

Usage:
  root template [flags]

Examples:
  Get yaml config template: gopssh template
  Save template as file:    gopssh template -o sample.yaml
  Get json config template: gopssh template -t json

Flags:
  -h, --help                 help for template
  -o, --output-path string   output file path
  -t, --type string          config template type, yaml or json (default "yaml")
```

### convert

`convert` 能够转换 `yaml` 文件和 `json` 文件。

- `-o`

  默认输出到控制台，也可以通过 `-o` 指定保存到文件。

- `-y`

  将 `yaml` 文件转化为 `json` 文件，指定 `yaml` 文件的路径，需要文件以 `.yaml` 结尾。

- `-j`

  将 `json` 文件转化为 `yaml` 文件，指定 `json` 文件的路径，需要文件以 `.json` 结尾。

通过 `gopssh convert -h` 可以查看详细的使用方法。

```bash
Convert json to yaml, or yaml to json

Usage:
  root convert [flags]

Examples:
  Convert YAML to JSON: gopssh convert -y sample.yaml
  Convert JSON to YAML: gopssh convert -j sample.json
  Convert and save:     gopssh convert -j sample.yaml -o convert.yaml

Flags:
  -h, --help                 help for convert
  -j, --json-path string     json file path
  -o, --output-path string   output file path
  -y, --yaml-path string     yaml file path
```

### version

`version` 可以查看当前 `gopssh` 的版本。

## 配置

配置分为 `global` 和 `groups` 两部分。

### global

`global` 提供全局的默认配置，可以被 `group` 下的对应的选项替代。

`global` 包含 `port`, `username` 和 `password` 三个部分。其中 `port` 有默认值 `22`， `username` 有默认值 `root`。

```yaml
global:
  port: 23
  username: root
  password: cm9vdCQxMjM=
```

### groups

`groups` 是一个列表，其中每一个值都是一个 `group`。

一个 `group` 下的 Ip 共享相同的 `port`、`username`、`password` 和 `labels`。其中对应的值有比 `global` 中更高的优先级。

一个 `group` 的每一项包含 `ips`、`port`、`username`、`password` 和 `labels`。其中 `ips` 为必填，是 ip 的列表，需要保证其中的值是一个合法的 Ip 地址；如果没有设置 `port`、`username`、`password` 则会使用 `global` 中对应的值；`labels`  是该组的标签，它是一个 map，通过指定 label，可以轻松地选定需要的主机。

```yaml
groups:
  - ips:
    - 192.168.8.84
    - 192.168.8.85
    port: 22
    username: mysql
    password: bXlzcWwkMTIz
    labels:
      app: mysql
      mysql: slave
```


### 示例

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
groups:
  # The 'groups..ips' is necessary
  - ips: 
    - 192.168.8.83
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
  # The sample without comment
  - ips: 
    - 192.168.8.84
    - 192.168.8.85
    port: 22
    username: mysql
    password: bXlzcWwkMTIz
    labels:
      app: mysql
      mysql: slave
  # The sample with 'global'
  - ips: 
    - 192.168.8.87
    # port: 23
    # username: root
    # password: cm9vdCQxMjM=
```