package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup //pointer
var mut sync.Mutex    // pointer

func main() {
	// go greeter("hello")
	// greeter("world")
	websitelist := []string{
		"https://lco.dev",
		"https://google.com",
		"https://github.com",
	}

	for _, web := range websitelist {

		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(500 * time.Millisecond)
// 		fmt.Println(s)

// 	}
// }

func getStatusCode(endpoint string) {

	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)

	}

}
