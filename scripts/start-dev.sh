#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
PID_DIR="$ROOT_DIR/.tmp"

mkdir -p "$PID_DIR"

start_backend() {
  if [[ -f "$PID_DIR/backend.pid" ]]; then
    echo "Backend already running (pid $(cat "$PID_DIR/backend.pid"))."
    return
  fi
  (cd "$ROOT_DIR/backend" && nohup go run ./cmd/server > "$PID_DIR/backend.log" 2>&1 & echo $! > "$PID_DIR/backend.pid")
  echo "Backend started (pid $(cat "$PID_DIR/backend.pid"))."
}

start_frontend() {
  if [[ -f "$PID_DIR/frontend.pid" ]]; then
    echo "Frontend already running (pid $(cat "$PID_DIR/frontend.pid"))."
    return
  fi
  (cd "$ROOT_DIR/frontend" && nohup python -m http.server 4173 > "$PID_DIR/frontend.log" 2>&1 & echo $! > "$PID_DIR/frontend.pid")
  echo "Frontend started (pid $(cat "$PID_DIR/frontend.pid"))."
}

start_backend
start_frontend

echo "Backend: http://localhost:8080"
echo "Frontend: http://localhost:4173/index.html"
