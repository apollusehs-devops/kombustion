package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// TopicRuleIotAnalyticsAction Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-iotanalyticsaction.html
type TopicRuleIotAnalyticsAction struct {
	ChannelName interface{} `yaml:"ChannelName"`
	RoleArn     interface{} `yaml:"RoleArn"`
}

// TopicRuleIotAnalyticsAction validation
func (resource TopicRuleIotAnalyticsAction) Validate() []error {
	errors := []error{}

	return errors
}
