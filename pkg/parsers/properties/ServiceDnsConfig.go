package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ServiceDnsConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-service-dnsconfig.html
type ServiceDnsConfig struct {
	NamespaceId   interface{} `yaml:"NamespaceId,omitempty"`
	RoutingPolicy interface{} `yaml:"RoutingPolicy,omitempty"`
	DnsRecords    interface{} `yaml:"DnsRecords"`
}

// ServiceDnsConfig validation
func (resource ServiceDnsConfig) Validate() []error {
	errors := []error{}

	return errors
}
