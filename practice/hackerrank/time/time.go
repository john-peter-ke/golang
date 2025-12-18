package main

import (
	"time"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func TimeConversion(s string) string {
	layout1 := "03:04:05PM"
	layout2 := "15:04:05"
	t, err := time.Parse(layout1, s)
	if err != nil {
		return ""
	}
	return t.Format(layout2)

}
