@echo off
pushd %~dp0
IF EXIST "./msvc/cl.exe"

REM install https://aka.ms/vs/16/release/vs_buildtools.exe
REM C:\TEMP\vs_buildtools.exe --quiet --wait --norestart --nocache `
REM    --installPath C:\BuildTools `
REM    --add Microsoft.VisualStudio.Workload.AzureBuildTools `
REM    --remove Microsoft.VisualStudio.Component.Windows10SDK.10240 `
REM    --remove Microsoft.VisualStudio.Component.Windows10SDK.10586 `
REM    --remove Microsoft.VisualStudio.Component.Windows10SDK.14393 `
REM    --remove Microsoft.VisualStudio.Component.Windows81SDK `
REM || IF "%ERRORLEVEL%"=="3010" EXIT 0

:endfile
popd