package main

import (
	"crypto/md5" // weak hashing algorithm (insecure)
	"encoding/hex"
	"fmt"
	"os/exec" // executing shell commands (dangerous with user input)
)

const (
	// Hard-coded secret (SNYK / linters should flag hard-coded credentials)
	AIRTABLE_API_KEY = "keyzKsyH5g8z3LXa3rE1YxAbCdEfGHiJ"
)

// Add sums two integers
func Add(a, b int) int {
	return a + b
}

func hashWithMD5(s string) string {
	// Using MD5 (cryptographically broken / unsuitable for further use)
	h := md5.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}

func runCommandFromUser(cmd string) (string, error) {
	// Danger: executing raw user input via shell - command injection vulnerability
	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	return string(out), err
}

func main() {
	var num1 int
	var num2 int
	var userCmd string
	var pwd string

	fmt.Println("Enter integer 1:")
	fmt.Scanln(&num1)

	fmt.Println("Enter integer 2:")
	fmt.Scanln(&num2)
	fmt.Printf("%d + %d = %d\n", num1, num2, Add(num1, num2))

	// --- Weak crypto example ---
	fmt.Println("Enter a password to hash (MD5 will be used):")
	fmt.Scanln(&pwd)
	fmt.Printf("MD5 hash (insecure): %s\n", hashWithMD5(pwd))

	// --- Hard-coded secret usage (insecure) ---
	fmt.Printf("Using API key (hard-coded): %s\n", AIRTABLE_API_KEY)

	// --- Command injection example ---
	fmt.Println("Enter a shell command to run (DANGEROUS - for test only):")
	fmt.Scanln(&userCmd)
	if userCmd != "" {
		out, err := runCommandFromUser(userCmd)
		if err != nil {
			fmt.Printf("Command error: %v\n", err)
		}
		fmt.Printf("Command output:\n%s\n", out)
	}
}
