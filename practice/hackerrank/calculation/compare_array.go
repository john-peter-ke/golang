package main

/*
 * Complete the 'compareTriplets' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func CompareTriplets(a []int32, b []int32) []int32 {
	const point int32 = 1
	var suma, sumb int32
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			suma = suma + point
		} else if a[i] < b[i] {
			sumb = sumb + point
		}
	}

	return []int32{suma, sumb}
}
