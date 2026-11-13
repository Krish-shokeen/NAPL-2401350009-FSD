
namespace FSharpStudentManagement

open Microsoft.Data.Sqlite

module Database =

    // SQLite connection string
    let connectionString = "Data Source=students.db"

    // Create connection
    let createConnection () =
        new SqliteConnection(connectionString)

    // Initialize database and create table
    let initializeAsync () =
        async {
            use connection = createConnection ()

            do! connection.OpenAsync()
                |> Async.AwaitTask

            use command = connection.CreateCommand()

            command.CommandText <-
                """
                CREATE TABLE IF NOT EXISTS Students (
                    Id INTEGER PRIMARY KEY AUTOINCREMENT,
                    Name TEXT NOT NULL,
                    Age INTEGER NOT NULL,
                    Course TEXT NOT NULL,
                    Status TEXT NOT NULL
                );
                """

            let! _ =
                command.ExecuteNonQueryAsync()
                |> Async.AwaitTask

            return ()
        }

    // CREATE: Add student
    let addStudentAsync (student: Student) =
        async {
            use connection = createConnection ()

            do! connection.OpenAsync()
                |> Async.AwaitTask

            use command = connection.CreateCommand()

            command.CommandText <-
                """
                INSERT INTO Students
                (Name, Age, Course, Status)
                VALUES
                ($name, $age, $course, $status);
                """

            command.Parameters.AddWithValue(
                "$name", student.Name) |> ignore

            command.Parameters.AddWithValue(
                "$age", student.Age) |> ignore

            command.Parameters.AddWithValue(
                "$course", student.Course) |> ignore

            command.Parameters.AddWithValue(
                "$status",
                Domain.statusToText student.Status) |> ignore

            let! _ =
                command.ExecuteNonQueryAsync()
                |> Async.AwaitTask

            return ()
        }

    // READ: Get all students
    let getStudentsAsync () =
        async {
            use connection = createConnection ()

            do! connection.OpenAsync()
                |> Async.AwaitTask

            use command = connection.CreateCommand()

            command.CommandText <-
                "SELECT Id, Name, Age, Course, Status FROM Students ORDER BY Id;"

            use! reader =
                command.ExecuteReaderAsync()
                |> Async.AwaitTask

            let students = ResizeArray<Student>()

            let mutable reading = true

            while reading do

                let! hasRow =
                    reader.ReadAsync()
                    |> Async.AwaitTask

                if hasRow then

                    let student = {
                        Id = Some (reader.GetInt32(0))
                        Name = reader.GetString(1)
                        Age = reader.GetInt32(2)
                        Course = reader.GetString(3)
                        Status =
                            Domain.textToStatus (
                                reader.GetString(4)
                            )
                    }

                    students.Add(student)

                else
                    reading <- false

            return students |> Seq.toList
        }

    // UPDATE: Update student
    let updateStudentAsync id (student: Student) =
        async {
            use connection = createConnection ()

            do! connection.OpenAsync()
                |> Async.AwaitTask

            use command = connection.CreateCommand()

            command.CommandText <-
                """
                UPDATE Students
                SET
                    Name = $name,
                    Age = $age,
                    Course = $course,
                    Status = $status
                WHERE Id = $id;
                """

            command.Parameters.AddWithValue(
                "$name", student.Name) |> ignore

            command.Parameters.AddWithValue(
                "$age", student.Age) |> ignore

            command.Parameters.AddWithValue(
                "$course", student.Course) |> ignore

            command.Parameters.AddWithValue(
                "$status",
                Domain.statusToText student.Status) |> ignore

            command.Parameters.AddWithValue(
                "$id", id) |> ignore

            let! affected =
                command.ExecuteNonQueryAsync()
                |> Async.AwaitTask

            return affected > 0
        }

    // DELETE: Delete student
    let deleteStudentAsync id =
        async {
            use connection = createConnection ()

            do! connection.OpenAsync()
                |> Async.AwaitTask

            use command = connection.CreateCommand()

            command.CommandText <-
                "DELETE FROM Students WHERE Id = $id;"

            command.Parameters.AddWithValue(
                "$id", id) |> ignore

            let! affected =
                command.ExecuteNonQueryAsync()
                |> Async.AwaitTask

            return affected > 0
        }
