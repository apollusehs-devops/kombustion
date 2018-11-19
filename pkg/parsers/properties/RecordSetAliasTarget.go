package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// RecordSetAliasTarget Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html
type RecordSetAliasTarget struct {
	DNSName              interface{} `yaml:"DNSName"`
	EvaluateTargetHealth interface{} `yaml:"EvaluateTargetHealth,omitempty"`
	HostedZoneId         interface{} `yaml:"HostedZoneId"`
}

// RecordSetAliasTarget validation
func (resource RecordSetAliasTarget) Validate() []error {
	errors := []error{}

	return errors
}
