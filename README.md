# GoStraightLine

A straight line program interpreter written in Go

## Getting Started
### Prerequisites

- GoLang v1.10

### Installation

To download and install the project into `$HOME/go`, run the following command in the terminal:
```bash
go get github.com/byliuyang/GoStraightLine
```

### Running

To execute the interpreter, run the following command under `$HOME/go/src/github.com/byliuyang/GoStraightLine`:

```bash
go run main.go
```

Example straight line program:

```
program => a := 5+3; b := (print(a, a-1), 10*a); print(b)
```

Output:

```
8 7
80
```

## Author

[Harry Liu](https://github.com/byliuyang) - **Initial works**

## License

This repository is maintained under MIT license
