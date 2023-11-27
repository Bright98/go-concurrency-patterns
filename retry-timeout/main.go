package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	var res *http.Response
	var err error
	timeout := time.Duration(3 * time.Second) //retry after 3 sec
	for i := 0; i < 4; i++ {                  //try connecting 4 times
		res, err = http.Get("example.com")
		if err == nil {
			break
		}
		fmt.Println("retrying...")
		<-time.After(timeout)
	}

	if err == nil {
		fmt.Println("Response: ", res)
	} else {
		fmt.Println("Error: ", err)
	}

}
