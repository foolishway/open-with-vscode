package main

import (
	"flag"
	"os"
	"os/exec"
	"path"
	"strings"
)

var source = flag.String("s", "", "github repository address")

func main() {
	flag.Parse()

	var pwd, _ = os.Getwd()
	if *source == "" {
		var t string = pwd
		if len(os.Args) > 1 {
			t = os.Args[1]
		}
		openWithVscode(t)
	} else {
		var gitRepo = *source

		var dirName string = (gitRepo)[strings.LastIndex((gitRepo), "/")+1:]
		dirName = dirName[:strings.LastIndex(dirName, ".")]

		// download git repo
		if len(flag.Args()) != 0 {
			var clonePath string = path.Clean(flag.Args()[0] + "/" + dirName)
			gitClone(gitRepo, clonePath)
			openWithVscode(clonePath)
		} else {
			gitClone(*source, "")
			openWithVscode("./" + dirName)
		}
	}
}

func gitClone(repo string, dir string) {
	var cmd *exec.Cmd
	if dir != "" {
		cmd = exec.Command("git", "clone", *source, dir)
	} else {
		cmd = exec.Command("git", "clone", *source)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func openWithVscode(path string) {
	var err = exec.Command("code", path).Run()
	if err != nil {
		panic(err)
	}
}
