func decodeCiphertext(encodedText string, rows int) string {
    if len(encodedText) == 0 {
        return ""
    }

    n := len(encodedText)
    cols := n / rows

    result := make([]byte, 0)

    for startCol := 0; startCol < cols; startCol++ {
        row, col := 0, startCol

        for row < rows && col < cols {
            index := row*cols + col
            result = append(result, encodedText[index])
            row++
            col++
        }
    }

    end := len(result) - 1
    for end >= 0 && result[end] == ' ' {
        end--
    }

    return string(result[:end+1])
}