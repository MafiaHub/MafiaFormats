package main

import (
	"bufio"
	"encoding/binary"
	"flag"
	"fmt"
	"os"

	"strings"

	jsoniter "github.com/json-iterator/go"
)

type wrappedTextBlob struct {
	ID   uint32 `json:"id"`
	Text string `json:"text"`
}

type wrappedFormat struct {
	Blobs []wrappedTextBlob `json:"blobs"`
}

func main() {
	filePath := flag.String("file", "", "path to the file")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Not enough arguments!")
		return
	}

	data := FormatTextDB{
		TextBlobs: []TextBlob{},
	}
	f, err := os.Open(*filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}

	binary.Read(f, binary.LittleEndian, &data.NumBlobs)
	binary.Read(f, binary.LittleEndian, &data.Unknown)

	out := wrappedFormat{
		Blobs: []wrappedTextBlob{},
	}

	for i := 0; i < int(data.NumBlobs); i++ {
		var blob TextBlob
		binary.Read(f, binary.LittleEndian, &blob.TextID)
		binary.Read(f, binary.LittleEndian, &blob.TextOffset)
		data.TextBlobs = append(data.TextBlobs, blob)
	}

	for _, v := range data.TextBlobs {
		f.Seek(int64(v.TextOffset), os.SEEK_SET)
		var text string = "lala"
		// binary.Read(text)
		r := bufio.NewReader(f)
		text, _ = r.ReadString(0)
		text = strings.TrimSuffix(text, "\x00")

		out.Blobs = append(out.Blobs, wrappedTextBlob{
			ID:   v.TextID,
			Text: text,
		})
	}

	output, _ := jsoniter.MarshalToString(out)
	fmt.Println(output)
}
