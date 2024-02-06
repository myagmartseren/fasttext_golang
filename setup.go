package fasttext

import (
	"fmt"
	"os/exec"
	"runtime"
)

func isLibfasttextInstalled() bool {
	// Check if libfasttext is installed by attempting to run a command
	// that verifies its presence. Adjust this according to your system.
	cmd := exec.Command("libfasttext", "--version")
	err := cmd.Run()
	return err == nil
}

func installLibfasttext() {
	fmt.Println("Installing libfasttext...")
	// Add your installation logic here.
	// This might involve downloading the library, compiling it, and
	// installing it in a location where your package can find it.
	// Ensure it's compatible with different operating systems.
	switch runtime.GOOS {
	case "linux":
		// Linux installation commands
	case "darwin":
		// macOS installation commands
	case "windows":
		// Windows installation commands
	default:
		fmt.Println("Unsupported operating system for libfasttext installation.")
	}
}
