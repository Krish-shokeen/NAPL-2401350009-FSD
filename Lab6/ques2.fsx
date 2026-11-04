
open System

// Recursive function to check prime number
let rec isPrime n divisor =
    if n <= 1 then
        false
    elif divisor * divisor > n then
        true
    elif n % divisor = 0 then
        false
    else
        isPrime n (divisor + 1)

// Taking input
printf "Enter a number: "
let n = int(Console.ReadLine())

// Checking prime
if isPrime n 2 then
    printfn "%d is a Prime Number" n
else
    printfn "%d is Not a Prime Number" n
