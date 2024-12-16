fun main() {
    val movesToSteps = mapOf("^" to (-1 to 0), "v" to (1 to 0), "<" to (0 to -1), ">" to (0 to 1))

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
        fun dfs(startPos: Pair<Int, Int>, direction: Pair<Int, Int>): Boolean {
            val nextPos = startPos.first + direction.first to startPos.second + direction.second

            when (map[nextPos.first][nextPos.second]) {
                "#" -> return false
                "." -> {
                    map[nextPos.first][nextPos.second] = map[startPos.first][startPos.second]
                    map[startPos.first][startPos.second] = "."
                    return true
                }

                "O" -> {
                    if (dfs(nextPos, direction)) {
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
            if (dfs(robotPos, dx to dy)) {
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
        fun dfs(p: Pair<Int, Int>, direction: Pair<Int, Int>): Boolean {
            val curr = map[p.first][p.second]
            when (curr) {
                "#" -> return false
                "." -> return true
            }
            val (nx, ny) = p.first + direction.first to p.second + direction.second
            if (direction == (1 to 0) || direction == (-1 to 0)) {
                if (curr == "[" && dfs(nx to ny, direction) && dfs(nx to ny + 1, direction)) {
                    map[nx][ny] = map[p.first][p.second]
                    map[p.first][p.second] = "."
                    map[nx][ny + 1] = map[p.first][p.second + 1]
                    map[p.first][p.second + 1] = "."
                    return true
                }
                if (curr == "]" && dfs(nx to ny, direction) && dfs(nx to ny - 1, direction)) {
                    map[nx][ny] = map[p.first][p.second]
                    map[p.first][p.second] = "."
                    map[nx][ny - 1] = map[p.first][p.second - 1]
                    map[p.first][p.second - 1] = "."
                    return true
                }
            } else {
                if (dfs(nx to ny, direction)) {
                    map[nx][ny] = map[p.first][p.second]
                    map[p.first][p.second] = "."
                    return true
                }
            }

            return false

        }

        fun dfsCheck(p: Pair<Int, Int>, direction: Pair<Int, Int>): Boolean {
            val curr = map[p.first][p.second]
            when (curr) {
                "#" -> return false
                "." -> return true
            }
            val (nx, ny) = p.first + direction.first to p.second + direction.second
            if (direction == (1 to 0) || direction == (-1 to 0)) {
                if (curr == "[" && dfsCheck(
                        nx to ny,
                        direction
                    ) && dfsCheck(nx to ny + 1, direction)
                ) return true
                if (curr == "]" && dfsCheck(
                        nx to ny,
                        direction
                    ) && dfsCheck(nx to ny - 1, direction)
                ) return true
            } else {
                if (dfs(nx to ny, direction)) {
                    return true
                }
            }
            return false

        }

        var robotPos = robotInitPos
        for (move in moves) {
            val (dx, dy) = movesToSteps[move]!!
            val nextPos = robotPos.first + dx to robotPos.second + dy
            if (dfsCheck(nextPos, dx to dy)) {
                dfs(nextPos, dx to dy)
                map[nextPos.first][nextPos.second] = "@"
                map[robotPos.first][robotPos.second] = "."
                robotPos = nextPos
            }
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

    val testInput = readInput("Day15_test")
    val (mapE, movesE, robotPosE) = parseInputPart1(testInput)
    val (mapE2, movesE2, robotPosE2) = parseInputPart2(testInput)


    check(part1(mapE, movesE, robotPosE) == 10092)
    check(part2(mapE2, movesE2, robotPosE2) == 9021)


    val (map, moves, robotPos) = parseInputPart1(readInput("Day15"))
    val (map2, moves2, robotPos2) = parseInputPart2(readInput("Day15"))

    part1(map, moves, robotPos).println()
    part2(map2, moves2, robotPos2).println()
}
