import java.io.File

fun main(args: Array<String>) {
    val input = File("aoc2023/src/input.txt").readLines()
    println("Solution A: ${Day4.solveA(input)}")
    println("Solution B: ${Day4.solveB(input)}")
}

object Day4 {
    fun solveA(input: List<String>): Long {
        val scores = input.map {
            val numbers = it.split(":")[1]
            val winning = numbers.split("|")[0].trim()
            val myNumbers = numbers.split("|")[1].trim()
        }
        return sum
    }

    fun solveB(input: List<String>): Long {
        return 0
    }
}
