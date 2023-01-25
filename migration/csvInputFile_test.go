///////////////////////////////////////////////////////////////////////////////
//	csvInputFile_test.go  -  Jan-24-2023  -  aldebap
//
//	Unit tests for CSV file as a data input source
////////////////////////////////////////////////////////////////////////////////

package migration

import (
	"fmt"
	"os"
	"testing"
)

//	Test_CSVFile_ValidateFormat test cases for validation of file fields format
func Test_CSVFile_ValidateFormat(t *testing.T) {

	//	a few test cases
	var testScenarios = []struct {
		scenario string
		input    JobInput
		output   string
	}{
		{scenario: "empty field list", input: JobInput{}, output: "File format need at least one field or file have a header"},
		{scenario: "missing field separator", input: JobInput{Header: true}, output: "Missing or invalid field separator"},
		{scenario: "invalid field separator", input: JobInput{Header: true, FieldSeparator: ",;:"}, output: "Missing or invalid field separator"},
		{scenario: "invalid field type", input: JobInput{FieldSeparator: ",", FieldList: []DataField{{
			Type: "xpto",
		}}}, output: "Invalid field type: xpto"},
		{scenario: "invalid start position", input: JobInput{FieldSeparator: ",", FieldList: []DataField{{
			Name:          "test",
			Type:          "string",
			StartPosition: 1,
		}}}, output: "Field start position must not be used for CSV files: test"},
		{scenario: "invalid end position", input: JobInput{FieldSeparator: ",", FieldList: []DataField{{
			Name:        "test",
			Type:        "string",
			EndPosition: 10,
		}}}, output: "Field end position must not be used for CSV files: test"},
		{scenario: "valid field list", input: JobInput{FieldSeparator: ",", FieldList: []DataField{
			{
				Name: "test_1",
				Type: "string",
			}, {
				Name: "test_2",
				Type: "string",
			},
		}}, output: ""},
	}

	t.Run(">>> validation of CSV file fields format", func(t *testing.T) {

		for _, test := range testScenarios {

			fmt.Printf("scenario: %s\n", test.scenario)

			testDataSource := NewCSVInputFile(test.input)

			//	validate the format
			got := ""
			want := test.output

			err := testDataSource.ValidateFormat()
			if err != nil {
				got = err.Error()
			}

			if want != got {
				t.Errorf("fail in ValidateFormat(): expected: %s result: %v", want, got)
			}
		}
	})
}

//	Test_CSVFile_ImportData test cases for data file importing
func Test_CSVFile_ImportData(t *testing.T) {

	t.Run(">>> validation data file importing - invalid file name", func(t *testing.T) {

		testDataSource := NewCSVInputFile(JobInput{FileName: "xpto.txt"})

		//	import data
		got := ""
		want := "fail opening data file: open xpto.txt: no such file or directory"

		_, err := testDataSource.ImportData(nil)
		if err != nil {
			got = err.Error()
		}

		if want != got {
			t.Errorf("fail in ImportData(): expected: %s result: %v", want, got)
		}
	})

	t.Run(">>> validation data file importing - empty file", func(t *testing.T) {

		const testFileName = "testData.txt"

		err := os.WriteFile(testFileName, []byte(""), 0644)
		if err != nil {
			t.Errorf("unexpected error creating test file: %s", err)
		}
		defer os.Remove(testFileName)

		testDataSource := NewCSVInputFile(JobInput{
			FileName:       testFileName,
			FieldSeparator: ",",
			FieldList: []DataField{
				{
					Name: "test_1",
					Type: "string",
				}, {
					Name: "test_2",
					Type: "string",
				},
			},
		})

		//	import data
		got := int64(0)
		want := int64(0)

		got, err = testDataSource.ImportData(nil)
		if err != nil {
			t.Errorf("unexpected error in ImportData(): %s", err)
		}

		if want != got {
			t.Errorf("fail in ImportData(): expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> validation data file importing - valid file", func(t *testing.T) {

		const testFileName = "testData.txt"

		err := os.WriteFile(testFileName, []byte("1,LINE#1\n2,LINE#2\n"), 0644)
		if err != nil {
			t.Errorf("unexpected error creating test file: %s", err)
		}
		defer os.Remove(testFileName)

		testDataSource := NewCSVInputFile(JobInput{
			FileName:       testFileName,
			FieldSeparator: ",",
			FieldList: []DataField{
				{
					Name: "test_1",
					Type: "string",
				}, {
					Name: "test_2",
					Type: "string",
				},
			},
		})

		//	import data
		got := int64(0)
		want := int64(2)

		got, err = testDataSource.ImportData(nil)
		if err != nil {
			t.Errorf("unexpected error in ImportData(): %s", err)
		}

		if want != got {
			t.Errorf("fail in ImportData(): expected: %d result: %d", want, got)
		}
	})
}
