package parser

import "encoding/json"

type JSONParser[T any] struct{}

func GetJSONParser[T any]() JSONParser[T] {
	return JSONParser[T]{}
}

func (p JSONParser[T]) Parse(data []byte) (T, error) {
	var result T
	err := json.Unmarshal(data, &result)
	return result, err
}
