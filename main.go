package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func readNumber(filename string) (int, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(strings.TrimSpace(string(data)))
}

func writeNumber(filename string, number int) error {
	return ioutil.WriteFile(filename, []byte(fmt.Sprintf("%d", number)), 0644)
}

func gitCommit() error {
	cmd := exec.Command("git", "add", "number.txt")
	if err := cmd.Run(); err != nil {
		return err
	}

	date := time.Now().Format("2006-01-02")
	commitMsg := fmt.Sprintf("Update number: %s", date)
	cmd = exec.Command("git", "commit", "-m", commitMsg)
	return cmd.Run()
}

func gitPush() error {
	cmd := exec.Command("git", "push")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error pushing to GitHub: %s\n", string(output))
		return err
	}
	fmt.Println("Changes pushed to GitHub successfully.")
	return nil
}

func main() {
	filename := "number.txt"

	number, err := readNumber(filename)
	if err != nil {
		fmt.Printf("Error reading number: %s\n", err)
		os.Exit(1)
	}

	number++
	err = writeNumber(filename, number)
	if err != nil {
		fmt.Printf("Error writing number: %s\n", err)
		os.Exit(1)
	}

	err = gitCommit()
	if err != nil {
		fmt.Printf("Error committing changes: %s\n", err)
		os.Exit(1)
	}

	err = gitPush()
	if err != nil {
		os.Exit(1)
	}
}
