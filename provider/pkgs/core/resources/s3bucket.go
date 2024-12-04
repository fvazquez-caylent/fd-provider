package resources

import (
	"github.com/pulumi/pulumi-go-provider/infer"
)

// S3Bucket representa un recurso S3 en el sistema de inferencia.
type S3Bucket struct {
	Name string `pulumi:"name"`
}

// S3BucketArgs contiene los argumentos necesarios para crear un S3Bucket.
type S3BucketArgs struct {
	Name string `pulumi:"name"`
}

// NewS3Bucket es la fábrica del recurso S3Bucket compatible con infer.
func NewS3Bucket() infer.InferredResource {
	return infer.Resource[S3Bucket, S3BucketArgs]()
}

