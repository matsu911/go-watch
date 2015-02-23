package gowatch

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"

	"code.google.com/p/go.crypto/ssh/terminal"
)

func termWidth() (width int) {
	width, _, err := terminal.GetSize(0)
	if err != nil {
		debug("Failed to get terminal size: %v.\n", err)
	}
	return
}

func header(s string, c color.Attribute) {
	headerColor := color.New(c).SprintfFunc()

	h := strings.Repeat("=", termWidth()-len(s)-2)
	h1 := headerColor("%s", h[0:4])
	h2 := headerColor("%s", h[4:])
	s = color.RedString(" %s ", s)

	h = h1 + s + h2

	fmt.Println(h)
}

func greenHeader(s string) {
	header(s, color.FgGreen)
}

func magentaHeader(s string) {
	header(s, color.FgMagenta)
}

func run(count int) {
	generate()
	lint()
	vet()
	test()
	coverage()
	magentaHeader(strconv.Itoa(count))
}

func printer(title, output string) {
	if output != "" {
		greenHeader(title)
		fmt.Print(output)
	}
}

func generate() {}

func lint() {
	output := runCommand("golint *.go")
	printer("Lint", output)
}

func vet() {
	output := runCommand("go vet")
	printer("Vet", output)
}

func test() {
	output := runCommand("go test -race -coverprofile='coverage'")
	// output := runCommand("go test -race -coverprofile='coverage' -tags debug")
	printer("Test", output)
}

func coverage() {
	output := runCommand("go tool cover -func='coverage'")
	// _ = runCommand("rm ./coverage")
	printer("Coverage", output)
}

func runCommand(command string) string {
	cmd := exec.Command("/bin/bash", "-c", command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		switch err.Error() {
		case "exit status 2":
			debug("Command returned an expected error: %v.\n", err)
		default:
			debug("Command returned: %v.\n", err)
			// panic(err)
		}
	}
	return string(out)
}
