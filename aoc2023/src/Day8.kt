import java.io.File

fun main(args: Array<String>) {
    val input = File("aoc2023/src/input.txt").readLines()
    //println("Solution A: ${Day8.solveA(input)}")
    println("Solution B: ${Day8.solveB(input)}")
}

object Day8 {
    fun solveA(input: List<String>): Long {
        val instructions = input.first().toCharArray()
        var currInstr = 0
        var noOfSteps: Long = 0
        val nodesList = input.subList(2, input.size)

        val nodes = nodesList.associate {
            it.split("=")[0].trim() to it.split("=")[1].trim(' ', '(', ')')
        }

        var currNode = "AAA"

        while (currNode != "ZZZ") {
            val left = nodes[currNode]!!.split(',')[0]
            val right = nodes[currNode]!!.split(',')[1].trim()

            when (instructions[currInstr]) {
                'L' -> currNode = left
                'R' -> currNode = right
            }

            currInstr = (currInstr+1) % instructions.size
            noOfSteps++
        }

        return noOfSteps
    }

    fun solveB(input: List<String>): Long {
        val instructions = input.first().toCharArray()
        val nodesList = input.subList(2, input.size)

        val nodes = nodesList.associate {
            it.split("=")[0].trim() to it.split("=")[1].trim(' ', '(', ')')
        }

        val starts = nodes.keys.filter { it.endsWith("A") }

        val steps = starts.map {

            var currInstr = 0
            var noOfSteps: Long = 0
            var currNode = it

            while (!currNode.endsWith("Z")) {
                val left = nodes[currNode]!!.split(',')[0]
                val right = nodes[currNode]!!.split(',')[1].trim()

                when (instructions[currInstr]) {
                    'L' -> currNode = left
                    'R' -> currNode = right
                }

                currInstr = (currInstr+1) % instructions.size
                noOfSteps++
            }

            return@map noOfSteps
        }

        return steps.reduce { acc, i -> lcm(acc, i) }
    }

    private fun gcd(a: Long, b: Long): Long {
        var dividend = a
        var divisor = b

        while (divisor != 0L) {
            val r = dividend % divisor
            dividend = divisor
            divisor = r
        }

        return dividend
    }

    private fun lcm(a: Long, b: Long): Long {
        return (a * b) / gcd(a, b)
    }
}
