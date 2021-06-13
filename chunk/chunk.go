package chunk

import (
	"encoding/base64"
	"fmt"
)

type (
	Chunk struct {
		hash string
		data []byte
	}

	File struct {
		chunkHashes []string
	}
)

func main() {
	fA := "abcdefghijklmno"

	fB := "abcdefghi123456"

	// chunked is all of our file
	chunked := map[string]Chunk{}

	// 4 byte
	chunkSize := 4

	fAByte := []byte(fA)
	fBByte := []byte(fB)

	datas := [][]byte{fAByte, fBByte}
	// files represent db data
	files := []File{}
	for _, data := range datas {
		f := File{}
		currentLen := 0
		currentStart := 0
		for i := range data {
			if currentLen == chunkSize {
				data := data[currentStart:i]
				hash := base64.StdEncoding.EncodeToString(data)
				c := Chunk{
					data: data,
					hash: hash,
				}
				chunked[hash] = c
				f.chunkHashes = append(f.chunkHashes, hash)
				currentLen = 0
				currentStart = i
			}
			currentLen++
		}
		// data remain
		data := data[currentStart:]
		hash := base64.StdEncoding.EncodeToString(data)
		c := Chunk{
			data: data,
			hash: hash,
		}
		chunked[hash] = c
		f.chunkHashes = append(f.chunkHashes, hash)

		files = append(files, f)
	}

	fmt.Println(files)

	// reverse from chunks

	for _, f := range files {
		//construct file data form chunk
		fileChunks := []byte{}

		for _, chunkHash := range f.chunkHashes {
			fileChunks = append(fileChunks, chunked[chunkHash].data...)
		}

		fmt.Println(string(fileChunks))
	}
}
