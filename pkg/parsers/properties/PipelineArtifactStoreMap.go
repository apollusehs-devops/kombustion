package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// PipelineArtifactStoreMap Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstoremap.html
type PipelineArtifactStoreMap struct {
	Region        interface{} `yaml:"Region"`
	ArtifactStore interface{} `yaml:"ArtifactStore"`
}

// PipelineArtifactStoreMap validation
func (resource PipelineArtifactStoreMap) Validate() []error {
	errors := []error{}

	return errors
}
