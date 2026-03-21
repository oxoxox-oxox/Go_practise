package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID        int
	Content   string
	Completed bool
	CreatedAt time.Time
}

type TaskList struct {
	Tasks  []Task
	nextID int
}

const fileName = "tasks.json"

// load tasks from file
func loadTasks() (*TaskList, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		//if file not exist, return empty task list
		if os.IsNotExist(err) {
			return &TaskList{Tasks: []Task{}, nextID: 1}, nil
		}
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	//calculate nextID: max ID + 1
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return &TaskList{Tasks: tasks, nextID: maxID + 1}, nil
}

// save task list to file
func (tl *TaskList) save() error {
	data, err := json.MarshalIndent(tl.Tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func (tl *TaskList) add(content string) {
	task := Task{
		ID:        tl.nextID,
		Content:   content,
		Completed: false,
		CreatedAt: time.Now(),
	}
	tl.Tasks = append(tl.Tasks, task)
	tl.nextID++
	if err := tl.save(); err != nil {
		fmt.Println("Error saving tasks:", err)
	} else {
		fmt.Println("Task added:")
	}
}

func (tl *TaskList) list() {
	if len(tl.Tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("ID\tstatus\tcontent")
	fmt.Println("--\t----\t\t----")
	for _, t := range tl.Tasks {
		status := "☐"
		if t.Completed {
			status = "√"
		}
		fmt.Printf("%d\t%s\t\t%s\n", t.ID, status, t.Content)
	}
}

func (tl *TaskList) complete(id int) {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			tl.Tasks[i].Completed = true
			if err := tl.save(); err != nil {
				fmt.Println("Error saving tasks:", err)
			} else {
				fmt.Println("Task marked as completed:")
			}
			return
		}
	}
	fmt.Println("Task not found")
}

func (tl *TaskList) delete(id int) {
	for i, t := range tl.Tasks {
		if t.ID == id {
			//delete the element in the slice\
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			if err := tl.save(); err != nil {
				fmt.Println("delete failed", err)
			} else {
				fmt.Println("Task deleted:")
			}
			return
		}
	}
	fmt.Println("Task not found")
}

func main() {
	//load tasks from file
	tl, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	//If no args, show help
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	cmd := os.Args[1]
	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("method: todo add <task content>")
			return
		}
		content := os.Args[2]
		tl.add(content)

	case "list":
		tl.list()

	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("method: todo complete <task id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("task ID must be a number")
			return
		}
		tl.complete(id)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("method: todo delete <task id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("task ID must be a number")
			return
		}
		tl.delete(id)

	default:
		fmt.Println("unknow command")
		printHelp()
	}
}

func printHelp() {
	fmt.Println(`todo tool - usage:
	todo add <task content>   add a new task
	todo list 				  list all tasks
	todo complete <task id>   mark a task as completed
	todo delete <task id>     delete a task`)
}
