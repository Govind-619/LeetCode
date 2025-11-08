func minimumOneBitOperations(n int) int {
    mult := 1
    ops := 0

    for shift := 30; shift >= 0; shift-- {
        if (1 << shift) & n == 0 {
            continue
        }

        ops += mult * ((1 << (shift+1)) - 1)
        mult *= -1
    }

    return ops
}