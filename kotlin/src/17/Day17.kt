package day17

import println
import readInput

fun main() {
    data class Input(val registers: List<Long>, val instructions: List<Int>)

    fun parseInput(input: List<String>): Input {
        val registers = List(3) {
            input[it].split(":")[1].trim().toLong()
        }
        val instructions = input.last().split(": ")[1].split(",").map { it.toInt() }
        return Input(registers, instructions)
    }

    class Computer(input: Input) {
        var instructionPtr = 0
        val instructions = input.instructions
        val registers = input.registers.toMutableList()
        val output = mutableListOf<Int>()

        private fun Int.literal(): Int = this
        private fun Int.combo() = if (this < 4) this.toLong() else registers[this - 4]

        fun run(): String {
            do {
                var skip = false
                val opcode = instructions[instructionPtr]
                val input = instructions[instructionPtr + 1]
                when (opcode) {
                    0 -> registers[0] = registers[0] shr input.combo().toInt();
                    1 -> registers[1] = registers[1] xor input.literal().toLong()
                    2 -> registers[1] = input.combo() % 8
                    3 -> {
                        if (registers[0] != 0L) {
                            instructionPtr = input.literal()
                            skip = true
                        }
                    }

                    4 -> registers[1] = registers[1] xor registers[2]
                    5 -> output.add((input.combo() % 8).toInt())
                    6 -> registers[1] = registers[0] shr input.combo().toInt()
                    7 -> registers[2] = registers[0] shr input.combo().toInt()
                }

            } while (skip || nextInstruction())

            return getOutput()
        }

        fun getOutput(): String = output.joinToString(",")

        private fun nextInstruction(): Boolean {
            instructionPtr += 2
            return instructionPtr <= instructions.lastIndex
        }


    }

    fun part1(input: List<String>): String {
        val computer = Computer(parseInput(input))
        return computer.run()
    }

    fun part2(input: List<String>): Long {
        val parsedInput = parseInput(input)
        val instructions = parsedInput.instructions.map { it.toLong() }
        fun find(instruction: List<Long>, answer: Long): Long? {
            if (instruction.isEmpty()) return answer
            for (t in 0..7) {
                val a = (answer shl 3) + t
                var b = a % 8
                b = b xor 1
                val c = a shr b.toInt()
                b = b xor 5
                b = b xor c
                if (b % 8 == instruction.last()) {
                    val subSol = find(instruction.dropLast(1), a) ?: continue
                    return subSol
                }
            }
            return null
        }

        return find(instructions, 0)!!;
    }


    check(part1(readInput("Day17_test")) == "4,6,3,5,6,3,5,2,1,0")

    part1(readInput("Day17")).println()
    part2(readInput("Day17")).println()
}
