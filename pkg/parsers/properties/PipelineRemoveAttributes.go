package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// PipelineRemoveAttributes Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-removeattributes.html
type PipelineRemoveAttributes struct {
	Name       interface{} `yaml:"Name,omitempty"`
	Next       interface{} `yaml:"Next,omitempty"`
	Attributes interface{} `yaml:"Attributes,omitempty"`
}

// PipelineRemoveAttributes validation
func (resource PipelineRemoveAttributes) Validate() []error {
	errors := []error{}

	return errors
}
