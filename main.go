/*
DreShellGUI

Author: Andrias Zelele
Date: 2026

Description:
DreShellGUI is a graphical shell interface built using the Fyne GUI toolkit.
The application allows users to run common system commands through a GUI
while displaying output similar to a traditional terminal.

Features:
- Built-in commands: help, clear, exit, pwd, cd
- Executes system commands using the OS shell (bash/zsh/sh on Unix, cmd on Windows)
- Prevents interactive commands that would freeze the GUI
- Scrollable terminal output
- Cross-platform support (Linux, macOS, Windows)

Note:
Some interactive terminal programs (nano, vim, sudo, etc.) are intentionally
blocked because they require a full terminal (TTY) environment.
*/

package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// currentDir stores the active working directory used by DreShellGUI.
// All commands executed by the shell run relative to this directory.
var currentDir string

// runCommand executes a command entered by the user.
// It handles built-in commands first and then runs system commands
// through the operating system shell.
func runCommand(input string) string {

	// Split the command string into individual parts (command + arguments)
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return ""
	}

	// List of commands that require an interactive terminal.
	// These commands are blocked to prevent the GUI from freezing.
	var blockedCommands = map[string]bool{
		"sudo":    true,
		"su":      true,
		"passwd":  true,
		"ssh":     true,
		"ftp":     true,
		"sftp":    true,
		"nano":    true,
		"vim":     true,
		"vi":      true,
		"top":     true,
		"htop":    true,
		"less":    true,
		"more":    true,
		"man":     true,
		"python":  true,
		"python3": true,
		"node":    true,
		"bash":    true,
		"sh":      true,
		"zsh":     true,
		"mysql":   true,
		"psql":    true,
		"sqlite3": true,
		"watch":   true,
	}

	// Prevent execution of blocked interactive commands
	if blockedCommands[parts[0]] {
		return parts[0] + " is not supported in DreShellGUI because it requires an interactive terminal."
	}

	// Handle built-in commands
	switch parts[0] {

	// Display help information for the user
	case "help":
		return strings.Join([]string{
			"DreShellGUI Help",
			"",
			"Built-in commands:",
			"  help   - show this help message",
			"  clear  - clear the terminal output",
			"  exit   - close DreShell GUI",
			"  pwd    - show current directory",
			"  cd     - change directory",
			"",
			"System commands:",
			"  DreShellGUI executes commands using your operating system's shell.",
			"",
			"  Linux/macOS: bash, zsh, or sh",
			"  Windows: cmd (Command Prompt)",
			"",
			"  This means commands may differ depending on your OS:",
			"    - 'ls' works on Linux/macOS",
			"    - 'dir' works on Windows",
			"",
			"  DreShellGUI does not translate commands between systems.",
			"Examples:",
			"  ls",
			"  echo Hello",
			"  mkdir test",
			"  cat file.txt",
			"",
			"Note:",
			"  Some interactive commands are disabled because they require",
			"  a full terminal (nano, vim, sudo, etc.).",
		}, "\n")

	// Exit the application
	case "exit":
		os.Exit(0)
		return ""

	// Display the current working directory
	case "pwd":
		return currentDir

	// Change directory command implementation
	case "cd":

		// If no directory is provided, go to the user's home directory
		if len(parts) == 1 {
			home, err := os.UserHomeDir()
			if err != nil {
				return "cd: unable to find home directory"
			}
			currentDir = home
			return "" + currentDir
		}

		target := parts[1]

		// Support "~" as shorthand for the home directory
		if target == "~" {
			home, err := os.UserHomeDir()
			if err != nil {
				return "cd: unable to find home directory"
			}
			currentDir = home
			return ""
		}

		// Convert relative paths to absolute paths
		if !filepath.IsAbs(target) {
			target = filepath.Join(currentDir, target)
		}

		// Clean the path (removes ../ and redundant separators)
		target = filepath.Clean(target)

		// Validate that the target exists
		info, err := os.Stat(target)
		if err != nil {
			return "cd: no such file or directory: " + target
		}

		// Ensure the target is actually a directory
		if !info.IsDir() {
			return "cd: not a directory: " + target
		}

		// Update the current working directory
		currentDir = target
		return ""
	}

	// Prepare the command execution
	var cmd *exec.Cmd

	// Windows uses cmd.exe
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", input)
	} else {
		// Unix-based systems use the user's default shell
		shell := os.Getenv("SHELL")
		if shell == "" {
			shell = "/bin/sh"
		}
		cmd = exec.Command(shell, "-c", input)
	}

	// Ensure the command runs in the current working directory
	cmd.Dir = currentDir

	// Buffers to capture stdout and stderr
	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	// Execute the command
	err := cmd.Run()

	// Remove trailing newline characters for cleaner output
	stdout := strings.TrimRight(out.String(), "\n")
	stderr := strings.TrimRight(errOut.String(), "\n")

	// If an error occurred, display stderr if available
	if err != nil {
		if stderr != "" {
			return stderr
		}
		return fmt.Sprintf("Error: %v", err)
	}

	// If no output was produced, notify the user
	if stdout == "" && stderr == "" {
		return "[command executed with no output]"
	}

	// If both stdout and stderr exist, return both
	if stdout != "" && stderr != "" {
		return stdout + "\n" + stderr
	}

	// Return whichever output exists
	if stdout != "" {
		return stdout
	}

	return stderr
}

