@echo off
set pwd=%cd%
cd %pwd%\src\Util
go build Util.go
go install
cd %pwd%\src\Zenit
go build Zenit.go
go install
cd %pwd%\src\Glaze
go build Glaze.go
go install
cd %pwd%\src\Entre
go build Entre.go
go install
cd %pwd%\src\Isas
go build Isas.go
go install
cd %pwd%
go build -o %pwd%\bin\menu.exe %pwd%\src\lunchMenu.go
