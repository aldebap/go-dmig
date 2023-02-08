///////////////////////////////////////////////////////////////////////////////
//	fixedPositionOutputFile.go  -  Fev-7-2023  -  aldebap
//
//	Implementation for a fixed position file as a pipeline step
////////////////////////////////////////////////////////////////////////////////

package migration

import (
	"fmt"
	"os"
)

//	attributes for a fixedPositionOutputFile pipeline step
type fixedPositionOutputFile struct {
	FileName  string
	Header    bool
	Trailer   bool
	FieldList []DataField

	NextStep DataPipelineStep
}

//	NewFixedPositionOutputFile create a new fixedPositionOutputFile
func NewFixedPositionOutputFile() DataPipelineStep {

	return &fixedPositionOutputFile{}
}

//	SetNextStep set the next step in data pipeline
func (s *fixedPositionOutputFile) SetNextStep(nextStep DataPipelineStep) {
	s.NextStep = nextStep
}

//	GetNextStep get the next step in data pipeline
func (s *fixedPositionOutputFile) GetNextStep() DataPipelineStep {
	return s.NextStep
}

//	ProcessRow generate a trace of the data row
func (s *fixedPositionOutputFile) ProcessRow(row map[string]string) (rowProcessed bool, err error) {

	var i uint

	for fieldName, fieldValue := range row {
		if i > 0 {
			fmt.Fprintf(os.Stdout, "; ")
		}
		fmt.Fprintf(os.Stdout, "%s = '%s'", fieldName, fieldValue)

		i++
	}
	fmt.Fprintf(os.Stdout, "\n")

	//	if available, invoke the next step in the pipeline
	if s.NextStep != nil {
		return s.NextStep.ProcessRow(row)
	}

	return true, nil
}
