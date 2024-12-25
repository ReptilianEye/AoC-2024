package day25

import println
import readInput

fun main() {
    fun parseInput(input: List<String>): Pair<List<List<Int>>, List<List<Int>>> {
        val keys = mutableListOf<List<Int>>()
        val locks = mutableListOf<List<Int>>()

        var i = 0
        while (i < input.size) {
            val isLock = input[i].all { it == '#' }
            val heights = (0..<5).fold(emptyList<Int>()) { acc, j ->
                val column = (1..5).fold("") { col, k ->
                    col + input[i + k][j]
                }
                acc + column.count { it == '#' }
            }
            if (isLock) {
                locks.add(heights)
            } else {
                keys.add(heights)
            }
            i += 8
        }

        return locks to keys
    }

    val maxHeight = 5
    fun fits(key: List<Int>, lock: List<Int>) = key.zip(lock).all { (k, l) -> (k + l) <= maxHeight }


    fun part1(locks: List<List<Int>>, keys: List<List<Int>>): Int {
        var sol = 0
        for (lock in locks) {
            for (key in keys) {
                if (fits(key, lock)) sol++
            }
        }
        return sol

    }


    val (locksE, keysE) = parseInput(readInput("Day25_test"))
    check(part1(locksE, keysE).also { it.println() } == 3)
    val (locks, keys) = parseInput(readInput("Day25"))
    part1(locks, keys).println()
}
