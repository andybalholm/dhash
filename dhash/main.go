package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/andybalholm/dhash"
)

func main() {
	switch len(os.Args) {
	case 2:
		img, err := loadImage(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dhash.New(img))

	case 3:
		img1, err := loadImage(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		img2, err := loadImage(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		hash1 := dhash.New(img1)
		hash2 := dhash.New(img2)

		fmt.Printf("%s: %v\n", os.Args[1], hash1)
		fmt.Printf("%s: %v\n", os.Args[2], hash2)
		fmt.Printf("Bits different: %d\n", dhash.Distance(hash1, hash2))

	default:
		fmt.Fprintf(os.Stderr, "Usage:\n\t%s file (to hash one image)\n\t%s file1 file2 (to compare two image)\n", os.Args[0], os.Args[0])
		os.Exit(1)
	}
}

func loadImage(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	return img, err
}
