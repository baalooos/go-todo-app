# go-todo-app

## Description

Simple todo app in go

## Installation

Clone and build from repository.

## Usage

### Init

```bash
go-todo-app init
```

### Add task

```bash
go-todo-app add -n "Task name" -d "Task description"
```

### List tasks

#### List by Id

```bash
go-todo-app list --id 3
```

#### List all tasks

```bash
go-todo-app list --all
```

### Delete task

```bash
go-todo-app delete --id 3
```

### Demo

```bash
go-todo-app
```

## Acknowledgements

### Library used

- [Cobra](https://github.com/spf13/cobra): Library for creating powerful modern CLI applications
- [Go-sqlite](https://github.com/glebarez/go-sqlite): Pure-Go SQLite driver for Golang's native database/sql package.
