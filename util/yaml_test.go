package util

import (
	"os"
	"testing"
)

type testStruct struct {
	A int `yaml:"a"`
	B int `yaml:"b"`
}

var testCases = []struct {
	name         string
	actionType   string
	filePath     string
	testData     testStruct
	expectedData testStruct
}{
	{
		name:       "test save yaml",
		actionType: "save",
		filePath:   "test.yaml",
		testData: testStruct{
			A: 1,
			B: 2,
		},
		expectedData: testStruct{
			A: 1,
			B: 2,
		},
	},
	{
		name:       "test load yaml",
		actionType: "load",
		filePath:   "test.yaml",
		testData: testStruct{
			A: 1,
			B: 2,
		},
		expectedData: testStruct{
			A: 1,
			B: 2,
		},
	},
}

func TestYaml(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.actionType {
			case "save":
				err := SaveToYaml(testCase.filePath, testCase.testData)
				if err != nil {
					t.Errorf("save yaml failed: %v", err)
				}
			case "load":
				var data testStruct
				err := LoadFromYaml(testCase.filePath, &data)
				if err != nil {
					t.Errorf("load yaml failed: %v", err)
				}
				if data != testCase.expectedData {
					t.Errorf("expected data: %v, actual data: %v", testCase.expectedData, data)
				}
				os.Remove(testCase.filePath)
			}
		})
	}
}
