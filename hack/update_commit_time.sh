#!/bin/bash
# Path: hack/update_commit_time.sh

# > Define constants
HACK_PATH=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
readonly HACK_PATH
readonly REPOSITORY_PATH=${HACK_PATH}/..

readonly GIT_BASE_PATH=${REPOSITORY_PATH}/.git
readonly GIT_LOG_DIR=${GIT_BASE_PATH}/logs/refs/heads
readonly GIT_COMMIT_TIME_FORMAT="+%Y-%m-%d %H:%m:%S"

readonly COMMAND_VERSION_FILE_PATH=${REPOSITORY_PATH}/cmd/version.go

readonly SHELL_LOG_TIME_FORMAT="%Y-%m-%d %H:%m:%S"

readonly GOLANG_COMMIT_TIME_VARIABLE_NAME="latestCommitDate"

# > Define variables
current_branch=""
latest_git_commit_time=""

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
  readonly bak_suffix=".bak"
  sed -i "${bak_suffix}" "s/${GOLANG_COMMIT_TIME_VARIABLE_NAME} = \".*\"/${GOLANG_COMMIT_TIME_VARIABLE_NAME} = \"${latest_git_commit_time}\"/g" "${COMMAND_VERSION_FILE_PATH}"
  # Clean up backup file
  rm -rf "${COMMAND_VERSION_FILE_PATH}"${bak_suffix}

  info "succeed to update variable ${GOLANG_COMMIT_TIME_VARIABLE_NAME} value as ${latest_git_commit_time} in file ${COMMAND_VERSION_FILE_PATH}"
}

# shellcheck disable=SC2181
function set_current_branch_latest_git_commit_time() {
  set_current_branch
  if [ $? -ne 0 ]; then
    error "failed to get current branch"
    return 1
  fi

  git_log_file_path="${GIT_LOG_DIR}/${current_branch}"
  if [ ! -f "${git_log_file_path}" ]; then
    error "can not find git log file: ${git_log_file_path}"
    return 1
  fi

  raw_time=$(tail <"${git_log_file_path}" -1 | awk '{print $5}')
  # Set the latest git commit time
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
update_git_commit_time_in_version_command
