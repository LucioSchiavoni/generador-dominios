@echo off
cls
echo Ejecutando el programa como administrador...
powershell Start-Process cmd -ArgumentList "/c go run main.go" -Verb RunAs
