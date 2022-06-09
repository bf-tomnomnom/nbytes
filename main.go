package main

import (
	"encoding/binary"
	"flag"
	"io"
	"log"
	"os"
)

func main() {
	var n int
	flag.IntVar(&n, "i", 8192, "byte to remove")
	flag.Parse()
	bb, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	for i, b := range bb {
		if (i % n) == 0 {
			continue
		}
		binary.Write(os.Stdout, binary.LittleEndian, b)
	}
}
