package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

type Node struct {
	Day      int
	Reminder string
	Next     *Node
}

type LinkedList struct {
	Head *Node
}

type Month struct {
	Name      string
	MonthDays int
	StartDay  int
	Reminders *LinkedList
}

type Reminder struct {
	Day      int    `json:"day"`
	Reminder string `json:"reminder"`
}

func (list *LinkedList) AddReminder(day int, reminder string) {
	list.Head = &Node{Day: day, Reminder: reminder, Next: list.Head}
}

func (list *LinkedList) EditReminder(day int, oldReminder, newReminder string) bool {
	for current := list.Head; current != nil; current = current.Next {
		if current.Day == day && current.Reminder == oldReminder {
			current.Reminder = newReminder
			return true
		}
	}
	return false
}

func (list *LinkedList) DeleteReminder(day int, reminder string) bool {
	var prev *Node
	for current := list.Head; current != nil; current = current.Next {
		if current.Day == day && current.Reminder == reminder {
			if prev == nil {
				list.Head = current.Next
			} else {
				prev.Next = current.Next
			}
			return true
		}
		prev = current
	}
	return false
}

func (list *LinkedList) GetReminders(day int) []string {
	var reminders []string
	for current := list.Head; current != nil; current = current.Next {
		if current.Day == day {
			reminders = append(reminders, current.Reminder)
		}
	}
	return reminders
}

func (m *Month) PrintReminders() {
	fmt.Println("--------- Reminders ---------")
	hasReminders := false
	for day := 1; day <= m.MonthDays; day++ {
		reminders := m.Reminders.GetReminders(day)
		if len(reminders) > 0 {
			hasReminders = true
			fmt.Printf("Day %d:\n", day)
			for _, reminder := range reminders {
				fmt.Printf("  + %s\n", reminder)
			}
		}
	}
	if !hasReminders {
		fmt.Println("No reminders for this month.")
	}
	fmt.Println("-----------------------------")
}

func (m *Month) PrintCalendar() {
	fmt.Println(" Sun Mon Tue Wed Thu Fri Sat ")
	position := m.StartDay
	for day := 1; day <= m.MonthDays; day++ {
		reminders := m.Reminders.GetReminders(day)
		if len(reminders) > 0 {
			fmt.Printf("(%2d)", day)
		} else {
			fmt.Printf(" %2d ", day)
		}
		position++
		if position == 7 {
			position = 0
			fmt.Println()
		}
	}
	if position != 0 {
		fmt.Println()
	}
	m.PrintReminders()
}

func (m *Month) SaveReminders() {
	file, err := os.Create("./data.bin")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)
	var reminders []Reminder
	for day := 1; day <= m.MonthDays; day++ {
		for _, reminder := range m.Reminders.GetReminders(day) {
			reminders = append(reminders, Reminder{
				Day:      day,
				Reminder: reminder,
			})
		}
	}
	jsonData, err := json.Marshal(reminders)
	if err != nil {
		fmt.Println("Error encoding reminders to JSON:", err)
		return
	}
	base64Data := base64.StdEncoding.EncodeToString(jsonData)
	if _, err := file.WriteString(base64Data); err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func (m *Month) LoadReminders() {
	file, err := os.Open("./data.bin")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error reading file info:", err)
		return
	}
	fileSize := fileInfo.Size()
	data := make([]byte, fileSize)
	if _, err := file.Read(data); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	jsonData, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		fmt.Println("Error decoding base64 data:", err)
		return
	}
	var reminders []Reminder
	if err := json.Unmarshal(jsonData, &reminders); err != nil {
		fmt.Println("Error decoding JSON data:", err)
		return
	}
	for _, reminder := range reminders {
		m.Reminders.AddReminder(reminder.Day, reminder.Reminder)
	}
}
