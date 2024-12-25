package day19

import println
import readInput

fun main() {
    fun parseInput(input: List<String>): Pair<List<String>, List<String>> {
        val towels = input[0].split(",").map { it.trim() }
        val queries = input.subList(2, input.size)
        return towels to queries
    }

    fun tryCreatingPattern(towels: List<String>, pattern: String, index: Int = 0): Boolean {
        if (index == pattern.length) return true
        var sol = false
        towels.forEach { towel ->
            if (towel.length > pattern.length - index) return@forEach
            if (pattern.substring(index, index + towel.length) == towel) {
                sol = sol || tryCreatingPattern(towels, pattern, index + towel.length)
            }
        }
        return sol
    }
    val cache = mutableMapOf<String,Long>()
    fun countAllPossible(towels: List<String>, pattern: String, index: Int = 0): Long {
        if (index == pattern.length) return 1

        if (pattern.substring(index) in cache) return cache[pattern.substring(index)]!!

        var sol = 0L
        towels.forEach { towel ->
            if (towel.length > pattern.length - index) return@forEach
            if (pattern.substring(index, index + towel.length) == towel) {
                sol += countAllPossible(towels, pattern, index + towel.length)
            }
        }
        cache[pattern.substring(index)] = sol
        return sol
    }


    fun part1(towels: List<String>, queries: List<String>): Int {
        return queries.fold(0) { acc, query ->
            acc + if (tryCreatingPattern(towels, query)) 1 else 0
        }
    }

    fun part2(towels: List<String>, queries: List<String>): Long {
        cache.clear()
        return queries.fold(0L) { acc, query ->
            acc + countAllPossible(towels, query)
        }
    }


    val (towelsE, queriesE) = parseInput(readInput("Day19_test"))
    check(part1(towelsE, queriesE).also { it.println() } == 6)
    check(part2(towelsE, queriesE).also { it.println() } == 16L)

    val (towels, queries) = parseInput(readInput("Day19"))
    part1(towels,queries).println()
    part2(towels,queries).println()
}
