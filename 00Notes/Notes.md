# Go Language Tutorial Notes

1. Go is a compiled language.
2. Go is suitable for cloud development.
3. Go is both object-oriented (Yes and No).
4. Missing `try`-`catch` - Use error handling instead.
5. Lexer does a lot of work; even if you don't write a semicolon, it won't give you an error because the lexer provides it.
6. `go mod init name` - Initializes a new module.
7. `fmt` package - Used for formatting.
8. Go automatically manages packages, cleaning or uncleaning the ones you need or don't need.
9. `go run main.go` - Command to run Go code.
10. `go help` - Displays help for Go commands.
11. Go is case-sensitive.
12. `go env GOPATH` - Command to find the Go path.
13. Lexer's job is to understand the grammar of the language, ensuring correct syntax and proper variable definition.
14. In Go, everything is a type.
15. Variable initialization and default values:

    ```go
    var number int
    fmt.Println(number)
    // Output will be 0; if not initialized, it does not give garbage value.
    // In the case of a string, it gives an empty space.
    ```

16. Inside any method, you can use the `:=` operator, but not outside.
17. `const Key = "jfhdjhfjd"` - Public constant (capital letter indicates public).
18. `bufio` - Package for buffered I/O.
19. `comma ok` or `comma err` - Used in multiple return values.
20. Input is of type string.
21. Time format: `fmt.Println(Time.Format("01-02-2006 15:04:05 Monday"))`.
22. `go env` - Displays Go environment information.
23. Building applications for different OS in PowerShell and Bash.

    ```powershell
    $env:GOOS = "windows"; go build
    $env:GOOS = "linux"; go build
    $env:GOOS = "darwin"; go build
    ```

    ```bash
    GOOS=darwin go build
    ```

24. Memory Management:

    - Memory allocation and deallocation happen automatically.
    - Methods `new()` and `make()` for memory allocation.
    - `new()`: Allocates memory but does not initialize.
    - `make()`: Allocates and initializes memory.

25. `var ptr *int` - If you initialize a pointer, its default value is `<nil>`.
26. In slice `fruitList[1:3]` - Example slice operation.
27. Removing a value from a slice.

    ```go
    var courses = []string{"reactjs", "python", "ruby", "swift", "c++"}
    courses = append(courses[:index], courses[index+1:]...)
    ```

28. No inheritance, no super or parent in Golang.
29. `++count` is invalid in Go.
30. Variadic function example:

    ```go
    func main() {
        proRes, message := proAdder(2, 4, 3, 4, 5, 1, 2, 3, 34)
        fmt.Println("Pro_Result is ", proRes)
        fmt.Println(message)
    }

    func proAdder(values ...int) (int, string) {
        total := 0
        for _, val := range values {
            total += val
        }
        return total, "Job Done"
    }
    ```

31. `defer` - Executes at the end; multiple defers execute in LIFO order.
32. `response.Body.close()` - Closes the connection.
33. `Response` is a type of `*http.Response`.
34. `omitempty` - If the value is null/nil, don't include that field.
35. Convert response from bytes to string.

**Module Management Commands:**

36. Initialization is crucial for version control and dependencies.
37. `go sum` is built to avoid attacks, checking repository integrity.
38. `go mod verify` - Checks if all modules are verified.
39. Data from the web comes in byte format.
40. `go mod tidy` - Synchronizes `go.mod` file with actual dependencies.
41. `go list all`, `go list -m all`, `go list -m -versions github.com/gorilla/mux` - List commands for dependencies.
42. `go mod why github.com/gorilla/mux` - Identifies module dependencies.
43. `go mod graph` - Displays a graph of all dependencies.
44. `go mod edit` - Edits the module file (optional).
45. `go mod vendor` - Builds the application using dependencies from the vendor directory.

46. **Concurrency vs. Parallelism:**
    - Concurrency: Tasks can start, run, and complete in overlapping time periods.
    - Parallelism: Tasks literally run at the same time (multicore processors).

47. Goroutines help achieve parallelism.
48. Goroutines have a flexible stack size (2KB), while threads managed by the OS have a fixed stack size (1MB).
49. Go slogan: "Do not communicate by sharing memory; instead, share memory by communicating."
50. `sync.Mutex` - Used for locking mechanisms to prevent race conditions.

**Additional Commands:**

51. `go run --race` - Checks for race conditions.
52. Channels are used for communication between Goroutines.
