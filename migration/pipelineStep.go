///////////////////////////////////////////////////////////////////////////////
//	pipelineStep.go  -  Jan-22-2023  -  aldebap
//
//	Interface of a data pipeline step
////////////////////////////////////////////////////////////////////////////////

package migration

type DataPipelineStep interface {
	SetNextStep(nextStep DataPipelineStep)
	GetNextStep() DataPipelineStep

	ProcessRow(row map[string]string) (rowProcessed bool, err error)
}
