package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ConfigurationAggregatorAccountAggregationSource Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-accountaggregationsource.html
type ConfigurationAggregatorAccountAggregationSource struct {
	AllAwsRegions interface{} `yaml:"AllAwsRegions,omitempty"`
	AccountIds    interface{} `yaml:"AccountIds"`
	AwsRegions    interface{} `yaml:"AwsRegions,omitempty"`
}

// ConfigurationAggregatorAccountAggregationSource validation
func (resource ConfigurationAggregatorAccountAggregationSource) Validate() []error {
	errors := []error{}

	return errors
}
