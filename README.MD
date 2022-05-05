# Gopssh

[TOC]

Gopssh (go parallel ssh tool) is an efficient and easy-to-use ssh tool written in Go language. It can execute commands, check connectivity, and upload files in batches through ssh connections.
This project was inspired by [pypssh](https://github.com/souloss/pypssh), and thanks to its author [souloss](https://github.com/souloss).

- [中文](./README_zh-CN.md)
- [English](./README.md)

## Introduction

Gopssh can execute commands and upload files for hosts with different labels in batches.

You can add host information to the default configuration file `~/.gopssh/inventory.yaml`, and add various labels to them, specify different hosts through labels, and perform various operations for these hosts in batches.

A simple configuration file is as follows, you can specify various labels in labels, it should be noted that the password needs to be encoded by base64.

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

Taking the above configuration file as an example, if you want to execute the command for all hosts with the label `mysql=master`, you can execute

```bash
gopssh execute -l mysql=master -c 'ls -l /'
```

Specify tags through `-l`, support `<key>=<val>` and `<key>!=<val>` Multiple tags are separated by `,`.

Use `-c` to specify the command to be executed. If you need to see more usage methods, you can use `gopssh --help`.

## Install

- Download prebuilt binaries

    The latest prebuilt binaries can be found at [Release](TODO) for this repository on GitHub.

- Manual compilation and installation

    After pulling the repository, you can build the binary by executing the script file `./hack/build.sh` in the repository directory.

    ```bash
    git clone https://github.com/XIniris/gopssh
    cd gopssh
    ./build.sh
    ```

    You can also build the binaries manually with `go build`

    ```bash
    git clone https://github.com/XIniris/gopssh
    cd gopssh
    go build -o gopssh
    ```

- an examination

    After getting the binary file, you can check whether the program can run normally through `version`

    ```bash
    gopssh version
    ```

## Usage

### execute

`execute` is used to execute a command and return the execution result (stdout and stderr).

Currently supported options are:

- `-c`
  
    Specify the command to be executed

- `-f`

    Specifies the configuration file to use, defaults to `$HOME/.gopssh/inventory.yaml`.

- `-l`

    Specifies the label of the host that needs to execute the command. `=` and `!=` are supported. Multiple labels are separated by `,`. If not specified, all hosts will be selected.

- `-n`

    Whether to use the cache, the default is used, add `-n` to reparse the configuration file.

Detailed usage can be viewed through `gopssh execute -h`.

```bash
Execute command and return result

Usage:
  root execute [flags]

Examples:
  Simple: gopssh execute -c 'ls -l'
  Specify config: gopssh execute -c 'ls -l' -f /sample.yaml
  Select host to execute: gopssh execute -c 'ls -l' -l app=mysql
  Execute without cache: gopssh execute -c 'ls -l' -n

Flags:
  -c, --command string command to execute
  -f, --config-file string config file path
  -h, --help help for execute
  -l, --labels string label to filter on, default select all host, supports '=', and '!=' (e.g. -l key1=value1,key2!=value2
  -n, --without-cache not use cache, default use cache
```

### upload

`upload` is used to upload local files to the specified host.

Mainly concerned with options, other options are similar to `execute`:

- `-i`

    The local file path to upload.

- `-o`

    The file needs to be uploaded to the remote path, and the file needs to not exist on the remote host.

Detailed usage can be viewed through `gopssh upload -h`.

```bash
Upload file to remote

Usage:
  root upload [flags]

Examples:
 Simple: gopssh upload -i sample.txt -o /tmp/upload.txt
  Specify config: gopssh execute -i sample.txt -o /tmp/upload.txt -f /sample.yaml
  Select host to execute: gopssh execute -i sample.txt -o /tmp/upload.txt -l app=mysql
  Execute without cache: gopssh execute -i sample.txt -o /tmp/upload.txt -n

Flags:
  -f, --config-file string config file path
  -h, --help help for upload
  -l, --labels string label to filter on, default select all host, supports '=', and '!=' (e.g. -l key1=value1,key2!=value2
  -o, --output-path string upload file download path
  -i, --upload-file string the file to upload
  -n, --without-cache not use cache, default use cache
```

### check

`check` is used to check the connectivity of the host.

Detailed usage can be viewed through `gopssh check -h`.

```bash
Check all IP ports in the configuration file for connectivity

Usage:
  root check [flags]

Examples:
  Simple: gopssh check
  Specify config: gopssh check -f config.yaml
  Select host to check: gopssh check -l app=mysql
  Check without cache: gopssh check -f config.yaml -n

Flags:
  -f, --config-file string config file path
  -h, --help help for check
  -l, --labels string label to filter on, default select all host, supports '=', and '!=' (e.g. -l key1=value1,key2!=value2
  -n, --without-cache not use cache, default use cache
```

### get

`get` is used to get host details including IP, port, username, password (base64 encoded) and label.

Detailed usage can be viewed through `gopssh get -h`.

```bash
Get connection instance of config file

Usage:
  root get [flags]

Examples:
  Simple: gopssh get
  Specify config: gopssh get -f /sample.yaml
  Select host to execute: gopssh get -l app=mysql
  Execute without cache: gopssh get -n

Flags:
  -f, --config-file string config file path
  -h, --help help for get
  -l, --labels string label to filter on, default select all host, supports '=', and '!=' (e.g. -l key1=value1,key2!=value2
  -n, --without-cache not use cache, default use cache
```

### base64

`base64` can base64 encode or decode a string from base64.

Support options:

- `-c`

  Specifies the string that needs to be encoded/decoded.

- `-d`

  Whether it is the decoding mode, the default is the encoding mode.

Detailed usage can be viewed through `gopssh base64 -h`.

```bash
Encode or decode content with base64

Usage:
  root base64 [flags]

Examples:
  Encode: gopssh base64 -c 'root$123'
  Decode: gopssh base64 -d -c 'cm9vdA=='

Flags:
  -c, --content string decode / encode data content (required)
  -d, --decode decode or encode, default is encode
  -h, --help help for base64
```

### template

`template` can get the default configuration file template.

Support options:

- `-o`

  The default output is to the console, or it can be specified to save to a file with `-o`.

- `-t`

  Specifies the format of the configuration file template, supports `yaml` and `json`

Detailed usage can be viewed through `gopssh template -h`.

```bash
Dump config template, yaml or json

Usage:
  root template [flags]

Examples:
  Get yaml config template: gopssh template
  Save template as file: gopssh template -o sample.yaml
  Get json config template: gopssh template -t ​​json

Flags:
  -h, --help help for template
  -o, --output-path string output file path
  -t, --type string config template type, yaml or json (default "yaml")
```

### convert

`convert` can convert `yaml` files and `json` files.

- `-o`

  The default output is to the console, or it can be specified to save to a file with `-o`.

- `-y`

  Convert the `yaml` file to a `json` file, specify the path of the `yaml` file, and the file needs to end with `.yaml`.

- `-j`

  Convert `json` file to `yaml` file, specify the path of `json` file, the file needs to end with `.json`.

Detailed usage can be viewed through `gopssh convert -h`.

```bash
Convert json to yaml, or yaml to json

Usage:
  root convert [flags]

Examples:
  Convert YAML to JSON: gopssh convert -y sample.yaml
  Convert JSON to YAML: gopssh convert -j sample.json
  Convert and save: gopssh convert -j sample.yaml -o convert.yaml

Flags:
  -h, --help help for convert
  -j, --json-path string json file path
  -o, --output-path string output file path
  -y, --yaml-path string yaml file path
```

### version

`version` can check the current version of `gopssh`.

## Configure

The configuration is divided into two parts, `global` and `groups`.

### global

`global` provides a global default configuration that can be overridden by the corresponding options under `group`.

`global` contains three parts `port`, `username` and `password`. where `port` has a default value of `22` and `username` has a default value of `root`.

```yaml
global:
  port: 23
  username: root
  password: cm9vdCQxMjM=
```

### groups

`groups` is a list where each value is a `group`.

IPs under a `group` share the same `port`, `username`, `password` and `labels`. The corresponding value has a higher priority than `global`.

Each item of a `group` contains `ips`, `port`, `username`, `password` and `labels`. Wherein `ips` is required, it is a list of ip, you need to ensure that the value is a legal IP address; if `port`, `username`, `password` are not set, the corresponding value in `global` will be used; `labels` is the label of the group, it is a map, by specifying the label, you can easily select the desired host.

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


### Example

```yaml
# The global setting will take effect for all hosts,
# if he does not set the corresponding value
# The 'global' is not necessary
global:
  # If 'global.port' is not set, the default port is 22
  port: 23
  # If 'global.port' is not set, the default username is root
  username: root
  # The 'global.password' have not default value,
  # please set 'global.password' or set 'password' for each host
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
    # Encode password with base64 with command:
    # echo -n 'mysql$123' | base64
    # gopssh base64 'mysql$123'
    # Decode base64 password with command:
    # echo -n 'bXlzcWwkMTIz' | base64 -d
    # gopssh base64 -d 'bXlzcWwkMTIz'
    password: bXlzcWwkMTIz
    labels:
      # Execute command 'ls -l /' for all hosts that have label app=mysql
      # gopssh execute -l app=mysql 'ls -l /'
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
