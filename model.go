package main

//"github.com/gabetucker2/gostack"

func runModel() {

	for x := 0; x < 255; x++ {
		for y := 0; y < 255; y++ {
			imgMatrix[x][y] = (imgMatrix[x][y] + 1) % 255
		}
	}

}
