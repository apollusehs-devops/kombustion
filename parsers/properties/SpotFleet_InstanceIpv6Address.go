package properties

	import "fmt"

type SpotFleet_InstanceIpv6Address struct {
	
	Ipv6Address interface{} `yaml:"Ipv6Address"`
}

func (resource SpotFleet_InstanceIpv6Address) Validate() []error {
	errs := []error{}
	
	if resource.Ipv6Address == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Ipv6Address'"))
	}
	return errs
}
