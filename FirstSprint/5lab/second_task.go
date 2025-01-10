package main

import "encoding/json"

type Stud struct {
	Name  string `json:"name"`
	Class string `json:"class"`
}

func mergeJSONData(jsonDataList ...[]byte) ([]byte, error) {
	var ss []Stud
	var s []Stud
	for _, jsonData := range jsonDataList {
		err := json.Unmarshal(jsonData, &s)
		if err != nil {
			return nil, err
		}
		for _, x := range s {
			ss = append(ss, x)
		}
	}
	jsonBytes, err1 := json.Marshal(ss)
	if err1 != nil {
		return nil, err1
	}
	return jsonBytes, nil
}
