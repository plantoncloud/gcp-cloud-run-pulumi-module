package main

import (
	gcpcloudrunv1 "buf.build/gen/go/plantoncloud/project-planton/protocolbuffers/go/project/planton/provider/gcp/gcpcloudrun/v1"
	"github.com/pkg/errors"
	"github.com/plantoncloud/gcp-cloud-run-pulumi-module/pkg"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/stackinput"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		stackInput := &gcpcloudrunv1.GcpCloudRunStackInput{}

		if err := stackinput.LoadStackInput(ctx, stackInput); err != nil {
			return errors.Wrap(err, "failed to load stack-input")
		}

		return pkg.Resources(ctx, stackInput)
	})
}
