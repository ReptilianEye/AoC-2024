val movesToSteps = mapOf("^" to (-1 to 0), "v" to (1 to 0), "<" to (0 to -1), ">" to (0 to 1))
fun main() {
    fun parseInputPart1(input: List<String>): Triple<MutableList<MutableList<String>>, List<String>, Pair<Int, Int>> {
        val map = mutableListOf<MutableList<String>>()
        var i = 0
        var robotPos = Pair(0, 0)
        while (i < input.size) {
            if (input[i].isEmpty()) break
            val row = input[i].split("").toMutableList()
            row.removeFirst()
            row.removeLast()
            map.add(row)
            if (row.contains("@")) {
                robotPos = Pair(i, row.indexOf("@"))
            }
            i++
        }
        i++
        val moves = mutableListOf<String>()
        while (i < input.size) {
            moves += input[i].split("").toList()
            i++;
        }

        return Triple(map, moves.filter { it != "" }, robotPos)
    }

    fun parseInputPart2(input: List<String>): Triple<MutableList<MutableList<String>>, List<String>, Pair<Int, Int>> {
        val map = mutableListOf<MutableList<String>>()
        var i = 0
        var robotPos = Pair(0, 0)
        while (i < input.size) {
            if (input[i].isEmpty()) break
            val extendRow = mutableListOf<String>()
            for (el in input[i].split("").toMutableList().filter { it != "" }) {
                when (el) {
                    "." -> {
                        extendRow.add(".")
                        extendRow.add(".")
                    }

                    "#" -> {
                        extendRow.add("#")
                        extendRow.add("#")
                    }

                    "O" -> {
                        extendRow.add("[")
                        extendRow.add("]")
                    }

                    "@" -> {
                        extendRow.add("@")
                        extendRow.add(".")
                    }
                }
            }
            map.add(extendRow)
            if (extendRow.contains("@")) {
                robotPos = Pair(i, extendRow.indexOf("@"))
            }
            i++
        }
        i++
        val moves = mutableListOf<String>()
        while (i < input.size) {
            moves += input[i].split("").toList().filter { it != "" }
            i++;
        }

        return Triple(map, moves, robotPos)
    }


    fun part1(map: MutableList<MutableList<String>>, moves: List<String>, robotInitPos: Pair<Int, Int>): Int {
        fun tryMoving(startPos: Pair<Int, Int>, direction: Pair<Int, Int>): Boolean {
            val nextPos = startPos.first + direction.first to startPos.second + direction.second

            when (map[nextPos.first][nextPos.second]) {
                "#" -> return false
                "." -> {
                    map[nextPos.first][nextPos.second] = map[startPos.first][startPos.second]
                    map[startPos.first][startPos.second] = "."
                    return true
                }

                "O" -> {
                    if (tryMoving(nextPos, direction)) {
                        map[nextPos.first][nextPos.second] = map[startPos.first][startPos.second]
                        map[startPos.first][startPos.second] = "."
                        return true
                    }
                    return false
                }
            }
            return false
        }

        var robotPos = robotInitPos
        for (move in moves) {
            val (dx, dy) = movesToSteps[move]!!
            val nextPos = robotPos.first + dx to robotPos.second + dy
            if (tryMoving(robotPos, dx to dy)) {
                robotPos = nextPos
            }
        }
        var sol = 0
        for (i in map.indices) {
            for (j in map[i].indices) {
                if (map[i][j] == "O") {
                    sol += 100 * i + j
                }
            }
        }
        return sol
    }

    fun part2(map: MutableList<MutableList<String>>, moves: List<String>, robotInitPos: Pair<Int, Int>): Int {
        fun tryMoving(startPos: Pair<Int, Int>, direction: Pair<Int, Int>): Boolean {
            val nextPos = startPos.first + direction.first to startPos.second + direction.second
            if (map[nextPos.first][nextPos.second] == "#") return false

            if (direction == (0 to 1) || direction == (0 to -1)) {
                when (map[nextPos.first][nextPos.second]) {
                    "." -> {
                        map[nextPos.first][nextPos.second] = map[startPos.first][startPos.second]
                        map[startPos.first][startPos.second] = "."
                        return true
                    }

                    "[", "]" -> {
                        if (tryMoving(nextPos, direction)) {
                            map[nextPos.first][nextPos.second] = map[startPos.first][startPos.second]
                            map[startPos.first][startPos.second] = "."
                            return true
                        }
                        return false
                    }
                }
            } else {
                val positionsToCheck = mutableListOf(startPos)
                if (map[startPos.first][startPos.second] == "[") {
                    positionsToCheck += startPos.first to startPos.second + 1
                } else if (map[startPos.first][startPos.second] == "]") {
                    positionsToCheck += startPos.first to startPos.second - 1
                }
                var canMoveAll = true
                for (pos in positionsToCheck) {
                    val posNext = pos.first + direction.first to pos.second + direction.second
                    when (map[posNext.first][posNext.second]) {
                        "#" -> {
                            canMoveAll = false
                            break
                        }

                        "." -> {
                            continue
                        }

                        "[", "]" -> {
                            canMoveAll = canMoveAll && tryMoving(posNext, direction)
                        }
                    }
                }
                if (canMoveAll) {
                    for (pos in positionsToCheck) {
                        val posNext = pos.first + direction.first to pos.second + direction.second
                        map[posNext.first][posNext.second] = map[pos.first][pos.second]
                        map[pos.first][pos.second] = "."
                    }
                    return true
                }
            }
            return false
        }

        var robotPos = robotInitPos
        for (move in moves) {
            val (dx, dy) = movesToSteps[move]!!
            val nextPos = robotPos.first + dx to robotPos.second + dy
            if (tryMoving(robotPos, dx to dy)) {
                robotPos = nextPos
            }
//            for (row in map) println(row.joinToString(""))
//            println()

        }

        var sol = 0
        for (i in map.indices) {
            for (j in map[i].indices) {
                if (map[i][j] == "[") {
                    sol += 100 * i + j
                }
            }
        }
        return sol
    }

// Or read a large test input from the `src/Day01_test.txt` file:
    val testInput = readInput("Day15_test")
    val (mapE, movesE, robotPosE) = parseInputPart1(testInput)
    val (mapE2, movesE2, robotPosE2) = parseInputPart2(testInput)

//    for (row in mapE2) println(row.joinToString(""))

    check(part1(mapE, movesE, robotPosE) == 10092)
    check(part2(mapE2, movesE2, robotPosE2) == 9021)

// Read the input from the `src/Day01.txt` file.
    val (map, moves, robotPos) = parseInputPart1(readInput("Day15"))
    val (map2, moves2, robotPos2) = parseInputPart2(readInput("Day15"))

    part1(map, moves, robotPos).println()
    part2(map2, moves2, robotPos2).println()
    for (row in map2) println(row.joinToString(""))
}
