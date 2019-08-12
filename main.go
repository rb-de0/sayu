package main

import (
	"flag"
	"image"
	"image/jpeg"
	"os"
	"strconv"

	"github.com/minodisk/go-fix-orientation/processor"
	"golang.org/x/image/draw"
)

const outputDir = "output"

func main() {
	// parse input
	flag.Parse()
	if len(flag.Args()) == 0 {
		panic("no arguments")
	}
	fileName := flag.Arg(0)
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	resizeParcentage, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		panic(err)
	}
	resize := 100.0 / float64(resizeParcentage)
	// fix orientation
	fixed, err := processor.Process(file)
	if err != nil {
		panic(err)
	}
	// resize image
	srcBounds := fixed.Bounds()
	width := int(float64(srcBounds.Dx()) / resize)
	height := int(float64(srcBounds.Dy()) / resize)
	outputImage := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(outputImage, outputImage.Bounds(), fixed, srcBounds, draw.Over, nil)
	// write image
	outputFileName := outputDir + "/" + fileName
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.Mkdir(outputDir, 0777); err != nil {
			panic(err)
		}
	}
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	jpeg.Encode(outputFile, outputImage, nil)
}
