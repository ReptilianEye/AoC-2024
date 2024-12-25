package day24

import println
import readInput

private data class Wire(
    var value: Int? = null,
    val left: String? = null,
    val right: String? = null,
    val op: String? = null
)

private var wires: Map<String, Wire> = emptyMap()
fun main() {
    fun parseOp(op: String) = when (op) {
        "XOR" -> Int::xor
        "AND" -> Int::and
        "OR" -> Int::or
        else -> null
    }

    fun parseInput(input: List<String>): Map<String, Wire> {
        val nodes = mutableMapOf<String, Wire>()
        var i = 0
        while (true) {
            val line = input[i]
            if (line.isEmpty()) break
            val (label, value) = line.split(": ")
            nodes[label] = Wire(value = value.toInt())
            i++
        }
        for (line in input.subList(i + 1, input.size)) {
            val query = line.split(" ")
            val (left, op, right, label) = listOf(query[0], query[1], query[2], query[4])

            nodes[label] = Wire(left = left, right = right, op = op)
        }
        return nodes
    }


    val usedNodes = mutableSetOf<String>()
    fun eval(label: String): Int {
        val (value, left, right, op) = wires[label]!!
        if (value != null) return value
        usedNodes.add(label)
        return parseOp(op!!)!!(eval(left!!), eval(right!!))
    }

    fun part1(n: Map<String, Wire>): Long {
        wires = n
        val zVals: List<Pair<String, Int>> =
            wires.keys.filter { l -> l.startsWith("z") }.fold<String, List<Pair<String, Int>>>(emptyList()) { acc, l ->
                acc + (l to eval(l))
            }.sortedByDescending { it.first }
        val bin = zVals.joinToString("") { it.second.toString() }
        return bin.toLong(2)
    }


    fun part2(n: Map<String, Wire>): String {
        wires = n
        var i = 0
        while (true) {
            if (!verifyZ("z".padLabel(i), i))
                break
            i++
        }
        println("Broke on $i")
        // qdg <-> z12
        // vvf <-> z19
        // dck <-> fgn
        // nvh <-> z37

        return listOf("qdg", "z12", "vvf", "z19", "dck", "fgn", "nvh", "z37").sorted().joinToString(",")
    }


    check(part1(parseInput(readInput("Day24_test"))).also { it.println() } == 4L)
    check(part1(parseInput(readInput("Day24_test1"))).also { it.println() } == 2024L)
    part1(parseInput(readInput("Day24"))).println()

    part2(parseInput(readInput("Day24"))).println()
}

private fun String.padLabel(n: Int) = this + n.toString().padStart(2, '0')

private fun makeDesc(l: String?, r: String?) = listOf(l ?: "", r ?: "").sorted()

private fun verifyZ(label: String, n: Int): Boolean {
//    println("verify z $label $n")
    val (_, left, right, op) = wires[label]!!
    if (op != "XOR") return false
    val desc = makeDesc(left, right)
    if (n == 0) return desc == listOf("x00", "y00")
    return desc.zip(desc.reversed()).any { (l, r) -> verifyIntermediateXor(l, n) && verifyCarryBit(r, n) }
}

private fun verifyIntermediateXor(label: String, n: Int): Boolean {
//    println("verify intermediate xor $label $n")
    val (_, left, right, op) = wires[label]!!
    if (op != "XOR") return false
    return makeDesc(left, right) == listOf("x".padLabel(n), "y".padLabel(n))
}

private fun verifyCarryBit(label: String, n: Int): Boolean {
//    println("carry bit $label $n")
    val (_, left, right, op) = wires[label]!!
    val desc = makeDesc(left, right)
    if (n == 1) return op == "AND" && desc == listOf("x00", "y00")
    return desc.zip(desc.reversed()).any { (l, r) -> verifyRecarry(l, n - 1) && verifyDirectCarry(r, n - 1) }
}

private fun verifyDirectCarry(label: String, n: Int): Boolean {
//    println("verifyDirectCarry $label $n")
    val (_, left, right, op) = wires[label]!!
    if (op != "AND") return false
    return makeDesc(left, right) == listOf("x".padLabel(n), "y".padLabel(n))
}

private fun verifyRecarry(label: String, n: Int): Boolean {
//    println("verify recarry $label $n")
    val (_, left, right, op) = wires[label]!!
    if (op != "AND") return false
    val desc = makeDesc(left, right)
    return desc.zip(desc.reversed()).any { (l, r) -> verifyIntermediateXor(l, n) && verifyCarryBit(r, n) }
}
