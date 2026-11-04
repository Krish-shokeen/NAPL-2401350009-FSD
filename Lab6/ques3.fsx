
open System
open System.Text.RegularExpressions

// 1. Reverse a string and check palindrome
let checkPalindrome (str: string) =
    let reversed = String(Array.rev (str.ToCharArray()))
    printfn "Reversed String: %s" reversed

    if str.Equals(reversed, StringComparison.OrdinalIgnoreCase) then
        printfn "It is a Palindrome"
    else
        printfn "It is Not a Palindrome"


// 2. Transform string using pattern matching
let transformString (str: string) =
    match str with
    | s when s.Length > 0 && s |> Seq.forall Char.IsLetter ->
        s.ToUpper()
    | s when s.Length > 0 && s |> Seq.forall Char.IsDigit ->
        s + s
    | _ -> str


// 3. Validate email using Regex
let validateEmail (email: string) =
    Regex.IsMatch(email, @"^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$")


// Main program
printfn "----- STRING PROCESSING -----"

// Task 1
printf "\nEnter a string to check palindrome: "
let str1 = Console.ReadLine()
checkPalindrome str1

// Task 2
printf "\nEnter a string to transform: "
let str2 = Console.ReadLine()
printfn "Transformed String: %s" (transformString str2)

// Task 3
printf "\nEnter an email address: "
let email = Console.ReadLine()

if validateEmail email then
    printfn "Valid Email Address"
else
    printfn "Invalid Email Address"
