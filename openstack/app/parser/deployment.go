package parser

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-yaml/yaml"
)

// Deployment is the struct type for Openshift App Deployment configuration
type Deployment struct {
	Kind       string `yaml:"kind"`
	APIVersion string `yaml:"apiVersion"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		Template struct {
			Metadata struct {
				Labels struct {
					Name string `yaml:"name"`
				} `yaml:"labels"`
			} `yaml:"metadata"`
			Spec struct {
				Containers []struct {
					Name  string `yaml:"name"`
					Image string `yaml:"image"`
					Ports []struct {
						ContainerPort int    `yaml:"containerPort"`
						Protocol      string `yaml:"protocol"`
					} `yaml:"ports"`
				} `yaml:"containers"`
			} `yaml:"spec"`
		} `yaml:"template"`
		Replicas int `yaml:"replicas"`
		Triggers []struct {
			Type              string `yaml:"type"`
			ImageChangeParams struct {
				Automatic      bool     `yaml:"automatic"`
				ContainerNames []string `yaml:"containerNames"`
				From           struct {
					Kind string `yaml:"kind"`
					Name string `yaml:"name"`
				} `yaml:"from"`
			} `yaml:"imageChangeParams,omitempty"`
		} `yaml:"triggers,omitempty"`
		Strategy struct {
			Type string `yaml:"type"`
		} `yaml:"strategy"`
		Paused               bool `yaml:"paused"`
		RevisionHistoryLimit int  `yaml:"revisionHistoryLimit"`
		MinReadySeconds      int  `yaml:"minReadySeconds"`
	} `yaml:"spec"`
}

// Deployment reads the yaml and returns the Deployment struct
func (p *Parser) Deployment(file string) error {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("error while reading yaml file %s : %v", file, err)
	}
	err = yaml.Unmarshal(yamlFile, &p.deployment)
	if err != nil {
		return fmt.Errorf("error during unmarshal of parser.Deployment : %v", err)
	}
	return nil
}

// GetImage returns the image used in 1st container
func (p *Parser) GetImage() string {
	var image string
	image = p.deployment.Spec.Template.Spec.Containers[0].Image
	return image
}

// GetImageForContainer returns the image used in a specific container name
func (p *Parser) GetImageForContainer(containerName string) string {
	var image string
	for _, container := range p.deployment.Spec.Template.Spec.Containers {
		if containerName == container.Name {
			image = container.Image
			break
		}
	}
	return image
}

// SetImage sets new image for 1st container
func (p *Parser) SetImage(newImage string) {
	p.deployment.Spec.Template.Spec.Containers[0].Image = newImage
}

// SetImageForContainer sets new image for specific container name
func (p *Parser) SetImageForContainer(newImage string, containerName string) {
	for index, container := range p.deployment.Spec.Template.Spec.Containers {
		if containerName == container.Name {
			p.deployment.Spec.Template.Spec.Containers[index].Image = newImage
			break
		}
	}
}

// WriteDeployment writes the parsed YAML into a file
func (p *Parser) WriteDeployment(outFile string) error {
	yamlData, err := yaml.Marshal(p.deployment)
	if err != nil {
		return fmt.Errorf("error while encoding YAML parser.Deployment data : %v", err)
	}
	yamlFile, err := os.Create(outFile)
	if err != nil {
		return fmt.Errorf("error while creating file - %s : %v", outFile, err)
	}
	defer yamlFile.Close()
	yamlFile.Write(yamlData)
	yamlFile.Close()
	return nil
}
