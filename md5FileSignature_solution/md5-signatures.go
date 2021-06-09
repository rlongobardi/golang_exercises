package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func calculateMD5Sign(fileToHash []byte) []byte {

	if len(fileToHash) == 0 {
		fmt.Printf("ERROR: file content is empty")
		return []byte{}
	}
	sign := md5.Sum(fileToHash)
	return sign[:]
}

func setupLogging() {
	out, err := os.OpenFile("md5-signature.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer out.Close()
	log.SetOutput(out)
}

func verifyMD5ForAGivenDirectory(fileToVerify io.Reader, directory string) int {
	count := 0
	scanner := bufio.NewScanner(fileToVerify)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		//fmt.Println("File content -> ", scanner.Text())

		content := strings.Fields(scanner.Text())
		if len(content) != 2 {
			fmt.Errorf("File read has a bad line, does not respect the format")
			log.Println("Error bad line, file cannot be processed")
		} else {
			md5Expected := content[0]
			fileLog := content[1]
			fmt.Printf("splitted content -> %s %s\n", md5Expected, fileLog)

			file, err := ioutil.ReadFile(directory + fileLog)
			if err != nil {
				log.Fatal("File was not found:", err)
			}

			md5Calculated := hex.EncodeToString(calculateMD5Sign(file))

			if len(md5Calculated) == 0 {
				fmt.Printf("Error: file was empty or with no content, filename =%s\n", fileLog)
				log.Println("Error: file was empty or with no content", fileLog)
			}

			if strings.Compare(md5Expected, md5Calculated) == 0 {
				fmt.Printf("Success MD5=%s is matching for FILENAME=%s \n", md5Calculated, fileLog)
				count++

			} else {
				fmt.Printf("Error, the two MD5 do not match md5Expected=%s md5Calculated=%x for the file %s\n", md5Expected, md5Calculated, fileLog)

			}
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("ERROR WHILE SCANNING:", err)
		log.Println("ERROR WHILE SCANNING:", err)
	}

	return count
}

func main() {

	setupLogging()
	directory := "nasa-logs/"
	file, err := os.Open(directory + "md5sum.txt")

	if err != nil {
		log.Fatal("Cannot read the main file containg the list of file, error: ", err)
	}

	defer file.Close()
	fmt.Println("Scanning files for MD5 checks..")

	totalVerifiedFiles := verifyMD5ForAGivenDirectory(file, directory)
	files, _ := ioutil.ReadDir(directory)
	fmt.Printf("MD5 scanning is completed, succesful files verified is = %d while the total files were %d\n", totalVerifiedFiles, len(files)-1)

}
