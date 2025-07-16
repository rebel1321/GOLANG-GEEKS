Best practice for error handling in Go
1.Check error immediately
res,err:=doSomething()
if err!=nil{
	log.Println("Error:", err)
	return
}


2.Return error instead of panic
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

3.Wrap errors with context
func loadConfig() error{
	err:=readFile()
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}
	return nil
}

4.Use sentinel errors sparingly
var ErrNotFound = errors.New("not found")

func getResource(id int) error{
	if id==0{
		return ErrNotFound
	}
	return nil
}

5.Leverage Custom Error Types
type MyError struct {
	code    int    // Error code
	Message string // Error message
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.code, e.Message)
}

func process() error {
	return &MyError{
		code:    404,
		Message: "Resource not found",
	}
}

6.Use errors.Is and errors.As for error inspection
7.Use defer with recover for panic recovery

8.Do not hide errors

9.Handle errors in Goroutines
When launching goroutines ,handles or propogates errors back to the main routine using channels


10.Return nil for success
11.Avoid overuse of panic
