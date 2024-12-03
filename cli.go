package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func printHelp() {
	fmt.Printf(`
Usage: %s [COMMAND] [ARGUMENTS]

Commands:
  --help                      Show this help message and exit.
  add [DAY] "TEXT"            Add a new reminder on the specified day.
  edit [DAY] [INDEX] "TEXT"   Edit an existing reminder for the specified day and index.
  delete [DAY] [INDEX]        Delete a reminder from the specified day and index.
  show                        Display the current month's calendar with all reminders.

Interactive Mode:
  If no command is provided, the program starts in interactive mode, allowing you to 
  add, edit, and delete reminders through a guided interface.

`, os.Args[0])
}

func handleCLI(args []string) {
	monthName, monthDays, startDay := GetCurrentMonthInfo()
	CurrentMonth := &Month{
		Name:      monthName,
		MonthDays: monthDays,
		StartDay:  startDay,
		Reminders: &LinkedList{},
	}
	CurrentMonth.LoadReminders()
	action := args[0]
	switch action {
	case "add":
		if len(args) < 3 {
			fmt.Println("\033[1;31mPlease provide a day and reminder text to add.\033[0m")
			return
		}
		day, err := strconv.Atoi(args[1])
		if err != nil || !ValidateDay(day, CurrentMonth.MonthDays) {
			fmt.Printf("\033[1;31mInvalid day. Please enter a day between 1 and %d.\033[0m\n", CurrentMonth.MonthDays)
			return
		}
		reminderText := strings.Join(args[2:], " ")
		CurrentMonth.Reminders.AddReminder(day, reminderText)
		fmt.Printf("\033[1;32mReminder added for day %d: %s\033[0m\n", day, reminderText)
		CurrentMonth.SaveReminders()
		fmt.Println("\033[1;32mReminders saved successfully!\033[0m")
	case "edit":
		if len(args) < 4 {
			fmt.Println("\033[1;31mPlease provide the day, reminder index, and new text to edit.\033[0m")
			return
		}
		day, err := strconv.Atoi(args[1])
		if err != nil || !ValidateDay(day, CurrentMonth.MonthDays) {
			fmt.Printf("\033[1;31mInvalid day. Please enter a day between 1 and %d.\033[0m\n", CurrentMonth.MonthDays)
			return
		}
		reminders := CurrentMonth.Reminders.GetReminders(day)
		index, _ := strconv.Atoi(args[2])
		newText := strings.Join(args[3:], " ")
		if index >= 1 && index <= len(reminders) {
			oldReminder := reminders[index-1]
			newReminder := strings.Join(args[3:], " ")
			if newReminder != "" && CurrentMonth.Reminders.EditReminder(day, oldReminder, newText) {
				fmt.Println("\033[1;32mReminder updated!\033[0m")
				CurrentMonth.SaveReminders()
			} else {
				fmt.Println("\033[1;31mFailed to update reminder.\033[0m")
			}
		} else {
			fmt.Println("\033[1;31mInvalid reminder selection.\033[0m")
		}
	case "delete":
		if len(args) < 3 {
			fmt.Println("\033[1;31mPlease provide the day and reminder index to delete.\033[0m")
			return
		}
		day, err := strconv.Atoi(args[1])
		if err != nil || !ValidateDay(day, CurrentMonth.MonthDays) {
			fmt.Printf("\033[1;31mInvalid day. Please enter a day between 1 and %d.\033[0m\n", CurrentMonth.MonthDays)
			return
		}
		reminders := CurrentMonth.Reminders.GetReminders(day)
		index, _ := strconv.Atoi(args[2])
		if index >= 1 && index <= len(reminders) {
			reminderToDelete := reminders[index-1]
			if CurrentMonth.Reminders.DeleteReminder(day, reminderToDelete) {
				fmt.Println("\033[1;32mReminder deleted!\033[0m")
				CurrentMonth.SaveReminders()
			} else {
				fmt.Println("\033[1;31mFailed to delete reminder.\033[0m")
			}
		} else {
			fmt.Println("\033[1;31mInvalid reminder selection.\033[0m")
		}
	case "show":
		fmt.Printf("\n\033[1;34m%s Reminder Manager - %d\033[0m\n", CurrentMonth.Name, time.Now().Year())
		CurrentMonth.PrintCalendar()
	default:
		fmt.Println("\033[1;31mInvalid action. Use '--help' for usage instructions.\033[0m")
		fmt.Println("Please try again.")
	}
}
