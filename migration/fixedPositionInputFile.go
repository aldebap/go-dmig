///////////////////////////////////////////////////////////////////////////////
//	fixedPositionInputFile.go  -  Jan-14-2023  -  aldebap
//
//	Implementation for a fixed position file as a data input source
////////////////////////////////////////////////////////////////////////////////

package migration

import (
	"bufio"
	"fmt"
	"os"
)

//	attributes for a migration job
type fixedPositionInputFile struct {
	FileName string
	Header   bool
	Trailer  bool
	Fields   []DataField
}

//	create a new FixedPositionInputFile
func NewFixedPositionInputFile(config JobInput) DataInputSource {

	return &fixedPositionInputFile{
		FileName: config.FileName,
		Header:   config.Header,
		Trailer:  config.Trailer,
	}
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
