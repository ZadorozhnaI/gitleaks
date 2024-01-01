package main

import (
        "fmt"
        "os"
        "os/exec"
        "strings"
)

func installGitleaks() {
        fmt.Println("Installing gitleaks...")
        cmd := exec.Command("sh", "-c", "curl -sSfL https://install.goreleaser.com/github.com/zricethezav/gitleaks.sh | sh")
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        err := cmd.Run()
        if err != nil {
                fmt.Println("Error installing gitleaks:", err)
                os.Exit(1)
        }
        fmt.Println("Gitleaks installed successfully.")
}

func runGitleaks() {
        cmd := exec.Command("gitleaks")
        output, err := cmd.CombinedOutput()
        if err != nil {
                fmt.Println("Error: Secrets found in the code. Commit rejected.")
                fmt.Println(string(output))
                os.Exit(1)
        }
        fmt.Println("No secrets found. Commit allowed.")
}
func main() {
        // Check if gitleaks is installed
        _, err := exec.LookPath("gitleaks")
        if err != nil {
                // Gitleaks not found, check if enable option is set
                output, _ := exec.Command("git", "config", "--get", "hooks.gitleaks.enable").CombinedOutput()
                enableGitleaks := strings.TrimSpace(string(output))
                if strings.ToLower(enableGitleaks) == "true" {
                        installGitleaks()
                }
        }

        runGitleaks()
}
