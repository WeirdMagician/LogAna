package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	start(os.Stdout)
}

func start(w io.Writer) {
	//var filesdata string
	var filesbyarr []byte
	if len(os.Args) == 1 {
		fmt.Fprintf(w, "No log files are provided")
	} else {
		for _, filename := range os.Args[1:] {
			data, err := ioutil.ReadFile(filename)
			if err != nil {
				fmt.Fprintf(w, "\n\t Error Occurred while reaading the file %q\n the error is: %v", filename, err)
			}
			filesbyarr = append(filesbyarr, data...)
			//filesdata +=

		}
		//	strings.SplitAfter(string(data), "[console.accounts]")
		fmt.Fprintf(w, "\n%s", string(filesbyarr))
	}

}
