package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type dataExample struct {
	UserID    int    `json:"userid"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags"`
}

func main() {
	fmt.Println("JSON DATA")
	// EncodeJson()
	// decodeJson()
	handleDataExample()

}

func EncodeJson() {
	myCourses := []course{
		{"ReactJs Bootcamp", 299, "Udemy", "hello123", []string{"Web", "JS"}},
		{"MERN Bootcamp", 199, "Udemy", "hello123", []string{"Web", "JS"}},
		{"Angular Bootcamp", 499, "Udemy", "hello123", []string{"Web", "JS"}},
	}

	//package this data as json data

	finalJson, _ := json.MarshalIndent(myCourses, "", "\t")
	fmt.Println(string(finalJson))
}

func decodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJs Bootcamp",
		"price": 299,
		"platform": "Udemy",
		"tags": ["Web","JS"]
}
	`)

	var MyCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &MyCourse)
		fmt.Printf("%#v\n", MyCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where u want to add data to key value pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v\n", k, v)
	}

}

func handleDataExample() {
	const myurl = "https://jsonplaceholder.typicode.com/todos/1"
	// var data map[string]interface{}
	var handleData dataExample

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	dataByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(dataByte, &handleData) // storing in to struct  called handledata
	// finalJson, _ := json.MarshalIndent(handleData, "", "\t")
	// for k, v := range data {
	// 	fmt.Printf("the key is %v and the value is %v\n", k, v)
	// }

	fmt.Println(handleData)

	// fmt.Println(string(dataByte))
}
