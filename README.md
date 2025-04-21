# Task Tracker CLI

## Overview
**Task Tracker CLI** is a simple command-line tool written in Go that helps you manage your tasks efficiently.  
You can add, list, and delete tasks with ease — all saved locally in a lightweight JSON file.

This project is ideal for practicing Go development, file I/O operations, and structuring CLI tools without using heavy external libraries.

## Project

The application should run from the command line, accept user actions and inputs as arguments, and store the tasks in a JSON file. The user should be able to:

- AddTask, Update, and Delete tasks
- Mark a task as in progress or done
- List all tasks
- List all tasks that are done
- List all tasks that are not done
- List all tasks that are in progress

**Constraints**

- Use positional arguments in command line to accept user inputs.
- Use a JSON file to store the tasks in the current directory.
- The JSON file should be created if it does not exist.
- Use the native file system module of your programming language to interact with the JSON file.
- Do not use any external libraries or frameworks to build this project.
- Ensure to handle errors and edge cases gracefully.

**Directory Structure**
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

---

## Prerequisites

- **Go 1.20+** installed
- (Optional) **Make** if you want to use provided `Makefile` shortcuts

You can check your Go version with:

```bash
go version
```

---

## Installation

First, clone the repository:

```bash
git clone https://github.com/jbeausoleil/task-tracker-cli.git
cd task-tracker-cli
```

### Building the project

You can build the binary manually:

```bash
go build -o task-cli
```

Or use the `Makefile`:

```bash
make build
```

This will create a `task-cli` binary in the project root.

---

## Usage

Once built, you can start using the CLI:

### Adding a task

```bash
./task-cli add "Buy groceries"
```

### Listing all tasks

```bash
./task-cli list
```

### Deleting a task

```bash
./task-cli delete <task_id>
```

> 💡 **Note:** Each task has a unique ID. You can find the task ID by running `list`.

---

## Configuration

- The tool automatically creates and uses a file (like `tasks.json`) to store tasks.
- No external databases, servers, or API keys required.

If you want to specify a custom file location, you can set a flag or environment variable (depending on if/when you extend the project).

---

## Make Commands

| Command       | Description              |
| ------------- | ------------------------- |
| `make build`  | Build the CLI binary       |
| `make clean`  | Remove the binary          |
| `make run`    | Build and run the project  |

---

## Development

If you want to modify or extend the CLI:

- Code is structured cleanly inside `internal/` packages
- All task-related logic is separated for easy maintainability
- Follow Go best practices (e.g., error handling, small functions)

Typical dev loop:

```bash
make clean
make build
./task-cli <command>
```

---

## Troubleshooting

- **Error:** `permission denied` when running `./task-cli`
    - Run: `chmod +x task-cli`
- **Error:** `command not found`
    - Make sure you're in the correct directory or provide the relative path `./task-cli`
- **JSON file issues:**
    - If the storage file is corrupted, delete it manually. A new one will be created automatically.

---

## Roadmap Ideas

- Add due dates and priorities
- Add task categories or tags
- Export/import tasks
- More advanced filtering (e.g., completed vs pending)

---

## License

This project is open source under the [MIT License](LICENSE).

---

# Quick Demo

```bash
$ ./task-cli add "Learn Go"
Task added!

$ ./task-cli list
1. [2025-04-20] Learn Go

$ ./task-cli delete 1
Task deleted!
```
