package handlers

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type LocalFileManipulation struct {
}

func (lfm LocalFileManipulation) folderExists(folderName string) bool {
	_, err := os.Stat(folderName)
	return os.IsExist(err)
}

func (lfm LocalFileManipulation) New() {
	lfm.createFolder("uploads")
}

func (lfm LocalFileManipulation) createFolder(path string) error {
	folderExists := lfm.folderExists(path)
	if !folderExists {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("no se pudo crear la carepta '%s': %w", path, err)
		}
	}
	return nil
}

func (lfm LocalFileManipulation) CreateCustomFolder(nameDir string) {
	path := fmt.Sprintf("uploads/%s", nameDir)
	lfm.createFolder(path)
}

func (lfm LocalFileManipulation) DeleteAllFolders() {
	filepath.Walk("uploads/", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != "uploads/" {
			os.RemoveAll(path)
		}
		return nil
	})
}

func (lfm LocalFileManipulation) GetFolders() []string {
	folderChannel := make(chan string)
	wg := &sync.WaitGroup{}
	go sendFolder(folderChannel, wg)
	wg.Wait()

	folders := processFolder(folderChannel, wg)
	return folders
}

func sendFolder(folderChannel chan string, wg *sync.WaitGroup) {
	folders, err := os.ReadDir("uploads")
	if err != nil {
		log.Println("error al obtener los folders:", err)
		close(folderChannel)
	}
	wg.Add(len(folders))
	for _, folder := range folders {
		folderChannel <- folder.Name()
	}
	close(folderChannel)
}

func processFolder(folderChannel chan string, wg *sync.WaitGroup) []string {
	defer wg.Done()
	var folders []string
	for folder := range folderChannel {
		folders = append(folders, folder)
	}
	return folders
}

func (lfm LocalFileManipulation) findFolder(folderChannel chan []fs.DirEntry, wg *sync.WaitGroup, searchedFolder, folder string) {
	if searchedFolder == folder {
		fmt.Println("xd")
		dirs, err := os.ReadDir(fmt.Sprintf("uploads/%s/", searchedFolder))
		if err != nil {
			fmt.Println("error al buscar el folder:", err)
			close(folderChannel)
		}
		folderChannel <- dirs
		close(folderChannel)

	}
}

func (lfm LocalFileManipulation) processFindFolder(folderChannel chan []fs.DirEntry, wg *sync.WaitGroup) []fs.DirEntry {
	defer wg.Done()
	dirs := <-folderChannel
	return dirs
}

func (lfm LocalFileManipulation) GetFolderContent(folderName string) {
	folderChannel := make(chan []fs.DirEntry)
	wg := &sync.WaitGroup{}
	folders := lfm.GetFolders()
	wg.Add(len(folders))

	for _, folder := range folders {
		go lfm.findFolder(folderChannel, wg, folderName, folder)
		wg.Wait()

		dirs := lfm.processFindFolder(folderChannel, wg)

		fmt.Println("aea")

		if dirs != nil {
			fmt.Println(dirs)
			// return dirs
		}
	}
}
