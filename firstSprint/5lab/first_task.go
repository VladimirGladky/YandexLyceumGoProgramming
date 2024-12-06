package main

import "encoding/json"

type Student struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func modifyJSON(jsonData []byte) ([]byte, error) {
	var students []Student
	err := json.Unmarshal(jsonData, &students)
	if err != nil {
		return nil, err
	}
	for i := range students {
		students[i].Grade += 1
	}
	jsonBytes, err1 := json.Marshal(students)
	if err1 != nil {
		return nil, err1
	}
	return jsonBytes, nil
}
