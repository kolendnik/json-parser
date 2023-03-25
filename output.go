package jsonparser

import (
	"fmt"
)

func ToStdout(w DataWriter) {
	for w.Next() {
		fmt.Println("{")
		for f, v := range w.GetRow() {
			fmt.Printf("\t%s: %v\n", f, v)
		}
		fmt.Println("}")
	}
}
