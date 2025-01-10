package main

import "encoding/json"

type Studs struct {
	Class string `json:"class"`
	Name  string `json:"name"`
}

func splitJSONByClass(jsonData []byte) (map[string][]byte, error) {
	preres := make(map[string][]Studs)
	var cur []Studs
	err := json.Unmarshal(jsonData, &cur)
	if err != nil {
		return nil, err
	}

	for _, stud := range cur {
		preres[stud.Class] = append(preres[stud.Class], stud)
	}

	res := make(map[string][]byte)
	for class, students := range preres {
		jsonBytes, err := json.Marshal(students)
		if err != nil {
			return nil, err
		}
		res[class] = jsonBytes
	}

	return res, nil
}
