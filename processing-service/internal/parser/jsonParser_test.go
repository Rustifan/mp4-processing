package parser_test

import (
	"reflect"
	"testing"

	"github.com/rustifan/mp4-processing/processing-service/internal/parser"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func GetJSONParser(t *testing.T) {

	personParser := parser.GetJSONParser[Person]()
	addressParser := parser.GetJSONParser[Address]()
	if reflect.TypeOf(personParser) != reflect.TypeOf(parser.JSONParser[Person]{}) {
		t.Errorf("Expected parser of type JSONParser[Person], got %T", personParser)
	}

	if reflect.TypeOf(addressParser) != reflect.TypeOf(parser.JSONParser[Address]{}) {
		t.Errorf("Expected parser of type JSONParser[Address], got %T", addressParser)
	}
}

func TestJSONParserParse_ValidData(t *testing.T) {
	personParser := parser.GetJSONParser[Person]()
	personJSON := []byte(`{"name":"Tvrtko","age":30}`)
	person, err := personParser.Parse(personJSON)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedPerson := Person{Name: "Tvrtko", Age: 30}
	if person != expectedPerson {
		t.Errorf("Expected %v, got %v", expectedPerson, person)
	}

	addressParser := parser.GetJSONParser[Address]()
	addressJSON := []byte(`{"street":"Ulica testa","city":"Mrkopalj","country":"Croatia"}`)
	address, err := addressParser.Parse(addressJSON)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedAddress := Address{Street: "Ulica testa", City: "Mrkopalj", Country: "Croatia"}
	if address != expectedAddress {
		t.Errorf("Expected %v, got %v", expectedAddress, address)
	}

	emptyPersonJSON := []byte(`{}`)
	emptyPerson, err := personParser.Parse(emptyPersonJSON)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedEmptyPerson := Person{}
	if emptyPerson != expectedEmptyPerson {
		t.Errorf("Expected %v, got %v", expectedEmptyPerson, emptyPerson)
	}
}

func TestJSONParserParse_InvalidData(t *testing.T) {
	personParser := parser.GetJSONParser[Person]()
	invalidJSON := []byte(`{"name":"Mirko","age":30`)
	_, err := personParser.Parse(invalidJSON)

	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}

	emptyInput := []byte("")
	_, err = personParser.Parse(emptyInput)

	if err == nil {
		t.Error("Expected error for empty input, got nil")
	}

	typeMismatchJSON := []byte(`{"name":"Stipa","age":"triDeset"}`)
	_, err = personParser.Parse(typeMismatchJSON)

	if err == nil {
		t.Error("Expected error for type mismatch, got nil")
	}
}

func TestJSONParserWithDifferentTypes(t *testing.T) {
	intParser := parser.GetJSONParser[int]()
	intJSON := []byte(`42`)
	intResult, err := intParser.Parse(intJSON)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if intResult != 42 {
		t.Errorf("Expected 42, got %v", intResult)
	}

	sliceParser := parser.GetJSONParser[[]string]()
	sliceJSON := []byte(`["apple","banana","cherry"]`)
	sliceResult, err := sliceParser.Parse(sliceJSON)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedSlice := []string{"apple", "banana", "cherry"}
	if !reflect.DeepEqual(sliceResult, expectedSlice) {
		t.Errorf("Expected %v, got %v", expectedSlice, sliceResult)
	}

	mapParser := parser.GetJSONParser[map[string]interface{}]()
	mapJSON := []byte(`{"key1":"value1","key2":42,"key3":true}`)
	mapResult, err := mapParser.Parse(mapJSON)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if mapResult["key1"] != "value1" || mapResult["key2"] != float64(42) || mapResult["key3"] != true {
		t.Errorf("Map values don't match expected values")
	}
}

func TestJSONParserWithNestedStructs(t *testing.T) {
	type Employee struct {
		Person  Person  `json:"person"`
		Address Address `json:"address"`
		Title   string  `json:"title"`
	}

	employeeParser := parser.GetJSONParser[Employee]()
	employeeJSON := []byte(`{
		"person": {"name":"Ivan","age":28},
		"address": {"street":"Ulica IX","city":"Prvic","country":"Croatia"},
		"title": "Software Engineer"
	}`)

	employee, err := employeeParser.Parse(employeeJSON)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedEmployee := Employee{
		Person:  Person{Name: "Ivan", Age: 28},
		Address: Address{Street: "Ulica IX", City: "Prvic", Country: "Croatia"},
		Title:   "Software Engineer",
	}

	if !reflect.DeepEqual(employee, expectedEmployee) {
		t.Errorf("Expected %v, got %v", expectedEmployee, employee)
	}
}
