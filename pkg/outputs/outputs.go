package outputs

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/aws/ekscluster"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

func PulumiOutputsToStackOutputsConverter(pulumiOutputs auto.OutputMap,
	input *ekscluster.EksClusterStackInput) *ekscluster.EksClusterStackOutputs {
	return &ekscluster.EksClusterStackOutputs{}
}
