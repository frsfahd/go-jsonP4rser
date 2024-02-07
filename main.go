package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func openFile(filepath string) (*os.File, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func main() {
	var (
		err   error
		input *os.File
		res   map[string]interface{}
		dec   *json.Decoder
	)

	if len(os.Args) > 1 {
		// if input file exist
		input, err = openFile(os.Args[1])
		// input.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		// if no input file, alternatively using the stdin
		input = os.Stdin
	}

	// create decoder based on input and decode it
	dec = json.NewDecoder(input)
	err = dec.Decode(&res)

	if err != nil {
		fmt.Printf("invalid json => %s\n", err.Error())
		return
	}

	// print the json parsed data to stdout
	for i, v := range res {
		fmt.Printf("%s : %v\n", i, v)
	}
	// fmt.Println(os.Args)

}
