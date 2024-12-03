package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func interactiveMode() {
	monthName, monthDays, startDay := GetCurrentMonthInfo()
	CurrentMonth := &Month{
		Name:      monthName,
		MonthDays: monthDays,
		StartDay:  startDay,
		Reminders: &LinkedList{},
	}
	CurrentMonth.LoadReminders()
	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV)
	go func() {
		select {
		case sig := <-sigChan:
			HandleExit(sig, CurrentMonth)
		}
	}()
	fmt.Printf("\n\033[1;34m%s Reminder Manager - %d\033[0m\n", CurrentMonth.Name, time.Now().Year())
	CurrentMonth.PrintCalendar()
	for {
	hStep1:
		fmt.Printf("\nEnter a day (1-%d) to add, edit, or delete a reminder, or 0 to exit: ", CurrentMonth.MonthDays)
		day, err := strconv.Atoi(ReadInput())
		if err != nil {
			fmt.Println("\033[1;31mInvalid input, please enter a number!\033[0m")
			goto hStep1
		}
		if day == 0 {
			CurrentMonth.SaveReminders()
			fmt.Println("\033[1;32mReminders saved successfully!\033[0m")
			return
		}
		if !ValidateDay(day, CurrentMonth.MonthDays) {
			fmt.Printf("\033[1;31mInvalid day. Please enter a day between 1 and %d.\033[0m\n", CurrentMonth.MonthDays)
			continue
		}
		fmt.Println("1. Add a reminder")
		fmt.Println("2. Edit a reminder")
		fmt.Println("3. Delete a reminder")
	hStep2:
		fmt.Printf("Choose an option: ")
		option, err := strconv.Atoi(ReadInput())
		if err != nil {
			fmt.Println("\033[1;31mInvalid input, please enter a number!\033[0m")
			goto hStep2
		}
		reminders := CurrentMonth.Reminders.GetReminders(day)
		HandleReminderAction(option, day, reminders, CurrentMonth)
		time.Sleep(1 * time.Second)
		ClearScreen()
		fmt.Printf("\n\033[1;34m%s Reminder Manager - %d\033[0m\n", CurrentMonth.Name, time.Now().Year())
		CurrentMonth.PrintCalendar()
	}
}
