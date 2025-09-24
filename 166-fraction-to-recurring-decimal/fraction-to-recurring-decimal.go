import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
    if numerator == 0 {
        return "0"
    }

    result := ""

    // Handle negative
    if (numerator < 0) != (denominator < 0) {
        result += "-"
    }

    // Work with positive values
    num := int64(abs(numerator))
    den := int64(abs(denominator))

    // Integer part
    result += strconv.FormatInt(num/den, 10)
    remainder := num % den
    if remainder == 0 {
        return result
    }

    result += "."

    // Fractional part
    m := make(map[int64]int)
    for remainder != 0 {
        if idx, ok := m[remainder]; ok {
            // Insert parentheses for repeating part
            return result[:idx] + "(" + result[idx:] + ")"
        }
        m[remainder] = len(result)

        remainder *= 10
        result += strconv.FormatInt(remainder/den, 10)
        remainder %= den
    }

    return result
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
