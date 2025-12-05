package main

/*
 * Complete the 'chessboardGame' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER x
 *  2. INTEGER y
 */

func ChessboardGame(x int32, y int32) string {
	// The game's winning/losing positions depend on the coordinates modulo 4.
	// This is a known property of this specific game.
	xMod4 := x % 4
	yMod4 := y % 4

	// If both x and y modulo 4 fall within the losing positions (1 or 2),
	// then the current player is in a losing state, making the other player win.
	if (xMod4 == 1 || xMod4 == 2) && (yMod4 == 1 || yMod4 == 2) {
		return "Second"
	}
	// Otherwise, the current player is in a winning state.
	return "First"
}
