package main 

import (
    "fmt"
    "image"
    "image/color"
    "image/png"
    "bufio"
    "os"
    "strconv"
)

func (w *World) updateMap(view int) {
    p := createPalette(view)
    w.Map = *image.NewRGBA(image.Rect(0, 0, len(w.Grid[view]), len(w.Grid[view][0])))
    for x := 0; x < w.Map.Bounds().Max.X; x++ {
	for y := 0; y < w.Map.Bounds().Max.Y; y++ {
	    c := p[int(w.Grid[view][x][y] * 31.0)]
	    fmt.Print(c)
	    w.Map.Set(x, y, c) 
	}
    }
}

func (w *World) saveMap() {
    f, err := os.Create("map.png")
    if err != nil {
	panic(err)
    }
    defer f.Close()
    png.Encode(f, &w.Map)
}

func createPalette(view int) (p color.Palette) {
    var path string
    if view == ELEVATION {
	path = "res/palettes/elevation.pal"
    } else if view == CLIMATE {
	path = "res/palettes/climate.pal"
    } else if view == POLITICAL {
	path = "res/palettes/political.pal"
    } else {
	path = "res/palettes/biome.pal"
    }
  
    hexColors, err := readLines(path)
    if err != nil {
	fmt.Print(path + "could not be opened.\n")
    }
    
    p = color.Palette{} 
    for i := 0; i < len(hexColors); i++ {
	r, _ := strconv.ParseUint(hexColors[i][1:3], 16, 8)
	g, _ := strconv.ParseUint(hexColors[i][3:5], 16, 8)
	b, _ := strconv.ParseUint(hexColors[i][5:7], 16, 8)
	var a uint8 = 255
	fmt.Print(hexColors[i][1:3] + "\n")
	fmt.Print(hexColors[i][3:5] + "\n")
	fmt.Print(hexColors[i][5:7] + "\n")
	fmt.Println(r, g, b, a)
	c := color.RGBA{uint8(r), uint8(g), uint8(b), a}
	p = append(p, c)
    }
    return
} 

func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
	return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
	lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}
