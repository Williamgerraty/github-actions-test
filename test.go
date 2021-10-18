package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// An example of how to test the Terraform module in examples/terraform-aws-s3-example using Terratest.
func TestTerraformAwsS3Example(t *testing.T) {
	t.Parallel()

	
	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t)

	// Run `terraform output` to get the value of an output variable
	bucketID := terraform.Output(t)

	// Verify that our Bucket has versioning enabled
	actualStatus := aws.GetS3BucketVersioning(t, "ap-southeast-2", "willjg227272727727271")
	expectedStatus := "Enabled"
	assert.Equal(t, expectedStatus, actualStatus)
}