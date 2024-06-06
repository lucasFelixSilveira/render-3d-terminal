package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type Point struct {
	x, y, z float64
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") // Windows
	if runtime.GOOS == "linux" {
		cmd = exec.Command("clear") // Linux
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func drawPoint(x, y int) {
	fmt.Printf("\033[%d;%dH*", y+1, x+1)
}

func rotatePoint(p *Point, angleX, angleY, angleZ float64) {
	cosX, sinX := math.Cos(angleX), math.Sin(angleX)
	cosY, sinY := math.Cos(angleY), math.Sin(angleY)
	cosZ, sinZ := math.Cos(angleZ), math.Sin(angleZ)

	// Rotate around X axis
	y1 := p.y*cosX - p.z*sinX
	z1 := p.y*sinX + p.z*cosX
	p.y, p.z = y1, z1

	// Rotate around Y axis
	x1 := p.x*cosY + p.z*sinY
	z2 := -p.x*sinY + p.z*cosY
	p.x, p.z = x1, z2

	// Rotate around Z axis
	x2 := p.x*cosZ - p.y*sinZ
	y2 := p.x*sinZ + p.y*cosZ
	p.x, p.y = x2, y2
}

func drawRotatingCube(angleX, angleY, angleZ float64) {
	clearScreen()

	halfSize := 10.0
	vertices := []Point{
		{-halfSize, -halfSize, -halfSize},
		{halfSize, -halfSize, -halfSize},
		{halfSize, halfSize, -halfSize},
		{-halfSize, halfSize, -halfSize},
		{-halfSize, -halfSize, halfSize},
		{halfSize, -halfSize, halfSize},
		{halfSize, halfSize, halfSize},
		{-halfSize, halfSize, halfSize},
	}

	for i := range vertices {
		rotatePoint(&vertices[i], angleX, angleY, angleZ)
	}

	for _, vertex := range vertices {
		screenX := int(vertex.x/(vertex.z+20)*40 + 80/2)
		screenY := int(vertex.y/(vertex.z+20)*20 + 25/2)
		drawPoint(screenX, screenY)
	}

	edges := [][]int{
		{0, 1}, {1, 2}, {2, 3}, {3, 0},
		{4, 5}, {5, 6}, {6, 7}, {7, 4},
		{0, 4}, {1, 5}, {2, 6}, {3, 7},
	}

	for _, edge := range edges {
		p1 := vertices[edge[0]]
		p2 := vertices[edge[1]]

		x1 := int(p1.x/(p1.z+20)*40 + 80/2)
		y1 := int(p1.y/(p1.z+20)*20 + 25/2)
		x2 := int(p2.x/(p2.z+20)*40 + 80/2)
		y2 := int(p2.y/(p2.z+20)*20 + 25/2)

		dx := int(math.Abs(float64(x2 - x1)))
		dy := int(math.Abs(float64(y2 - y1)))
		sx, sy := 1, 1
		if x1 >= x2 {
			sx = -1
		}
		if y1 >= y2 {
			sy = -1
		}
		err := dx - dy

		for {
			drawPoint(x1, y1)
			if x1 == x2 && y1 == y2 {
				break
			}
			e2 := 2 * err
			if e2 > -dy {
				err -= dy
				x1 += sx
			}
			if e2 < dx {
				err += dx
				y1 += sy
			}
		}
	}
}

func main() {
	angleX, angleY, angleZ := 0.0, 0.0, 0.0
	delay := 100 * time.Millisecond

	for {
		drawRotatingCube(angleX, angleY, angleZ)
		time.Sleep(delay)
		angleX += 0.05
		angleY += 0.05
		angleZ += 0.05
	}
}
