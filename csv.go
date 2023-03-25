package jsonparser

import (
	"encoding/csv"
	"os"
)

func ToCsv(dw DataWriter) error {

	f, err := os.Create("result.csv")
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)

	w.Write(dw.GetHeaders())
	w.Flush()

	for dw.Next() {
		var tmp []string
		for _, v := range dw.GetRow() {
			tmp = append(tmp, ToString(v))
		}

		err := w.Write(tmp)
		if err != nil {
			return err
		}
		w.Flush()
	}
	return nil
}
