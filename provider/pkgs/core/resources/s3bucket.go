package resources

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type S3Bucket struct {
	pulumi.CustomResourceState

	Name pulumi.StringOutput `pulumi:"name"`
}

func NewS3Bucket(ctx *pulumi.Context, name string, args *S3BucketArgs, opts ...pulumi.ResourceOption) (*S3Bucket, error) {
	bucket, err := pulumi.NewCustomResource(ctx, name, &S3BucketArgs{}, opts...)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

type S3BucketArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}
