
open System

printfn "Enter marks of four subjects out of 100:"

let m1 = int(Console.ReadLine())
let m2 = int(Console.ReadLine())
let m3 = int(Console.ReadLine())
let m4 = int(Console.ReadLine())

let total = m1 + m2 + m3 + m4
let percentage = float total / 4.0

let division =
    match percentage with
    | p when p >= 60.0 -> "First Division"
    | p when p >= 50.0 -> "Second Division"
    | p when p >= 40.0 -> "Third Division"
    | _ -> "Fail"

printfn "Total Marks = %d" total
printfn "Percentage = %.2f%%" percentage
printfn "Division = %s" division
