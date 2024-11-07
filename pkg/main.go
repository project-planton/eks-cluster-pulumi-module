package pkg

import (
	eksclusterv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/aws/ekscluster/v1"
	"fmt"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *eksclusterv1.EksClusterStackInput) error {
	// Create a variable with descriptive name for the API resource in the input
	eksCluster := stackInput.Target

	awsCredential := stackInput.AwsCredential

	// Create AWS provider using the credentials from the input
	provider, err := aws.NewProvider(ctx,
		"classic-provider",
		&aws.ProviderArgs{
			AccessKey: pulumi.String(awsCredential.AccessKeyId),
			SecretKey: pulumi.String(awsCredential.SecretAccessKey),
			Region:    pulumi.String(awsCredential.Region),
		})
	if err != nil {
		return errors.Wrap(err, "failed to create AWS provider")
	}

	// Prepare SubnetIds and SecurityGroupIds as pulumi.StringArray
	subnetIds := pulumi.ToStringArray(eksCluster.Spec.Subnets)
	securityGroupIds := pulumi.ToStringArray(eksCluster.Spec.SecurityGroups)

	// Create EKS cluster
	eksClusterResource, err := eks.NewCluster(ctx,
		"eksCluster",
		&eks.ClusterArgs{
			Name:    pulumi.String(eksCluster.Metadata.Name),
			RoleArn: pulumi.String(eksCluster.Spec.RoleArn),
			VpcConfig: &eks.ClusterVpcConfigArgs{
				VpcId:            pulumi.String(eksCluster.Spec.VpcId),
				SubnetIds:        subnetIds,
				SecurityGroupIds: securityGroupIds,
			},
		},
		pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "failed to create EKS cluster")
	}

	// Create managed node group
	managedNodeGroup, err := eks.NewNodeGroup(ctx,
		"eksNodeGroup",
		&eks.NodeGroupArgs{
			ClusterName:   eksClusterResource.Name,
			NodeGroupName: pulumi.String(fmt.Sprintf("%s-node-group", eksCluster.Metadata.Name)),
			NodeRoleArn:   pulumi.String(eksCluster.Spec.NodeRoleArn),
			SubnetIds:     pulumi.ToStringArray(eksCluster.Spec.Subnets),
			InstanceTypes: pulumi.StringArray{
				pulumi.String(eksCluster.Spec.InstanceType),
			},
			ScalingConfig: &eks.NodeGroupScalingConfigArgs{
				DesiredSize: pulumi.Int(eksCluster.Spec.DesiredSize),
				MaxSize:     pulumi.Int(eksCluster.Spec.MaxSize),
				MinSize:     pulumi.Int(eksCluster.Spec.MinSize),
			},
		},
		pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "failed to create EKS node group")
	}

	// Export outputs
	ctx.Export("eksClusterName", eksClusterResource.Name)
	ctx.Export("eksNodeGroupName", managedNodeGroup.NodeGroupName)

	return nil
}
