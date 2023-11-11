package conf

import (
	"errors"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// ReverseConfig represents a reverse configuration
type ReverseConfig struct {
	Kind    string          `yaml:"kind"`
	Name    string          `yaml:"name"`
	Source  ReverseSource   `yaml:"source"`
	Targets []ReverseTarget `yaml:"targets"`
}

// NewReverseConfigFromYAML parse config yaml and return it. support multiple yaml document in one file.
func NewReverseConfigFromYAML(path string) ([]*ReverseConfig, error) {
	ret := make([]*ReverseConfig, 0)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	for {
		item := new(ReverseConfig)
		err = decoder.Decode(&item)
		if errors.Is(err, io.EOF) {
			// read last document
			break
		}

		if item == nil {
			// empty document
			continue
		}

		if err != nil {
			// other error
			return nil, err
		}

		ret = append(ret, item)
	}

	return ret, nil
}

// ReverseSource represents a reverse source which should be a database connection
type ReverseSource struct {
	Database string `yaml:"database"`
	ConnStr  string `yaml:"conn_str"`
}

// ReverseTarget represents a reverse target
type ReverseTarget struct {
	Type          string   `yaml:"type"`
	IncludeTables []string `yaml:"include_tables"`
	ExcludeTables []string `yaml:"exclude_tables"`
	TableMapper   string   `yaml:"table_mapper"`
	ColumnMapper  string   `yaml:"column_mapper"`
	TemplatePath  string   `yaml:"template_path"`
	Template      string   `yaml:"template"`
	MultipleFiles bool     `yaml:"multiple_files"`
	OutputDir     string   `yaml:"output_dir"`
	TablePrefix   string   `yaml:"table_prefix"`
	Language      string   `yaml:"language"`
	TableName     bool     `yaml:"table_name"`
	ColumnName    bool     `yaml:"column_name"`

	Funcs     map[string]string `yaml:"funcs"`
	Formatter string            `yaml:"formatter"`
	Importter string            `yaml:"importter"`
	ExtName   string            `yaml:"ext_name"`
}
