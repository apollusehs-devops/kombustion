package parsers


import (
  "github.com/KablamoOSS/kombustion/types"
  "github.com/KablamoOSS/kombustion/parsers/outputs"
)

func GetParsers_outputs() map[string]types.ParserFunc {
	return map[string]types.ParserFunc{
		
		"AWS::ApiGateway::Account": outputs.ParseApiGatewayAccount,
		
		"AWS::ApiGateway::ApiKey": outputs.ParseApiGatewayApiKey,
		
		"AWS::ApiGateway::Authorizer": outputs.ParseApiGatewayAuthorizer,
		
		"AWS::ApiGateway::BasePathMapping": outputs.ParseApiGatewayBasePathMapping,
		
		"AWS::ApiGateway::ClientCertificate": outputs.ParseApiGatewayClientCertificate,
		
		"AWS::ApiGateway::Deployment": outputs.ParseApiGatewayDeployment,
		
		"AWS::ApiGateway::DocumentationPart": outputs.ParseApiGatewayDocumentationPart,
		
		"AWS::ApiGateway::DocumentationVersion": outputs.ParseApiGatewayDocumentationVersion,
		
		"AWS::ApiGateway::DomainName": outputs.ParseApiGatewayDomainName,
		
		"AWS::ApiGateway::GatewayResponse": outputs.ParseApiGatewayGatewayResponse,
		
		"AWS::ApiGateway::Method": outputs.ParseApiGatewayMethod,
		
		"AWS::ApiGateway::Model": outputs.ParseApiGatewayModel,
		
		"AWS::ApiGateway::RequestValidator": outputs.ParseApiGatewayRequestValidator,
		
		"AWS::ApiGateway::Resource": outputs.ParseApiGatewayResource,
		
		"AWS::ApiGateway::RestApi": outputs.ParseApiGatewayRestApi,
		
		"AWS::ApiGateway::Stage": outputs.ParseApiGatewayStage,
		
		"AWS::ApiGateway::UsagePlan": outputs.ParseApiGatewayUsagePlan,
		
		"AWS::ApiGateway::UsagePlanKey": outputs.ParseApiGatewayUsagePlanKey,
		
		"AWS::ApiGateway::VpcLink": outputs.ParseApiGatewayVpcLink,
		
		"AWS::AppSync::ApiKey": outputs.ParseAppSyncApiKey,
		
		"AWS::AppSync::DataSource": outputs.ParseAppSyncDataSource,
		
		"AWS::AppSync::GraphQLApi": outputs.ParseAppSyncGraphQLApi,
		
		"AWS::AppSync::GraphQLSchema": outputs.ParseAppSyncGraphQLSchema,
		
		"AWS::AppSync::Resolver": outputs.ParseAppSyncResolver,
		
		"AWS::ApplicationAutoScaling::ScalableTarget": outputs.ParseApplicationAutoScalingScalableTarget,
		
		"AWS::ApplicationAutoScaling::ScalingPolicy": outputs.ParseApplicationAutoScalingScalingPolicy,
		
		"AWS::Athena::NamedQuery": outputs.ParseAthenaNamedQuery,
		
		"AWS::AutoScaling::AutoScalingGroup": outputs.ParseAutoScalingAutoScalingGroup,
		
		"AWS::AutoScaling::LaunchConfiguration": outputs.ParseAutoScalingLaunchConfiguration,
		
		"AWS::AutoScaling::LifecycleHook": outputs.ParseAutoScalingLifecycleHook,
		
		"AWS::AutoScaling::ScalingPolicy": outputs.ParseAutoScalingScalingPolicy,
		
		"AWS::AutoScaling::ScheduledAction": outputs.ParseAutoScalingScheduledAction,
		
		"AWS::AutoScalingPlans::ScalingPlan": outputs.ParseAutoScalingPlansScalingPlan,
		
		"AWS::Batch::ComputeEnvironment": outputs.ParseBatchComputeEnvironment,
		
		"AWS::Batch::JobDefinition": outputs.ParseBatchJobDefinition,
		
		"AWS::Batch::JobQueue": outputs.ParseBatchJobQueue,
		
		"AWS::CertificateManager::Certificate": outputs.ParseCertificateManagerCertificate,
		
		"AWS::Cloud9::EnvironmentEC2": outputs.ParseCloud9EnvironmentEC2,
		
		"AWS::CloudFormation::CustomResource": outputs.ParseCloudFormationCustomResource,
		
		"AWS::CloudFormation::Stack": outputs.ParseCloudFormationStack,
		
		"AWS::CloudFormation::WaitCondition": outputs.ParseCloudFormationWaitCondition,
		
		"AWS::CloudFormation::WaitConditionHandle": outputs.ParseCloudFormationWaitConditionHandle,
		
		"AWS::CloudFront::CloudFrontOriginAccessIdentity": outputs.ParseCloudFrontCloudFrontOriginAccessIdentity,
		
		"AWS::CloudFront::Distribution": outputs.ParseCloudFrontDistribution,
		
		"AWS::CloudFront::StreamingDistribution": outputs.ParseCloudFrontStreamingDistribution,
		
		"AWS::CloudTrail::Trail": outputs.ParseCloudTrailTrail,
		
		"AWS::CloudWatch::Alarm": outputs.ParseCloudWatchAlarm,
		
		"AWS::CloudWatch::Dashboard": outputs.ParseCloudWatchDashboard,
		
		"AWS::CodeBuild::Project": outputs.ParseCodeBuildProject,
		
		"AWS::CodeCommit::Repository": outputs.ParseCodeCommitRepository,
		
		"AWS::CodeDeploy::Application": outputs.ParseCodeDeployApplication,
		
		"AWS::CodeDeploy::DeploymentConfig": outputs.ParseCodeDeployDeploymentConfig,
		
		"AWS::CodeDeploy::DeploymentGroup": outputs.ParseCodeDeployDeploymentGroup,
		
		"AWS::CodePipeline::CustomActionType": outputs.ParseCodePipelineCustomActionType,
		
		"AWS::CodePipeline::Pipeline": outputs.ParseCodePipelinePipeline,
		
		"AWS::Cognito::IdentityPool": outputs.ParseCognitoIdentityPool,
		
		"AWS::Cognito::IdentityPoolRoleAttachment": outputs.ParseCognitoIdentityPoolRoleAttachment,
		
		"AWS::Cognito::UserPool": outputs.ParseCognitoUserPool,
		
		"AWS::Cognito::UserPoolClient": outputs.ParseCognitoUserPoolClient,
		
		"AWS::Cognito::UserPoolGroup": outputs.ParseCognitoUserPoolGroup,
		
		"AWS::Cognito::UserPoolUser": outputs.ParseCognitoUserPoolUser,
		
		"AWS::Cognito::UserPoolUserToGroupAttachment": outputs.ParseCognitoUserPoolUserToGroupAttachment,
		
		"AWS::Config::ConfigRule": outputs.ParseConfigConfigRule,
		
		"AWS::Config::ConfigurationRecorder": outputs.ParseConfigConfigurationRecorder,
		
		"AWS::Config::DeliveryChannel": outputs.ParseConfigDeliveryChannel,
		
		"AWS::DAX::Cluster": outputs.ParseDAXCluster,
		
		"AWS::DAX::ParameterGroup": outputs.ParseDAXParameterGroup,
		
		"AWS::DAX::SubnetGroup": outputs.ParseDAXSubnetGroup,
		
		"AWS::DMS::Certificate": outputs.ParseDMSCertificate,
		
		"AWS::DMS::Endpoint": outputs.ParseDMSEndpoint,
		
		"AWS::DMS::EventSubscription": outputs.ParseDMSEventSubscription,
		
		"AWS::DMS::ReplicationInstance": outputs.ParseDMSReplicationInstance,
		
		"AWS::DMS::ReplicationSubnetGroup": outputs.ParseDMSReplicationSubnetGroup,
		
		"AWS::DMS::ReplicationTask": outputs.ParseDMSReplicationTask,
		
		"AWS::DataPipeline::Pipeline": outputs.ParseDataPipelinePipeline,
		
		"AWS::DirectoryService::MicrosoftAD": outputs.ParseDirectoryServiceMicrosoftAD,
		
		"AWS::DirectoryService::SimpleAD": outputs.ParseDirectoryServiceSimpleAD,
		
		"AWS::DynamoDB::Table": outputs.ParseDynamoDBTable,
		
		"AWS::EC2::CustomerGateway": outputs.ParseEC2CustomerGateway,
		
		"AWS::EC2::DHCPOptions": outputs.ParseEC2DHCPOptions,
		
		"AWS::EC2::EIP": outputs.ParseEC2EIP,
		
		"AWS::EC2::EIPAssociation": outputs.ParseEC2EIPAssociation,
		
		"AWS::EC2::EgressOnlyInternetGateway": outputs.ParseEC2EgressOnlyInternetGateway,
		
		"AWS::EC2::FlowLog": outputs.ParseEC2FlowLog,
		
		"AWS::EC2::Host": outputs.ParseEC2Host,
		
		"AWS::EC2::Instance": outputs.ParseEC2Instance,
		
		"AWS::EC2::InternetGateway": outputs.ParseEC2InternetGateway,
		
		"AWS::EC2::LaunchTemplate": outputs.ParseEC2LaunchTemplate,
		
		"AWS::EC2::NatGateway": outputs.ParseEC2NatGateway,
		
		"AWS::EC2::NetworkAcl": outputs.ParseEC2NetworkAcl,
		
		"AWS::EC2::NetworkAclEntry": outputs.ParseEC2NetworkAclEntry,
		
		"AWS::EC2::NetworkInterface": outputs.ParseEC2NetworkInterface,
		
		"AWS::EC2::NetworkInterfaceAttachment": outputs.ParseEC2NetworkInterfaceAttachment,
		
		"AWS::EC2::NetworkInterfacePermission": outputs.ParseEC2NetworkInterfacePermission,
		
		"AWS::EC2::PlacementGroup": outputs.ParseEC2PlacementGroup,
		
		"AWS::EC2::Route": outputs.ParseEC2Route,
		
		"AWS::EC2::RouteTable": outputs.ParseEC2RouteTable,
		
		"AWS::EC2::SecurityGroup": outputs.ParseEC2SecurityGroup,
		
		"AWS::EC2::SecurityGroupEgress": outputs.ParseEC2SecurityGroupEgress,
		
		"AWS::EC2::SecurityGroupIngress": outputs.ParseEC2SecurityGroupIngress,
		
		"AWS::EC2::SpotFleet": outputs.ParseEC2SpotFleet,
		
		"AWS::EC2::Subnet": outputs.ParseEC2Subnet,
		
		"AWS::EC2::SubnetCidrBlock": outputs.ParseEC2SubnetCidrBlock,
		
		"AWS::EC2::SubnetNetworkAclAssociation": outputs.ParseEC2SubnetNetworkAclAssociation,
		
		"AWS::EC2::SubnetRouteTableAssociation": outputs.ParseEC2SubnetRouteTableAssociation,
		
		"AWS::EC2::TrunkInterfaceAssociation": outputs.ParseEC2TrunkInterfaceAssociation,
		
		"AWS::EC2::VPC": outputs.ParseEC2VPC,
		
		"AWS::EC2::VPCCidrBlock": outputs.ParseEC2VPCCidrBlock,
		
		"AWS::EC2::VPCDHCPOptionsAssociation": outputs.ParseEC2VPCDHCPOptionsAssociation,
		
		"AWS::EC2::VPCEndpoint": outputs.ParseEC2VPCEndpoint,
		
		"AWS::EC2::VPCGatewayAttachment": outputs.ParseEC2VPCGatewayAttachment,
		
		"AWS::EC2::VPCPeeringConnection": outputs.ParseEC2VPCPeeringConnection,
		
		"AWS::EC2::VPNConnection": outputs.ParseEC2VPNConnection,
		
		"AWS::EC2::VPNConnectionRoute": outputs.ParseEC2VPNConnectionRoute,
		
		"AWS::EC2::VPNGateway": outputs.ParseEC2VPNGateway,
		
		"AWS::EC2::VPNGatewayRoutePropagation": outputs.ParseEC2VPNGatewayRoutePropagation,
		
		"AWS::EC2::Volume": outputs.ParseEC2Volume,
		
		"AWS::EC2::VolumeAttachment": outputs.ParseEC2VolumeAttachment,
		
		"AWS::ECR::Repository": outputs.ParseECRRepository,
		
		"AWS::ECS::Cluster": outputs.ParseECSCluster,
		
		"AWS::ECS::Service": outputs.ParseECSService,
		
		"AWS::ECS::TaskDefinition": outputs.ParseECSTaskDefinition,
		
		"AWS::EFS::FileSystem": outputs.ParseEFSFileSystem,
		
		"AWS::EFS::MountTarget": outputs.ParseEFSMountTarget,
		
		"AWS::EMR::Cluster": outputs.ParseEMRCluster,
		
		"AWS::EMR::InstanceFleetConfig": outputs.ParseEMRInstanceFleetConfig,
		
		"AWS::EMR::InstanceGroupConfig": outputs.ParseEMRInstanceGroupConfig,
		
		"AWS::EMR::SecurityConfiguration": outputs.ParseEMRSecurityConfiguration,
		
		"AWS::EMR::Step": outputs.ParseEMRStep,
		
		"AWS::ElastiCache::CacheCluster": outputs.ParseElastiCacheCacheCluster,
		
		"AWS::ElastiCache::ParameterGroup": outputs.ParseElastiCacheParameterGroup,
		
		"AWS::ElastiCache::ReplicationGroup": outputs.ParseElastiCacheReplicationGroup,
		
		"AWS::ElastiCache::SecurityGroup": outputs.ParseElastiCacheSecurityGroup,
		
		"AWS::ElastiCache::SecurityGroupIngress": outputs.ParseElastiCacheSecurityGroupIngress,
		
		"AWS::ElastiCache::SubnetGroup": outputs.ParseElastiCacheSubnetGroup,
		
		"AWS::ElasticBeanstalk::Application": outputs.ParseElasticBeanstalkApplication,
		
		"AWS::ElasticBeanstalk::ApplicationVersion": outputs.ParseElasticBeanstalkApplicationVersion,
		
		"AWS::ElasticBeanstalk::ConfigurationTemplate": outputs.ParseElasticBeanstalkConfigurationTemplate,
		
		"AWS::ElasticBeanstalk::Environment": outputs.ParseElasticBeanstalkEnvironment,
		
		"AWS::ElasticLoadBalancing::LoadBalancer": outputs.ParseElasticLoadBalancingLoadBalancer,
		
		"AWS::ElasticLoadBalancingV2::Listener": outputs.ParseElasticLoadBalancingV2Listener,
		
		"AWS::ElasticLoadBalancingV2::ListenerCertificate": outputs.ParseElasticLoadBalancingV2ListenerCertificate,
		
		"AWS::ElasticLoadBalancingV2::ListenerRule": outputs.ParseElasticLoadBalancingV2ListenerRule,
		
		"AWS::ElasticLoadBalancingV2::LoadBalancer": outputs.ParseElasticLoadBalancingV2LoadBalancer,
		
		"AWS::ElasticLoadBalancingV2::TargetGroup": outputs.ParseElasticLoadBalancingV2TargetGroup,
		
		"AWS::Elasticsearch::Domain": outputs.ParseElasticsearchDomain,
		
		"AWS::Events::Rule": outputs.ParseEventsRule,
		
		"AWS::GameLift::Alias": outputs.ParseGameLiftAlias,
		
		"AWS::GameLift::Build": outputs.ParseGameLiftBuild,
		
		"AWS::GameLift::Fleet": outputs.ParseGameLiftFleet,
		
		"AWS::Glue::Classifier": outputs.ParseGlueClassifier,
		
		"AWS::Glue::Connection": outputs.ParseGlueConnection,
		
		"AWS::Glue::Crawler": outputs.ParseGlueCrawler,
		
		"AWS::Glue::Database": outputs.ParseGlueDatabase,
		
		"AWS::Glue::DevEndpoint": outputs.ParseGlueDevEndpoint,
		
		"AWS::Glue::Job": outputs.ParseGlueJob,
		
		"AWS::Glue::Partition": outputs.ParseGluePartition,
		
		"AWS::Glue::Table": outputs.ParseGlueTable,
		
		"AWS::Glue::Trigger": outputs.ParseGlueTrigger,
		
		"AWS::GuardDuty::Detector": outputs.ParseGuardDutyDetector,
		
		"AWS::GuardDuty::Filter": outputs.ParseGuardDutyFilter,
		
		"AWS::GuardDuty::IPSet": outputs.ParseGuardDutyIPSet,
		
		"AWS::GuardDuty::Master": outputs.ParseGuardDutyMaster,
		
		"AWS::GuardDuty::Member": outputs.ParseGuardDutyMember,
		
		"AWS::GuardDuty::ThreatIntelSet": outputs.ParseGuardDutyThreatIntelSet,
		
		"AWS::IAM::AccessKey": outputs.ParseIAMAccessKey,
		
		"AWS::IAM::Group": outputs.ParseIAMGroup,
		
		"AWS::IAM::InstanceProfile": outputs.ParseIAMInstanceProfile,
		
		"AWS::IAM::ManagedPolicy": outputs.ParseIAMManagedPolicy,
		
		"AWS::IAM::Policy": outputs.ParseIAMPolicy,
		
		"AWS::IAM::Role": outputs.ParseIAMRole,
		
		"AWS::IAM::User": outputs.ParseIAMUser,
		
		"AWS::IAM::UserToGroupAddition": outputs.ParseIAMUserToGroupAddition,
		
		"AWS::Inspector::AssessmentTarget": outputs.ParseInspectorAssessmentTarget,
		
		"AWS::Inspector::AssessmentTemplate": outputs.ParseInspectorAssessmentTemplate,
		
		"AWS::Inspector::ResourceGroup": outputs.ParseInspectorResourceGroup,
		
		"AWS::IoT::Certificate": outputs.ParseIoTCertificate,
		
		"AWS::IoT::Policy": outputs.ParseIoTPolicy,
		
		"AWS::IoT::PolicyPrincipalAttachment": outputs.ParseIoTPolicyPrincipalAttachment,
		
		"AWS::IoT::Thing": outputs.ParseIoTThing,
		
		"AWS::IoT::ThingPrincipalAttachment": outputs.ParseIoTThingPrincipalAttachment,
		
		"AWS::IoT::TopicRule": outputs.ParseIoTTopicRule,
		
		"AWS::KMS::Alias": outputs.ParseKMSAlias,
		
		"AWS::KMS::Key": outputs.ParseKMSKey,
		
		"AWS::Kinesis::Stream": outputs.ParseKinesisStream,
		
		"AWS::KinesisAnalytics::Application": outputs.ParseKinesisAnalyticsApplication,
		
		"AWS::KinesisAnalytics::ApplicationOutput": outputs.ParseKinesisAnalyticsApplicationOutput,
		
		"AWS::KinesisAnalytics::ApplicationReferenceDataSource": outputs.ParseKinesisAnalyticsApplicationReferenceDataSource,
		
		"AWS::KinesisFirehose::DeliveryStream": outputs.ParseKinesisFirehoseDeliveryStream,
		
		"AWS::Lambda::Alias": outputs.ParseLambdaAlias,
		
		"AWS::Lambda::EventSourceMapping": outputs.ParseLambdaEventSourceMapping,
		
		"AWS::Lambda::Function": outputs.ParseLambdaFunction,
		
		"AWS::Lambda::Permission": outputs.ParseLambdaPermission,
		
		"AWS::Lambda::Version": outputs.ParseLambdaVersion,
		
		"AWS::Logs::Destination": outputs.ParseLogsDestination,
		
		"AWS::Logs::LogGroup": outputs.ParseLogsLogGroup,
		
		"AWS::Logs::LogStream": outputs.ParseLogsLogStream,
		
		"AWS::Logs::MetricFilter": outputs.ParseLogsMetricFilter,
		
		"AWS::Logs::SubscriptionFilter": outputs.ParseLogsSubscriptionFilter,
		
		"AWS::OpsWorks::App": outputs.ParseOpsWorksApp,
		
		"AWS::OpsWorks::ElasticLoadBalancerAttachment": outputs.ParseOpsWorksElasticLoadBalancerAttachment,
		
		"AWS::OpsWorks::Instance": outputs.ParseOpsWorksInstance,
		
		"AWS::OpsWorks::Layer": outputs.ParseOpsWorksLayer,
		
		"AWS::OpsWorks::Stack": outputs.ParseOpsWorksStack,
		
		"AWS::OpsWorks::UserProfile": outputs.ParseOpsWorksUserProfile,
		
		"AWS::OpsWorks::Volume": outputs.ParseOpsWorksVolume,
		
		"AWS::RDS::DBCluster": outputs.ParseRDSDBCluster,
		
		"AWS::RDS::DBClusterParameterGroup": outputs.ParseRDSDBClusterParameterGroup,
		
		"AWS::RDS::DBInstance": outputs.ParseRDSDBInstance,
		
		"AWS::RDS::DBParameterGroup": outputs.ParseRDSDBParameterGroup,
		
		"AWS::RDS::DBSecurityGroup": outputs.ParseRDSDBSecurityGroup,
		
		"AWS::RDS::DBSecurityGroupIngress": outputs.ParseRDSDBSecurityGroupIngress,
		
		"AWS::RDS::DBSubnetGroup": outputs.ParseRDSDBSubnetGroup,
		
		"AWS::RDS::EventSubscription": outputs.ParseRDSEventSubscription,
		
		"AWS::RDS::OptionGroup": outputs.ParseRDSOptionGroup,
		
		"AWS::Redshift::Cluster": outputs.ParseRedshiftCluster,
		
		"AWS::Redshift::ClusterParameterGroup": outputs.ParseRedshiftClusterParameterGroup,
		
		"AWS::Redshift::ClusterSecurityGroup": outputs.ParseRedshiftClusterSecurityGroup,
		
		"AWS::Redshift::ClusterSecurityGroupIngress": outputs.ParseRedshiftClusterSecurityGroupIngress,
		
		"AWS::Redshift::ClusterSubnetGroup": outputs.ParseRedshiftClusterSubnetGroup,
		
		"AWS::Route53::HealthCheck": outputs.ParseRoute53HealthCheck,
		
		"AWS::Route53::HostedZone": outputs.ParseRoute53HostedZone,
		
		"AWS::Route53::RecordSet": outputs.ParseRoute53RecordSet,
		
		"AWS::Route53::RecordSetGroup": outputs.ParseRoute53RecordSetGroup,
		
		"AWS::S3::Bucket": outputs.ParseS3Bucket,
		
		"AWS::S3::BucketPolicy": outputs.ParseS3BucketPolicy,
		
		"AWS::SDB::Domain": outputs.ParseSDBDomain,
		
		"AWS::SES::ConfigurationSet": outputs.ParseSESConfigurationSet,
		
		"AWS::SES::ConfigurationSetEventDestination": outputs.ParseSESConfigurationSetEventDestination,
		
		"AWS::SES::ReceiptFilter": outputs.ParseSESReceiptFilter,
		
		"AWS::SES::ReceiptRule": outputs.ParseSESReceiptRule,
		
		"AWS::SES::ReceiptRuleSet": outputs.ParseSESReceiptRuleSet,
		
		"AWS::SES::Template": outputs.ParseSESTemplate,
		
		"AWS::SNS::Subscription": outputs.ParseSNSSubscription,
		
		"AWS::SNS::Topic": outputs.ParseSNSTopic,
		
		"AWS::SNS::TopicPolicy": outputs.ParseSNSTopicPolicy,
		
		"AWS::SQS::Queue": outputs.ParseSQSQueue,
		
		"AWS::SQS::QueuePolicy": outputs.ParseSQSQueuePolicy,
		
		"AWS::SSM::Association": outputs.ParseSSMAssociation,
		
		"AWS::SSM::Document": outputs.ParseSSMDocument,
		
		"AWS::SSM::MaintenanceWindowTask": outputs.ParseSSMMaintenanceWindowTask,
		
		"AWS::SSM::Parameter": outputs.ParseSSMParameter,
		
		"AWS::SSM::PatchBaseline": outputs.ParseSSMPatchBaseline,
		
		"AWS::ServiceCatalog::CloudFormationProvisionedProduct": outputs.ParseServiceCatalogCloudFormationProvisionedProduct,
		
		"AWS::ServiceDiscovery::Instance": outputs.ParseServiceDiscoveryInstance,
		
		"AWS::ServiceDiscovery::PrivateDnsNamespace": outputs.ParseServiceDiscoveryPrivateDnsNamespace,
		
		"AWS::ServiceDiscovery::PublicDnsNamespace": outputs.ParseServiceDiscoveryPublicDnsNamespace,
		
		"AWS::ServiceDiscovery::Service": outputs.ParseServiceDiscoveryService,
		
		"AWS::StepFunctions::Activity": outputs.ParseStepFunctionsActivity,
		
		"AWS::StepFunctions::StateMachine": outputs.ParseStepFunctionsStateMachine,
		
		"AWS::WAF::ByteMatchSet": outputs.ParseWAFByteMatchSet,
		
		"AWS::WAF::IPSet": outputs.ParseWAFIPSet,
		
		"AWS::WAF::Rule": outputs.ParseWAFRule,
		
		"AWS::WAF::SizeConstraintSet": outputs.ParseWAFSizeConstraintSet,
		
		"AWS::WAF::SqlInjectionMatchSet": outputs.ParseWAFSqlInjectionMatchSet,
		
		"AWS::WAF::WebACL": outputs.ParseWAFWebACL,
		
		"AWS::WAF::XssMatchSet": outputs.ParseWAFXssMatchSet,
		
		"AWS::WAFRegional::ByteMatchSet": outputs.ParseWAFRegionalByteMatchSet,
		
		"AWS::WAFRegional::IPSet": outputs.ParseWAFRegionalIPSet,
		
		"AWS::WAFRegional::Rule": outputs.ParseWAFRegionalRule,
		
		"AWS::WAFRegional::SizeConstraintSet": outputs.ParseWAFRegionalSizeConstraintSet,
		
		"AWS::WAFRegional::SqlInjectionMatchSet": outputs.ParseWAFRegionalSqlInjectionMatchSet,
		
		"AWS::WAFRegional::WebACL": outputs.ParseWAFRegionalWebACL,
		
		"AWS::WAFRegional::WebACLAssociation": outputs.ParseWAFRegionalWebACLAssociation,
		
		"AWS::WAFRegional::XssMatchSet": outputs.ParseWAFRegionalXssMatchSet,
		
		"AWS::WorkSpaces::Workspace": outputs.ParseWorkSpacesWorkspace,
		
	}
}
