import java.io.File

fun main(args: Array<String>) {
    val input = File("aoc2023/src/input.txt").readLines()
    println("Solution A: ${SolutionA().solveA(input)}")
    println("Solution B: ${SolutionB().solveB(input)}")
}

class SolutionA {
    fun solveA(input: List<String>): Int {
        val calValues = input.map {
            var first = -1
            var last = -1

            val charArray = it.toCharArray()

            charArray.forEachIndexed() { i, char ->
                if (char.isDigit()) {
                    if (first == -1) {
                        first = i
                    }
                    last = i
                }
            }

            return@map (charArray[first].toString() + charArray[last].toString()).toInt()
        }

        return calValues.sum()
    }
}

class SolutionB {
    fun solveB(input: List<String>): Int {
        val digitStrings = mapOf(
            "zero" to "0",
            "one" to "1",
            "two" to "2",
            "three" to "3",
            "four" to "4",
            "five" to "5",
            "six" to "6",
            "seven" to "7",
            "eight" to "8",
            "nine" to "9",
        )

        val calValues = input.map {
            var first = ""
            var last = ""

            val charArray = it.toCharArray()

            charArray.forEachIndexed() { i, char ->
                var value = ""
                if (char.isDigit()) {
                    value = charArray[i].toString()
                } else {
                    val remaining = it.substring(i, it.length)
                    digitStrings.keys.forEach { digitString ->
                        if (remaining.startsWith(digitString)) {
                            value = digitStrings[digitString].toString()
                        }
                    }
                }
                if (value.isNotEmpty()) {
                    if (first.isEmpty()) {
                        first = value
                    }
                    last = value
                }
            }

            return@map (first + last).toInt()
        }

        return calValues.sum()
    }
}