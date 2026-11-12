
open System

// Base Class
[<AbstractClass>]
type Shape() =
    abstract member Area : unit -> float

// Derived Class 1
type Circle(radius: float) =
    inherit Shape()

    override this.Area() =
        Math.PI * radius * radius

// Derived Class 2
type Rectangle(width: float, height: float) =
    inherit Shape()

    override this.Area() =
        width * height

// Creating Objects
let circle = Circle(5.0)
let rectangle = Rectangle(4.0, 6.0)

// Calling Methods
printfn "Circle Area: %.2f" (circle.Area())
printfn "Rectangle Area: %.2f" (rectangle.Area())

// Demonstrating Polymorphism
let shapes: Shape list =
    [ circle :> Shape; rectangle :> Shape ]

printfn "\nPolymorphism:"

for shape in shapes do
    printfn "Area: %.2f" (shape.Area())
