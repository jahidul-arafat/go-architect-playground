package main

import "fmt"

// Task:
/*
create a ServiceMeshOperator struct with both reference and non-reference fields,
along with three methods that demonstrate potential issues when handling pointers and values.
Here's the implementation

Simulation Output:
=================================================
Simulating some Struct value/reference bugs .....

Initial Operator State:
Operator: FirstServiceMashOperator (v1.0)
Namespaces: [Default]
ConfigMap: map[]
IsUpdated: false
Deployment: 0

Updating Config (with Value receiver):
Operator: FirstServiceMashOperator (v1.0)
Namespaces: [Default]
ConfigMap: map[policy:strict]
IsUpdated: false
Deployment: 0

Appending the Namespaces and other fields (with Pointer Receiver: )
Operator: FirstServiceMashOperator (v1.0)
Namespaces: [Default Second]
ConfigMap: map[policy:strict]
IsUpdated: true
Deployment: 1

*/

/*
Potential Issues (if value receiver, instead of pointer receiver)
===================================================================

1. Inconsistent Behavior: UpdateConfig partially works (for reference types) but fails to update non-reference fields, leading to confusion.
2. Unintended Side Effects: Users of UpdateConfig might expect all fields to be updated, but only the ConfigMap is affected.
3. Performance Issues: PrintOperatorInfo and UpdateConfig create a copy of the entire struct each time they're called, which can be inefficient for large structs.
4. Concurrency Problems: If used in concurrent operations, UpdateConfig might lead to race conditions as it's not actually modifying the original struct.
5. Difficulty in Testing: The effects of UpdateConfig on non-reference fields can't be observed or tested on the original struct.
*/

type ServiceMashOperator struct {
	// define the non-reference field
	// These non-reference fields are the major source of Bugs
	Name            string            // non-reference field
	Version         string            // non-reference field
	Namespace       []string          // reference field (slice)	// slice of string -> default referencing to the array of strings
	ConfigMap       map[string]string // reference field (map)
	IsModified      bool              // non-reference field
	DeploymentCount int               // non-reference field

}

// Methods of struct "ServiceMashOperator"

// Method-01: printing the operator information
// this is not a silo function, its a value receiver
// this will receive the struct object as a value and will print the object's information
func (smo ServiceMashOperator) PrintOperatorInfo() {
	fmt.Printf("Operator: %s (v%s)\n"+
		"Namespaces: %v\n"+
		"ConfigMap: %v\n"+
		"IsUpdated: %v\n"+
		"Deployment: %d\n",
		smo.Name, smo.Version, smo.Namespace,
		smo.ConfigMap, smo.IsModified, smo.DeploymentCount)
}

// Method-2: Update the default configs
// use a value receiver and check how this introduces bugs-> leading to unexpected behavior
func (smo ServiceMashOperator) UpdateConfig(key, value string) {
	smo.ConfigMap[key] = value
	smo.IsModified = true // will not have impact since struct object is received as a value
	smo.DeploymentCount++ // will not have impact
}

// Method-3: Initialize Operator
// a pointer receiver instead of value received to minimize the performace bottlececka and nullify the modification bugs
func (smo *ServiceMashOperator) InitializeOperator(namespace string) {
	smo.Namespace = append(smo.Namespace, namespace)
	smo.IsModified = true
	smo.DeploymentCount++
}

func main() {
	fmt.Println("Simulating some Struct value/reference bugs .....")

	// lets create a service mash operator and observe its behavior
	operator := ServiceMashOperator{
		Name:            "FirstServiceMashOperator",
		Version:         "1.0",
		Namespace:       []string{"Default"},
		ConfigMap:       make(map[string]string), // make() will initialize the map, else there will be a "panic" due to being 'nil'
		IsModified:      false,                   // default is false even if not initialized
		DeploymentCount: 0,                       // default is 0 even if not initialized
	}

	// print the operator
	fmt.Println("\nInitial Operator State: ")
	operator.PrintOperatorInfo()

	// Update the configmap with some (key,value) and look for inconsistencies or bugs introduced
	fmt.Println("\nUpdating Config (with Value receiver):")
	operator.UpdateConfig("policy", "strict")
	operator.PrintOperatorInfo() // isModified should be True and DeploymentCount should be 1
	// but due to value receiver none of the non-reference field has had the modifications reflected
	// Lets Append the Namespaces to make the operator available to both the "Default" and "Second" namespaces
	// pointer receiver: changes will reflect in both the non-reference and reference fields
	fmt.Println("\nAppending the Namespaces and other fields (with Pointer Receiver: )")
	operator.InitializeOperator("Second")
	operator.PrintOperatorInfo()
}
