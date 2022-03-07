package test

import (
	"os"
	"testing"
	"crypto/tls"
	"strings"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	// "github.com/stretchr/testify/assert"
)

var uniqueId = random.UniqueId()

var terraformBinary = "/usr/local/bin/terraform"

var workingDir = "../examples/build"

func TestVmss(t *testing.T) {
	t.Parallel()

	//os.Setenv("SKIP_reterraform_deploy", "true")
	//os.Setenv("SKIP_terraform_redeploy", "true")
	//os.Setenv("SKIP_terraform_destroy", "true")

	if tfbin := os.Getenv("TF_CLI_PATH"); tfbin != "" {
		terraformBinary = tfbin
	}

	if tfdir := os.Getenv("TERRATEST_WORKING_DIR"); tfdir != "" {
		workingDir = tfdir
	}

	terraformVars := map[string]interface{}{
	}

	terraformOptions := setupTesting(t, workingDir, terraformBinary, terraformVars)

	// Destroy the infra after testing is finished
	defer test_structure.RunTestStage(t, "terraform_destroy", func(){
		terraform_destroy(t, workingDir)
	})

	// Deploy using Terraform
	test_structure.RunTestStage(t, "terraform_deploy", func() {
		deployUsingTerraform(t, workingDir)
	})

	// Redeploy using Terraform and ensure idempotency
	test_structure.RunTestStage(t, "terraform_redeploy", func() {
		redeployUsingTerraform(t, workingDir)
	})

	rsaCertPEM := terraform.Output(t, terraformOptions, "cert_host1")
	rsaKeyPEM := terraform.Output(t, terraformOptions, "key_host1")

	t.Run("Test generated server cert", func(t *testing.T){
		testCerts(t, rsaCertPEM, rsaKeyPEM)
	})

}

func testCerts(t *testing.T, rsaCertPEM string, rsaKeyPEM string) {

	_, err := tls.X509KeyPair([]byte(rsaKeyPEM), []byte(rsaCertPEM))
	if err == nil {
		t.Fatalf("X509KeyPair didn't return an error when arguments were switched")
	}
	if subStr := "been switched"; !strings.Contains(err.Error(), subStr) {
		t.Fatalf("Expected %q in the error when switching arguments to X509KeyPair, but the error was %q", subStr, err)
	}

	_, err = tls.X509KeyPair([]byte(rsaCertPEM), []byte(rsaCertPEM))
	if err == nil {
		t.Fatalf("X509KeyPair didn't return an error when both arguments were certificates")
	}
	if subStr := "certificate"; !strings.Contains(err.Error(), subStr) {
		t.Fatalf("Expected %q in the error when both arguments to X509KeyPair were certificates, but the error was %q", subStr, err)
	}

}
