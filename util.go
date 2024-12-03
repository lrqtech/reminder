package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	return strings.TrimSpace(input)
}

func ValidateDay(day, maxDay int) bool {
	return day >= 1 && day <= maxDay
}

func HandleExit(sig os.Signal, m *Month) {
	if sig != nil {
		fmt.Println("\nExiting gracefully...")
		m.SaveReminders()
		os.Exit(0)
	}
}

func GetCurrentMonthInfo() (string, int, int) {
	now := time.Now()
	monthName := now.Month().String()
	lastDayOfCurrentMonth := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, -1).Day()
	startDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC).Weekday()
	return monthName, lastDayOfCurrentMonth, int(startDay)
}

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error clearing screen:", err)
	}
}

func HandleReminderAction(action int, day int, reminders []string, m *Month) {
	switch action {
	case 1:
		fmt.Print("Enter reminder: ")
		reminder := ReadInput()
		if reminder != "" {
			m.Reminders.AddReminder(day, reminder)
			fmt.Println("\033[1;32mReminder added!\033[0m")
		} else {
			fmt.Println("\033[1;31mReminder cannot be empty.\033[0m")
		}
	case 2:
		if len(reminders) == 0 {
			fmt.Println("\033[1;31mNo reminders found for this day.\033[0m")
			return
		}
		fmt.Println("Existing reminders:")
		for idx, reminder := range reminders {
			fmt.Printf("%d. %s\n", idx+1, reminder)
		}
	eS:
		fmt.Print("Enter the number of the reminder to edit: ")
		reminderIdx, _ := strconv.Atoi(ReadInput())
		if reminderIdx >= 1 && reminderIdx <= len(reminders) {
			oldReminder := reminders[reminderIdx-1]
			fmt.Print("Enter new reminder: ")
			newReminder := ReadInput()
			if newReminder != "" && m.Reminders.EditReminder(day, oldReminder, newReminder) {
				fmt.Println("\033[1;32mReminder updated!\033[0m")
			} else {
				fmt.Println("\033[1;31mFailed to update reminder.\033[0m")
			}
		} else {
			fmt.Println("\033[1;31mInvalid reminder selection.\033[0m")
			goto eS
		}
	case 3:
		if len(reminders) == 0 {
			fmt.Println("\033[1;31mNo reminders found for this day.\033[0m")
			return
		}
		fmt.Println("Existing reminders:")
		for idx, reminder := range reminders {
			fmt.Printf("%d. %s\n", idx+1, reminder)
		}
	dS:
		fmt.Print("Enter the number of the reminder to delete: ")
		reminderIdx, _ := strconv.Atoi(ReadInput())
		if reminderIdx >= 1 && reminderIdx <= len(reminders) {
			reminderToDelete := reminders[reminderIdx-1]
			if m.Reminders.DeleteReminder(day, reminderToDelete) {
				fmt.Println("\033[1;32mReminder deleted!\033[0m")
			} else {
				fmt.Println("\033[1;31mFailed to delete reminder.\033[0m")
			}
		} else {
			fmt.Println("\033[1;31mInvalid reminder selection.\033[0m")
			goto dS
		}
	default:
		fmt.Println("\033[1;31mInvalid option.\033[0m")
	}
}
