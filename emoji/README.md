# Challenge

This program has a flag inside, but reverse engineering might be an overkill. The program can run in Ubuntu and it will ask you for a password with the following characteristics:

* Its entropy must be 497.246559
* If it includes Unicode emojis, each of them must be unique
* It must *not* contain ASCII characters

# Build the executable

```bash
go build main.go
```