@echo off
set APP=guestbook.exe
set ADDR=:8081
set PWD=%cd%\..
set GO_PWD=%PWD%\..\..
echo "Golanger Web Framework"
echo "Golanger is a lightweight framework for writing web applications in Golang."
set GOPATH=%GO_PWD%\framework;%PWD%\src\add-on;%PWD%
cd %PWD%\src

if exist %APP% del %APP%

echo "Building %APP%"
go build .

if exist src.exe (
    rename src.exe %APP%
    echo "Runing %APP%"
    %APP% -addr=%ADDR%
)