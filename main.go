package main

import (
	"github.com/pkg/errors"
	"github.com/project-planton/eks-cluster-pulumi-module/pkg"
	eksclusterv1 "github.com/project-planton/project-planton/apis/go/project/planton/provider/aws/ekscluster/v1"
	"github.com/project-planton/pulumi-module-golang-commons/pkg/stackinput"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		stackInput := &eksclusterv1.EksClusterStackInput{}

		if err := stackinput.LoadStackInput(ctx, stackInput); err != nil {
			return errors.Wrap(err, "failed to load stack-input")
		}

		return pkg.Resources(ctx, stackInput)
	})
}
