@echo off
:menu
cls
echo ======================================
echo        PORT KILLER
echo ======================================
echo 0 - View current port usage
echo 1 - Kill a process by port
echo 2 - Check port is Used or not
echo q - Quit
echo ======================================
set /p choice="Please enter your choice (0, 1, 2, or q): "

if /i "%choice%"=="0" goto view_ports
if /i "%choice%"=="1" goto kill_port
if /i "%choice%"=="2" goto check_port
if /i "%choice%"=="q" goto exit
echo Invalid choice, please try again.
pause
goto menu

:view_ports
go run main.go
pause
goto menu

:kill_port
set /p port="Enter the port number to kill (1-65535): "
echo %port%| findstr /R "^[0-9]*$" >nul
if %errorlevel% neq 0 (
    echo Invalid port number, must be numeric.
    pause
    goto menu
)
if %port% lss 1 (
    echo Port must be between 1 and 65535.
    pause
    goto menu
)
if %port% gtr 65535 (
    echo Port must be between 1 and 65535.
    pause
    goto menu
)
go run main.go -c %port%
pause
goto menu

:check_port
set /p port="Enter the port number to check (1-65535): "
echo %port%| findstr /R "^[0-9]*$" >nul
if %errorlevel% neq 0 (
    echo Invalid port number, must be numeric.
    pause
    goto menu
)
if %port% lss 1 (
    echo Port must be between 1 and 65535.
    pause
    goto menu
)
if %port% gtr 65535 (
    echo Port must be between 1 and 65535.
    pause
    goto menu
)
go run main.go %port%
pause
goto menu

:exit
echo Goodbye!
exit /b