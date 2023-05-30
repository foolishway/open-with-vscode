package main

import "testing"

func TestIsFromGithub(t *testing.T) {
	github := "https://github.com/foolishway/open-with-vscode.git"
	p := "/path"

	isGithub := isFromGithub(github)

	if !isGithub {
		t.Errorf("expect true but false")
	}

	isGithub = isFromGithub(p)

	if isGithub {
		t.Errorf("expect false but true")
	}
}
