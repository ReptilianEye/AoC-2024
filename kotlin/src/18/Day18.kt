package day18

import println
import readInput

fun main() {
    val steps = listOf(0 to 1, 1 to 0, 0 to -1, -1 to 0)
    fun parseInput(input: List<String>): List<Pair<Int, Int>> =
        input.map { it.split(",").map { it.toInt() } }.map { it[0] to it[1] }


    fun part1(grid: MutableList<MutableList<String>>, input: List<Pair<Int, Int>>, test: Boolean = false): Int {

        input.subList(0, if (test) 12 else 1024).forEach {
            val (y, x) = it
            grid[x][y] = "#"
        }

        val start = 0 to 0
        val end = grid.lastIndex to grid.first().lastIndex

        fun bfs(start: Pair<Int, Int>): Int {
            val queue = ArrayDeque(listOf(start to 0))
            val visited = mutableSetOf<Pair<Int, Int>>()

            while (queue.isNotEmpty()) {
                val (pos, marker) = queue.removeFirst()
                if (pos == end) {
                    return marker
                }
                steps.forEach { (dx, dy) ->
                    val nextPos = pos.first + dx to pos.second + dy
                    if (nextPos.first !in grid.indices || nextPos.second !in grid.first().indices)
                        return@forEach

                    if (grid[nextPos.first][nextPos.second] != "#" && nextPos !in visited) {
                        visited.add(nextPos)
                        queue.add(nextPos to marker + 1)
                    }
                }
            }
            return Int.MAX_VALUE
        }
        return bfs(start)
    }

    fun part2(grid: MutableList<MutableList<String>>, input: List<Pair<Int, Int>>, test: Boolean = false): Pair<Int,Int> {
        val start = 0 to 0
        val end = grid.lastIndex to grid.first().lastIndex


        fun bfs(start: Pair<Int, Int>): Pair<Int, List<Pair<Int, Int>>> {
            val queue = ArrayDeque(listOf(Triple(start, 0, mutableListOf(start))))
            val visited = mutableSetOf<Pair<Int, Int>>()

            while (queue.isNotEmpty()) {
                val (pos, marker, path) = queue.removeFirst()
                if (pos == end) {
                    return marker to path
                }
                steps.forEach { (dx, dy) ->
                    val nextPos = pos.first + dx to pos.second + dy
                    if (nextPos.first !in grid.indices || nextPos.second !in grid.first().indices)
                        return@forEach

                    if (grid[nextPos.first][nextPos.second] != "#" && nextPos !in visited) {
                        visited.add(nextPos)
                        path.add(nextPos)
                        queue.add(Triple(nextPos, marker + 1, path.toMutableList()))
                    }
                }
            }
            return Int.MAX_VALUE to emptyList()
        }

        var inputIdx = if (test) 12 else 1024
        input.toList().subList(0, inputIdx).forEach { (y, x) ->
            grid[x][y] = "#"
        }

        while (inputIdx <= input.lastIndex) {
            val (_, path) = bfs(start)
            if (path.isEmpty()) return input[inputIdx - 1]
            while (inputIdx <= input.lastIndex) {
                val (y, x) = input[inputIdx]
                grid[x][y] = "#"
                inputIdx++
                if (x to y in path) break
            }

        }

        return 0 to 0
    }


    check(
        part1(
            MutableList(7) { MutableList(7) { "." } },
            parseInput(readInput("Day18_test")), true
        ).also { it.println() } == 22)
    check(part2(MutableList(7) { MutableList(7) { "." } }, parseInput(readInput("Day18_test")), true).also { it.println() } == 6 to 1)


    part1(MutableList(71) { MutableList(71) { "." } }, parseInput(readInput("Day18")).subList(0, 1024)).println()
    part2(MutableList(71) { MutableList(71) { "." } }, parseInput(readInput("Day18"))).println()
}
