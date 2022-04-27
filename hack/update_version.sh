#!/bin/bash
# Path: hack/update_version.sh

# > Define constants
HACK_PATH=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
readonly HACK_PATH
readonly REPOSITORY_PATH=${HACK_PATH}/..
readonly GIT_BASE_PATH=${REPOSITORY_PATH}/.git

readonly SHELL_LOG_TIME_FORMAT="%Y-%m-%d %H:%m:%S"

readonly COMMAND_VERSION_FILE_PATH=${REPOSITORY_PATH}/cmd/version/version.go

readonly GOLANG_VERSION_CONSTANT_PREFIX="version = "
readonly GOLANG_COMMIT_TIME_CONSTANT_PREFIX="latestCommitDate = "

# > Define variables
current_branch=""
latest_git_commit_time=""

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
# DESC: Update file version.go constant version as input
# ARGS: $1 (optional): Version
# OUTS: None
# shellcheck disable=SC2181
function update_version_in_version_command() {
  if [ -z "$1" ]; then
    info "no version input, do not update version"
    return
  fi
  # Update version
  sed -i "" "s/${GOLANG_VERSION_CONSTANT_PREFIX}\".*\"/${GOLANG_VERSION_CONSTANT_PREFIX}\"$1\"/g" "${COMMAND_VERSION_FILE_PATH}"
  if [ $? -ne 0 ]; then
    error "failed to update version"
    return 1
  fi

  info "succeed to update version as $1 in file ${COMMAND_VERSION_FILE_PATH}"
}

# DESC: Update file version.go constant latestCommitDate as lastest git commit time of current branch
# ARGS: None
# OUTS: None
# shellcheck disable=SC2181
function update_git_commit_time_in_version_command() {
  is_git_base_dir_exist
  if [ $? -eq 1 ]; then
    warning "please run this script in a git repository"
    return 1
  fi

  # Get the latest commit time
  set_current_branch_latest_git_commit_time
  if [ $? -ne 0 ]; then
    error "failed to get latest git commit time"
    return 1
  fi

  # Update the latest commit time in version.go
  # readonly bak_suffix=".bak"
  # sed -i "${bak_suffix}" "s/${GOLANG_COMMIT_TIME_VARIABLE_NAME} = \".*\"/${GOLANG_COMMIT_TIME_VARIABLE_NAME} = \"${latest_git_commit_time}\"/g" "${COMMAND_VERSION_FILE_PATH}"
  sed -i "" "s/${GOLANG_COMMIT_TIME_CONSTANT_PREFIX}\".*\"/${GOLANG_COMMIT_TIME_CONSTANT_PREFIX}\"${latest_git_commit_time}\"/g" "${COMMAND_VERSION_FILE_PATH}"
  if [ $? -ne 0 ]; then
    error "failed to update latest git commit time"
    return 1
  fi
  # # Clean up backup file
  # rm -rf "${COMMAND_VERSION_FILE_PATH}"${bak_suffix}

  info "succeed to update commit time as ${latest_git_commit_time} in file ${COMMAND_VERSION_FILE_PATH}"
}

# shellcheck disable=SC2181
function set_current_branch_latest_git_commit_time() {
  set_current_branch
  if [ $? -ne 0 ]; then
    error "failed to get current branch"
    return 1
  fi

  git_log_file_path="${GIT_BASE_PATH}/logs/refs/heads/${current_branch}"
  if [ ! -f "${git_log_file_path}" ]; then
    error "can not find git log file: ${git_log_file_path}"
    return 1
  fi

  raw_time=$(tail <"${git_log_file_path}" -1 | awk '{print $5}')
  # Set the latest git commit time
  readonly GIT_COMMIT_TIME_FORMAT="+%Y-%m-%d %H:%m:%S"
  latest_git_commit_time="$(date -r "${raw_time}" "${GIT_COMMIT_TIME_FORMAT}")"

  info "branch ${current_branch} latest git commit time is ${latest_git_commit_time}"
}

function set_current_branch() {
  git_head_file_path="${GIT_BASE_PATH}/HEAD"
  if [ ! -f "${git_head_file_path}" ]; then
    error "can not find ${git_head_file_path}"
    return 1
  fi

  ref=$(awk <"${git_head_file_path}" '{print $2}')
  # Set the current branch
  current_branch="${ref##*/}"

  info "current branch is ${current_branch}"
}

function is_git_base_dir_exist() {
  if [ ! -d "${GIT_BASE_PATH}" ]; then
    warning "${REPOSITORY_PATH} is not a git repository"
    return 1
  fi
  return 0
}

# > Execute main function
function main() {
  update_git_commit_time_in_version_command
  update_version_in_version_command "$1"
}

main "$@"
