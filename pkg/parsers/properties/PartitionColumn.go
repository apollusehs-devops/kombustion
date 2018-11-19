package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// PartitionColumn Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-column.html
type PartitionColumn struct {
	Comment interface{} `yaml:"Comment,omitempty"`
	Name    interface{} `yaml:"Name"`
	Type    interface{} `yaml:"Type,omitempty"`
}

// PartitionColumn validation
func (resource PartitionColumn) Validate() []error {
	errors := []error{}

	return errors
}
