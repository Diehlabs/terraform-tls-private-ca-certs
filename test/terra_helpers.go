package test

import(
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

func setupTesting(t *testing.T, workingDir string, terraformBinary string, terraformVars map[string]interface{}) (*terraform.Options) {

	testDataExists := test_structure.IsTestDataPresent(t, test_structure.FormatTestDataPath(workingDir, "TerraformOptions.json"))

	if (testDataExists) {
		logger.Logf(t, "Found and loaded test data in %s", workingDir)
		return test_structure.LoadTerraformOptions(t, workingDir)
	} else {

		terraformOptions := &terraform.Options{
			TerraformDir:    workingDir,
			TerraformBinary: terraformBinary,
			Vars: terraformVars,
		}

		test_structure.SaveTerraformOptions(t, workingDir, terraformOptions)

		logger.Logf(t, "Saved test data in %s so it can be reused later", workingDir)

		return terraformOptions
	}
}

func deployUsingTerraform(t *testing.T, workingDir string) {
	terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)
	terraform.InitAndApply(t, terraformOptions)
}

func redeployUsingTerraform(t *testing.T, workingDir string) {
	terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)
	terraform.ApplyAndIdempotent(t, terraformOptions)
}

func terraform_destroy(t *testing.T, workingDir string) {
	terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)
	terraform.Destroy(t, terraformOptions)
	test_structure.CleanupTestDataFolder(t, workingDir)
}
