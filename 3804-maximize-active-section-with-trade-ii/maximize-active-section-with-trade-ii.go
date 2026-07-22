func maxActiveSectionsAfterTrade(s string, queries [][]int) []int {
    n := len(s)
    totalOnes := 0
    
    for i := 0; i < n; i++ {
        if s[i] == '1' {
            totalOnes++
        }
    }
    
    typeArr := []int{}
    start := []int{}
    endIdx := []int{}
    
    for i := 0; i < n; {
        j := i
        for j < n && s[j] == s[i] {
            j++
        }
        val := 0
        if s[i] == '1' {
            val = 1
        }
        typeArr = append(typeArr, val)
        start = append(start, i)
        endIdx = append(endIdx, j-1)
        i = j
    }
    
    N := len(typeArr)
    
    posToSeg := make([]int, n)
    for i := 0; i < N; i++ {
        for j := start[i]; j <= endIdx[i]; j++ {
            posToSeg[j] = i
        }
    }
    
    ans := make([]int, N)
    for i := 1; i < N-1; i++ {
        if typeArr[i] == 1 {
            ans[i] = (endIdx[i-1] - start[i-1] + 1) + (endIdx[i+1] - start[i+1] + 1)
        }
    }
    
    logTable := make([]int, N+1)
    for i := 2; i <= N; i++ {
        logTable[i] = logTable[i/2] + 1
    }
    
    K := logTable[N] + 1
    st := make([][]int, K)
    for i := 0; i < K; i++ {
        st[i] = make([]int, N)
    }
    
    for i := 0; i < N; i++ {
        st[0][i] = ans[i]
    }
    
    for j := 1; j < K; j++ {
        for i := 0; i+(1<<j) <= N; i++ {
            a := st[j-1][i]
            b := st[j-1][i+(1<<(j-1))]
            if b > a {
                st[j][i] = b
            } else {
                st[j][i] = a
            }
        }
    }
    
    queryRMQ := func(L, R int) int {
        if L > R {
            return 0
        }
        j := logTable[R-L+1]
        a := st[j][L]
        b := st[j][R-(1<<j)+1]
        if b > a {
            return b
        }
        return a
    }
    
    maxFn := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }
    
    eval := func(i, L, R, segL, segR int) int {
        if i <= segL || i >= segR {
            return 0
        }
        if typeArr[i] == 0 {
            return 0
        }
        
        leftLen := 0
        if i-1 == segL {
            leftLen = maxFn(0, endIdx[i-1]-L+1)
        } else {
            leftLen = endIdx[i-1] - start[i-1] + 1
        }
        
        rightLen := 0
        if i+1 == segR {
            rightLen = maxFn(0, R-start[i+1]+1)
        } else {
            rightLen = endIdx[i+1] - start[i+1] + 1
        }
        
        return leftLen + rightLen
    }
    
    res := make([]int, len(queries))
    for idx, q := range queries {
        L := q[0]
        R := q[1]
        
        segL := posToSeg[L]
        segR := posToSeg[R]
        
        if segR-segL < 2 {
            res[idx] = totalOnes
            continue
        }
        
        maxGain := 0
        maxGain = maxFn(maxGain, eval(segL+1, L, R, segL, segR))
        maxGain = maxFn(maxGain, eval(segR-1, L, R, segL, segR))
        
        if segL+2 <= segR-2 {
            maxGain = maxFn(maxGain, queryRMQ(segL+2, segR-2))
        }
        
        res[idx] = totalOnes + maxGain
    }
    
    return res
}