package main

//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//// Constants for probabilities
//const (
//	P_L = 0.3 // Probability of library compromise
//	P_I = 0.2 // Probability of image tampering
//	P_E = 0.4 // Probability of runtime conditions being met
//)
//
//// Deployment struct
//type Deployment struct {
//	Name       string
//	Namespace  string
//	Containers []string
//}
//
//// Function to simulate outcomes
//func simulateVulnerabilities(deployments []Deployment, images map[string]float64) (float64, int) {
//	rand.Seed(time.Now().UnixNano())
//	totalPods := 0
//	vulnerablePods := 0
//
//	for _, deployment := range deployments {
//		for _, container := range deployment.Containers {
//			totalPods++
//			imageRisk := images[container]
//			if rand.Float64() < imageRisk && rand.Float64() < P_E {
//				vulnerablePods++
//			}
//		}
//	}
//
//	probability := float64(vulnerablePods) / float64(totalPods)
//	return probability, vulnerablePods
//}
//
//func main() {
//	// Simulate deployments and images
//	deployments := []Deployment{
//		{"dep1", "namespace1", []string{"image1", "image2"}},
//		{"dep2", "namespace1", []string{"image3", "image4"}},
//		{"dep3", "namespace2", []string{"image5", "image6", "image7"}},
//		{"dep4", "namespace3", []string{"image8"}},
//		{"dep5", "namespace4", []string{"image9", "image10"}},
//	}
//
//	// Image risks based on tampering probability
//	images := map[string]float64{
//		"image1": 0.2, "image2": 0.2, "image3": 0.2, "image4": 0.2,
//		"image5": 0.2, "image6": 0.2, "image7": 0.2, "image8": 0.2,
//		"image9": 0.2, "image10": 0.2,
//	}
//
//	// Calculate probabilities
//	probability, vulnerablePods := simulateVulnerabilities(deployments, images)
//
//	// Output results
//	fmt.Printf("Probability of cluster compromise: %.2f\n", probability)
//	fmt.Printf("Expected number of vulnerable pods: %d\n", vulnerablePods)
//}
