package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// AppDataSource Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-datasource.html
type AppDataSource struct {
	Arn          interface{} `yaml:"Arn,omitempty"`
	DatabaseName interface{} `yaml:"DatabaseName,omitempty"`
	Type         interface{} `yaml:"Type,omitempty"`
}

// AppDataSource validation
func (resource AppDataSource) Validate() []error {
	errors := []error{}

	return errors
}
