package properties

	import "fmt"

type InstanceFleetConfig_InstanceTypeConfig struct {
	
	
	
	
	
	
	BidPrice interface{} `yaml:"BidPrice,omitempty"`
	BidPriceAsPercentageOfOnDemandPrice interface{} `yaml:"BidPriceAsPercentageOfOnDemandPrice,omitempty"`
	InstanceType interface{} `yaml:"InstanceType"`
	WeightedCapacity interface{} `yaml:"WeightedCapacity,omitempty"`
	Configurations interface{} `yaml:"Configurations,omitempty"`
	EbsConfiguration *InstanceFleetConfig_EbsConfiguration `yaml:"EbsConfiguration,omitempty"`
}

func (resource InstanceFleetConfig_InstanceTypeConfig) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	if resource.InstanceType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InstanceType'"))
	}
	return errs
}
