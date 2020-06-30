FROM ubuntu:18.04
# escape=`

# Use the latest Windows Server Core image with .NET Framework 4.8.
FROM mcr.microsoft.com/dotnet/framework/sdk:4.8-windowsservercore-ltsc2019

# Restore the default Windows shell for correct batch processing.
SHELL ["cmd", "/S", "/C"]
COPY ./src ./src
RUN gcc -v ./src/macserial/macserial.c -o ./src/macserial.prog
ENTRYPOINT [ "python", "./src/main.py" ]