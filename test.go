package main

import (
	"fmt"

	"github.com/lucas-assuncao/cplusplus-cplex-example/pkg/gopack_cplusplus_swig"
)

// This example illustrates how C++ classes can be used from Go using SWIG.

/*func SwigGo() {
	// ----- Object creation -----

	fmt.Println("Creating some objects:")
	c := gopack_cplusplus.NewCircle(10)
	fmt.Println("   Created circle", c)
	s := gopack_cplusplus.NewSquare(10)
	fmt.Println("   Created square", s)

	// ----- Access a static member -----

	fmt.Println("\nA total of", gopack_cplusplus.GetShapeNshapes(), "shapes were created")

	// ----- Member data access -----

	// Notice how we can do this using functions specific to
	// the 'Circle' class.
	c.SetX(20)
	c.SetY(30)

	// Now use the same functions in the base class
	var shape gopack_cplusplus.Shape = s
	shape.SetX(-10)
	shape.SetY(5)

	fmt.Println("\nHere is their current position:")
	fmt.Println("    Circle = (", c.GetX(), " ", c.GetY(), ")")
	fmt.Println("    Square = (", s.GetX(), " ", s.GetY(), ")")

	// ----- Call some methods -----

	fmt.Println("\nHere are some properties of the shapes:")
	shapes := []gopack_cplusplus.Shape{c, s}
	for i := 0; i < len(shapes); i++ {
		fmt.Println("   ", shapes[i])
		fmt.Println("        area      = ", shapes[i].Area())
		fmt.Println("        perimeter = ", shapes[i].Perimeter())
		fmt.Println("        type = ", shapes[i].GetName())
	}

	// Notice how the area() and perimeter() functions really
	// invoke the appropriate virtual method on each object.

	// ----- Delete everything -----

	fmt.Println("\nGuess I'll clean up now")

	// Note: this invokes the virtual destructor
	// You could leave this to the garbage collector
	gopack_cplusplus.DeleteCircle(c)
	gopack_cplusplus.DeleteSquare(s)

	fmt.Println(gopack_cplusplus.GetShapeNshapes(), " shapes remain")
	fmt.Println("Goodbye")
}*/

func main() {
	//SwigGo()
	input := "lucas"
	output := gopack_cplusplus_swig.CallLib(input)
	fmt.Println(output)
}
