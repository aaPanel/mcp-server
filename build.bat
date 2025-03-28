@echo off
setlocal enabledelayedexpansion

REM Set Go commands and environment variables
set GOCMD=go
set GOBUILD=%GOCMD% build
set GOCLEAN=%GOCMD% clean

REM Set build architecture
set GOARCH=amd64
set GOOS=windows

REM Set paths
set BASE_PATH=%CD%
set BUILD_PATH=%BASE_PATH%\build
set MAIN_PATH=%BASE_PATH%\main.go
set BIN_NAME=mcp-btpanel.exe

if "%1"=="build" goto build
if "%1"=="clean" goto clean
if "%1"=="" goto help

:build
    echo Building %BIN_NAME%...
    if not exist %BUILD_PATH% mkdir %BUILD_PATH%
    cd %BASE_PATH%
    set GOARCH=%GOARCH%
    set GOOS=%GOOS%
    %GOBUILD% -trimpath -ldflags "-s -w" -o %BUILD_PATH%\%BIN_NAME% %MAIN_PATH%
    if %ERRORLEVEL% EQU 0 (
        echo Build completed successfully.
        echo Binary location: %BUILD_PATH%\%BIN_NAME%
    ) else (
        echo Build failed.
    )
    goto end

:clean
    echo Cleaning...
    if exist %BUILD_PATH% rmdir /s /q %BUILD_PATH%
    %GOCLEAN%
    echo Clean completed.
    goto end

:help
    echo Usage:
    echo   build.bat build    - Build the binary
    echo   build.bat clean    - Clean build artifacts
    goto end

:end
endlocal