package main

import (
    "strconv"
    "strings"
)

type Spreadsheet struct {
    cellValues map[string]int
}

func Constructor(rows int) Spreadsheet {
    return Spreadsheet{cellValues: make(map[string]int)}
}

func (s *Spreadsheet) SetCell(cell string, value int) {
    s.cellValues[cell] = value
}

func (s *Spreadsheet) ResetCell(cell string) {
    // either delete or set to zero
    // delete is fine; lookup in getValue treats missing as zero
    delete(s.cellValues, cell)
}

func (s *Spreadsheet) GetValue(formula string) int {
    // formula starts with "="
    expr := formula[1:]
    parts := strings.Split(expr, "+")
    sum := 0
    for _, part := range parts {
        part = strings.TrimSpace(part)
        if part == "" {
            continue
        }
        // check if it's an integer or a cell reference
        // integer if first character is digit
        r := part[0]
        if r >= '0' && r <= '9' {
            // parse integer
            if v, err := strconv.Atoi(part); err == nil {
                sum += v
            } // else maybe invalid, but per problem, should be valid
        } else {
            // assume cell reference; if not set, get zero
            if v, ok := s.cellValues[part]; ok {
                sum += v
            }
            // else add 0
        }
    }
    return sum
}
