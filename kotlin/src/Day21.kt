import kotlin.math.min

fun main() {
    val steps = mapOf(0 to 1 to '>', 1 to 0 to 'v', 0 to -1 to '<', -1 to 0 to '^')
    val dirKeypad =
        listOf(
            listOf('#', '#', '#', '#', '#'),
            listOf('#', '#', '^', 'A', '#'),
            listOf('#', '<', 'v', '>', '#'),
            listOf('#', '#', '#', '#', '#')
        )
    val numKeypad =
        listOf(
            listOf('#', '#', '#', '#', '#'),
            listOf('#', '7', '8', '9', '#'),
            listOf('#', '4', '5', '6', '#'),
            listOf('#', '1', '2', '3', '#'),
            listOf('#', '#', '0', 'A', '#'),
            listOf('#', '#', '#', '#', '#')
        )


    fun fillPaths(keyboard: List<List<Char>>): Map<Pair<Char, Char>, List<String>> {
        fun dfs(start: Pair<Int, Int>, map: List<List<Char>>): Map<Pair<Char, Char>, List<String>> {
            val startEl = map[start.first][start.second]
            val pathsFromStart = mutableMapOf<Pair<Char, Char>, List<String>>().withDefault { listOf() }

            fun dfsRec(pos: Pair<Int, Int>, path: String) {
                val curr = map[pos.first][pos.second]
                val pathsForCurr = pathsFromStart.getValue(startEl to curr).toMutableList()
                var shortestInPaths = pathsForCurr.minOfOrNull { it.length } ?: Int.MAX_VALUE
                if (pathsForCurr.isNotEmpty() && shortestInPaths < path.length) {
                    return
                }
                shortestInPaths = min(shortestInPaths, path.length)
                pathsForCurr.add(path)
                pathsFromStart[startEl to curr] = pathsForCurr.filter { it.length <= shortestInPaths }
                steps.forEach { (step, marker) ->
                    val (dx, dy) = step
                    val nextPos = pos.first + dx to pos.second + dy
                    val nextEl = map[nextPos.first][nextPos.second]
                    if (nextEl != '#') {
                        dfsRec(nextPos, path + marker)
                    }
                }
            }
            dfsRec(start, "")
            return pathsFromStart
        }

        fun isSorted(s: String): Boolean {
            val sList = s.toList()
            val sSorted = sList.sorted()
            return sList == sSorted || sList == sSorted.reversed()
        }

        val paths = mutableMapOf<Pair<Char, Char>, List<String>>()
        for (i in keyboard.indices) {
            for (j in keyboard[0].indices) {
                if (keyboard[i][j] != '#') {
                    val pathsFromI = dfs(i to j, keyboard)
                    paths += pathsFromI
                        .mapValues {
                            it.value.filter { p -> isSorted(p) }
                        }
                }
            }
        }
        return paths.mapValues { (_, paths) -> paths.map { it + "A" } }
    }

    var numPaths: Map<Pair<Char, Char>, List<String>> = emptyMap()
    var dirPaths: Map<Pair<Char, Char>, List<String>> = emptyMap()

    fun simulateSingleQuery(query: String, paths: Map<Pair<Char, Char>, List<String>>): List<String> {
        var currentCodes = listOf("")
        var prev = 'A'
        for (el in query) {
            val thisStepResult = mutableListOf<String>()
            for (prevExpanded in currentCodes) {
                for (path in paths[prev to el]!!) {
                    thisStepResult.add(prevExpanded + path)
                }
            }
            prev = el
            currentCodes = thisStepResult
        }
        return currentCodes
    }

    val cache: MutableMap<Pair<String, Int>, Long> = mutableMapOf()
    fun bestLength(seq: String, depth: Int): Long {
        if (cache.containsKey(seq to depth)) {
            return cache[seq to depth]!!
        }
        if (depth == 1) {
            return ("A$seq" zip seq).sumOf { (a, b) -> dirPaths[a to b]!!.first().length }.toLong()
        }
        return ("A$seq" zip seq).fold(0L) { acc, (x, y) ->
            acc + dirPaths[x to y]!!.minOf { subseq -> bestLength(subseq, depth - 1) }
        }.also { cache[seq to depth] = it }
    }

    fun solve(queries: List<String>, robotsCount: Int): Long {
        var sol = 0L
        for (query in queries) {
            val length = simulateSingleQuery(query, numPaths).minOf { bestLength(it, robotsCount) }
            sol += length * query.filter { it.isDigit() }.toLong()
        }

        return sol

    }

    fun part1(queries: List<String>) = solve(queries, 2)

    fun part2(queries: List<String>) = solve(queries, 25)


    numPaths = fillPaths(numKeypad)
    dirPaths = fillPaths(dirKeypad)
    check(part1(readInput("Day21_test")).also { it.println() } == 126384L)

    part1(readInput("Day21")).println()
    part2(readInput("Day21")).println()
}
