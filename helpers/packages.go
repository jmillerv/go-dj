package helpers

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/term"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func PackageIsInstalled(packageName string) bool {
	cmd := exec.Command("dpkg", "-s", packageName) // Use dpkg to check package status
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		// If dpkg -s fails, assume the package is not installed
		return false
	}
	return strings.Contains(out.String(), "Status: install ok installed")
}

func InstallPackage(packageName string) error {
	cmd := exec.Command("sudo", "apt", "install", "-y", packageName) // -y for non-interactive install

	// Handle sudo password securely (see previous responses for details)
	// Example using environment variable (still not ideal, but better than hardcoding)
	sudoPass := os.Getenv("SUDO_PASSWORD")
	if sudoPass != "" {
		cmd.Stdin = strings.NewReader(passwordPrompt("enter your password, or exit this and run"+
			" `sudo apt install libasound2-dev libudev-dev pkg-config` in your terminal and then rerun go-dj") + "\n")
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to install %s: %w\nOutput: %s", packageName, err, string(out))
	}
	return nil
}

func passwordPrompt(label string) string {
	var s string
	for {
		_, err := fmt.Fprint(os.Stderr, label+" ")
		if err != nil {
			log.WithError(err).Error("error producing label")
			return ""
		}
		b, _ := term.ReadPassword(int(syscall.Stdin))
		s = string(b)
		if s != "" {
			break
		}
	}
	log.Println()
	return s
}
