///////////////////////////////////////////////////////////////////////////////
//	traceData.go  -  Jan-22-2023  -  aldebap
//
//	Implementation for a trace step for data pipeline
////////////////////////////////////////////////////////////////////////////////

package migration

import (
	"fmt"
	"os"
)

//	attributes for a migration job
type traceDataPipelineStep struct {
	Trace bool

	NextStep DataPipelineStep
}

//	NewTraceDataStep create a new traceDataPipelineStep
func NewTraceDataStep(trace bool) DataPipelineStep {

	return &traceDataPipelineStep{
		Trace: trace,
	}
}

//	SetNextStep set the next step in data pipeline
func (s *traceDataPipelineStep) SetNextStep(nextStep DataPipelineStep) {
	s.NextStep = nextStep
}

//	GetNextStep get the next step in data pipeline
func (s *traceDataPipelineStep) GetNextStep() DataPipelineStep {
	return s.NextStep
}

//	ProcessRow generate a trace of the data row
func (s *traceDataPipelineStep) ProcessRow(row map[string]string) (rowProcessed bool, err error) {

	var i uint

	fmt.Fprintf(os.Stdout, "[trace] fields: ")
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
