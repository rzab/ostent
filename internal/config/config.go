package config

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/influxdata/telegraf/plugins/inputs"
	"github.com/influxdata/telegraf/plugins/outputs"
	"github.com/influxdata/telegraf/plugins/serializers"

	"github.com/influxdata/toml"
	"github.com/influxdata/toml/ast"

	"github.com/ostrost/ostent/internal"
	internal_models "github.com/ostrost/ostent/internal/models"
)

var config = struct {
	UnmarshalTable func(*ast.Table, interface{}) error
}{UnmarshalTable: toml.UnmarshalTable}

var (
	// envVarRe is a regex to find environment variables in the config file
	envVarRe = regexp.MustCompile(`\$\w+`)
)

// Config specifies the URL/user/password for the database that the agent
// will be logging to, as well as all the plugins that the user has
// specified
type Config struct {
	Agent   *AgentConfig
	Inputs  []*internal_models.RunningInput
	Outputs []*internal_models.RunningOutput
}

func NewConfig() *Config {
	c := &Config{
		Agent: &AgentConfig{
			Interval:      internal.Duration{Duration: 10 * time.Second},
			FlushInterval: internal.Duration{Duration: 10 * time.Second},
		},
	}
	return c
}

type AgentConfig struct {
	// Interval at which to gather information
	Interval internal.Duration

	// FlushInterval is the Interval at which to flush data
	FlushInterval internal.Duration
}

// trimBOM trims the Byte-Order-Marks from the beginning of the file.
// this is for Windows compatability only.
// see https://github.com/influxdata/telegraf/issues/1378
func trimBOM(f []byte) []byte {
	return bytes.TrimPrefix(f, []byte("\xef\xbb\xbf"))
}

func parse(contents []byte) (*ast.Table, error) {
	// ugh windows why
	contents = trimBOM(contents)

	for _, dword := range envVarRe.FindAll(contents, -1) {
		if val := os.Getenv(string(dword[1:])); val != "" {
			contents = bytes.Replace(contents, dword, []byte(val), 1)
		}
	}
	return toml.Parse(contents)
}

func (c *Config) LoadConfig() error {
	path := "internal config"
	tbl, err := parse([]byte(`
[agent]
  interval = "1s"
  flushInterval = "1s"
[[inputs.system_ostent]]
  interval = "1s"
# [[outputs.file]]
[[outputs.ostent]]
`))
	if err != nil {
		return err
	}

	// Parse agent table:
	if val, ok := tbl.Fields["agent"]; ok {
		subTable, ok := val.(*ast.Table)
		if !ok {
			return fmt.Errorf("%s: invalid configuration", path)
		}
		if err = config.UnmarshalTable(subTable, c.Agent); err != nil {
			log.Printf("Could not parse [agent] config\n")
			return fmt.Errorf("Error parsing %s, %s", path, err)
		}
	}

	// Parse all the rest of the plugins:
	for name, val := range tbl.Fields {
		subTable, ok := val.(*ast.Table)
		if !ok {
			return fmt.Errorf("%s: invalid configuration", path)
		}
		switch name {
		case "outputs":
			for pluginName, pluginVal := range subTable.Fields {
				switch pluginSubTable := pluginVal.(type) {
				case *ast.Table:
					if err = c.addOutput(pluginName, pluginSubTable); err != nil {
						return fmt.Errorf("Error parsing %s, %s", path, err)
					}
				case []*ast.Table:
					for _, t := range pluginSubTable {
						if err = c.addOutput(pluginName, t); err != nil {
							return fmt.Errorf("Error parsing %s, %s", path, err)
						}
					}
				default:
					return fmt.Errorf("Unsupported config format: %s, file %s",
						pluginName, path)
				}
			}
		case "inputs", "plugins":
			for pluginName, pluginVal := range subTable.Fields {
				switch pluginSubTable := pluginVal.(type) {
				case *ast.Table:
					if err = c.addInput(pluginName, pluginSubTable); err != nil {
						return fmt.Errorf("Error parsing %s, %s", path, err)
					}
				case []*ast.Table:
					for _, t := range pluginSubTable {
						if err = c.addInput(pluginName, t); err != nil {
							return fmt.Errorf("Error parsing %s, %s", path, err)
						}
					}
				default:
					return fmt.Errorf("Unsupported config format: %s, file %s",
						pluginName, path)
				}
			}
		}
	}
	return err
}

func (c *Config) addOutput(name string, table *ast.Table) error {
	creator, ok := outputs.Outputs[name]
	if !ok {
		return fmt.Errorf("Undefined but requested output: %s", name)
	}
	output := creator()

	// If the output has a SetSerializer function, then this means it can write
	// arbitrary types of output, so build the serializer and set it.
	switch t := output.(type) {
	case serializers.SerializerOutput:
		serializer, err := buildSerializer(name, table)
		if err != nil {
			return err
		}
		if serializer == nil {
			return fmt.Errorf("Serializer is nil")
		}
		t.SetSerializer(serializer)
	}

	ro := internal_models.NewRunningOutput(name, output)
	c.Outputs = append(c.Outputs, ro)
	return nil
}

func (c *Config) addInput(name string, table *ast.Table) error {
	creator, ok := inputs.Inputs[name]
	if !ok {
		return fmt.Errorf("Undefined but requested input: %s", name)
	}
	input := creator()
	pluginConfig, err := buildInput(name, table)
	if err != nil {
		return err
	}
	rp := &internal_models.RunningInput{
		Name:   name,
		Input:  input,
		Config: pluginConfig,
	}
	c.Inputs = append(c.Inputs, rp)
	return nil
}

// buildSerializer grabs the necessary entries from the ast.Table for creating
// a serializers.Serializer object, and creates it, which can then be added onto
// an Output object.
func buildSerializer(name string, tbl *ast.Table) (serializers.Serializer, error) {
	return serializers.NewSerializer(&serializers.Config{
		DataFormat: "graphite",
	})
}

func buildInput(name string, tbl *ast.Table) (*internal_models.InputConfig, error) {
	cp := &internal_models.InputConfig{Name: name}
	if node, ok := tbl.Fields["interval"]; ok {
		if kv, ok := node.(*ast.KeyValue); ok {
			if str, ok := kv.Value.(*ast.String); ok {
				dur, err := time.ParseDuration(str.Value)
				if err != nil {
					return nil, err
				}
				cp.Interval = dur
			}
		}
	}
	return cp, nil
}