package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

//Logentity is the struct to define components of a log
type Logentity struct {
	LogFile   string
	TypeOfLog string
	LogTime   string
	RequestID string
	Container string
	Classname string
	Value     string
}

//Data consists of whole output
var Data []Logentity
var delimeter = "\\[|\\]\\[|\\]:\\[|\\]:"
var reg = regexp.MustCompile(delimeter)

func main() {
	filename := os.Args[1:2]
	data, err := ioutil.ReadFile(filename[0])
	if err != nil {
		fmt.Printf("err value: %v", err)
	}
	filedata := string(data)
	fileslice := strings.Split(filedata, "[console.accounts]")
	for _, v := range fileslice {
		//fmt.Println(v)
		result := reg.Split(v, 7)
		//fmt.Printf(strings.Join(result, "*************"))
		if len(result) == 1 {
			continue
		}
		oneenity := Logentity{"[console.accounts]", result[1], result[2], result[3], result[4], result[5], result[6]}
		Data = append(Data, oneenity)
	}
	res2B, _ := json.Marshal(Data)
	mode := int(0777)

	ioutil.WriteFile("out.txt", res2B, os.FileMode(mode))
}
