// (1) Function with two default parameters
fun greet(name: String = "Guest", greeting: String = "Hello") {
    println("$greeting, $name!")
}

// (2) Function used with named arguments
fun createProfile(name: String = "Unknown", age: Int = 18, city: String = "NA") {
    println("Name: $name, Age: $age, City: $city")
}

// (3) Vararg function
fun printNumbers(vararg nums: Int) {
    print("Numbers: ")
    for (n in nums) print("$n ")
    println()
}

// (4) Higher-order function taking a lambda (Int) -> Int
fun applyOperation(list: List<Int>, operation: (Int) -> Int): List<Int> {
    val result = mutableListOf<Int>()
    for (item in list) {
        result.add(operation(item))
    }
    return result
}

fun main() {
    println("--- Default Parameters ---")
    greet()                      // both defaults used
    greet("Rahul")               // one default used
    greet("Priya", "Welcome")    // no defaults used

    println("\n--- Named Arguments ---")
    createProfile(city = "Delhi", name = "Amit", age = 21)  // out of order
    createProfile(age = 25, name = "Neha")                  // city uses default

    println("\n--- Vararg ---")
    printNumbers(1, 2, 3)
    printNumbers(10, 20, 30, 40, 50)

    println("\n--- Higher-Order Function ---")
    val numbers = listOf(1, 2, 3, 4, 5)
    val squared = applyOperation(numbers) { it * it }
    val doubled = applyOperation(numbers) { it * 2 }
    println("Original: $numbers")
    println("Squared : $squared")
    println("Doubled : $doubled")
}