package cmd

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"gopkg.in/yaml.v2"
)

const (
	PackageConfigFileName = "package_config.yaml"
	ToolsDir              = "tools"
	AnsiblePackagesDir    = "ansilbe_packages"
)

type PackageConfig struct {
	Tools           []string               `yaml:"tools"`
	AnsiblePackages []AnsiblePackageConfig `yaml:"ansiblePackages"`
}

type AnsiblePackageConfig struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

func downloadArtifacts(dir string) {
	fmt.Println(dir)
	// Read the YAML file
	packageConfigPath := filepath.Join(dir, PackageConfigFileName)
	yamlFile, err := os.ReadFile(packageConfigPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Parse the YAML file
	var packageCfg PackageConfig
	err = yaml.Unmarshal(yamlFile, &packageCfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(packageCfg)
	downloadAnsiblePackages(dir, packageCfg)
	//downloadTools(dir, packageCfg)
}

func downloadAnsiblePackages(dir string, packageCfg PackageConfig) {
	ansiblePackagesDirPath := filepath.Join(dir, AnsiblePackagesDir)
	if _, err := os.Stat(ansiblePackagesDirPath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(ansiblePackagesDirPath, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	for index, element := range packageCfg.AnsiblePackages {
		fmt.Println("At index", index, "value is", element.Url)
		downloadFromUrl(element.Url, ansiblePackagesDirPath)
	}
}

func downloadTools(dir string, packageCfg PackageConfig) {
	toolsDirPath := filepath.Join(dir, ToolsDir)
	if _, err := os.Stat(toolsDirPath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(toolsDirPath, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	for index, element := range packageCfg.Tools {
		fmt.Println("At index", index, "value is", element)
		downloadFromUrl(element, toolsDirPath)
	}
}

func downloadFromUrl(url string, path string) {
	s := spinner.New(spinner.CharSets[31], 100*time.Millisecond)
	s.Start()
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	fmt.Println("Downloading", url, "to", fileName)

	downloadPath := filepath.Join(path, fileName)
	fmt.Println(downloadPath)
	// TODO: check file existence first with io.IsExist
	output, err := os.Create(downloadPath)
	if err != nil {
		fmt.Println("Error while creating", downloadPath, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}

	fmt.Println(n, "bytes downloaded.")
	s.Stop()
}
