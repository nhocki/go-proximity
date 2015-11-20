package proximity

// TODO: Make this pipeline and not multiple queries
// Get all elements from a redis sorted set that are in one of the intervals.
func query(client Client, set string, intervals []Int64arr) ([]string, error) {
	var results []string
	for _, interval := range intervals {
		responses, err := client.ZRangeByScore(set, float64(interval[0]), float64(interval[1]))
		if err != nil {
			return nil, err
		}
		results = append(results, responses...)
	}
	return results, nil
}
