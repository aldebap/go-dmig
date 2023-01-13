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

type MigrationJob struct {
}

//	attributes used to configure a migration
type DataMigration struct {
	Description string `yaml:"description"`
	Author      string `yaml:"author"`
	Date        string `yaml:"date"`

	//JobList []MigrationJob
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
