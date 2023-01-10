package parser_test

import (
	"testing"

	"github.com/jchen42703/create-fullstack/internal/parser"
)

func TestYamlToTemplateCfg(t *testing.T) {

	// yaml file where it only contains fullstack key
	fstack, err := parser.YamlToTemplateCfg("./test_files/fullstack_only.yaml")

	if err != nil {
		t.Fatalf("YamlToAugmentCfg failed with err: %s", err)
	}

	if fstack.FullstackCfg.OutputDirectoryPath != "./fullstack_app_name" {
		t.Fatalf("not the same output dir")
	}

	if !fstack.FullstackCfg.AuthOpts.UsernamePasswordOpts.EmailVerification {
		t.Fatalf("email verification does not match")
	}

	if !fstack.FullstackCfg.AuthOpts.UsernamePasswordOpts.UsernameIsEmail {
		t.Fatalf("user is email does not match")
	}

	// yaml file where it only contains ui key
	ui_config, err := parser.YamlToTemplateCfg("./test_files/ui_only.yaml")

	if err != nil {
		t.Fatalf("YamlToAugmentCfg failed with err: %s", err)
	}

	if ui_config.UiCfg.OutputDirectoryPath != "./fullstack_app_name/ui" {
		t.Fatalf("not the same output dir")
	}

	if ui_config.UiCfg.Base != "nextjs" {
		t.Fatalf("Base is not the same")
	}

	if ui_config.UiCfg.Language != "typescript" {
		t.Fatalf("Language is not correct, expected: ")
	}

	if ui_config.UiCfg.AugmentOpts.AddTailwind.Version != "default" {
		t.Fatalf("Incorrect version for tailwind")
	}

	if ui_config.UiCfg.AugmentOpts.AddScss.Version != "default" {
		t.Fatalf("Incorrect version for scss")
	}

	if ui_config.UiCfg.AugmentOpts.AddStyledComponents.Version != "default" {
		t.Fatalf("Incorrect version for styled components")
	}

	if ui_config.UiCfg.AugmentOpts.HuskyOpts.CommitLint.Version != "default" {
		t.Fatalf("Incorrect husky commit lint version")
	}

	if !ui_config.UiCfg.AugmentOpts.AddDockerfile {
		t.Fatalf("Dockerfile option not set correctly")
	}

	if ui_config.UiCfg.AugmentOpts.AddCi != "git_workflows" {
		t.Fatalf("Incorrect CI selection")
	}

	// yaml file where it only contains api key
	api_config, err := parser.YamlToTemplateCfg("./test_files/backend_only.yaml")

	if err != nil {
		t.Fatalf("YamlToAugmentCfg failed with err: %s", err)
	}

	if api_config.ApiCfg.OutputDirectoryPath != "./fullstack_app_name/api" {
		t.Fatalf("Incorrect output dir")
	}

	if api_config.ApiCfg.Base != "echo" {
		t.Fatalf("Incorrect base")
	}

	if api_config.ApiCfg.Language != "go" {
		t.Fatalf("Incorrect language")
	}

	if api_config.ApiCfg.Databases.SQL.DbType != "postgres" {
		t.Fatalf("Database is not postgres")
	}

	if api_config.ApiCfg.Databases.SQL.StartupScript != "" {
		t.Fatalf("startup script should be empty")
	}

}
