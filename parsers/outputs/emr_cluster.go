package outputs
import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
)

func ParseEMRCluster(name string, data string) (cf types.ValueMap, err error) {
	
	var resource, output types.ValueMap
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	
	cf = types.ValueMap{
		name: types.ValueMap{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-EMRCluster-" + name,
				},
			},
		},
	}

	
	output = types.ValueMap{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "MasterPublicDNS"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-EMRCluster-" + name + "-MasterPublicDNS",
			},
		},
	}
	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}
	cf[name+"MasterPublicDNS"] = output
	

	return
}
