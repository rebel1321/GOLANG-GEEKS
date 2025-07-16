package main

import "fmt"

type customError struct {
	Code    int
	Message string
}

func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func check(code int) (string,error) {
	if code == 404 {
		return "", &customError{
			Code:    404,
			Message: "Not Found",
		}
	}
	return "success", nil
}

func main() {
	if msg,err:=check(404);err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(msg)
	}
}