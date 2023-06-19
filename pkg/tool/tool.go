package tool

import "encoding/json"

func StructToMap(s interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func MapToStruct(m map[string]interface{}, s interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, s)
	if err != nil {
		return err
	}

	return nil
}
