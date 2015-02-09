::go get github.com/lxn/walk
::go get github.com/akavel/rsrc
rsrc -manifest crypt_tool.exe.manifest -o rsrc.syso

gox -ldflags="-H windowsgui" -os="windows"
::go build -ldflags="-H windowsgui"
pause
