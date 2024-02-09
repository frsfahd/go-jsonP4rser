package main

import (
	"bytes"
	"errors"
	"testing"
)

func Test_openFile(t *testing.T) {
	data := []struct {
		name     string
		filepath string
		wantErr  error
	}{
		{name: "valid file", filepath: "tests/step1/valid.json", wantErr: nil},
		{name: "invalid file", filepath: "tests/step1/valid1.json", wantErr: ErrFileNotFound},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			f, err := openFile(d.filepath)
			f.Close()
			if d.wantErr != nil {
				if !errors.Is(err, ErrFileNotFound) {
					t.Errorf("got error: %s\nwant error: %s\n", err, d.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("got error: %v\nwant: success open file\n", err)
				}
			}
		})
	}
}

func Test_decodeJson(t *testing.T) {
	data := []struct {
		name     string
		jsonData []byte
		wantErr  error
	}{
		{
			name:     "valid_json",
			jsonData: []byte(`{"key":"value"}`),
			wantErr:  nil,
		},
		{
			name:     "invalid_json: syntax_error",
			jsonData: []byte(`{"key":"value",}`),
			wantErr:  ErrInvalidSyntax,
		},
	}

	for _, d := range data {

		t.Run(d.name, func(t *testing.T) {
			var tmp map[string]interface{}

			json := bytes.NewReader(d.jsonData)
			err := decodeJson(json, &tmp)

			if d.wantErr != nil {
				if !errors.Is(err, d.wantErr) {
					t.Errorf("got error: %s\nwant error: %s", err.Error(), d.wantErr.Error())
				}
			} else {
				if err != nil {
					t.Errorf("got error: %s\nwant: encoding success", err.Error())
				}
			}

		})
	}
}
