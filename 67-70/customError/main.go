package main

import "fmt"

type CustomError struct {
	message string
	code    int
}

func (cu *CustomError) Error() string {
	return cu.message
}

func login(password string) error {
	if password != "123456" {
		return &CustomError{
			message: "password don't match",
			code:    401,
		}
	}
	return nil
}

func main() {
	// fmt.Println("custom error")
	err := login("1244")
	if err != nil {
		// fmt.Println(err, "code", err.(*CustomError).code)
		// customErr := err.(*CustomError)

		// fmt.Println(customErr)

		if cuserror, ok := err.(*CustomError); ok {
			fmt.Println(cuserror.code, cuserror.message)
		}

		// fmt.Println("main func")
	}

	users := map[int]string{
		1: "sujon",
	}

	fmt.Println(users[1])
}
