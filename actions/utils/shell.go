package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func RunCleanupShellCommands(appName string, appPath string) {
	runShellCommand("go", []string{"mod", "init", appName}, appPath)
	runShellCommand("go", []string{"install", "github.com/google/wire/cmd/wire@latest"}, appPath)
	runShellCommand("go", []string{"get", "./..."}, appPath)
	// runShellCommand("go", []string{"get", "github.com/google/wire"}, appPath)
	wirePath := appPath + "/internal/wire"
	runShellCommand("wire", []string{}, wirePath)
	runShellCommand("go", []string{"generate", "./..."}, appPath)
	runShellCommand("go", []string{"fmt", "./..."}, appPath)
}

func runShellCommand(app string, args []string, workingDirectory string) {
	fmt.Printf("Running: %v %v at %v\n", app, strings.Join(args, " "), workingDirectory)
	cmd := exec.Command(app, args...)
	cmd.Dir = workingDirectory
	err := cmd.Run()

	if err != nil {
		// implement an exponential backoff here
		fmt.Printf("Error running %v %v: %v\n", app, strings.Join(args, " "), err.Error())
		return
	}
}
