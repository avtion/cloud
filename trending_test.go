package main

import (
	"testing"
)

func Test_githubTrending_getProjects(t *testing.T) {
	projects, err := newGithubTrending().getProjects()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(projects)
}
