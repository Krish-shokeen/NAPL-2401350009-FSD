import java.util.Scanner

class Student(
    private var rollNo: Int,
    private var name: String,
    private var course: String,
    private var percentage: Double
) {
    // Accept / update student data from user input
    fun acceptData(sc: Scanner) {
        print("Enter Roll No: ")
        rollNo = sc.nextInt()
        sc.nextLine()
        print("Enter Name: ")
        name = sc.nextLine()
        print("Enter Course: ")
        course = sc.nextLine()
        print("Enter Percentage: ")
        percentage = sc.nextDouble()
    }

    // Display all details in a formatted manner
    fun displayData() {
        println("--------------------------------")
        println("Roll No    : $rollNo")
        println("Name       : $name")
        println("Course     : $course")
        println("Percentage : %.2f%%".format(percentage))
        println("--------------------------------")
    }
}

fun main() {
    val sc = Scanner(System.`in`)

    val s1 = Student(1, "Default1", "NA", 0.0)
    val s2 = Student(2, "Default2", "NA", 0.0)

    println("Enter details for Student 1:")
    s1.acceptData(sc)
    println("\nEnter details for Student 2:")
    s2.acceptData(sc)

    println("\nStudent 1 Details:")
    s1.displayData()
    println("\nStudent 2 Details:")
    s2.displayData()
}