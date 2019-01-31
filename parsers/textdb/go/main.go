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
	TextBlob
	Text string `json:"text"`
}

type wrappedFormat struct {
	FormatTextDB
	Blobs []wrappedTextBlob `json:"blobs"`
}

func main() {
	filePath := flag.String("file", "", "path to the file")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Not enough arguments!")
		return
	}

	data := wrappedFormat{
		Blobs: []wrappedTextBlob{},
	}
	f, err := os.Open(*filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}

	binary.Read(f, binary.LittleEndian, &data.NumBlobs)
	binary.Read(f, binary.LittleEndian, &data.Unknown)

	for i := 0; i < int(data.NumBlobs); i++ {
		var blob wrappedTextBlob
		binary.Read(f, binary.LittleEndian, &blob.TextID)
		binary.Read(f, binary.LittleEndian, &blob.TextOffset)
		pos, _ := f.Seek(0, os.SEEK_CUR)

		f.Seek(int64(blob.TextOffset), os.SEEK_SET)
		var text string
		r := bufio.NewReader(f)
		text, _ = r.ReadString(0)
		text = strings.TrimSuffix(text, "\x00")
		blob.Text = text
		f.Seek(pos, os.SEEK_SET)

		data.Blobs = append(data.Blobs, blob)
	}

	output, _ := jsoniter.MarshalToString(data)
	fmt.Println(output)
}
