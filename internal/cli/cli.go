package cli

import (
	"buildtool/internal/config"
	"buildtool/internal/core"
	"buildtool/internal/scriptcompiler"
	"fmt"
	"os"

	"github.com/shadowdara/finder/pub/argparser"
)

func ParseArgs() {
	// Main Command
	root := argparser.NewCommand("bt", "a simple build helping tool", false)

	root.String("tag", "", "Tag you want to add to a Git Repo", false, "t")
	root.Bool("bin", false, "Add the script shortcuts to your terminal path", false)
	root.Bool("compile", false, "Compile the Scripts from the Config File", false)

	cmd := root.Parse(os.Args[1:])

	// Add a Tag to a Git Repo
	res := cmd.GetString("tag")
	if res != "" {
		core.Addtag(res)
		return
	}

	var config = config.LoadConfig()

	// Compile scripts
	if cmd.GetBool("compile") == true {
		scriptcompiler.CompileScripts(config)
		return
	}

	// add script Path
	if cmd.GetBool("bin") == true {
		scriptcompiler.Addpath()
		return
	}

	switch cmd {
	default:
		fmt.Println("Buildtool")
		cmd.PrintHelp()
	}
}
