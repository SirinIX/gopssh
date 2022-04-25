#!/bin/bash
# Path: hack/build.sh

# > Define constants
HACK_PATH=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
readonly HACK_PATH
readonly REPOSITORY_PATH=${HACK_PATH}/..
readonly UPDAT_VERSION_FILE_PATH=${HACK_PATH}/update_version.sh

readonly SHELL_LOG_TIME_FORMAT="%Y-%m-%d %H:%m:%S"

# > Define variables
version="0.0.1"

# > Define util functions
function info() {
  log_time=$(date +"${SHELL_LOG_TIME_FORMAT}")
  # GRAMMAR: \033[${COLOR};40m${TEXT}\033[0m
  # Color: Black 30, Red 31, Green 32, Brown 33, Blue 34, Purple 35, Cyan 36, White 37
  echo -e "\033[34;49mINFO\033[0m[${log_time}] $*"
}

function warning() {
  log_time=$(date +"${SHELL_LOG_TIME_FORMAT}")
  echo -e "\033[33;49mWARN\033[0m[${log_time}] $*"
}

function error() {
  log_time=$(date +"${SHELL_LOG_TIME_FORMAT}")
  echo -e "\033[31;49mERRO\033[0m[${log_time}] $*"
}

# > Define main functions
function update_version() {
  sh "${UPDAT_VERSION_FILE_PATH}" "${version}"
}

# > Execute main function
function main() {
  update_version
}

main "$@"
