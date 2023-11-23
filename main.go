package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"daily-journal/calendar"
	"daily-journal/help"
	"daily-journal/util"
)

func init() {
	util.ClearScreen()
	fmt.Println("Go Daily Journal")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	newDir := filepath.Join(homeDir, ".go-journal")

	if _, err := os.Stat(newDir); os.IsNotExist(err) {
		errDir := os.MkdirAll(newDir, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
		fmt.Println("Created directory: " + newDir)
	}
}

func main() {
	helpFlag := flag.Bool("h", false, "show help")
	flag.Parse()

	if *helpFlag {
		help.BasicHelp()
		return
	}

	now := time.Now()
	year := now.Year()
	month := int(now.Month())
	selectedDay := now.Day()
	calendar.PrintCalendar(year, month, selectedDay)
	calendar.HandleCalendar(year, month, selectedDay)
}
