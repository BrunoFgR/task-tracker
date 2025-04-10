# Task Tracker
Task tracker is a CLI app to track your tasks and manage your to-do list.

Sample solution for the [task-tracker](https://roadmap.sh/projects/task-tracker) challenge from [roadmap.sh](https://roadmap.sh/).

## How to run

Clone the repository and run the following command:

```bash
git clone https://github.com/BrunoFgR/task-tracker
cd task-tracker
```

Run the following command to build and run the project:

```bash
# To make the build script executable
chmod +x build.sh

# To build the project
./build.sh

# To see the list of available commands
./task-tracker --help

# To add a task
./task-tracker add "Buy groceries"

# To update a task
./task-tracker update 1 "Buy groceries and cook dinner"

# To delete a task
./task-tracker delete 1

# To mark a task as in progress/done/todo
./task-tracker mark-in-progress 1
./task-tracker mark-done 1
./task-tracker mark-todo 1

# To list all tasks
./task-tracker list
./task-tracker list done
./task-tracker list todo
./task-tracker list in-progress
```
