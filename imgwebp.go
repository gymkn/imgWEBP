package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/deepteams/webp"
)

func convert(inputDir, outputDir string, deleteOld bool, quality int) {
	files, err := os.ReadDir(inputDir) // List files in the current directory
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Files found: ", len(files))
	fmt.Println("Starts converting...")

	for _, file := range files {
		fullFilename := file.Name()
		filename := strings.TrimSuffix(fullFilename, filepath.Ext(fullFilename))

		if file.IsDir() || strings.HasPrefix(fullFilename, ".") {
			continue
		}

		in, err := os.Open(filepath.Join(inputDir, fullFilename))
		if err != nil {
			log.Fatal(err)
		}
		defer in.Close()

		img, _, err := image.Decode(in)
		if err != nil {
			log.Fatal(err)
		}

		out, err := os.Create(filepath.Join(outputDir, filename+".webp"))
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		err = webp.Encode(out, img, &webp.EncoderOptions{
			Quality: float32(quality),
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	// delete if user wants to
	if deleteOld {
		if outputDir == inputDir {
			for _, file := range files {
				os.Remove(filepath.Join(inputDir, file.Name()))
			}
		} else {
			os.RemoveAll(inputDir)
		}
	}
}

func main() {
	inputDir := flag.String("input", ".", "Directory with files")
	outputDir := flag.String("output", "", "Alternative directory for output")
	deleteOld := flag.Bool("delete", false, "Delete old files")
	quality := flag.Int("quality", 80, "Quality of picture after processing from 1-100")
	flag.Parse()

	if *outputDir == "" {
		*outputDir = *inputDir
	}
	*inputDir = filepath.Clean(*inputDir)
	*outputDir = filepath.Clean(*outputDir)
	convert(*inputDir, *outputDir, *deleteOld, *quality)
	fmt.Println("Done")
}
