// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package bedrock

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	bedrock_sdkv2 "github.com/aws/aws-sdk-go-v2/service/bedrock"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory: newCustomModelDataSource,
			Name:    "Custom Model",
		},
		{
			Factory: newCustomModelsDataSource,
			Name:    "Custom Models",
		},
		{
			Factory: newFoundationModelDataSource,
			Name:    "Foundation Model",
		},
		{
			Factory: newFoundationModelsDataSource,
			Name:    "Foundation Models",
		},
		{
			Factory: newInferenceProfilesDataSource,
			Name:    "Inference Profiles",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newCustomModelResource,
			Name:    "Custom Model",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "job_arn",
			},
		},
		{
			Factory: newModelInvocationLoggingConfigurationResource,
			Name:    "Model Invocation Logging Configuration",
		},
		{
			Factory: newProvisionedModelThroughputResource,
			Name:    "Provisioned Model Throughput",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "provisioned_model_arn",
			},
		},
		{
			Factory: newResourceGuardrail,
			Name:    "Guardrail",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "guardrail_arn",
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Bedrock
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*bedrock_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return bedrock_sdkv2.NewFromConfig(cfg,
		bedrock_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
