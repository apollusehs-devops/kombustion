package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// KinesisStream Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-stream.html
type KinesisStream struct {
	Type       string                  `yaml:"Type"`
	Properties KinesisStreamProperties `yaml:"Properties"`
	Condition  interface{}             `yaml:"Condition,omitempty"`
	Metadata   interface{}             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}             `yaml:"DependsOn,omitempty"`
}

// KinesisStream Properties
type KinesisStreamProperties struct {
	Name                 interface{} `yaml:"Name,omitempty"`
	RetentionPeriodHours interface{} `yaml:"RetentionPeriodHours,omitempty"`
	ShardCount           interface{} `yaml:"ShardCount"`
	StreamEncryption     interface{} `yaml:"StreamEncryption,omitempty"`
	Tags                 interface{} `yaml:"Tags,omitempty"`
}

// NewKinesisStream constructor creates a new KinesisStream
func NewKinesisStream(properties KinesisStreamProperties, deps ...interface{}) KinesisStream {
	return KinesisStream{
		Type:       "AWS::Kinesis::Stream",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseKinesisStream parses KinesisStream
func ParseKinesisStream(
	name string,
	data string,
) (
	source string,
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	transform types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-resources"

	// Resources
	var resource KinesisStream
	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	if validateErrs := resource.Properties.Validate(); len(errors) > 0 {
		errors = append(errors, validateErrs...)
		return
	}

	resources = types.TemplateObject{name: resource}

	// Outputs

	outputs = types.TemplateObject{
		name: types.TemplateObject{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-KinesisStream-" + name,
				},
			},
		},
	}

	return
}

// ParseKinesisStream validator
func (resource KinesisStream) Validate() []error {
	return resource.Properties.Validate()
}

// ParseKinesisStreamProperties validator
func (resource KinesisStreamProperties) Validate() []error {
	errors := []error{}
	return errors
}
