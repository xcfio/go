package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/sys/windows"
)

const script = "irm https://raw.githubusercontent.com/xcfio/settings/main/install.ps1 | iex"

func exit(code int) {
	fmt.Print("\nPress Enter to exit...")
	bufio.NewReader(os.Stdin).ReadString('\n')
	os.Exit(code)
}

func main() {
	if !windows.GetCurrentProcessToken().IsElevated() {
		fmt.Fprintln(os.Stderr, "This program must be run as Administrator.")
		exit(1)
	}

	shell := "pwsh"
	if _, err := exec.LookPath(shell); err != nil {
		shell = "powershell"
	}

	if _, err := exec.LookPath(shell); err != nil {
		fmt.Fprintln(os.Stderr, "No suitable PowerShell found. Please install PowerShell Core from https://learn.microsoft.com/en-us/powershell/scripting/install/install-powershell-on-windows.")
		exit(1)
	}

	cmd := exec.Command(shell, "-NoProfile", "-c", script)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		exit(1)
	}
	exit(0)
}
