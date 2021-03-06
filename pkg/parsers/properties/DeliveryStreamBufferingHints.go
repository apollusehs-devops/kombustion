package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// DeliveryStreamBufferingHints Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-bufferinghints.html
type DeliveryStreamBufferingHints struct {
	IntervalInSeconds interface{} `yaml:"IntervalInSeconds"`
	SizeInMBs         interface{} `yaml:"SizeInMBs"`
}

// DeliveryStreamBufferingHints validation
func (resource DeliveryStreamBufferingHints) Validate() []error {
	errors := []error{}

	return errors
}
