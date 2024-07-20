package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
	"unicode/utf8"

	"github.com/atotto/clipboard"
)

// There are some more advanced clipboard actions https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpeclip/30688d09-96b6-46f8-af18-ea1998bb7987
// but Not sure if we need them
func main() {
	args := os.Args
	if len(args) == 1 {
		install(os.Args[0])
		return
	}
	if len(args) > 2 {
		logMessage(fmt.Sprintf("Unexpected number of arguments %d", len(args)))
		return
	}
	
	// This is somehow slightly exiting
	// [D:\CopyToClipboard\copytoclipboard.exe, C:\Users\Alexander\Downloads\gimp-2.10.36-setup-1.exe]
	fileName := os.Args[1]
	if fileName == "-u" {
		uninstall()
		return
	}

	readable, err := isReadableText(fileName)
	if err != nil{
		logMessage(fmt.Sprintf("Unable to open file: %v", err))
		return
	}
	if !readable {
		return
	}

	b, err := os.ReadFile(fileName)
	if err != nil{
		logError(err)
		return
	}

	err = clipboard.WriteAll(string(b))
	if err != nil {
		logError(err)
		return
	}
}

func logError(err error){
	logMessage(err.Error())
}

func logMessage(message string){
	// Going to also write to standard output in case the user is in an actual command line
	fmt.Println(message)
	logLocation := ""
	// Can't imagine this ever happens
	if len(os.Args) > 0 {
		logLocation = os.Args[0]
	}else{
		logLocation = "./"
	}
	logLocation = filepath.Dir(logLocation)
	logLocation = filepath.Join(logLocation, "logs.txt")
	message = fmt.Sprintf("%s: %s\n", time.Now().Format(time.DateTime), message)
	file, err := os.OpenFile(logLocation, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(message)
	if err != nil{
		fmt.Println(err)
		return
	}
}

// Feel like instead of checking the file, finding out if its a compatible text type would be better
func handleExtension(ext string)(accepted bool){
	switch ext{
	case ".txt":
		return true
	case ".json":
		return true
	case ".html":
		return true
	case ".yaml":
		return true
	case ".yml":
		return true
	}
	return
}

// Silly like read the file twice, but don't feel like fixing right now
func isReadableText(filePath string) (bool, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return false, err
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    buffer := make([]byte, 1024) // Read first 1024 bytes
    bytesRead, err := reader.Read(buffer)
    if err != nil {
        return false, err
    }

    return utf8.Valid(buffer[:bytesRead]), nil
}