package golang

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/redhat-developer/app-services-sdk-go/internal/apigen/generator/common"

	"github.com/redhat-developer/app-services-sdk-go/internal/apigen/openapi"
)

type GoGen struct {
	Config *common.Config
}

// Generate generates an API client in Golang
func (g *GoGen) Generate() error {
	packageName := fmt.Sprintf("%vclient", g.Config.ClientMetadata.APIGroup)
	apiVersion := fmt.Sprintf("api%v", g.Config.ClientMetadata.APIVersion)
	outputPath := fmt.Sprintf("%v/%v/client", g.Config.ClientMetadata.APIGroup, apiVersion)
	inputFile, err := openapi.GetFileName(g.Config.Input)
	if err != nil {
		return err
	}
	_, err = os.Stat(inputFile)
	if err != nil {
		return err
	}

	cmdArgs := []string{
		common.CmdName,
		"generate",
		"-g", "go",
		"-i", inputFile,
		"-o", outputPath,
		"-p", "generateInterfaces=true",
		"--ignore-file-override", g.Config.IgnoreFileOverride,
		"--git-user-id", "redhat-developer",
		"--git-repo-id", fmt.Sprintf("%v/%v", "app-services-sdk-go", packageName),
		"--package-name", packageName,
	}
	if g.Config.TemplatesDir != "" {
		cmdArgs = append(cmdArgs, "-t", g.Config.TemplatesDir)
	}

	err = exec.Command("npx", cmdArgs...).Run()
	if err != nil {
		return err
	}

	err = generateMocks(outputPath)
	if err != nil {
		return err
	}

	// when all is done move the OpenAPI file to .openapi
	return os.Rename(inputFile, path.Join(".openapi", inputFile))
}

func generateMocks(output string) error {
	err := os.Setenv("GO111MODULE", "on")
	if err != nil {
		return err
	}
	err = exec.Command("go", "get", "github.com/matryer/moq").Run()
	if err != nil {
		return err
	}
	defaultMockApiFile := path.Join(output, "default_api_mock.go")

	err = exec.Command("rm", "-f", defaultMockApiFile).Run()
	if err != nil {
		return err
	}
	err = exec.Command("moq", "-out", defaultMockApiFile, output, "DefaultApi").Run()
	if err != nil {
		return err
	}
	return exec.Command("go", "mod", "tidy").Run()
}
