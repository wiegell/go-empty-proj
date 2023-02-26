package main

import "fmt"

type teststruct struct {
	 content string
}

func helloworld() string {
	testvar := teststruct{content: "Hello World"}
	testvar2 := 2;
	return testvar.content + fmt.Sprint(testvar2) + "!!"
}

func main() {
	fmt.Println(helloworld())
}
