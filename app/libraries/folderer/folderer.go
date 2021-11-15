package folderer

import (
	"fmt"
	"io/fs"
	"io/ioutil"

	"github.com/rest-api/app/libraries/logger"
)

var (
	folderF FolderF
)

type FolderFI interface {
}

type FolderF []fs.FileInfo

func OpenFolder(folderPath string) (FolderF, error) {
	filesInfo, err := ioutil.ReadDir(folderPath)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return nil, err
	}
	folderF = FolderF(filesInfo)
	for _, f := range filesInfo {
		fmt.Println(f.Name())
	}
	return folderF, nil
}

func (fo FolderF) GetFilesPathsInFolder() []string {
	files := []fs.FileInfo(fo)
	filesNames := []string{}
	for _, f := range files {
		filesNames = append(filesNames, f.Name())
	}
	return filesNames
}
