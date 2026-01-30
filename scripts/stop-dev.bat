@echo off
setlocal

echo Stopping services on ports 8080 and 4173 ...
powershell -NoProfile -Command ^
  "$ports = 8080,4173; " ^
  "$pids = Get-NetTCPConnection -LocalPort $ports -ErrorAction SilentlyContinue | " ^
  "Select-Object -ExpandProperty OwningProcess -Unique; " ^
  "if ($pids) { Stop-Process -Id $pids -Force } else { Write-Host 'No processes found.' }"

echo Done.
