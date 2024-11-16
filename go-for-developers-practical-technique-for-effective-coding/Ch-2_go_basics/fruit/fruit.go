package fruit

import "fmt"

const (
	// public constant - Exported outside the hosting package
	Apple  = "apple"
	Banana = "banana"

	// private constant - Not exported outside the hosting package
	cherry     = "cherry"
	strawberry = "strawberry"
)

// Print Function <Print()> - exported outside the hosing package
func Print() {
	fmt.Println(Apple, Banana, cherry, strawberry)

}
