import java.io.File

fun main(args: Array<String>) {
    val input = File("aoc2023/src/input.txt").readLines()
    println("Solution A: ${Day3.solveA(input)}")
    println("Solution B: ${Day3.solveB(input)}")
}

object Day3 {
    fun solveA(input: List<String>): Int {
        var part = mutableListOf<Int>()
        var partVal = ""
        val parts = mutableListOf<Int>()

        input.forEachIndexed { i, line ->
            line.toCharArray().forEachIndexed { j, c ->
                if (c.isDigit()) {
                    part += j
                    partVal += c
                    if (partVal == "274") {
                        println("here")
                    }
                }
                if (!c.isDigit() || j == line.toCharArray().size.dec()) {
                    if (part.isNotEmpty()) {
                        if (isPart(input, i, part)) {
                            parts.add(partVal.toInt())
                        }
                        part = mutableListOf<Int>()
                        partVal = ""
                    }
                }
            }
        }

        return parts.sum()
    }

    private fun isPart(lines: List<String>, i: Int, partInds: List<Int>): Boolean {
        return (isPartUp(lines, i, partInds)
            || isPartDown(lines, i, partInds)
            || isPartLeft(lines, i, partInds.first())
            || isPartRight(lines, i, partInds.last()))
    }

    private fun isSpecialChar(c: Char): Boolean {
        return !c.isLetterOrDigit() && c != '.'
    }

    private fun isPartUp(lines: List<String>, i: Int, partInds: List<Int>): Boolean {
        if (i > 0) {
            if (partInds.first() > 0) {
                if (isSpecialChar(lines[i.dec()].toCharArray()[partInds.first().dec()])) {
                    return true
                }
            }
            if (partInds.last() < lines[i].length.dec()) {
                if (isSpecialChar(lines[i.dec()].toCharArray()[partInds.last().inc()])) {
                    return true
                }
            }
            partInds.forEach { partInd ->
                if (isSpecialChar(lines[i.dec()].toCharArray()[partInd])) {
                    return true
                }
            }
        }

        return false
    }

    private fun isPartDown(lines: List<String>, i: Int, partInds: List<Int>): Boolean {
        if (i < lines.size.dec()) {
            if (partInds.first() > 0) {
                if (isSpecialChar(lines[i.inc()].toCharArray()[partInds.first().dec()])) {
                    return true
                }
            }
            if (partInds.last() < lines[i].length.dec()) {
                if (isSpecialChar(lines[i.inc()].toCharArray()[partInds.last().inc()])) {
                    return true
                }
            }
            partInds.forEach { partInd ->
                if (isSpecialChar(lines[i.inc()].toCharArray()[partInd])) {
                    return true
                }
            }
        }

        return false
    }

    private fun isPartLeft(lines: List<String>, i: Int, j: Int): Boolean {
        if (j > 0) {
            if (isSpecialChar(lines[i].toCharArray()[j.dec()])) {
                return true
            }
        }

        return false
    }

    private fun isPartRight(lines: List<String>, i: Int, j: Int): Boolean {
        if (j < lines[i].length.dec()) {
            if (isSpecialChar(lines[i].toCharArray()[j.inc()])) {
                return true
            }
        }

        return false
    }

    fun solveB(input: List<String>): Int {
        val gearWithPart = mutableMapOf<String, Int>()
        val parts = mutableListOf<Int>()

        input.forEachIndexed { i, line ->
            var part = mutableListOf<Int>()
            var partVal = ""
            line.toCharArray().forEachIndexed { j, c ->
                if (c.isDigit()) {
                    part += j
                    partVal += c
                }
                if (!c.isDigit() || j == line.toCharArray().size.dec()) {
                    if (part.isNotEmpty()) {
                        val gearPos = gearPos(input, i, part)
                        if (gearPos.isNotEmpty()) {
                            if (gearWithPart.containsKey(gearPos)) {
                                parts.add(gearWithPart[gearPos]!! * partVal.toInt())
                            } else {
                                gearWithPart[gearPos] = partVal.toInt()
                            }
                        }
                        part = mutableListOf<Int>()
                        partVal = ""
                    }
                }
            }
        }

        return parts.sum()
    }

    private fun isGear(c: Char): Boolean {
        return c == '*'
    }

    private fun gearPos(lines: List<String>, i: Int, partInds: List<Int>): String {
        val posUp = gearPosUp(lines, i, partInds)
        if (posUp.isNotEmpty()) {
            return posUp
        }
        val posDown = gearPosDown(lines, i, partInds)
        if (posDown.isNotEmpty()) {
            return posDown
        }
        val posLeft = gearPosLeft(lines, i, partInds.first())
        if (posLeft.isNotEmpty()) {
            return posLeft
        }
        val posRight = gearPosRight(lines, i, partInds.last())
        if (posRight.isNotEmpty()) {
            return posRight
        }

        return ""
    }

    private fun gearPosUp(lines: List<String>, i: Int, partInds: List<Int>): String {
        if (i > 0) {
            if (partInds.first() > 0) {
                if (isGear(lines[i.dec()].toCharArray()[partInds.first().dec()])) {
                    return toDigitString(i.dec(), partInds.first().dec())
                }
            }
            if (partInds.last() < lines[i].length.dec()) {
                if (isGear(lines[i.dec()].toCharArray()[partInds.last().inc()])) {
                    return toDigitString(i.dec(), partInds.last().inc())
                }
            }
            partInds.forEach { partInd ->
                if (isGear(lines[i.dec()].toCharArray()[partInd])) {
                    return toDigitString(i.dec(), partInd)
                }
            }
        }

        return ""
    }

    private fun gearPosDown(lines: List<String>, i: Int, partInds: List<Int>): String {
        if (i < lines.size.dec()) {
            if (partInds.first() > 0) {
                if (isGear(lines[i.inc()].toCharArray()[partInds.first().dec()])) {
                    return toDigitString(i.inc(), partInds.first().dec())
                }
            }
            if (partInds.last() < lines[i].length.dec()) {
                if (isGear(lines[i.inc()].toCharArray()[partInds.last().inc()])) {
                    return toDigitString(i.inc(), partInds.last().inc())
                }
            }
            partInds.forEach { partInd ->
                if (isGear(lines[i.inc()].toCharArray()[partInd])) {
                    return toDigitString(i.inc(), partInd)
                }
            }
        }

        return ""
    }

    private fun gearPosLeft(lines: List<String>, i: Int, j: Int): String {
        if (j > 0) {
            if (isGear(lines[i].toCharArray()[j.dec()])) {
                return toDigitString(i, j.dec())
            }
        }

        return ""
    }

    private fun gearPosRight(lines: List<String>, i: Int, j: Int): String {
        if (j < lines[i].length.dec()) {
            if (isGear(lines[i].toCharArray()[j.inc()])) {
                return toDigitString(i, j.inc())
            }
        }

        return ""
    }

    private fun toDigitString(a: Int, b: Int): String {
        return "$a,$b"
    }
}
