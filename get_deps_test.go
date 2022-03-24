package main

import (
	"os"
	"testing"
)

func TestParseVanagonOutputNix(t *testing.T) {
	dat, err := os.ReadFile("testfiles/pxp-agent_el-8-x86_64.json")
	if err != nil {
		t.Errorf("Couldn't read test file")
	}
	contentString := string(dat)
	gems, err := parseVanagonOutput(contentString, "foo", "bar")
	if err != nil {
		t.Fatalf("Error parsing vanagon string. Err: %s", err)
	}
	a := gems
	_ = a
	defer os.RemoveAll(LOCKFILE_DIR)
	_, err = buildGemFile("foo", "bar", &gems)
	if err != nil {
		t.Fatalf("Error creating gemfile. Err: %s", err)
	}
}

func TestParseVanagonOutputWin(t *testing.T) {
	dat, err := os.ReadFile("testfiles/pxp-agent_windowsfips-2012r2-x64.json")
	if err != nil {
		t.Errorf("Couldn't read test file")
	}
	contentString := string(dat)
	gems, err := parseVanagonOutput(contentString, "foo", "bar")
	if err != nil {
		t.Fatalf("Error parsing vanagon string. Err: %s", err)
	}
	a := gems
	_ = a
	defer os.RemoveAll(LOCKFILE_DIR)
	_, err = buildGemFile("foo", "bar", &gems)
	if err != nil {
		t.Fatalf("Error creating gemfile. Err: %s", err)
	}
}
