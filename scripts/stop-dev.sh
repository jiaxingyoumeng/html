#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
PID_DIR="$ROOT_DIR/.tmp"

stop_process() {
  local name="$1"
  local pid_file="$PID_DIR/$name.pid"
  if [[ -f "$pid_file" ]]; then
    local pid
    pid=$(cat "$pid_file")
    if kill "$pid" >/dev/null 2>&1; then
      echo "Stopped $name (pid $pid)."
    else
      echo "$name process not running (pid $pid)."
    fi
    rm -f "$pid_file"
  else
    echo "$name pid file not found."
  fi
}

stop_process backend
stop_process frontend
