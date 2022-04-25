#!/bin/bash
# Path: hack/build.sh

# > Define constants
HACK_PATH=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
readonly HACK_PATH
readonly REPOSITORY_PATH=${HACK_PATH}/..
readonly UPDAT_VERSION_FILE_PATH=${HACK_PATH}/update_version.sh

readonly SHELL_LOG_TIME_FORMAT="%Y-%m-%d %H:%m:%S"

# > Define variables
version=""
go_os=$(go env GOOS)
go_arch=$(go env GOARCH)

module_name=""

# > Define util functions
function info() {
  log_time=$(date +"${SHELL_LOG_TIME_FORMAT}")
  # GRAMMAR: \033[${COLOR};40m${TEXT}\033[0m
  # Color: Black 30, Red 31, Green 32, Brown 33, Blue 34, Purple 35, Cyan 36, White 37
  echo -e "\033[36;49mINFO\033[0m[${log_time}] $*"
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
  "${UPDAT_VERSION_FILE_PATH}" "${version}"
}

# shellcheck disable=SC2181
function build_binary() {
  set_module_name

  GOOS=${go_os} GOARCH=${go_arch} go build -o "${REPOSITORY_PATH}"/"${module_name}"
  if [ $? -ne 0 ]; then
    logger "failed tp build binary file"
    exit 1
  fi

  info "succeed to build binary file ${module_name}"
}

function set_module_name() {
  go_mod_file_path="${REPOSITORY_PATH}/go.mod"

  if [ -f "${go_mod_file_path}" ]; then
    module_name=$(grep "module" "${go_mod_file_path}" | awk '{print $2}')
  else
    module_name=$(basename "$(cd "${REPOSITORY_PATH}" && pwd)")
  fi

  info "module name is ${module_name}"
}

# > Execute main function
function main() {
  update_version
  build_binary
}

main "$@"
