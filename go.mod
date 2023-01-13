module github.com/aldebap/go-dmig

go 1.17

require (
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require github.com/aldebap/go-dmig/migration v0.0.0-unpublished

replace github.com/aldebap/go-dmig/migration v0.0.0-unpublished => ./migration
