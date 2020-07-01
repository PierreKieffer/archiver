package archiver

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

func ZipArchiver(basePath, outputPath string) error {

	log.Println("INFO : Compression " + basePath + " to " + outputPath + " ...")

	// Init archive buffer
	archiveBuffer, err := os.Create(outputPath)
	if err != nil {
		return errors.New("ERROR : " + err.Error())
	}
	defer archiveBuffer.Close()

	// Create archive
	archiveWriter := zip.NewWriter(archiveBuffer)

	// Add files to archive
	err = GenerateArchive(archiveWriter, basePath+"/", "")
	if err != nil {
		return errors.New("ERROR : " + err.Error())
	}

	err = archiveWriter.Close()
	if err != nil {
		return errors.New("ERROR : " + err.Error())
	}

	log.Println("INFO : Compression " + basePath + " to " + outputPath + " complete")
	return nil
}

func GenerateArchive(archiveWriter *zip.Writer, folderPath, subDir string) error {

	// Read dir
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return errors.New("ERROR : " + err.Error())
	}

	for _, file := range files {

		if !file.IsDir() {
			filePath := folderPath + file.Name()
			data, err := ioutil.ReadFile(filePath)
			if err != nil {
				return errors.New("ERROR : " + err.Error())
			}

			// Add file to archive
			f, err := archiveWriter.Create(subDir + file.Name())
			if err != nil {
				return errors.New("ERROR : " + err.Error())
			}
			_, err = f.Write(data)
			if err != nil {
				return errors.New("ERROR : " + err.Error())
			}

		} else if file.IsDir() {
			// Recursive

			// Update base folder path
			newFolderPath := folderPath + "/" + file.Name() + "/"

			// Update sub directory name
			newSubDir := subDir + file.Name() + "/"

			GenerateArchive(archiveWriter, newFolderPath, newSubDir)
		}
	}

	return nil
}
