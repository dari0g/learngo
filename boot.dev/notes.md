## Chapter 1 & 2

- `package main` lets the Go compiler know that we want this code to compile and run as a standalone program, as opposed to being a library that's imported by other programs.
- `import fmt` imports the `fmt` (formatting) package. The formatting package exists in Go's standard library and lets us do things like print text to the console.
- `func main()` defines the `main` function. main is the name of the function that acts as the entry point for a Go program.
- go enforces strong and static typing.
- go runtime cleanups unused memory at runtime.
- syntax for creating a variable (long way) `var <name> <datatype> = <value>`
- syntax for creating a variable (short way) `<name> := <value>`
- when creating a string using different data types we use fmt.Sprintf()
  - %v default
  - %s string
  - %d int
  - %f decimal

## Chapter 3

- function
  - syntax := `func <func-name>(<param1>, <param2>) <return-data00type> { code }`
  -
