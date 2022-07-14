package main

import (
	"crypto/md5"
	"fmt"
	"image/png"
	"io"
	"os"

	"identicon/internal"
)

func main() {
	bytes, hashErr := hash()
	if hashErr != nil {
		fmt.Println(hashErr)
		os.Exit(1)
	}

	generateErr := generate(bytes)
	if generateErr != nil {
		fmt.Println(generateErr)
		os.Exit(1)
	}
}

func generate(input []uint8) error {
	identicon := internal.NewIdenticon(input)
	image := identicon.Image()
	err := png.Encode(os.Stdout, image)
	return err
}

func hash() ([]uint8, error) {
	digest := md5.New()
	_, err := io.Copy(digest, os.Stdin)
	if err != nil {
		return nil, err
	}

	bytes := make([]uint8, 0)
	bytes = digest.Sum(bytes)
	return bytes, nil
}
