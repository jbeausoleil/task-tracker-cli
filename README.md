**Project Requirements**

The application should run from the command line, accept user actions and inputs as arguments, and store the tasks in a JSON file. The user should be able to:

- Add, Update, and Delete tasks
- Mark a task as in progress or done
- List all tasks
- List all tasks that are done
- List all tasks that are not done
- List all tasks that are in progress 

**Project Constraints**
- Use positional arguments in command line to accept user inputs.
- Use a JSON file to store the tasks in the current directory.
- The JSON file should be created if it does not exist.
- Use the native file system module of your programming language to interact with the JSON file.
- Do not use any external libraries or frameworks to build this project.
- Ensure to handle errors and edge cases gracefully.

**Directory Structure**
```
taskcli/
├── cmd/                   # Commands like add, list, etc.
│   ├── add.go             # Handles "add" command
│   ├── list.go            # Handles "list" command
├── internal/              # Business logic, task management
│   └── tasks/
│       └── manager.go     # TaskManager, file I/O, logic
├── data/                  # (Optional) Data storage like JSON or text files
│   └── tasks.json
├── go.mod
├── go.sum
└── main.go                # Parses CLI args and dispatches commands
```

