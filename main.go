package main

import (
	// import Root Command
	"github.com/teamssix/cf/command"
	// import modules with _
	_ "github.com/teamssix/cf/command/misc"
	_ "github.com/teamssix/cf/command/scan"
)

func main() {
	Command.Execute()
}
