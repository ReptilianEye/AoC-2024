fun main() {
    fun parseInput(input: List<String>) = input.map { it.toLong() }
    fun Long.mix(other: Long) = this xor other
    fun Long.prune() = this % 16777216

    fun nextSecretNumber(prevSecretNumber: Long): Long {
        var secretNumber = (prevSecretNumber * 64).mix(prevSecretNumber).prune()
        secretNumber = (secretNumber / 32).mix(secretNumber).prune()
        secretNumber = (secretNumber * 2048).mix(secretNumber).prune()
        return secretNumber
    }

    fun part1(initSecretNumbers: List<Long>): Long {
        var sol = 0L
        for (secretNumber in initSecretNumbers) {
            var prev = secretNumber
            repeat(2000) {
                prev = nextSecretNumber(prev)
            }
            sol += prev
        }
        return sol

    }

    fun part2(initSecretNumbers: List<Long>): Int {
        fun generatePrices(secretNum: Long): Pair<List<Int>, List<Int>> {
            val sellerPrice = mutableListOf<Int>()
            var prev = secretNum
            repeat(2000) {
                sellerPrice.add((prev % 10).toInt())
                prev = nextSecretNumber(prev)
            }
            val pricesDiff = (sellerPrice zip (listOf(-1) + sellerPrice)).map { (p1, p2) -> p1 - p2 }.toMutableList()
            pricesDiff[0] = Int.MIN_VALUE

            return sellerPrice to pricesDiff
        }

        val pricesDiffs = initSecretNumbers.fold(listOf<Pair<List<Int>, List<Int>>>()) { pd, secret ->
            val (p, d) = generatePrices(
                secret
            )
            pd + (p to d)
        }
        val seqsForSellers = pricesDiffs.fold(listOf<Map<List<Int>, Int>>()) { acc, (price, diff) ->
            val sellerSeqs = mutableMapOf<List<Int>, Int>()
            (diff.indices.filter { it > 0 } zip diff.indices.filter { it >= 4 }).map { (s, e) ->
                val seq = diff.slice(s..e)
                if (seq !in sellerSeqs)
                    sellerSeqs[seq] = price[e]
            }
            acc + sellerSeqs
        }
        val seqs = seqsForSellers.fold(setOf<List<Int>>()) { acc, map -> acc + map.keys }

        val sol = seqs.maxOf { seq ->
            seqsForSellers.fold(0) { acc, sellerSeqs ->
                acc + sellerSeqs.getOrDefault(seq, 0)
            }
        }

        return sol
    }


    check(part1(parseInput(readInput("Day22_test"))).also { it.println() } == 37327623L)
    check(part2(parseInput(readInput("Day22_test1"))).also { it.println() } == 23)

    part1(parseInput(readInput("Day22"))).println()
    part2(parseInput(readInput("Day22"))).println()
}
