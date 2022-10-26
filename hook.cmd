@echo off
setlocal enabledelayedexpansion
echo %date%  %time%
git fetch --all
git reset --hard origin/master
git pull --rebase origin master

set port=14814
for /f "tokens=1-5" %%a in ('netstat -ano^|findstr ":%port%"') do (
    echo %%a %%b %%c %%d %%e
    if "%%e%" == "" (
        taskkill /f /pid %%d
    ) else (
        taskkill /f /pid %%e
    )
)

go run main.go
