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

```
task-tracker-cli/
├── cmd/               # CLI interface layer: parses args, routes commands
│   ├── root.go        # Entry point and command dispatcher
│   ├── task.go        # CLI handlers for task-related commands (e.g., add, list)
│   └── version.go     # CLI version command
│
├── internal/          # Business logic and domain-specific code (not CLI-aware)
│   └── task/          
│       ├── model.go   # Defines the Task struct and related data models
│       ├── service.go # High-level logic and operations for tasks
│       └── store.go   # In-memory or file-based task storage
│
├── main.go            # Application entry point, delegates to cmd.Execute()
```
