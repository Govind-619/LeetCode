func finalValueAfterOperations(operations []string) int {
    x := 0
    for _, v := range operations {
        if v[1] == '+' {  // check the middle character
            x++
        } else {
            x--
        }
    }
    return x
}
