package main

import (
	"fmt"
	//"os"
	"io/ioutil"
	"net/http"
	"io"
	
)


func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Azure AD Pod Identity + Keyvault Flex Volume mount demo")
}

func getFileContent(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/Users/marcintr/Work/demon.yaml")
	if err != nil {
		io.WriteString(w, "Provided file is not mounted, check your Identity")
	} else {
		io.WriteString(w, string(data))
	}
	
}

func main() {
	data, err := ioutil.ReadFile("/Users/marcintr/Work/demo.yaml")
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Print(string(data))
	http.HandleFunc("/", hello)
	http.HandleFunc("/file", getFileContent)
	http.ListenAndServe(":8080", nil)
}
