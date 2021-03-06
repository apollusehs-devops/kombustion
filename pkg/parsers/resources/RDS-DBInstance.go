package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// RDSDBInstance Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html
type RDSDBInstance struct {
	Type       string                  `yaml:"Type"`
	Properties RDSDBInstanceProperties `yaml:"Properties"`
	Condition  interface{}             `yaml:"Condition,omitempty"`
	Metadata   interface{}             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}             `yaml:"DependsOn,omitempty"`
}

// RDSDBInstance Properties
type RDSDBInstanceProperties struct {
	AllocatedStorage                   interface{} `yaml:"AllocatedStorage,omitempty"`
	AllowMajorVersionUpgrade           interface{} `yaml:"AllowMajorVersionUpgrade,omitempty"`
	AutoMinorVersionUpgrade            interface{} `yaml:"AutoMinorVersionUpgrade,omitempty"`
	AvailabilityZone                   interface{} `yaml:"AvailabilityZone,omitempty"`
	BackupRetentionPeriod              interface{} `yaml:"BackupRetentionPeriod,omitempty"`
	CharacterSetName                   interface{} `yaml:"CharacterSetName,omitempty"`
	CopyTagsToSnapshot                 interface{} `yaml:"CopyTagsToSnapshot,omitempty"`
	DBClusterIdentifier                interface{} `yaml:"DBClusterIdentifier,omitempty"`
	DBInstanceClass                    interface{} `yaml:"DBInstanceClass"`
	DBInstanceIdentifier               interface{} `yaml:"DBInstanceIdentifier,omitempty"`
	DBName                             interface{} `yaml:"DBName,omitempty"`
	DBParameterGroupName               interface{} `yaml:"DBParameterGroupName,omitempty"`
	DBSnapshotIdentifier               interface{} `yaml:"DBSnapshotIdentifier,omitempty"`
	DBSubnetGroupName                  interface{} `yaml:"DBSubnetGroupName,omitempty"`
	DeleteAutomatedBackups             interface{} `yaml:"DeleteAutomatedBackups,omitempty"`
	DeletionProtection                 interface{} `yaml:"DeletionProtection,omitempty"`
	Domain                             interface{} `yaml:"Domain,omitempty"`
	DomainIAMRoleName                  interface{} `yaml:"DomainIAMRoleName,omitempty"`
	EnableIAMDatabaseAuthentication    interface{} `yaml:"EnableIAMDatabaseAuthentication,omitempty"`
	EnablePerformanceInsights          interface{} `yaml:"EnablePerformanceInsights,omitempty"`
	Engine                             interface{} `yaml:"Engine,omitempty"`
	EngineVersion                      interface{} `yaml:"EngineVersion,omitempty"`
	Iops                               interface{} `yaml:"Iops,omitempty"`
	KmsKeyId                           interface{} `yaml:"KmsKeyId,omitempty"`
	LicenseModel                       interface{} `yaml:"LicenseModel,omitempty"`
	MasterUserPassword                 interface{} `yaml:"MasterUserPassword,omitempty"`
	MasterUsername                     interface{} `yaml:"MasterUsername,omitempty"`
	MonitoringInterval                 interface{} `yaml:"MonitoringInterval,omitempty"`
	MonitoringRoleArn                  interface{} `yaml:"MonitoringRoleArn,omitempty"`
	MultiAZ                            interface{} `yaml:"MultiAZ,omitempty"`
	OptionGroupName                    interface{} `yaml:"OptionGroupName,omitempty"`
	PerformanceInsightsKMSKeyId        interface{} `yaml:"PerformanceInsightsKMSKeyId,omitempty"`
	PerformanceInsightsRetentionPeriod interface{} `yaml:"PerformanceInsightsRetentionPeriod,omitempty"`
	Port                               interface{} `yaml:"Port,omitempty"`
	PreferredBackupWindow              interface{} `yaml:"PreferredBackupWindow,omitempty"`
	PreferredMaintenanceWindow         interface{} `yaml:"PreferredMaintenanceWindow,omitempty"`
	PromotionTier                      interface{} `yaml:"PromotionTier,omitempty"`
	PubliclyAccessible                 interface{} `yaml:"PubliclyAccessible,omitempty"`
	SourceDBInstanceIdentifier         interface{} `yaml:"SourceDBInstanceIdentifier,omitempty"`
	SourceRegion                       interface{} `yaml:"SourceRegion,omitempty"`
	StorageEncrypted                   interface{} `yaml:"StorageEncrypted,omitempty"`
	StorageType                        interface{} `yaml:"StorageType,omitempty"`
	Timezone                           interface{} `yaml:"Timezone,omitempty"`
	UseDefaultProcessorFeatures        interface{} `yaml:"UseDefaultProcessorFeatures,omitempty"`
	DBSecurityGroups                   interface{} `yaml:"DBSecurityGroups,omitempty"`
	EnableCloudwatchLogsExports        interface{} `yaml:"EnableCloudwatchLogsExports,omitempty"`
	ProcessorFeatures                  interface{} `yaml:"ProcessorFeatures,omitempty"`
	Tags                               interface{} `yaml:"Tags,omitempty"`
	VPCSecurityGroups                  interface{} `yaml:"VPCSecurityGroups,omitempty"`
}

// NewRDSDBInstance constructor creates a new RDSDBInstance
func NewRDSDBInstance(properties RDSDBInstanceProperties, deps ...interface{}) RDSDBInstance {
	return RDSDBInstance{
		Type:       "AWS::RDS::DBInstance",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseRDSDBInstance parses RDSDBInstance
func ParseRDSDBInstance(
	name string,
	data string,
) (
	source string,
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	transform types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-resources"

	// Resources
	var resource RDSDBInstance
	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	if validateErrs := resource.Properties.Validate(); len(errors) > 0 {
		errors = append(errors, validateErrs...)
		return
	}

	resources = types.TemplateObject{name: resource}

	// Outputs

	outputs = types.TemplateObject{
		name: types.TemplateObject{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-RDSDBInstance-" + name,
				},
			},
		},
	}

	return
}

// ParseRDSDBInstance validator
func (resource RDSDBInstance) Validate() []error {
	return resource.Properties.Validate()
}

// ParseRDSDBInstanceProperties validator
func (resource RDSDBInstanceProperties) Validate() []error {
	errors := []error{}
	return errors
}
