import (
    "unicode"
    "slices"
)

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
    type coupon struct {
        code string
        line string
    }

    valid := []coupon{}

    lineOrder := map[string]int{
        "electronics": 0,
        "grocery":     1,
        "pharmacy":    2,
        "restaurant":  3,
    }

    n := len(code)

    for i := 0; i < n; i++ {
        if !isActive[i] || code[i] == "" {
            continue
        }

        if _, ok := lineOrder[businessLine[i]]; !ok {
            continue
        }

        ok := true
        for _, v := range code[i] {
            if !unicode.IsLetter(v) && !unicode.IsDigit(v) && v != '_' {
                ok = false
                break
            }
        }

        if ok {
            valid = append(valid, coupon{
                code: code[i],
                line: businessLine[i],
            })
        }
    }

    slices.SortFunc(valid, func(a, b coupon) int {
        if lineOrder[a.line] != lineOrder[b.line] {
            return lineOrder[a.line] - lineOrder[b.line]
        }
        if a.code < b.code {
            return -1
        }
        if a.code > b.code {
            return 1
        }
        return 0
    })

    out := make([]string, 0, len(valid))
    for _, c := range valid {
        out = append(out, c.code)
    }

    return out
}
