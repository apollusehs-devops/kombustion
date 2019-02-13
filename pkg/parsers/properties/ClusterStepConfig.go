package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ClusterStepConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-stepconfig.html
type ClusterStepConfig struct {
	ActionOnFailure interface{} `yaml:"ActionOnFailure,omitempty"`
	Name            interface{} `yaml:"Name"`
	HadoopJarStep   interface{} `yaml:"HadoopJarStep"`
}

// ClusterStepConfig validation
func (resource ClusterStepConfig) Validate() []error {
	errors := []error{}

	return errors
}