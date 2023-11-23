package util

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"golang.org/x/term"
)

func GetOneLetterInput() (string, error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
    if err != nil {
        return "", err
    }
    defer term.Restore(int(os.Stdin.Fd()), oldState)

    b := make([]byte, 1)
    _, err = os.Stdin.Read(b)
    if err != nil {
		return "", err
    }

	return string(b), nil
}

func CreateDirectory(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		errDir := os.MkdirAll(path, 0755)
		if errDir != nil {
			fmt.Println(err)
		}
	}
}

func CreateJournalPath(year, month, day int) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	journalPath := filepath.Join(homeDir, ".go-journal", fmt.Sprintf("%d", year), fmt.Sprintf("%d-%d", month, day))
	return journalPath
}

func DirectoryFileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "darwin", "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Unsupported operating system for screen clearing")
	}
}