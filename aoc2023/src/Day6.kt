import java.io.File

fun main(args: Array<String>) {
    val input = File("aoc2023/src/input.txt").readLines()
    val raceTimes = input[0].split(":")[1].trim().split("\\s+".toRegex())
    val raceDists = input[1].split(":")[1].trim().split("\\s+".toRegex())

    println("Solution A: ${Day6.solveA(raceTimes, raceDists)}")
    println("Solution B: ${Day6.solveB(input)}")
}

object Day6 {
    fun solveA(times: List<String>, dists: List<String>): Int {
        val possibleWins = mutableListOf<Int>()

        times.forEachIndexed { i, time ->
            val t = time.toInt()
            val d = dists[i].toInt()
            val buttonPressTimes = mutableListOf<Int>()

            for (buttonPress in 1..t) {
                val travelTime = t - buttonPress
                val attemptDist = buttonPress * travelTime
                if (attemptDist > d) {
                    buttonPressTimes.add(buttonPress)
                }
            }
            possibleWins.add(buttonPressTimes.size)
        }

        return possibleWins.reduce { acc, i -> acc * i }
    }

    fun solveB(input: List<String>): Int {
        val t = input[0].split(":")[1].replace("\\s".toRegex(), "").toBigDecimal()
        val d = input[1].split(":")[1].replace("\\s".toRegex(), "").toBigDecimal()

        val buttonPressTimes = mutableListOf<Int>()

        for (buttonPress in 1..t.toInt()) {
            val bp = buttonPress.toBigDecimal()
            val travelTime = t - bp
            val attemptDist = bp * travelTime
            if (attemptDist > d) {
                buttonPressTimes.add(buttonPress)
            }
        }

        return buttonPressTimes.size
    }
}
