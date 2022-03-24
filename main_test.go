package main

import (
	"os"
	"testing"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func TestGetEnvVar(t *testing.T) {
	os.Setenv("INPUT_SNYKTOKEN", "foo")
	defer os.Unsetenv("INPUT_SNYKTOKEN")

	os.Setenv("INPUT_SNYKORG", "org")
	defer os.Unsetenv("INPUT_SNYKORG")

	os.Setenv("GITHUB_WORKSPACE", "/var/foo")
	defer os.Unsetenv("GITHUB_WORKSPACE")

	os.Setenv("INPUT_NOMONITOR", "true")
	defer os.Unsetenv("INPUT_NOMONITOR")

	os.Setenv("INPUT_SKIPPLATFORMS", "foo,bar,baz")
	defer os.Unsetenv("INPUT_SKIPPLATFORMS")

	os.Setenv("INPUT_SKIPPROJECTS", "abc,def,hij")
	defer os.Unsetenv("INPUT_SKIPPROJECTS")

	os.Setenv("INPUT_URLSTOREPLACE", "artifactory.delivery.puppetlabs.net,%s/xart,builds.delivery.puppetlabs.net,%s/xbuild")
	defer os.Unsetenv("INPUT_URLSTOREPLACE")

	os.Setenv("INPUT_NEWHOST", "localhost")
	defer os.Unsetenv("INPUT_NEWHOST")

	os.Setenv("INPUT_SVDEBUG", "true")
	defer os.Unsetenv("INPUT_SVDEBUG")

	os.Setenv("INPUT_BRANCH", "foobarbaz123")
	defer os.Unsetenv("INPUT_BRANCH")

	conf, err := getEnvVar()
	if err != nil {
		t.Errorf("Error running getEnvVar: %s", err)
	}
	if conf.SnykToken != "foo" {
		t.Errorf("Snyk token value wrong. Got: %s", conf.SnykToken)
	}
	if conf.SnykOrg != "org" {
		t.Errorf("Snyk org value wrong. Got: %s", conf.SnykOrg)
	}
	if conf.GithubWorkspace != "/var/foo" {
		t.Errorf("workspace value wrong. Got: %s", conf.GithubWorkspace)
	}
	if !conf.NoMonitor {
		t.Errorf("NoMonitor should have been true")
	}
	plat_tvals := []string{"foo", "bar", "baz"}
	for _, v := range plat_tvals {
		if !contains(conf.SkipPlatforms, v) {
			t.Errorf("SkipPlatforms did not have val: %s", v)
		}
	}
	proj_tvals := []string{"abc", "def", "hij"}
	for _, v := range proj_tvals {
		if !contains(conf.SkipProjects, v) {
			t.Errorf("SkipProjects did not have val: %s", v)
		}
	}
	expected_urls := map[string]string{
		"artifactory.delivery.puppetlabs.net": "%s/xart",
		"builds.delivery.puppetlabs.net":      "%s/xbuild",
	}
	for k, v := range expected_urls {
		if val, ok := conf.UrlsToReplace[k]; ok {
			if val != v {
				t.Errorf("UrlsToReplace val doesn't match: Got: %s:%s", k, val)
			}
		} else {
			t.Errorf("Couldn't find key %s in UrlsToReplace", k)
		}
	}
	if conf.ProxyHost != "localhost" {
		t.Errorf("proxyhost value wrong. Got: %s", conf.ProxyHost)
	}
	if !conf.Debug {
		t.Errorf("Debug should have been true")
	}
	if conf.Branch != "foobarbaz1" {
		t.Errorf("branch value wrong. Got: %s", conf.Branch)
	}

}
