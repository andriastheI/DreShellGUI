# DreShellGUI

**Author:** Andrias Zelele\
**Language:** Go\
**GUI Toolkit:** Fyne\
**Project Type:** Graphical Shell Interface

------------------------------------------------------------------------

## Overview

DreShellGUI is a graphical terminal interface written in **Go** using
the **Fyne GUI framework**.\
It allows users to run many common system commands from a graphical
window while displaying the output in a terminal‑style view.

Instead of using a traditional terminal, DreShellGUI provides:

-   A **command input field**
-   A **scrollable output terminal**
-   Support for **common shell commands**
-   A few **built‑in commands** for convenience

Commands are executed through the **system shell**:

  Operating System   Shell Used
  ------------------ --------------------------
  Linux              User shell (bash/zsh/sh)
  macOS              User shell (zsh/bash)
  Windows            cmd.exe

------------------------------------------------------------------------

# Features

### Built‑in Commands

  Command   Description
  --------- -------------------------------------
  `help`    Displays help information
  `clear`   Clears the terminal output
  `exit`    Closes the DreShellGUI application
  `pwd`     Shows the current working directory
  `cd`      Changes the working directory

### System Commands

Most normal shell commands work:

Example:

    ls
    echo Hello World
    mkdir test
    cat file.txt

Commands run through the system shell just like a normal terminal.

------------------------------------------------------------------------

# Commands That Are Disabled

Some programs require a **fully interactive terminal (TTY)** and would
freeze the GUI.\
These commands are intentionally blocked:

    sudo
    su
    passwd
    ssh
    nano
    vim
    top
    htop
    less
    more
    man
    python
    node
    mysql
    psql
    sqlite3
    watch

------------------------------------------------------------------------

# Requirements

You need the following installed:

### 1. Go

Go version **1.20 or newer** is recommended.

Check installation:

    go version

Install Go if needed:

https://go.dev/dl/

------------------------------------------------------------------------

### 2. Fyne Toolkit

Install the Fyne dependency:

    go get fyne.io/fyne/v2

------------------------------------------------------------------------

# Running the Program

### Step 1 --- Clone or download the project

    git clone https://github.com/andriastheI/DreShellGUI
    cd DreShellGUI

or place the `main.go` file inside a folder.

------------------------------------------------------------------------

### Step 2 --- Initialize the Go module (if needed)

    go mod init dreshellgui

------------------------------------------------------------------------

### Step 3 --- Install dependencies

    go mod tidy

------------------------------------------------------------------------

### Step 4 --- Run the program

    go run main.go

The DreShellGUI window should appear.

------------------------------------------------------------------------

# Building the Application

You can compile the program into a standalone executable.

### Linux

    go build

This creates:

    ./DreShellGUI

Run with:

    ./DreShellGUI

------------------------------------------------------------------------

### Windows

    go build

Output:

    DreShellGUI.exe

------------------------------------------------------------------------

### macOS

    go build

Output:

    DreShellGUI

------------------------------------------------------------------------

# Cross‑Platform Builds

Go can build binaries for other systems.

### Build macOS binary from Linux

    GOOS=darwin GOARCH=amd64 go build

### Build Windows binary

    GOOS=windows GOARCH=amd64 go build

------------------------------------------------------------------------

# Example Usage

    pwd
    cd ..
    ls
    echo Hello
    mkdir project

------------------------------------------------------------------------

# Project Structure

    DreShellGUI/
    │
    ├── main.go
    └── README.md

------------------------------------------------------------------------

# Future Improvements

Possible enhancements:

-   Command history (↑ ↓ keys)
-   Colored shell prompt
-   Dark terminal theme
-   Tab command autocomplete
-   Real terminal streaming output
-   PTY support for interactive commands

------------------------------------------------------------------------

# License

This project is for educational and personal use.