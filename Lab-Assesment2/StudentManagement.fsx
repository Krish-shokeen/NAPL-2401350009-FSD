
#r "nuget: Microsoft.Data.Sqlite, 8.0.11"

#load "Domain.fs"
#load "Database.fs"

open System
open FSharpStudentManagement


// Read non-empty string
let readNonEmpty prompt =

    let rec read () =

        printf "%s" prompt

        let value = Console.ReadLine()

        if String.IsNullOrWhiteSpace(value) then
            printfn "This field cannot be empty."
            read ()
        else
            value.Trim()

    read ()

// Read positive integer
let readPositiveInt prompt =

    let rec read () =

        printf "%s" prompt

        match Int32.TryParse(Console.ReadLine()) with

        | true, value when value > 0 ->
            value

        | _ ->
            printfn "Enter a valid positive number."
            read ()

    read ()

// Read student status
let readStatus () =

    printfn "\nChoose Student Status:"
    printfn "1. Active"
    printfn "2. Graduated"
    printfn "3. Suspended"

    match readNonEmpty "Enter choice: " with

    | "1" -> Active
    | "2" -> Graduated
    | "3" -> Suspended

    | _ ->
        printfn "Invalid choice. Defaulting to Active."
        Active

// CREATE STUDENT

let createStudent () =

    {
        Id = None

        Name =
            readNonEmpty "Enter student name: "

        Age =
            readPositiveInt "Enter student age: "

        Course =
            readNonEmpty "Enter course name: "

        Status =
            readStatus ()
    }

// DISPLAY STUDENTS

let printStudents students =

    printfn "\n=============================================================="

    printfn "ID   | NAME                 | AGE | COURSE             | STATUS"

    printfn "=============================================================="

    if List.isEmpty students then

        printfn "No student records found."

    else

        students
        |> List.iter (Domain.describe >> printfn "%s")

    printfn "=============================================================="

// FUNCTIONAL PROGRAMMING
// MAP, FILTER AND PATTERN MATCHING

let functionalReport students =

    // FILTER: Get only active students
    let activeStudents =

        students
        |> List.filter (fun student ->
            student.Status = Active
        )

    // MAP: Convert student names into uppercase
    let studentNames =

        students
        |> List.map (fun student ->
            student.Name.ToUpperInvariant()
        )

    // PATTERN MATCHING: Count student records
    let summary =

        match List.length students with

        | 0 ->
            "No student records available."

        | 1 ->
            "One student record available."

        | count ->
            sprintf "%d student records available." count

    printfn "\n========== FUNCTIONAL REPORT =========="

    printfn "%s" summary

    printfn "\nTotal active students: %d"
        (List.length activeStudents)

    printfn "\nStudent names using MAP:"

    studentNames
    |> List.iter (printfn "- %s")

    printfn "\nActive students using FILTER:"

    printStudents activeStudents

let rec menu () =

    printfn "\n=========================================="

    printfn "       STUDENT MANAGEMENT SYSTEM"

    printfn "=========================================="

    printfn "1. Add Student"

    printfn "2. View All Students"

    printfn "3. Update Student"

    printfn "4. Delete Student"

    printfn "5. Functional Programming Report"

    printfn "0. Exit"

    printfn "=========================================="

    let choice = readNonEmpty "\nEnter your choice: "

    match choice with

    // ADD STUDENT
    | "1" ->

        let student = createStudent ()

        Database.addStudentAsync student
        |> Async.RunSynchronously

        printfn "\nStudent added successfully!"

        menu ()

    // VIEW STUDENTS
    | "2" ->

        let students =

            Database.getStudentsAsync ()
            |> Async.RunSynchronously

        printStudents students

        menu ()

    // UPDATE STUDENT
    | "3" ->

        let id =
            readPositiveInt "Enter student ID to update: "

        printfn "\nEnter updated student details:"

        let student = createStudent ()

        let result =

            Database.updateStudentAsync id student
            |> Async.RunSynchronously

        if result then

            printfn "\nStudent updated successfully!"

        else

            printfn "\nStudent ID not found."

        menu ()

    // DELETE STUDENT
    | "4" ->

        let id =
            readPositiveInt "Enter student ID to delete: "

        let result =

            Database.deleteStudentAsync id
            |> Async.RunSynchronously

        if result then

            printfn "\nStudent deleted successfully!"

        else

            printfn "\nStudent ID not found."

        menu ()

    // FUNCTIONAL REPORT
    | "5" ->

        let students =

            Database.getStudentsAsync ()
            |> Async.RunSynchronously

        functionalReport students

        menu ()

    // EXIT
    | "0" ->

        printfn "\nThank you for using Student Management System!"

    // INVALID OPTION
    | _ ->

        printfn "\nInvalid choice. Please try again."

        menu ()

//starting 

try

    Database.initializeAsync ()
    |> Async.RunSynchronously

    printfn "\nDatabase initialized successfully!"

    printfn "Student Management System Started."

    menu ()

with

| ex ->

    printfn "\nApplication Error: %s" ex.Message
