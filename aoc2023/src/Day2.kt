import java.io.File

fun main(args: Array<String>) {
    val input = File("aoc2023/src/input.txt").readLines()
    println("Solution A: ${Day2.solveA(input)}")
    println("Solution B: ${Day2.solveB(input)}")
}

object Day2 {
    fun solveA(input: List<String>): Int {
        val maxValues = mapOf(
            "red" to 12,
            "green" to 13,
            "blue" to 14,
        )

        val games = input.map { game ->
            val gameId = game.split(":")
            val rounds = gameId[1].split(";")

            rounds.forEach { round ->
                val cubes = round.split(",")
                cubes.forEach { cube ->
                    val values = cube.trim().split(" ")
                    if (values[0].toInt() > maxValues[values[1]]!!) {
                        return@map 0
                    }
                }
            }

            return@map gameId[0].split(" ")[1].toInt()
        }

        return games.sum()
    }

    fun solveB(input: List<String>): Int {
        val games = input.map { game ->
            val gameId = game.split(":")
            val minCubes = mutableMapOf(
                "red" to 0,
                "green" to 0,
                "blue" to 0,
            )

            val rounds = gameId[1].split(";")
            rounds.forEach { round ->
                val cubes = round.split(",")
                cubes.forEach { cube ->
                    val values = cube.trim().split(" ")
                    if (values[0].toInt() > minCubes[values[1]]!!) {
                        minCubes[values[1]] = values[0].toInt()
                    }
                }
            }

            return@map (minCubes["red"]!!.times(minCubes["green"]!!).times(minCubes["blue"]!!))
        }

        return games.sum()
    }
}
