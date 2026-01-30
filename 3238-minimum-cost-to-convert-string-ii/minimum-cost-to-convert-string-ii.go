type TrieNode struct {
    children [26]*TrieNode
    id       int 
}

func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
    const inf = 1e15
    root := &TrieNode{id: -1}
    strToId := make(map[string]int)
    idCount := 0

    getID := func(s string) int {
        if id, ok := strToId[s]; ok { return id }
        strToId[s] = idCount
        insert(root, s, idCount)
        idCount++
        return strToId[s]
    }

    for i := range original {
        getID(original[i])
        getID(changed[i])
    }

    dist := make([][]int64, idCount)
    for i := range dist {
        dist[i] = make([]int64, idCount)
        for j := range dist[i] {
            if i == j { dist[i][j] = 0 } else { dist[i][j] = inf }
        }
    }

    for i := range original {
        u, v := strToId[original[i]], strToId[changed[i]]
        dist[u][v] = min(dist[u][v], int64(cost[i]))
    }

    for k := 0; k < idCount; k++ {
        for i := 0; i < idCount; i++ {
            if dist[i][k] == inf { continue }
            for j := 0; j < idCount; j++ {
                dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
            }
        }
    }

    memo := make([]int64, len(source))
    for i := range memo { memo[i] = -1 }

    var dfs func(int) int64
    dfs = func(i int) int64 {
        if i == len(source) { return 0 }
        if memo[i] != -1 { return memo[i] }

        var res int64 = inf
        
        if source[i] == target[i] {
            res = dfs(i + 1)
        }

        currS, currT := root, root
        for j := i; j < len(source); j++ {
            sChar, tChar := source[j]-'a', target[j]-'a'
            currS = currS.children[sChar]
            currT = currT.children[tChar]
            
            if currS == nil || currT == nil { break }
            
            if currS.id != -1 && currT.id != -1 {
                d := dist[currS.id][currT.id]
                if d < inf {
                    sub := dfs(j + 1)
                    if sub != inf {
                        res = min(res, d + sub)
                    }
                }
            }
        }

        memo[i] = res
        return res
    }

    ans := dfs(0)
    if ans >= inf { return -1 }
    return ans
}

func insert(root *TrieNode, s string, id int) {
    curr := root
    for i := 0; i < len(s); i++ {
        c := s[i] - 'a'
        if curr.children[c] == nil {
            curr.children[c] = &TrieNode{id: -1}
        }
        curr = curr.children[c]
    }
    curr.id = id
}