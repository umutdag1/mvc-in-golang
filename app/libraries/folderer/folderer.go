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

type FolderF []fs.FileInfo

func OpenFolder(folderPath string) (FolderF, error) {
	logger.InfoLogger.Println(fmt.Sprintf("Opening Folder %v", folderPath))
	filesInfo, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}
	folderF = FolderF(filesInfo)
	logger.InfoLogger.Println(fmt.Sprintf("Opened Folder %v", folderPath))
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
