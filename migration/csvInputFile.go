///////////////////////////////////////////////////////////////////////////////
//	csvInputFile.go  -  Jan-14-2023  -  aldebap
//
//	Implementation for a fixed position file as a data input source
////////////////////////////////////////////////////////////////////////////////

package migration

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

//	attributes for a CSV Input file
type csvInputFile struct {
	FileName       string
	FieldSeparator string
	Header         bool
	FieldList      []DataField
}

//	NewCSVInputFile create a new csvInputFile
func NewCSVInputFile(config JobInput) DataInputSource {

	return &csvInputFile{
		FileName:       config.FileName,
		FieldSeparator: config.FieldSeparator,
		Header:         config.Header,
		FieldList:      config.FieldList,
	}
}

//	ValidateFormat validate file fields format
func (f *csvInputFile) ValidateFormat() error {

	//	there must be at least one field
	if len(f.FieldList) == 0 && f.Header == false {
		return errors.New("File format need at least one field or file have a header")
	}

	//	there must a field separator
	if len(f.FieldSeparator) != 1 {
		return errors.New("Missing or invalid field separator")
	}

	//	validate file fields format
	for _, field := range f.FieldList {

		//	validate the field type
		_, found := data_field_type[field.Type]
		if !found {
			return errors.New("Invalid field type: " + field.Type)
		}

		//	validate start position
		if field.StartPosition != 0 {
			return errors.New("Field start position must not be used for CSV files: " + field.Name)
		}

		//	validate end position
		if field.EndPosition != 0 {
			return errors.New("Field end position must not be used for CSV files: " + field.Name)
		}
	}

	return nil
}

//	ImportData open fixed position file and import its data
func (f *csvInputFile) ImportData(nextStep DataPipelineStep) (rowsProcessed int64, err error) {

	//	 open fixed position file
	dataFile, err := os.Open(f.FileName)
	if err != nil {
		return 0, errors.New("fail opening data file: " + err.Error())
	}
	defer dataFile.Close()

	//	read config file line by line
	var dataRow []byte
	var rowValue map[string]string

	rowValue = make(map[string]string)

	rowsProcessed = 0
	dataFileReader := bufio.NewReader(dataFile)

	for {
		dataRow, _, err = dataFileReader.ReadLine()
		if err != nil {
			break
		}
		//	if file have a reader, ignores it
		if rowsProcessed == 0 && f.Header {
			continue
		}

		//	extract fields from input line
		values := strings.Split(string(dataRow), string(f.FieldSeparator))

		for i, field := range f.FieldList {
			if i < len(values) {
				rowValue[field.Name] = values[i]
			} else {
				rowValue[field.Name] = ""
			}
		}

		//	if available, invoke the next step in the pipeline
		if nextStep != nil {
			nextStep.ProcessRow(rowValue)
		}

		rowsProcessed++
	}

	return rowsProcessed, nil
}
