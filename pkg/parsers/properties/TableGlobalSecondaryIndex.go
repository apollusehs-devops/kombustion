package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// TableGlobalSecondaryIndex Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html
type TableGlobalSecondaryIndex struct {
	IndexName             interface{} `yaml:"IndexName"`
	ProvisionedThroughput interface{} `yaml:"ProvisionedThroughput,omitempty"`
	Projection            interface{} `yaml:"Projection"`
	KeySchema             interface{} `yaml:"KeySchema"`
}

// TableGlobalSecondaryIndex validation
func (resource TableGlobalSecondaryIndex) Validate() []error {
	errors := []error{}

	return errors
}
