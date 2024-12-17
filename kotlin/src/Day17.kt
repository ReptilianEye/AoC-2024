import kotlin.math.pow

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

        fun run(lookingFor: List<Int>? = null): String {
            var lookForIdx = 0
            do {
                var skip = false
                val opcode = instructions[instructionPtr]
                val input = instructions[instructionPtr + 1]
                when (opcode) {
                    0 -> registers[0] = (registers[0] / 2.toDouble().pow(input.combo().toInt())).toLong()
                    1 -> registers[1] = registers[1] xor input.literal().toLong()
                    2 -> registers[1] = input.combo() % 8
                    3 -> {
                        if (registers[0] != 0L) {
                            instructionPtr = input.literal()
                            skip = true
                        }
                    }

                    4 -> registers[1] = registers[1] xor registers[2]
                    5 -> {
                        output.add((input.combo() % 8).toInt())
                        if (lookingFor != null) {
                            if (lookForIdx >= lookingFor.size) {
                                return ""
                            }
                            if (output.last() != lookingFor[lookForIdx]) {
                                return ""
                            }
                            lookForIdx++
                        }
                    }

                    6 -> registers[1] = (registers[0] / 2.toDouble().pow(input.combo().toInt())).toLong()
                    7 -> registers[2] = (registers[0] / 2.toDouble().pow(input.combo().toInt())).toLong()
                }
//                println(output)

            } while (skip || nextInstruction())

            return getOutput()
        }

        fun getOutput(): String = output.joinToString(",")

        private fun nextInstruction(): Boolean {
            instructionPtr += 2
            return instructionPtr <= instructions.lastIndex
        }

        private fun Int.literal(): Int = this
        private fun Int.combo() = if (this < 4) this.toLong() else registers[this - 4]
    }

    fun part1(input: List<String>): String {
        val computer = Computer(parseInput(input))
        return computer.run()
    }

    fun part2(input: List<String>): Long {
        val parsedInput = parseInput(input)
        val initInstructions = parsedInput.instructions.joinToString(",")
//        36969000000
        var betterA = 36969000000
        betterA = 0
        while (true) {
            val computer = Computer(
                Input(
                    listOf(betterA, parsedInput.registers[1], parsedInput.registers[2]),
                    parsedInput.instructions
                )
            )
            val res = computer.run(parsedInput.instructions)
            if (res == initInstructions) {
                return betterA
            }
//            if (betterA % 1000000 == 0L)
//            println(betterA)
            betterA++

        }
    }



    check(part1(readInput("Day17_test")) == "4,6,3,5,6,3,5,2,1,0")
//    check(part2(readInput("Day17_test1")) == 117440L)


    part1(readInput("Day17")).println()
    part2(readInput("Day17")).println()
}
