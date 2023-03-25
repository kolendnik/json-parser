package jsonparser

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrInputNotJson error = errors.New("input is not a json object")
)

func ParseInput(input []byte) (DataWriter, error) {
	switch string(input[0:1]) {
	case "[":
		var multiple Multiple = NewMultiple()

		err := json.Unmarshal(input, &multiple.Rows)
		if err != nil {
			return nil, err
		}
		multiple.size = len(multiple.Rows)

		multiple.Headers = extractHeaders(multiple.Rows[0])

		return &multiple, nil

	case "{":
		var single Single = NewSingle()

		err := json.Unmarshal(input, &single.data)
		if err != nil {
			return nil, err
		}

		single.headers = extractHeaders(single.data)
		return &single, nil
	default:
		fmt.Println("error: invalid json input: ")
		fmt.Println(input)
		return nil, ErrInputNotJson
	}
}
