package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

const (
	//task statuses
	def       = "To Do"
	ongoing   = "In Progress"
	completed = "Completed"
	remaining = "To Do"

	// file path
	filePath = "tasks.json"
)

type Task struct {
	Id          int
	Description string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}

func storeInFile(tasks []Task) {
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	fmt.Println("JSON data:", string(jsonData))
	if err != nil || jsonData == nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = os.WriteFile(filePath, jsonData, 0777)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}

func (t Task) AddTask() {
	// task := Task{
	// 	Id:          1,
	// 	Description: "Complete the Go project",
	// 	Status:      "In Progress",
	// 	CreatedAt:   "2025-04-05",
	// 	UpdatedAt:   "",
	// }

	var tasks []Task

	data, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	} else {
		json.Unmarshal(data, &tasks)
	}

	if len(tasks) > 0 {
		tasks = append(tasks, t)
		storeInFile(tasks)
	}

}

// delete a task  given its id, refactor it so that it takes a parameter
// and returns the updated task list
func DeleteTask() {
	var tasksHolder []Task
	data, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	} else {
		json.Unmarshal(data, &tasksHolder)
	}

	fmt.Println("Enter the task ID to delete:")
	var id int
	_, err = fmt.Scan(&id)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	deleteIdx := slices.IndexFunc(tasksHolder, func(t Task) bool {
		return t.Id == id
	})

	if deleteIdx == -1 {
		fmt.Println("Task not found")
		return
	}

	lastIdx := len(tasksHolder) - 1
	if deleteIdx == lastIdx {
		tasksHolder = tasksHolder[:lastIdx]
	} else {
		tasksHolder[deleteIdx], tasksHolder[lastIdx] = tasksHolder[lastIdx], tasksHolder[deleteIdx]
		tasksHolder = tasksHolder[:lastIdx]
	}

	storeInFile(tasksHolder)
}

func main() {
	// task := Task{
	// 	Id:          1,
	// 	Description: "Complete the Go project",
	// 	Status:      "In Progress",
	// 	CreatedAt:   "2025-04-05",
	// 	UpdatedAt:   "",
	// }

	// task.AddTask()

	DeleteTask()

}
