package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// DatasetVariable Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-variable.html
type DatasetVariable struct {
	DoubleValue                interface{} `yaml:"DoubleValue,omitempty"`
	StringValue                interface{} `yaml:"StringValue,omitempty"`
	VariableName               interface{} `yaml:"VariableName"`
	OutputFileUriValue         interface{} `yaml:"OutputFileUriValue,omitempty"`
	DatasetContentVersionValue interface{} `yaml:"DatasetContentVersionValue,omitempty"`
}

// DatasetVariable validation
func (resource DatasetVariable) Validate() []error {
	errors := []error{}

	return errors
}
