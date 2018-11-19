package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// InstanceBlockDeviceMapping Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-mapping.html
type InstanceBlockDeviceMapping struct {
	DeviceName  interface{} `yaml:"DeviceName"`
	VirtualName interface{} `yaml:"VirtualName,omitempty"`
	NoDevice    interface{} `yaml:"NoDevice,omitempty"`
	Ebs         interface{} `yaml:"Ebs,omitempty"`
}

// InstanceBlockDeviceMapping validation
func (resource InstanceBlockDeviceMapping) Validate() []error {
	errors := []error{}

	return errors
}
