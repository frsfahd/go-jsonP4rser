package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var (
	err   error
	input *os.File
	res   map[string]interface{}
	dec   *json.Decoder
)

var ErrFileNotFound = errors.New("file not found")

func openFile(filepath string) (*os.File, error) {
	file, err := os.Open(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrFileNotFound
		}
		return nil, err
	}
	return file, nil
}

func decodeJson(*os.File) error {
	// create decoder based on input and decode it
	dec = json.NewDecoder(input)
	err = dec.Decode(&res)
	return err
}

func main() {

	if len(os.Args) > 1 {
		// if input file exist
		input, err = openFile(os.Args[1])
		// input.Close()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		// if no input file, alternatively using the stdin
		input = os.Stdin
	}

	// decode JSON

	err = decodeJson(input)

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
