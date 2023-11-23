package journal

import (
	"bufio"
	"daily-journal/models"
	"daily-journal/util"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func writeJournal(path string, jounal models.Journal) {
	output := ""
	output += "# TODO\n"

	for _, todo := range jounal.TodayTodoList {
		output += fmt.Sprintf("%s\n", todo)
	}

	output += "# DONE\n"

	for _, done := range jounal.TodayDoneList {
		output += fmt.Sprintf("%s\n", done)
	}

	output += "# NOTES\n"
	for _, note := range jounal.TodayNotes {
		output += fmt.Sprintf("%s\n", note)
	}

	output += "# TOMORROW\n"
	for _, todo := range jounal.TomoTodoList {
		output += fmt.Sprintf("%s\n", todo)
	}

	// Remove the last blank line
	output = strings.TrimSuffix(output, "\n")

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.WriteString(output)

}

func parseJournal(path string) models.Journal {
	journal := models.Journal{}
	sharpCount := 0
	
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return journal
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Remove blank lines
		line = strings.TrimSpace(line)
		if (strings.HasPrefix(line, "#")) {
			sharpCount++
			continue
		} else {
			switch sharpCount {
			case 1:
				journal.TodayTodoList = append(journal.TodayTodoList, line)
			case 2:
				journal.TodayDoneList = append(journal.TodayDoneList, line)
			case 3:
				journal.TodayNotes = append(journal.TodayNotes, line)
			case 4:
				journal.TomoTodoList = append(journal.TomoTodoList, line)
			}
		}
	}

	return journal

}

func initJournal(path string) {
	// Create empty journal data
	journal := models.Journal{
		TodayTodoList: []string{},
		TodayDoneList: []string{},
		TodayNotes:    []string{},
		TomoTodoList:  []string{},
	}

	writeJournal(path, journal)
}

func IsJournalEmpty(path string) bool {
	// Returns false if the number of lines in the journal is less than 5
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return false 
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	return lineCount < 5
}

func OpenJournal(year, month, day int) {
	// Check if the directory for the year exists
	homeDir, err := os.UserHomeDir()
	var journalPath string
	isJournalExists := false
	if err != nil {
		fmt.Println(err)
		return
	}

	yearDir := filepath.Join(homeDir, ".go-journal", fmt.Sprintf("%d", year))
	if !util.DirectoryFileExists(yearDir) {
		util.CreateDirectory(yearDir)
	} 
	// Check if the journal file exists in the directory
	journalPath = filepath.Join(yearDir, fmt.Sprintf("%d-%d", month, day))
	if util.DirectoryFileExists(journalPath) {
		isJournalExists = true
	}
	

	if (!isJournalExists) {
		// Create the journal file
		file, err := os.Create(journalPath)
		if err != nil {
			fmt.Println(err)
			return
		}
		file.Close()

		initJournal(journalPath)
		fmt.Println("Created journal file: " + journalPath)
	}
}