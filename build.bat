@echo off
set pwd=%cd%
cd %pwd%\src\Util
go build Util.go
go install
cd %pwd%\src\Zenit
go build Zenit.go
go install
cd %pwd%
go build -o %pwd%\bin\menu.exe %pwd%\src\lunchMenu.go
