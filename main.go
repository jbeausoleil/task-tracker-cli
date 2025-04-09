package main

import (
	"fmt"
	"os"
)

// Goal:
//Add, Update, and Delete tasks
//Mark a task as in progress or done
//List all tasks
//List all tasks that are done
//List all tasks that are not done
//List all tasks that are in progress
//-----
//Each task should have the following properties:
//id: A unique identifier for the task
//description: A short description of the task
//status: The status of the task (todo, in-progress, done)
//createdAt: The date and time when the task was created
//updatedAt: The date and time when the task was last updated

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Print(arg)
	}
}

func addTask() {}

func deleteTask() {}

func updateTask() {}

// pass arguments into list tasks for list all (default), list all done, list all not done, list all in progress
func listTasks() {}
