func mapWordWeights(words []string, weights []int) string {
	var result []byte

	for _, word := range words {
		sumWeight := 0

		for _, ch := range word {
			sumWeight += weights[ch-'a']
		}

		value := sumWeight % 26

		result = append(result, byte('z'-value))
	}

	return string(result)
}