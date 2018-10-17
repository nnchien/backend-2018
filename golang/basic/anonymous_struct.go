package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := []byte(`[
		{"English": "Hello", "Vietnam": "Xin chào"},
		{"English": "Doctor", "Vietnam": "Bác sĩ"},
		{"English": "Strange", "Vietnam": "Xa lạ"}
		]`)

	var translations []struct {
		English string
		Vietnam string
	}

	_ = json.Unmarshal(jsonData, &translations)
	for _, t := range translations {
		fmt.Printf("%s -> %s\n", t.English, t.Vietnam)
	}
}
