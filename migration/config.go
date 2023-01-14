///////////////////////////////////////////////////////////////////////////////
//	config.go  -  Jan-12-2023  -  aldebap
//
//	Parse migration config files
////////////////////////////////////////////////////////////////////////////////

package migration

import (
	"bufio"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
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

//	attributes for a data field
type DataField struct {
}

//	attributes for a migration job
type JobInput struct {
	Description string      `yaml:"description"`
	Type        string      `yaml:"type"`
	FileName    string      `yaml:"file_name"`
	Header      bool        `yaml:"header"`
	Trailer     bool        `yaml:"trailer"`
	Fields      []DataField `yaml:"fields"`
}

//	attributes for a migration job
type MigrationJob struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Input       JobInput `yaml:"input"`
	Trace       bool     `yaml:"trace"`
}

//	attributes used to configure a migration
type DataMigration struct {
	Description string         `yaml:"description"`
	Author      string         `yaml:"author"`
	Date        string         `yaml:"date"`
	JobList     []MigrationJob `yaml:"jobs"`
}

//	LoadConfigFile load a migration config file return a DataMigration
func LoadConfigFile(reader *bufio.Reader) (*DataMigration, error) {

	//	read config file line by line
	var configData string

	for {
		bufLine, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		configData += string(bufLine) + "\n"
	}

	//	unmarshal yaml configuration
	dataMig := &DataMigration{}

	err := yaml.Unmarshal([]byte(configData), dataMig)
	if err != nil {
		return nil, err
	}

	return dataMig, nil
}

//	PerformMigration perform a migration configured by the DataMigration object
func (dmig *DataMigration) PerformMigration() error {

	fmt.Fprintf(os.Stdout, ">>> Starting Migration: %s\n", dmig.Description)

	for _, job := range dmig.JobList {

		fmt.Fprintf(os.Stdout, "Migration Job: %s\n", job.Name)
	}

	return nil
}
