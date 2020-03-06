@echo off
set pwd=%cd%
go build -o %pwd%\bin\menu.exe %pwd%\src\lunchMenu.go
