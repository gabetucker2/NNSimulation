package main

func runModel() {

	for x := 0; x < W; x++ {
		for y := 0; y < H; y++ {
			imgMatrix[x][y] = (imgMatrix[x][y] + 20) % 255
		}
	}

}
