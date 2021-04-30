package utils

import (
	"encoding/json"
)

func Form2model(formStruct, obj interface{}) error {
	formJson, err := json.Marshal(formStruct)
	if err != nil {
		return err
	}
	err = json.Unmarshal(formJson, obj)
	if err != nil {
		return err
	}

	return nil
}
