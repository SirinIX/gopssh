#!/bin/bash
# Path: hack/build.sh

# > Define constants
HACK_PATH=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
readonly HACK_PATH
readonly REPOSITORY_PATH=${HACK_PATH}/..
readonly UPDAT_COMMIT_TIME_FILE_PATH=${REPOSITORY_PATH}/update_commit_time.sh

readonly SHELL_LOG_TIME_FORMAT="%Y-%m-%d %H:%m:%S"

# > Define variables

# > Define util functions
function info() {
  log_time=$(date +"${SHELL_LOG_TIME_FORMAT}")
  # GRAMMAR: \033[${COLOR};40m${TEXT}\033[0m
  # Color: Black 30, Red 31, Green 32, Brown 33, Blue 34, Purple 35, Cyan 36, White 37
  echo -e "[${log_time}] [\033[32;49mINFO\033[0m] $*"
}

function warning() {
  log_time=$(date +"${SHELL_LOG_TIME_FORMAT}")
  echo -e "[${log_time}] [\033[33;49mWARNING\033[0m] $*"
}

function error() {
  log_time=$(date +"${SHELL_LOG_TIME_FORMAT}")
  echo -e "[${log_time}] [\033[31;49mERROR\033[0m] $*"
}

# > Define main functions
function update_commit_time() {
    sh "${UPDAT_COMMIT_TIME_FILE_PATH}"
}

# > Execute main function
function main() {
    update_commit_time
}

main "$@"