// appendOutput adds the executed command and its result to the terminal output widget.
func appendOutput(output *widget.RichText, command, result string) {

	output.Segments = append(output.Segments,
		// Display the command prompt and command
		&widget.TextSegment{
			Text: fmt.Sprintf("%s$ %s", currentDir, command),
			Style: widget.RichTextStyle{
				ColorName: theme.ColorNamePrimary,
			},
		},

		// Display the command result
		&widget.TextSegment{
			Text: result,
			Style: widget.RichTextStyle{
				ColorName: theme.ColorNameForeground,
			},
		},
	)

	// Refresh the widget so new output appears
	output.Refresh()
}

// clearOutput resets the terminal display by removing all text segments.
func clearOutput(output *widget.RichText) {

	output.Segments = []widget.RichTextSegment{
		&widget.TextSegment{
			Text: "",
			Style: widget.RichTextStyle{
				ColorName: theme.ColorNameForeground,
			},
		},
	}

	output.Refresh()
}

// main initializes and runs the DreShellGUI application.
func main() {

	// Determine the starting working directory
	startDir, err := os.Getwd()
	if err != nil {
		startDir = "."
	}
	currentDir = startDir

	// Create the Fyne application
	myApp := app.New()

	// Apply the default light theme
	myApp.Settings().SetTheme(theme.LightTheme())

	// Create the main window
	myWindow := myApp.NewWindow("DreShell GUI")

	// Set initial window size
	myWindow.Resize(fyne.NewSize(800, 500))

	// Terminal output area using RichText for colored segments
	output := widget.NewRichText(
		&widget.TextSegment{
			Text: "Welcome to DreShell GUI\nType 'help' to see available commands.",
			Style: widget.RichTextStyle{
				ColorName: theme.ColorNameForeground,
			},
		},
	)

	// Enable text wrapping
	output.Wrapping = fyne.TextWrapWord

	// Scroll container for terminal output
	scrollOutput := container.NewScroll(output)

	// Input field where the user types commands
	input := widget.NewEntry()
	input.SetPlaceHolder("Type a command and press Enter...")

	// Function executed when the user runs a command
	runAction := func() {

		command := strings.TrimSpace(input.Text)

		// Ignore empty input
		if command == "" {
			return
		}

		// Clear command resets the terminal output
		if command == "clear" {
			clearOutput(output)
			input.SetText("")
			return
		}

		// Execute command and capture result
		result := runCommand(command)

		// Append command and result to the terminal display
		appendOutput(output, command, result)

		// Clear input field after execution
		input.SetText("")

		// Automatically scroll to the bottom
		scrollOutput.ScrollToBottom()
	}

	// Trigger command execution when Enter is pressed
	input.OnSubmitted = func(string) {
		runAction()
	}

	// Layout container holding the command input
	topBar := container.NewBorder(nil, nil, nil, nil, input)

	// Main layout: input on top, terminal output below
	content := container.NewBorder(topBar, nil, nil, nil, scrollOutput)

	// Attach layout to the window
	myWindow.SetContent(content)

	// Launch the GUI application
	myWindow.ShowAndRun()
}
