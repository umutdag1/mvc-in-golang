package filer

import (
	"fmt"
	"os"
	"time"

	"github.com/rest-api/app/libraries/logger"
)

var (
	filerF FilerF
)

type FilerFI interface {
	WriteFile(interface{})
}

type FilerF os.File

func OpenFile(filePath, name, extension string) (FilerF, error) {
	fileName := fmt.Sprintf("%v-%v.%v", time.Now().UTC().String(), name, extension)
	targetPath := fmt.Sprintf("%v", filePath)
	logger.InfoLogger.Println(fmt.Sprintf("Opening File %v", targetPath+"/"+fileName))
	isExist, _ := Exists(targetPath)
	if !isExist {
		err := os.Mkdir(targetPath, 0755)
		if err != nil {
			logger.ErrorLogger.Println(err.Error())
			return filerF, err
		}
	}
	f, err := os.OpenFile(targetPath+"/"+fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return filerF, err
	}
	filerF = FilerF(*f)
	logger.InfoLogger.Println(fmt.Sprintf("Opened File %v", targetPath+"/"+fileName))

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

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) || err != nil {
		return false, err
	}
	return true, nil
}
