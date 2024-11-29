package main

import (
	"fmt"
	"runtime"
)

var (
	// Version is the version of the binary in the vXXX.XXX.XXX format.
	Version string

	// GitCommit is the hash of the HEAD commit.
	GitCommit string

	// BuildDate is the date on which the binary was built.
	BuildDate string
)

func main() {
	versionInfo := fmt.Sprintf(`
Test Binary
-----------
Version: %s,
Git Commit Hash: %s,
Build Date: %s,
Go version: %s,
Operating System: %s,
Architecture:: %s,
`,
		Version,
		GitCommit,
		BuildDate,
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
	)
	fmt.Println(versionInfo)
}
