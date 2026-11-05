
open System

// Two different lists
let list1 = [1; 2; 3; 4; 5]
let list2 = [10; 20; 30; 40]

// 1. List.map - Square each element
let square1 = List.map (fun x -> x * x) list1
let square2 = List.map (fun x -> x * x) list2

printfn "MAP - Squared Lists:"
printfn "%A" square1
printfn "%A" square2

// 2. List.filter - Select even numbers
let even1 = List.filter (fun x -> x % 2 = 0) list1
let even2 = List.filter (fun x -> x % 2 = 0) list2

printfn "\nFILTER - Even Numbers:"
printfn "%A" even1
printfn "%A" even2

// 3. List.fold - Calculate sum
let sum1 = List.fold (fun acc x -> acc + x) 0 list1
let sum2 = List.fold (fun acc x -> acc + x) 0 list2

printfn "\nFOLD - Sum of Lists:"
printfn "%d" sum1
printfn "%d" sum2
