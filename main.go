package main

import (
	"awesomeProject/file"
	"awesomeProject/pipeline"
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	// Уже отсортированное
	jobNames := []string{"extra_volume_getter", "wrong_sum_volume_getter", "volume_validator", "coef_validator"}

	// Спаршенный файл
	//myFile := getValidFile()
	myFile := getInvalidFile()

	newPipeline := pipeline.NewPipeline(myFile, jobNames)

	err := newPipeline.StartJobs()
	if err != nil {
		panic(err)
	}

	for i, row := range myFile.FileRows {
		for _, rowErr := range row.ErrorValue {
			fmt.Println(fmt.Sprintf("[%d] Row error: %s", i, rowErr))
		}

		for _, coefComment := range row.Coefficient.ErrorComment {
			fmt.Println(fmt.Sprintf("[%d] Coef comment: %s", i, coefComment))
		}

		for _, nameComment := range row.Name.ErrorComment {
			fmt.Println(fmt.Sprintf("[%d] Name comment: %s", i, nameComment))
		}

		for _, volumeComment := range row.Volume.ErrorInfo.ErrorComment {
			fmt.Println(fmt.Sprintf("[%d] Volume comment: %s", i, volumeComment))
		}
	}
}

func getValidFile() *file.ParsedFile {
	firstRow := file.Row{
		Mutex: sync.Mutex{},
		ID:    0,
		Coefficient: file.CoefficientColumn{
			ID:        1,
			Value:     15,
			ErrorInfo: file.ErrorInfo{},
		},
		Name: file.NameColumn{
			ID:        2,
			Value:     "mega promo",
			ErrorInfo: file.ErrorInfo{},
		},
		Volume: file.VolumeColumn{
			ID:        3,
			Value:     101,
			ErrorInfo: file.ErrorInfo{},
		},
		ErrorValue: []string{},
	}

	secondRow := file.Row{
		Mutex: sync.Mutex{},
		ID:    1,
		Coefficient: file.CoefficientColumn{
			ID:        1,
			Value:     rand.Intn(5),
			ErrorInfo: file.ErrorInfo{},
		},
		Name: file.NameColumn{
			ID:        2,
			Value:     "mega promo",
			ErrorInfo: file.ErrorInfo{},
		},
		Volume: file.VolumeColumn{
			ID:        3,
			Value:     11,
			ErrorInfo: file.ErrorInfo{},
		},
		ErrorValue: []string{},
	}

	return &file.ParsedFile{
		FileRows: []file.Row{firstRow, secondRow},
	}
}

func getInvalidFile() *file.ParsedFile {
	fileRows := make([]file.Row, 0, 100)
	for i := 0; i < 100; i++ {
		row := file.Row{
			Mutex: sync.Mutex{},
			ID:    0,
			Coefficient: file.CoefficientColumn{
				ID:        1,
				Value:     rand.Intn(5),
				ErrorInfo: file.ErrorInfo{},
			},
			Name: file.NameColumn{
				ID:        2,
				Value:     "mega promo",
				ErrorInfo: file.ErrorInfo{},
			},
			Volume: file.VolumeColumn{
				ID:        3,
				Value:     10 + int64(i),
				ErrorInfo: file.ErrorInfo{},
			},
			ErrorValue: []string{},
		}
		fileRows = append(fileRows, row)
	}

	return &file.ParsedFile{
		FileRows: fileRows,
	}
}
