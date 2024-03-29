///////////////////////////////////////////////////////////////////////////////
//	config.go  -  Jan-12-2023  -  aldebap
//
//	Parse migration config files
////////////////////////////////////////////////////////////////////////////////

package migration

import (
	"bufio"

	"gopkg.in/yaml.v3"
)

//	attributes for a data field
type DataField struct {
	Name          string `yaml:"name"`
	Type          string `yaml:"type"`
	StartPosition int16  `yaml:"start"`
	EndPosition   int16  `yaml:"end"`
}

//	attributes for a migration job
type JobInput struct {
	Description    string      `yaml:"description"`
	Type           string      `yaml:"type"`
	FileName       string      `yaml:"file_name"`
	FieldSeparator string      `yaml:"field_separator"`
	Header         bool        `yaml:"header"`
	Trailer        bool        `yaml:"trailer"`
	FieldList      []DataField `yaml:"fields"`
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
