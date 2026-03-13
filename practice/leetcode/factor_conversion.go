package leetcode

type conversion struct {
	from   string
	to     string
	factor float64
}

type node struct {
	unit   string
	factor float64
}

func factorConvertMaps(conversions []conversion, a, b string) float64 {
	graph := make(map[string]map[string]float64)
	for _, c := range conversions {
		if graph[c.from] == nil {
			graph[c.from] = make(map[string]float64)
		}
		if graph[c.to] == nil {
			graph[c.to] = make(map[string]float64)
		}
		graph[c.from][c.to] = c.factor
		graph[c.to][c.from] = 1 / c.factor

	}

	queue := []node{{unit: a, factor: 1.0}}
	visited := map[string]bool{a: true}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.unit == b {
			return curr.factor
		}

		for neighbor, nextFactor := range graph[curr.unit] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, node{
					unit:   neighbor,
					factor: curr.factor * nextFactor,
				})
			}
		}
	}

	return 0
}

type edge struct {
	To     string
	Weight float64
}

func factorConvertWithList(conversions [][]interface{}, start, end string) float64 {
	// 1. Build Adjacency List
	graph := make(map[string][]edge)
	for _, conv := range conversions {
		u, v := conv[0].(string), conv[1].(string)
		val := conv[2].(float64)
		graph[u] = append(graph[u], edge{v, val})
		graph[v] = append(graph[v], edge{u, 1 / val})
	}

	queue := []node{{start, 1.0}}
	visited := map[string]bool{start: true}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.unit == end {
			return curr.factor
		}

		for _, edge := range graph[curr.unit] {
			if !visited[edge.To] {
				visited[edge.To] = true
				queue = append(queue, node{edge.To, curr.factor * edge.Weight})
			}
		}
	}
	return 0 // Return -1 if no path exists
}

func factorConvertMatrix(conversions [][]interface{}, start, end string) float64 {
	// 1. Map units to indices
	unitSet := make(map[string]int)
	units := []string{}
	for _, c := range conversions {
		for i := 0; i < 2; i++ {
			u := c[i].(string)
			if _, exists := unitSet[u]; !exists {
				unitSet[u] = len(units)
				units = append(units, u)
			}
		}
	}

	n := len(units)
	// 2. Initialize Matrix with 0.0 (or -1.0 for no path)
	matrix := make([][]float64, n)
	for i := range matrix {
		matrix[i] = make([]float64, n)
		matrix[i][i] = 1.0 // Unit to itself is always 1
	}

	// 3. Fill direct conversions
	for _, c := range conversions {
		u, v := unitSet[c[0].(string)], unitSet[c[1].(string)]
		val := c[2].(float64)
		matrix[u][v] = val
		matrix[v][u] = 1.0 / val
	}

	// 4. Floyd-Warshall: Find all-pairs paths
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				// If a path i->j doesn't exist but i->k and k->j do
				if matrix[i][j] == 0 && matrix[i][k] != 0 && matrix[k][j] != 0 {
					matrix[i][j] = matrix[i][k] * matrix[k][j]
				}
			}
		}
	}

	startIdx, sOk := unitSet[start]
	endIdx, eOk := unitSet[end]
	if !sOk || !eOk || matrix[startIdx][endIdx] == 0 {
		return -1.0
	}
	return matrix[startIdx][endIdx]
}
