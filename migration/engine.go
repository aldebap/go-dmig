///////////////////////////////////////////////////////////////////////////////
//	engine.go  -  Jan-15-2023  -  aldebap
//
//	Migration engine
////////////////////////////////////////////////////////////////////////////////

package migration

import (
	"errors"
	"fmt"
	"os"
)

//	constants for input/output types
const (
	FIXED_POSITION_FILE = 1
)

var (
	io_type = map[string]uint8{
		"FixedPositionFile": FIXED_POSITION_FILE,
	}
)

//	constants for data field types
const (
	INTEGER = 1
	STRING  = 2
)

var (
	data_field_type = map[string]uint8{
		"integer": INTEGER,
		"string":  STRING,
	}
)

//	PerformMigration perform a migration configured by the DataMigration object
func (dmig *DataMigration) PerformMigration() error {

	fmt.Fprintf(os.Stdout, ">>> Starting Migration: %s\n", dmig.Description)

	for _, job := range dmig.JobList {

		fmt.Fprintf(os.Stdout, "\nMigration Job: %s\n", job.Name)

		//check job's input type
		inputType, found := io_type[job.Input.Type]
		if !found {
			return errors.New("Invalid job's input type: " + job.Input.Type)
		}

		var input DataInputSource

		switch inputType {
		case FIXED_POSITION_FILE:
			input = NewFixedPositionInputFile(job.Input)
		}

		err := input.ValidateFormat()
		if err != nil {
			return err
		}

		rowsProcessed, err := input.ImportData()
		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "Job finished: %d rows processed\n", rowsProcessed)
	}

	return nil
}
