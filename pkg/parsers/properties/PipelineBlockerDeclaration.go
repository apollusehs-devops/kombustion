package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// PipelineBlockerDeclaration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-blockers.html
type PipelineBlockerDeclaration struct {
	Name interface{} `yaml:"Name"`
	Type interface{} `yaml:"Type"`
}

// PipelineBlockerDeclaration validation
func (resource PipelineBlockerDeclaration) Validate() []error {
	errors := []error{}

	return errors
}
