package calendar

import (
	"fmt"
	"time"

	"daily-journal/journal"
	"daily-journal/util"
)

func showCalendarHelp() {
	fmt.Println()
	fmt.Println("h: left, j: down, k: up, l: right, q: quit")
	fmt.Println("u: previous month, o: next month, i: open selected day journal")
}

func PrintCalendar(year, month, selectedDay int) {
	// Get the first day of the specified month
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)

	// Get the last day of the specified month
	lastDay := firstDay.AddDate(0, 1, -1)

	// Print the header
	fmt.Printf("%s %d\n", firstDay.Month().String(), year)
	fmt.Println()
	fmt.Println("Sun Mon Tue Wed Thu Fri Sat")

	// Print the days
	for day := firstDay; !day.After(lastDay); day = day.AddDate(0, 0, 1) {
		// Print spaces for the days before the first day of the month
		journalPath := util.CreateJournalPath(day.Year(), int(day.Month()), day.Day())
		if day.Weekday() != time.Sunday && day.Day() == 1 {
			for i := 0; i < int(day.Weekday()); i++ {
				fmt.Print("    ")
			}
		}

     	// Print the day with red color if it's the selected day
		if day.Day() == selectedDay {
			fmt.Printf("\x1b[31;1m%3d\x1b[0m ", day.Day()) 
		// Print the days with green color if there is a journal for the day
		} else if (util.DirectoryFileExists(journalPath) && !journal.IsJournalEmpty(journalPath)) {
			fmt.Printf("\x1b[32;1m%3d\x1b[0m ", day.Day())
		} else {
			fmt.Printf("%3d ", day.Day())
		}


		// Start a new line for the next week
		if day.Weekday() == time.Saturday {
			fmt.Println()
		}
	}

	fmt.Println("")
	showCalendarHelp()
}

func getLastDayOfMonth(year, month int) int {
	// get the first day of the specified month
	firstOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)

	// get the last day of the specified month
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	return lastOfMonth.Day()
}

func HandleCalendar(year, month, selectedDay int) {

	c, err := util.GetOneLetterInput()
	isValidMove := false

	if err != nil {
		fmt.Println(err)
		return
	}
	
	switch c{
	case "h":
		if (selectedDay > 1) {
			selectedDay--
			isValidMove = true
		}
	case "l":
		if (selectedDay < getLastDayOfMonth(year, month)) {
			selectedDay++
			isValidMove = true
		}
	case "j":
		if (selectedDay <= getLastDayOfMonth(year, month) - 7) {
			selectedDay += 7
			isValidMove = true
		}
	case "k":
		if (selectedDay > 7) {
			selectedDay -= 7
			isValidMove = true
		}
	case "u":
		if (month > 1) {
			month--
		} else {
			month = 12
			year--
		}
		selectedDay = 1
		isValidMove = true
	case "o":
		if (month < 12) {
			month++
		} else {
			month = 1
			year++
		}
		selectedDay = 1
		isValidMove = true
	case "i":
		journal.OpenJournal(year, month, selectedDay)
	case "q":
		return
	default:
		isValidMove = false
	}

	if isValidMove {
		util.ClearScreen()
		PrintCalendar(year, month, selectedDay)
	}

	HandleCalendar(year, month, selectedDay)
	
}