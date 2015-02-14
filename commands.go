package gowatch

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"

	"code.google.com/p/go.crypto/ssh/terminal"
)

func termWidth() (width int) {
	width, _, _ = terminal.GetSize(0)
	return
}

func header(s string) {
	headerColor := color.New(color.FgMagenta).SprintfFunc()
	if s != "" {
		s = " " + s + " "
		headerColor = color.New(color.FgGreen).SprintfFunc()
	}

	h := strings.Repeat("=", termWidth()-len(s))
	h1 := headerColor("%s", h[0:4])
	h2 := headerColor("%s", h[4:])
	s = color.RedString("%s", s)

	h = h1 + s + h2

	fmt.Println(h)
}

func run() {
	generate()
	lint()
	vet()
	test()
	coverage()
	header("")
}

func generate() {}

func lint() {
	header("Lint")
	runCommand("golint *.go")
}

func vet() {
	header("Vet")
	runCommand("go vet")
}

func test() {
	header("Test")
	runCommand("")
}

func coverage() {
	header("Coverage")
	runCommand("go test -race -coverprofile='coverage'")
}

func runCommand(command string) {
	cmd := exec.Command("/bin/bash", "-c", command)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error!")
		return
	}
	fmt.Print(string(out))
}
