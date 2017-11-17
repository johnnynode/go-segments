package main

import (
	"os"
	"fmt"
  "strings"
)

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	os.Unsetenv("FOO")

	fmt.Println()
	// 以下是循环出的环境变量
	for _, e := range os.Environ() {
		pair := strings.Split(e,"=")
		fmt.Println(pair[0])
	}
}

/*
$ go run Basic/environment-variables.go
FOO: 1
BAR:

TERM_PROGRAM
ANDROID_HOME
TERM
SHELL
TMPDIR
GRADLE_HOME
GOBIN
Apple_PubSub_Socket_Render
TERM_PROGRAM_VERSION
ANDROID_PLATFORM_TOOLS
ANDROID_BT_CLI
JRE_HOME
ANDROID_TOOLS
USER
CLASS_PATH
SSH_AUTH_SOCK
__CF_USER_TEXT_ENCODING
VSCODE_NODE_CACHED_DATA_DIR_922
GOOGLE_API_KEY
PATH
PWD
JAVA_HOME
LANG
XPC_FLAGS
XPC_SERVICE_NAME
SHLVL
HOME
GOROOT
VSCODE_NLS_CONFIG
LOGNAME
VSCODE_IPC_HOOK
GOPATH
VSCODE_PID
_
*/