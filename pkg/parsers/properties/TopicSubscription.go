package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// TopicSubscription Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-subscription.html
type TopicSubscription struct {
	Endpoint interface{} `yaml:"Endpoint"`
	Protocol interface{} `yaml:"Protocol"`
}

// TopicSubscription validation
func (resource TopicSubscription) Validate() []error {
	errors := []error{}

	return errors
}
