@echo off
setlocal

cd /d "%~dp0\.."

if "%PUBMED_API_KEY%"=="" (
  set "PUBMED_API_KEY=b7972bf7641ea4b41fe0193baae92207b308"
)

echo Starting backend on http://localhost:8080 ...
start "xiaozhi-backend" cmd /c "cd /d backend && set PUBMED_API_KEY=%PUBMED_API_KEY% && go run ./cmd/server"

echo Starting frontend on http://localhost:4173/index.html ...
start "xiaozhi-frontend" cmd /c "cd /d frontend && python -m http.server 4173"

echo.
echo Backend:  http://localhost:8080
echo Frontend: http://localhost:4173/index.html
echo.
echo Use scripts\stop-dev.bat to stop services.
