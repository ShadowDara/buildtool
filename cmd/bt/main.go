package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/shadowdara/finder/pub/argparser"
)

func main() {
	fmt.Println("Buildtool")

	// Main Command
	root := argparser.NewCommand("bt", "a simple build helping tool", false)
	root.String("tag", "", "Tag you want to add to a Git Repo", false, "t")

	cmd := root.Parse(os.Args[1:])

	res := cmd.GetString("tag")
	if res != "" {
		addtag(res)
		return
	}

	switch cmd {
	default:
		cmd.PrintHelp()
	}

}

// Function to add a Tag to a git repo and push it
func addtag(tag string) {
	cmd := exec.Command("git", "tag", tag)

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}

	cmd = exec.Command("git", "push", "origin", tag)

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
