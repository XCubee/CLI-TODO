package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var task []Task

type Task struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

func newAddTask() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("📝 Write your task: ")

	// Clear any leftover input
	for scanner.Scan() {
		name := strings.TrimSpace(scanner.Text())

		// Validate input is not empty
		if name == "" {
			fmt.Print("📝 Write your task: ")
			continue
		}

		newTask := Task{
			ID:          int64(len(task) + 1),
			Description: name,
			Status:      false,
		}

		task = append(task, newTask)
		fmt.Printf("✅ Task added successfully: [%d] %s\n", newTask.ID, newTask.Description)
		break
	}
}

func listtask() {
	if len(task) == 0 {
		fmt.Println("NO TASKS MENTIONED")
		return
	}
	for _, task1 := range task {
		fmt.Printf("[%d] %s - Completed: %v\n", task1.ID, task1.Description, task1.Status)
	}
}

func completedTask() {
	var id int
	fmt.Print("Please provide the task ID: ")
	fmt.Scan(&id)

	for i, t := range task {
		if int(t.ID) == id {
			task[i].Status = true
			fmt.Println("Task marked as completed!")
			return
		}
	}
	fmt.Println("Task not found!")
}

func deleteTask() {
	var id int
	fmt.Println("Enter the task ID you want to delete:")
	fmt.Scan(&id)

	index := -1
	for i, t := range task {
		if int(t.ID) == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Task not found!")
		return
	}

	task = append(task[:index], task[index+1:]...) // Correct way to remove an item
	fmt.Println("Task deleted successfully!")
}

func saveTasksToFile() {
	data, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	err = ioutil.WriteFile("tasks.json", data, 0644) // Corrected file permissions
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Tasks saved to file.")
}

func loadTasksFromFile() {
	file, err := ioutil.ReadFile("tasks.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(file, &task)
	if err != nil {
		fmt.Println("Error unmarshaling data:", err)
	}
}

func clearTasks() {
	fmt.Print("Are you sure you want to delete all tasks? (y/n): ")
	var confirm string
	fmt.Scan(&confirm)

	if confirm == "y" || confirm == "Y" {
		task = []Task{} // reinitializes the slice to an empty slice
		fmt.Println("All tasks have been cleared.")
	} else {
		fmt.Println("Operation canceled.")
	}
}

func editTask() {
	fmt.Println("Give the id of the task you want to edit ")
	var id int
	fmt.Scan(&id)
	for i, t := range task {
		if id == int(t.ID) {
			reader := bufio.NewScanner(os.Stdin)
			fmt.Println("Please give the edited task: ")
			reader.Scan()
			name := reader.Text()
			task[i].Description = name // Update the description directly
			fmt.Println("Task edited successfully!")
			return
		}
	}
	fmt.Println("Task not found.")
}

func main() {
	loadTasksFromFile()

	fmt.Println("Welcome to your CLI To-Do List!")

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Completed")
		fmt.Println("4. Delete a task")
		fmt.Println("5. Edit a task")
		fmt.Println("6. Clear all tasks")
		fmt.Println("7. Save Task to file")
		fmt.Println("8. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			newAddTask()
		case 2:
			listtask()
		case 3:
			completedTask()
		case 4:
			deleteTask()
		case 5:
			editTask()
		case 6:
			clearTasks()
		case 7:
			saveTasksToFile()
		case 8:
			saveTasksToFile()
			fmt.Println("Tasks saved successfully!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
