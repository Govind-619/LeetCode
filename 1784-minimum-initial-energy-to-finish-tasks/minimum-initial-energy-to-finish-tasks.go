func minimumEffort(tasks [][]int) int {

    sort.Slice(tasks, func(i, j int) bool {
        return (tasks[i][1] - tasks[i][0]) > (tasks[j][1] - tasks[j][0])
    })

    answer := 0
    energy := 0

    for _, task := range tasks {

        actual := task[0]
        minimum := task[1]

        if energy < minimum {

            need := minimum - energy

            answer += need
            energy += need
        }

        energy -= actual
    }

    return answer
}