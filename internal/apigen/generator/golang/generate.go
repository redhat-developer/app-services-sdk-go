package golang

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
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
	if err != nil {
		fmt.Fprintln(os.Stderr, string(out))
		return err
	} else {
		fmt.Println(string(out))
	}

	doc, err := openapi3.NewLoader().LoadFromFile(inputFile)
	if err != nil {
		return err
	}

	if len(doc.Tags) == 0 {
		if err = generateMocks("default", outputPath); err != nil {
			return err
		}
	}
	for _, tag := range doc.Tags {
		if err = generateMocks(tag.Name, outputPath); err != nil {
			return err
		}
	}

	return nil
}

// generate mocks for each API interface present
func generateMocks(tag string, output string) error {
	err := os.Setenv("GO111MODULE", "on")
	if err != nil {
		return err
	}
	out, err := exec.Command("go", "get", "github.com/matryer/moq").CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, string(out))
		return err
	} else {
		fmt.Println(string(out))
	}
	mockFilePath := path.Join(output, fmt.Sprintf("%v_api_mock.go", tag))

	out, err = exec.Command("rm", "-f", mockFilePath).CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, string(out))
		return err
	} else {
		fmt.Println(string(out))
	}
	interfaceName := fmt.Sprintf("%vApi", strings.Title(strings.ToLower(tag)))

	out, err = exec.Command("moq", "-out", mockFilePath, output, interfaceName).CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, string(out))
		return err
	} else {
		fmt.Println(string(out))
	}
	out, err = exec.Command("go", "mod", "tidy").CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, string(out))
		return err
	} else {
		fmt.Println(string(out))
	}
	return nil
}
