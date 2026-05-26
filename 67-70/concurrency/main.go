package main

import (
	"fmt"
	"time"
)

func main() {
	var start = time.Now()

	go uploadFile()
	go savetoDb()
	go sendEmail()
	time.Sleep(2 * time.Second)
	fmt.Println("time taken", time.Since(start))

}

func uploadFile() {
	fmt.Println("uploading file...")
	time.Sleep(1 * time.Second)
	fmt.Println("File upload done")
}

func savetoDb() {
	fmt.Println("savetodb.....")
	time.Sleep(1 * time.Second)
	fmt.Println("save to db done")
}

func sendEmail() {
	fmt.Println("sending email....")
	time.Sleep(1 * time.Second)
	fmt.Println("email send done")
}
