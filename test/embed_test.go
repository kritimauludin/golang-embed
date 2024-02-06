package test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

/**
Embed File ke String

● Embed file bisa kita lakukan ke variable dengan tipe data string
● Secara otomatis isi file akan dibaca sebagai text dan masukkan ke variable tersebut
*/

//go:embed version.txt
var version string //text

//go:embed version.txt
var version2 string

func TestString(t *testing.T) {
	fmt.Println(version)
	fmt.Println(version2)
}

/**
Embed File ke []byte

● Selain ke tipe data String, embed file juga bisa dilakukan ke variable tipe data []byte
● Ini cocok sekali jika kita ingin melakukan embed file dalam bentuk binary, seperti gambar dan
lain-lain
*/

//go:embed logo.jpg
var logo []byte //gambar/audio/video

func TestByte(t *testing.T) {
	err := os.WriteFile("logo_new.jpg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

/**
Embed Multiple Files

● Kadang ada kebutuhan kita ingin melakukan embed beberapa file sekaligus
● Hal ini juga bisa dilakukan menggunakan embed package
● Kita bisa menambahkan komentar //go:embed lebih dari satu baris
● Selain itu variable nya bisa kita gunakan tipe data embed.FS
*/

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))

}

/**
Path Matcher

● Selain manual satu per satu, kita bisa mengguakan patch matcher untuk membaca multiple file
yang kita inginkan
● Ini sangat cocok ketika misal kita punya pola jenis file yang kita inginkan untuk kita baca
● Caranya, kita perlu menggunakan path matcher seperti pada package function path.Match
https://golang.org/pkg/path/#Match
*/

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := os.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
