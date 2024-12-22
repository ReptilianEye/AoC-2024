import java.util.*
import kotlin.math.abs

typealias Coordinates = Pair<Int, Int>

fun main() {
    val steps = listOf(0 to 1, 1 to 0, 0 to -1, -1 to 0)
    fun parseInput(input: List<String>) =
        input.map { it.map { it.toString() } }


    fun solve(map: List<List<String>>, shortcut: Int): Int {
        var start = 0 to 0
        for (i in map.indices) {
            for (j in map[i].indices) {
                if (map[i][j] == "S") {
                    start = i to j
                }
            }
        }
        fun Coordinates.manhattan(other: Coordinates) = abs(this.first - other.first) + abs(this.second - other.second)

        fun bfs(): List<Coordinates> {
            val visited = mutableSetOf<Coordinates>()
            val queue: Queue<Pair<Coordinates, MutableList<Coordinates>>> = LinkedList()
            queue.add((start to mutableListOf(start)))

            while (queue.isNotEmpty()) {
                val (pos, path) = queue.poll()
                if (map[pos.first][pos.second] == "E") {
                    return path
                }
                steps.forEach { (dx, dy) ->
                    val nextPos = pos.first + dx to pos.second + dy
                    if (nextPos !in visited && map[nextPos.first][nextPos.second] != "#") {
                        visited.add(nextPos)
                        queue.add(nextPos to (path + nextPos).toMutableList())
                    }
                }
            }
            return emptyList()
        }

        val shortestPath = bfs()
        val savedTimes = mutableListOf<Int>()
        for (i in shortestPath.indices) {
            for (j in i + 1..shortestPath.lastIndex) {
                val a = shortestPath[i]
                val b = shortestPath[j]
                val dist = a.manhattan(b)
                if (dist <= shortcut) {
                    val save = j - i - dist
                    if (save > 0) savedTimes.add(save)
                }
            }
        }
        return savedTimes.filter { it >= 100 }.size
    }

    fun part1(map: List<List<String>>) = solve(map, 2)

    fun part2(map: List<List<String>>) = solve(map, 20)


    check(part1(parseInput(readInput("Day20_test"))).also { it.println() } == 0)

    part1(parseInput(readInput("Day20"))).println()
    part2(parseInput(readInput("Day20"))).println()
}
