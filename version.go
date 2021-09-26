package gowebbase

import (
	"fmt"
	"runtime"
)

const Version = "0.0.1"

// BuildDate returns the date the binary was built
var BuildDate = ""

// GitCommit returns the git commit that was compiled. This will be filled in by the compiler.
var GitCommit string

// GoVersion returns the version of the go runtime used to compile the binary
var GoVersion = runtime.Version()

// OsArch returns the os and arch used to build the binary
var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)


func GetVersionInfo() string{
	return fmt.Sprintf("Version: %s\nBuildData: %s\nGoVersion: %s, OsArch: %s\n\n", Version, BuildDate, GoVersion, OsArch)
}
