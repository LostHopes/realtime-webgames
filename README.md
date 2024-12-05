Simple event-streaming interface for building games in go.

# System requirements

1. golang >= 1.22
2. docker (optional)
3. git (optional)

# Building and running the backend

Running without building
```bash
go run src/main.go
```

Build and then run
```bash
cd src
go build .
chmod +x ./realtime-games
./realtime-games
```

# Questions

1. How to store data? Should I?

# TODO

- [ ] create tic-tac-toe game
- [ ] move login into separate files