import java.util.Scanner

// Task 2: Employee bonus using the when expression
fun calculateBonus(deptCode: String, years: Int, salary: Double): Double {
    val percent = when {
        deptCode == "IT" && years > 5 -> 20
        deptCode == "HR" && years > 3 -> 15
        deptCode == "FIN" && years > 4 -> 12
        else -> 10
    }
    return salary * percent / 100
}

// Task 3: Average of the digits of an integer
fun digitAverage(num: Int): Double {
    var n = Math.abs(num)
    if (n == 0) return 0.0
    var sum = 0
    var count = 0
    while (n > 0) {
        sum += n % 10
        count++
        n /= 10
    }
    return sum.toDouble() / count
}

fun main() {
    // Task 1
    println("Hello Kotlin!!!")

    // Task 2
    println("\n--- Employee Bonus ---")
    println("IT, 6 yrs, salary 50000  -> Bonus: ${calculateBonus("IT", 6, 50000.0)}")
    println("HR, 4 yrs, salary 40000  -> Bonus: ${calculateBonus("HR", 4, 40000.0)}")
    println("FIN, 5 yrs, salary 45000 -> Bonus: ${calculateBonus("FIN", 5, 45000.0)}")
    println("OPS, 2 yrs, salary 30000 -> Bonus: ${calculateBonus("OPS", 2, 30000.0)}")

    // Task 3
    println("\n--- Digit Average ---")
    print("Enter an integer: ")
    val sc = Scanner(System.`in`)
    val number = sc.nextInt()
    println("Average of digits = ${digitAverage(number)}")
}