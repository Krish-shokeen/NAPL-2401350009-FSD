
namespace FSharpStudentManagement

// Student status using Discriminated Union
type StudentStatus =
    | Active
    | Graduated
    | Suspended

// Student record
type Student = {
    Id: int option
    Name: string
    Age: int
    Course: string
    Status: StudentStatus
}

// Domain operations
module Domain =

    // Convert status to string
    let statusToText status =
        match status with
        | Active -> "Active"
        | Graduated -> "Graduated"
        | Suspended -> "Suspended"

    // Convert string to status
    let textToStatus value =
        match value with
        | "Active" -> Active
        | "Graduated" -> Graduated
        | "Suspended" -> Suspended
        | _ -> Active

    // Format student details
    let describe student =

        let idText =
            match student.Id with
            | Some id -> string id
            | None -> "New"

        sprintf "%-4s | %-20s | %-3d | %-18s | %s"
            idText
            student.Name
            student.Age
            student.Course
            (statusToText student.Status)
