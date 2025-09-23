func titleToNumber(columnTitle string) int {
    final := 0

    // iterate through all characters
    for _, v := range columnTitle {
        num := int(v % 65) + 1  // take remainder of current char and add 1 becuase column number in excel starts from 1
        final = (final * 26) + num
    }

    return final
}