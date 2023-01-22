///////////////////////////////////////////////////////////////////////////////
//	fixedPositionInputFile_test.go  -  Jan-19-2023  -  aldebap
//
//	Unit tests for a fixed position file as a data input source
////////////////////////////////////////////////////////////////////////////////

package migration

import (
	"fmt"
	"os"
	"testing"
)

//	Test_ValidateFormat test cases for validation of file fields format
func Test_ValidateFormat(t *testing.T) {

	//	a few test cases
	var testScenarios = []struct {
		scenario string
		input    JobInput
		output   string
	}{
		{scenario: "empty field list", input: JobInput{}, output: "File format need at least one field"},
		{scenario: "invalid field type", input: JobInput{FieldList: []DataField{{
			Type: "xpto",
		}}}, output: "Invalid field type: xpto"},
		{scenario: "missing start position", input: JobInput{FieldList: []DataField{{
			Name: "test",
			Type: "string",
		}}}, output: "Required field start position: test"},
		{scenario: "missing end position", input: JobInput{FieldList: []DataField{{
			Name:          "test",
			Type:          "string",
			StartPosition: 1,
		}}}, output: "Required field end position: test"},
		{scenario: "start position greater than end", input: JobInput{FieldList: []DataField{{
			Name:          "test",
			Type:          "string",
			StartPosition: 5,
			EndPosition:   1,
		}}}, output: "Field start position greater than end position: test"},
		{scenario: "start position overlap", input: JobInput{FieldList: []DataField{
			{
				Name:          "test",
				Type:          "string",
				StartPosition: 2,
				EndPosition:   3,
			}, {
				Name:          "test",
				Type:          "string",
				StartPosition: 1,
				EndPosition:   5,
			},
		}}, output: "Field #1 position overlapping with field #2"},
		{scenario: "end position overlap", input: JobInput{FieldList: []DataField{
			{
				Name:          "test_1",
				Type:          "string",
				StartPosition: 2,
				EndPosition:   6,
			}, {
				Name:          "test_2",
				Type:          "string",
				StartPosition: 5,
				EndPosition:   9,
			},
		}}, output: "Field #1 position overlapping with field #2"},
		{scenario: "more position overlap", input: JobInput{FieldList: []DataField{
			{
				Name:          "test_1",
				Type:          "string",
				StartPosition: 1,
				EndPosition:   3,
			}, {
				Name:          "test_2",
				Type:          "string",
				StartPosition: 4,
				EndPosition:   9,
			}, {
				Name:          "test_3",
				Type:          "string",
				StartPosition: 6,
				EndPosition:   12,
			},
		}}, output: "Field #2 position overlapping with field #3"},
		{scenario: "valid field list", input: JobInput{FieldList: []DataField{
			{
				Name:          "test_1",
				Type:          "string",
				StartPosition: 1,
				EndPosition:   3,
			}, {
				Name:          "test_2",
				Type:          "string",
				StartPosition: 4,
				EndPosition:   9,
			},
		}}, output: ""},
	}

	t.Run(">>> validation of file fields format", func(t *testing.T) {

		for _, test := range testScenarios {

			fmt.Printf("scenario: %s\n", test.scenario)

			testDataSource := NewFixedPositionInputFile(test.input)

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

//	Test_ImportData test cases for data file importing
func Test_ImportData(t *testing.T) {

	t.Run(">>> validation data file importing - invalid file name", func(t *testing.T) {

		testDataSource := NewFixedPositionInputFile(JobInput{FileName: "xpto.txt"})

		//	import data
		got := ""
		want := "fail opening data file: open xpto.txt: no such file or directory"

		_, err := testDataSource.ImportData()
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

		testDataSource := NewFixedPositionInputFile(JobInput{
			FileName: testFileName,
			FieldList: []DataField{
				{
					Name:          "test_1",
					Type:          "string",
					StartPosition: 1,
					EndPosition:   3,
				}, {
					Name:          "test_2",
					Type:          "string",
					StartPosition: 4,
					EndPosition:   9,
				},
			},
		})

		//	import data
		got := int64(0)
		want := int64(0)

		got, err = testDataSource.ImportData()
		if err != nil {
			t.Errorf("unexpected error in ImportData(): %s", err)
		}

		if want != got {
			t.Errorf("fail in ImportData(): expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> validation data file importing - valid file", func(t *testing.T) {

		const testFileName = "testData.txt"

		err := os.WriteFile(testFileName, []byte("001LINE#1\n002LINE#2\n"), 0644)
		if err != nil {
			t.Errorf("unexpected error creating test file: %s", err)
		}
		defer os.Remove(testFileName)

		testDataSource := NewFixedPositionInputFile(JobInput{
			FileName: testFileName,
			FieldList: []DataField{
				{
					Name:          "test_1",
					Type:          "string",
					StartPosition: 1,
					EndPosition:   3,
				}, {
					Name:          "test_2",
					Type:          "string",
					StartPosition: 4,
					EndPosition:   9,
				},
			},
		})

		//	import data
		got := int64(0)
		want := int64(2)

		got, err = testDataSource.ImportData()
		if err != nil {
			t.Errorf("unexpected error in ImportData(): %s", err)
		}

		if want != got {
			t.Errorf("fail in ImportData(): expected: %d result: %d", want, got)
		}
	})
}
