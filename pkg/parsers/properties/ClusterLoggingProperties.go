package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ClusterLoggingProperties Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html
type ClusterLoggingProperties struct {
	BucketName  interface{} `yaml:"BucketName"`
	S3KeyPrefix interface{} `yaml:"S3KeyPrefix,omitempty"`
}

// ClusterLoggingProperties validation
func (resource ClusterLoggingProperties) Validate() []error {
	errors := []error{}

	return errors
}
