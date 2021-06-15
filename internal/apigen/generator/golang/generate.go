package golang

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/redhat-developer/app-services-sdk-go/internal/apigen/generator/common"
)

type GoGen struct {
	Config *common.Config
}

// Generate generates an API client in Golang
func (g *GoGen) Generate() (err error) {
	packageName := fmt.Sprintf("%vclient", g.Config.ClientMetadata.APIGroup)
	apiVersion := fmt.Sprintf("api%v", g.Config.ClientMetadata.APIVersion)
	outputPath := fmt.Sprintf("%v/%v/client", g.Config.ClientMetadata.APIGroup, apiVersion)
	inputFile := g.Config.Input
	if _, err = os.Stat(inputFile); err != nil {
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

	out, err := exec.Command("npx", cmdArgs...).CombinedOutput()
	fmt.Fprintln(os.Stderr, string(out))
	if err != nil {
		return err
	}

	// err = generateMocks(outputPath)
	// if err != nil {
	// 	return err
	// }

	return nil
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
