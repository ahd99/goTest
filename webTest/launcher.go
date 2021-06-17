package main

import (
	"fmt"

	"aheydari.ir/gotest/webtest/wiki"
	//"aheydari.ir/goutils/utility/system"
	"github.com/a-heydari/goutils/system"
)

func main()  {
	//save_load_file_test()
	fmt.Println(system.SystemInfo())

	startServer()
}

func startServer() {
	//wiki.StartGeneralServer(8080)
	wiki.StartFileViewServer()
}

// save_load_file_test test function
func save_load_file_test()  {
	var page *wiki.Page = &wiki.Page{Title: "hello", Body: []byte("This body of hello page")}
	err := page.Save();
	if err != nil {
		fmt.Println("Error saving file ", err)
		return
	}
	
	p, err := wiki.Load("hello")
	if err != nil {
		fmt.Println("Error loading file ", err)
		return
	}
	fmt.Println(string(p.Body))
}