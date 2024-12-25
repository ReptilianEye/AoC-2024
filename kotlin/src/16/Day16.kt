package day16

import println
import readInput
import kotlin.math.min

fun main() {
    operator fun Pair<Int, Int>.plus(pair: Pair<Int, Int>) = this.first + pair.first to this.second + pair.second

    val moves = listOf(0 to 1, 1 to 0, 0 to -1, -1 to 0)

    fun part1(map: List<String>): Int {
        val start = map.size - 2 to 1
        val end = 1 to map[0].length - 2

        val cache = mutableMapOf<Triple<Int, Int, Int>, Int>()
        var bestScore = Int.MAX_VALUE
        fun dfs(p: Pair<Int, Int>, currentMove: Int, score: Int, visited: MutableSet<Pair<Int, Int>>) {
            val (x, y) = p
            if (p == end) {
                bestScore = min(bestScore, score)
                return
            }
            if (score > bestScore) return
            if (Triple(x, y, currentMove) in cache && cache[Triple(x, y, currentMove)]!! < score)
                return
            cache[Triple(x, y, currentMove)] = score

            visited.add(p)

            listOf(0, 1, 3).forEach { move ->
                val nextMove = (currentMove + move) % 4
                val (nx, ny) = p + moves[nextMove]
                if (map[nx][ny] != '#' && (nx to ny) !in visited) {
                    visited.add(nx to ny)
                    dfs(
                        nx to ny,
                        nextMove,
                        score + 1 + if (nextMove != currentMove) 1000 else 0,
                        visited
                    )
                    visited.remove(nx to ny)
                }
            }
        }
        dfs(start, 0, 0, mutableSetOf())
        return bestScore
    }

    fun part2(map: List<String>): Int {
        val start = map.size - 2 to 1
        val end = 1 to map[0].length - 2

        val cache = mutableMapOf<Triple<Int, Int, Int>, Int>()
        val bestScore = part1(map)
        val bestPaths = mutableListOf<List<Pair<Int, Int>>>()
        fun dfs(
            p: Pair<Int, Int>,
            currentMove: Int,
            score: Int,
            visited: MutableSet<Pair<Int, Int>>,
            path: MutableList<Pair<Int, Int>>
        ) {
            val (x, y) = p
            if (p == end) {
                if (bestScore == score) {
                    bestPaths.add(path.toList())
                }
                return
            }
            if (score > bestScore) return
            if (Triple(x, y, currentMove) in cache && cache[Triple(x, y, currentMove)]!! < score)
                return
            cache[Triple(x, y, currentMove)] = score

            listOf(0, 1, 3).forEach { move ->
                val nextMove = (currentMove + move) % 4
                val (nx, ny) = p + moves[nextMove]
                if (map[nx][ny] != '#' && (nx to ny) !in visited) {
                    visited.add(nx to ny)
                    path.add(nx to ny)
                    dfs(
                        nx to ny,
                        nextMove,
                        score + 1 + if (nextMove != currentMove) 1000 else 0,
                        visited,
                        path
                    )
                    path.removeLast()
                    visited.remove(nx to ny)
                }
            }
        }
        dfs(start, 0, 0, mutableSetOf(), mutableListOf(start))
        val nodesOnBestPair = bestPaths.flatten().toSet()

        return nodesOnBestPair.size
    }


    check(part1(readInput("Day16_test")).also { it.println() } == 7036)
    check(part1(readInput("Day16_test1")).also { it.println() } == 11048)
    check(part2(readInput("Day16_test")).also { it.println() } == 45)
    check(part2(readInput("Day16_test1")).also { it.println() } == 64)

    part1(readInput("Day16")).println()
    part2(readInput("Day16")).println()
}

