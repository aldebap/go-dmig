///////////////////////////////////////////////////////////////////////////////
//	fixedPositionInputFile.go  -  Jan-14-2023  -  aldebap
//
//	Implementation for a fixed position file as a data input source
////////////////////////////////////////////////////////////////////////////////

package migration

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

//	attributes for a migration job
type fixedPositionInputFile struct {
	FileName  string
	Header    bool
	Trailer   bool
	FieldList []DataField
}

//	create a new FixedPositionInputFile
func NewFixedPositionInputFile(config JobInput) DataInputSource {

	return &fixedPositionInputFile{
		FileName:  config.FileName,
		Header:    config.Header,
		Trailer:   config.Trailer,
		FieldList: config.FieldList,
	}
}

//	ValidateFormat validate file fields format
func (f *fixedPositionInputFile) ValidateFormat() error {

	for _, field := range f.FieldList {

		//	validate the field type
		_, found := data_field_type[field.Type]
		if !found {
			return errors.New("Invalid field type: " + field.Type)
		}

		//	validate start position
		if field.StartPosition == 0 {
			return errors.New("Required field start position: " + field.Name)
		}

		//	validate end position
		if field.EndPosition == 0 {
			return errors.New("Required field end position: " + field.Name)
		}

		//	validate start and end positions
		if field.StartPosition > field.EndPosition {
			return errors.New("Field start position greater than end position: " + field.Name)
		}
	}

	//	check for overlaping fields

	return nil
}

//	ImportData open fixed position file and import its data
func (f *fixedPositionInputFile) ImportData() (rowsProcessed int64, err error) {

	//	 open fixed position file
	dataFile, err := os.Open(f.FileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[error] opening data file: %s\n", err.Error())
		os.Exit(-1)
	}
	defer dataFile.Close()

	//	read config file line by line
	var dataRow []byte

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

		fmt.Fprintf(os.Stdout, "[trace] input rows: %s\n", string(dataRow))

		rowsProcessed++
	}

	return rowsProcessed, nil
}
