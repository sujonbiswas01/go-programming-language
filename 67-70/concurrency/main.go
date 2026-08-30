package main

import (
	"fmt"
	"sync"
	"time"
)

// একাধিক কাজকে এমনভাবে manage করা, যাতে কাজগুলো একই সময়ে এগোতে পারে।
// Concurrent = একই সময়ে একাধিক কাজের progress হওয়া।

var wg sync.WaitGroup
var uploadUrl string

func main() {
	var start = time.Now()
	// wg.Add(1)
	// go uploadFile()
	// uploadFile()
	wg.Go(uploadFile)
	// wg.Add(1)
	// go savetoDb()
	wg.Go(savetoDb)
	// wg.Add(1)
	// go sendEmail()
	wg.Go(sendEmail)
	wg.Wait()
	fmt.Println("file url : ", uploadUrl)
	fmt.Println("time taken", time.Since(start))

}

func uploadFile() {

	fmt.Println("uploading file...")
	time.Sleep(1 * time.Second)
	fmt.Println("File upload done")
	uploadUrl = "https://sujonbiswas01/image"

	// return uploadUrl
	// wg.Done()
}

func savetoDb() {
	fmt.Println("savetodb.....")
	time.Sleep(1 * time.Second)
	fmt.Println("save to db done")
	// wg.Done()
}

func sendEmail() {
	fmt.Println("sending email....")
	time.Sleep(1 * time.Second)
	fmt.Println("email send done")
	// wg.Done()
}
