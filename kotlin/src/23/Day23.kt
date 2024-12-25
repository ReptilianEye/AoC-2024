package day23

import println
import readInput

fun main() {
    fun parseInput(input: List<String>) = input.map { it.split("-") }
        .fold(mutableMapOf<String, MutableList<String>>().withDefault { mutableListOf() }) { acc, nodes ->
            (listOf(0, 1) zip listOf(1, 0)).forEach { (l, r) ->
                acc[nodes[l]] = (acc.getValue(nodes[l]) + nodes[r]).toMutableList()
            }
            acc
        }

    var graph: Map<String, List<String>> = emptyMap()

    fun findClique(n: Int): Set<Set<String>> {
        val solutions = mutableSetOf<Set<String>>()
        fun clique(nodes: Set<String>) {
            if (nodes.size == n) {
                solutions.add(nodes)
                return
            }
            for ((node, nbours) in graph) {
                if (node !in nodes && nodes.all { nbours.contains(it) }) {
                    clique(nodes + node)
                }
            }
        }
        clique(emptySet())
        return solutions
    }

    fun findMaxClique(): Set<String> {
        var solution = setOf<String>()
        fun bronKerbosh(r: Set<String>, p: Set<String>, x: Set<String>) {
            var R = r.toSet()
            val P = p.toMutableSet()
            val X = x.toMutableSet()

            if (P.isEmpty() && X.isEmpty()) {
                if (solution.size < R.size) {
                    solution = R
                }
                return
            }
            val u = (P + X).maxBy { graph.getValue(it).count { q -> q in P } }
            (P - graph.getValue(u).toSet()).forEach { v ->
                val nbours = graph.getValue(v).toSet()
                bronKerbosh(R + v, P intersect nbours, X intersect nbours)
                P.remove(v)
                X.add(v)
            }
        }
        bronKerbosh(emptySet(), graph.keys, emptySet())
        return solution
    }


    fun part1(g: Map<String, List<String>>): Int {
        graph = g
        val clique = findClique(3)
        return clique.toList().count { it -> it.any { it.startsWith('t') } }
    }

    fun part2(g: Map<String, List<String>>): String {
        graph = g
        val maxClique = findMaxClique()
        return maxClique.sorted().joinToString(",")
    }

    check(part1(parseInput(readInput("Day23_test"))).also { it.println() } == 7)
    part1(parseInput(readInput("Day23"))).println()

    check(part2(parseInput(readInput("Day23_test"))).also { it.println() } == "co,de,ka,ta")
    part2(parseInput(readInput("Day23"))).println()
}
