package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Constants for probabilities
const (
	P_L = 0.3 // Probability of library compromise
	P_I = 0.2 // Probability of image tampering
	P_E = 0.4 // Probability of runtime vulnerability activation
)

// Deployment struct represents a Kubernetes deployment
type Deployment struct {
	Name       string
	Namespace  string
	Containers []string
}

// ImageInfo struct represents metadata for container images
type ImageInfo struct {
	ImageName    string
	TamperedProb float64 // Probability of tampering
}

// LibraryInfo struct represents external libraries used by operators
type LibraryInfo struct {
	Name        string
	Compromised bool
}

// Vulnerability struct represents a vulnerability scenario
type Vulnerability struct {
	DeploymentName string `yaml:"deployment_name"`
	Namespace      string `yaml:"namespace"`
	ContainerImage string `yaml:"container_image"`
	Reason         string `yaml:"reason"`
}

// Operator struct represents the Kubernetes operator
type Operator struct {
	Libraries     []LibraryInfo
	Deployments   []Deployment
	ImageRegistry map[string]ImageInfo
}

// CheckLibraryIntegrity verifies the integrity of the operator's libraries
func (op *Operator) CheckLibraryIntegrity() bool {
	for _, lib := range op.Libraries {
		if lib.Compromised {
			fmt.Printf("Library '%s' is compromised!\n", lib.Name)
			return false
		}
	}
	return true
}

// SimulateImageTampering simulates whether an image is tampered with
func (op *Operator) SimulateImageTampering(imageName string) bool {
	image, exists := op.ImageRegistry[imageName]
	if !exists {
		fmt.Printf("Image '%s' not found in the registry!\n", imageName)
		return false
	}
	return rand.Float64() < image.TamperedProb
}

// SimulateRuntimeCondition checks if runtime conditions trigger vulnerabilities
func SimulateRuntimeCondition() bool {
	return rand.Float64() < P_E
}

// AnalyzeDeploymentSecurity analyzes each deployment for vulnerabilities
func (op *Operator) AnalyzeDeploymentSecurity() []Vulnerability {
	var vulnerabilities []Vulnerability

	for _, deployment := range op.Deployments {
		for _, container := range deployment.Containers {
			if op.SimulateImageTampering(container) && SimulateRuntimeCondition() {
				vulnerability := Vulnerability{
					DeploymentName: deployment.Name,
					Namespace:      deployment.Namespace,
					ContainerImage: container,
					Reason:         "Image tampered and runtime condition triggered",
				}
				vulnerabilities = append(vulnerabilities, vulnerability)
			}
		}
	}
	return vulnerabilities
}

// WriteYAMLReport writes the vulnerability report to a YAML file
func WriteYAMLReport(fileName string, vulnerabilities []Vulnerability) error {
	data, err := yaml.Marshal(&vulnerabilities)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML: %v", err)
	}
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write YAML file: %v", err)
	}
	fmt.Printf("YAML report successfully written to %s\n", fileName)
	return nil
}

func main() {
	// Seed for randomness
	rand.Seed(time.Now().UnixNano())

	// Simulated data
	imageRegistry := map[string]ImageInfo{
		"image1": {ImageName: "image1", TamperedProb: P_I},
		"image2": {ImageName: "image2", TamperedProb: P_I},
		"image3": {ImageName: "image3", TamperedProb: P_I},
		"image4": {ImageName: "image4", TamperedProb: P_I},
	}

	deployments := []Deployment{
		{"dep1", "namespace1", []string{"image1", "image2"}},
		{"dep2", "namespace2", []string{"image3"}},
		{"dep3", "namespace3", []string{"image4", "image1"}},
	}

	libraries := []LibraryInfo{
		{Name: "library1", Compromised: false},
		{Name: "library2", Compromised: false},
	}

	// Initialize operator
	operator := Operator{
		Libraries:     libraries,
		Deployments:   deployments,
		ImageRegistry: imageRegistry,
	}

	// Step 1: Check library integrity
	if !operator.CheckLibraryIntegrity() {
		fmt.Println("Operator libraries are compromised. Exiting...")
		return
	}

	// Step 2: Analyze deployments for vulnerabilities
	vulnerabilities := operator.AnalyzeDeploymentSecurity()

	// Step 3: Write YAML report
	fileName := "vulnerability_report.yaml"
	if err := WriteYAMLReport(fileName, vulnerabilities); err != nil {
		fmt.Printf("Error writing YAML report: %v\n", err)
		return
	}

	// Optional: Print vulnerabilities to console
	fmt.Println("\n--- Security Analysis Report ---")
	for _, v := range vulnerabilities {
		fmt.Printf("Deployment: %s | Namespace: %s | Image: %s | Reason: %s\n",
			v.DeploymentName, v.Namespace, v.ContainerImage, v.Reason)
	}
	fmt.Println("--------------------------------")
}
