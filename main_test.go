package main_test

import (
	"os/exec"
	"regexp"
	"testing"
)

func Test(t *testing.T) {
	if err := exec.Command("go", "build", "main.go").Run(); err != nil {
		panic("build failed: " + err.Error())
	}

	expectedRegexp := regexp.MustCompile(`\Afatal error: concurrent map writes\n\ngoroutine .*`)

	for {
		output, err := exec.Command("./main").CombinedOutput()
		if err == nil {
			panic("expected error, got output: " + string(output))
		}

		if expectedRegexp.Match(output) {
			continue
		}

		t.Fatalf("unexpected output: \n%s", output)
	}
}
