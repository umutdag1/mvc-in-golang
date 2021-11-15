package filer

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/rest-api/app/libraries/logger"
)

var (
	filerF FilerF
)

type FilerFI interface {
	WriteFile(interface{})
}

type FilerF os.File

func CreateFile(targetPath, name, extension string) error {
	fileName := fmt.Sprintf("%v.%v", name, extension)
	logger.InfoLogger.Println(fmt.Sprintf("Creating File %v", targetPath+"/"+fileName))
	isExist, _ := Exists(targetPath)
	if !isExist {
		err := os.Mkdir(targetPath, 0755)
		if err != nil {
			logger.ErrorLogger.Println(err.Error())
			return err
		}
	}
	_, err := os.Create(targetPath + "/" + fileName)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return err
	}
	logger.InfoLogger.Println(fmt.Sprintf("Created File %v", targetPath+"/"+fileName))

	return nil
}

func OpenFile(filePath string) (FilerF, error) {
	logger.InfoLogger.Println(fmt.Sprintf("Opening File %v", filePath))
	if isExist, _ := Exists(filePath); !isExist {
		filePathPartials := strings.Split(filePath, "/")
		folderPath := strings.Join(filePathPartials[:len(filePathPartials)-1], "/")
		fileNameAndExt := strings.Split(filePathPartials[len(filePathPartials)-1], ".")
		if err := CreateFile(folderPath, fileNameAndExt[0], fileNameAndExt[1]); err != nil {
			logger.ErrorLogger.Println(err.Error())
			return filerF, err
		}
	}
	f, err := os.OpenFile(filePath, os.O_WRONLY, 0644)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return filerF, err
	}
	filerF = FilerF(*f)
	logger.InfoLogger.Println(fmt.Sprintf("Opened File %v", filePath))

	return filerF, nil
}

func (fi *FilerF) WriteFile(data []byte) error {
	f := os.File(*fi)
	logger.InfoLogger.Println(fmt.Sprintf("Writing File %v", f.Name()))
	err := os.WriteFile(f.Name(), data, 0644)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return err
	}
	logger.InfoLogger.Println(fmt.Sprintf("Written File %v", f.Name()))
	return nil
}

func (fi *FilerF) ReadFile() ([]byte, error) {
	f := os.File(*fi)
	fmt.Println(f.Name())
	dataByte, err := ioutil.ReadFile(f.Name())
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return nil, err
	}
	return dataByte, err
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) || err != nil {
		return false, err
	}
	return true, nil
}
