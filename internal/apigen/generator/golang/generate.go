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
	inputFile := g.Config.Input
	if _, err = os.Stat(inputFile); err != nil {
		return err
	}

	clientMetadata := g.Config.ClientMetadata
	packageName := clientMetadata.PackageName()
	outputPath := clientMetadata.OutputPath()

	fmt.Fprintln(os.Stderr, "deleting old API client before generation")
	if err = os.RemoveAll(outputPath); err != nil {
		fmt.Fprintln(os.Stderr, "failed to remove old client code")
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

	var tags []string = []string{}
	for _, path := range doc.Paths {
		tags = appendOperationTags(tags, path.Get)
		tags = appendOperationTags(tags, path.Post)
		tags = appendOperationTags(tags, path.Put)
		tags = appendOperationTags(tags, path.Patch)
		tags = appendOperationTags(tags, path.Delete)
	}

	if len(tags) == 0 {
		_ = generateMocks("default", outputPath)
		return nil
	}

	tags = unique(tags)
	for _, tag := range tags {
		_ = generateMocks(tag, outputPath)
	}

	return nil
}

// generate mocks for each API interface present
func generateMocks(tag string, output string) error {
	err := os.Setenv("GO111MODULE", "on")
	if err != nil {
		return err
	}
	out, err := exec.Command("go", "get", "-u", "github.com/matryer/moq").CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, string(out))
		return err
	} else {
		fmt.Println(string(out))
	}

	filePrefix := strings.ToLower(strings.Join(strings.Split(tag, " "), "_"))

	mockFilePath := path.Join(output, fmt.Sprintf("%v_api_mock.go", filePrefix))
	fmt.Println(mockFilePath)

	out, err = exec.Command("rm", "-f", mockFilePath).CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, string(out))
		return err
	} else {
		fmt.Println(string(out))
	}

	tagChunks := strings.Split(tag, " ")
	for index := range tagChunks {
		tagChunks[index] = strings.Title(tagChunks[index])
	}
	interfaceName := fmt.Sprintf("%vApi", strings.Join(tagChunks, ""))

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

// returns a slice with duplicate values removed
func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func appendOperationTags(tags []string, operation *openapi3.Operation) []string {
	if operation == nil {
		return tags
	}
	return append(tags, operation.Tags...)
}
