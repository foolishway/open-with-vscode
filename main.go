package main

import (
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

func main() {
	var pwd, _ = os.Getwd()
	var target string = pwd
	if len(os.Args) > 1 {
		target = os.Args[1]
	}

	isGithub := isFromGithub(target)

	if isGithub {
		var dirName string = (target)[strings.LastIndex((target), "/")+1:]
		dirName = dirName[:strings.LastIndex(dirName, ".")]
		var clonePath = "./" + dirName
		if len(os.Args) >= 2 {
			clonePath = path.Clean(os.Args[2] + "/" + dirName)
		}
		gitClone(target, clonePath)
		openWithVscode(clonePath)
	} else {
		openWithVscode(target)
	}
}

func isFromGithub(path string) bool {
	isGithub, err := regexp.MatchString(`^git@|https:\/\/`, path)

	if err != nil {
		panic(err)
	}

	return isGithub
}

func gitClone(repo string, dir string) {
	var cmd *exec.Cmd
	if dir != "" {
		cmd = exec.Command("git", "clone", repo, dir)
	} else {
		cmd = exec.Command("git", "clone", repo)
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
