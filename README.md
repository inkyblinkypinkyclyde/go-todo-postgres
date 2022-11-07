# Go To-Do app

This is a simple todo app built in order to learn how to code in go.

### Built with
* Go
* PostgreSQL
* Fiber

## Getting Started

To get started first clone the git repository:

```bash
git clone https://github.com/inkyblinkypinkyclyde/go-todo-postgres

```


Then cd into the directory, create the database and connect it to the provided schema:

```bash
cd go-todo-postgres 
createdb go-todo
psql -d go-todo -f db/go-todo.sql
```

Then run the app with:


```bash
go run server.go
```

You can then visit 

http://localhost:3000 to view the app

## Planned updates

* Adding multiple users
* Adding multiple projects for each user