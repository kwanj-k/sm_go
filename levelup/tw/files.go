package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func file() {

	mydata := []byte("All the data I wish to write to a file\n")

	// the WriteFile method returns an error if unsuccessful
	err := ioutil.WriteFile("test.data", mydata, 0777)
	// handle this error
	if err != nil {
		// print it out
		fmt.Println(err)
	}

	data, err := ioutil.ReadFile("test.data")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))

	f, err := os.OpenFile("test.data", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString("new data that wasn't there originally\n"); err != nil {
		panic(err)
	}

	data, err = ioutil.ReadFile("test.data")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))
}