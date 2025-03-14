package article

import (
	"fmt"
	"time"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

type ArticleRepository interface {
	Merge(file1 string, file2 string) error
	Split(file string, pagePerSplit int) error
	Compressed(file string) error
}

type Service struct {}

func NewService() *Service {
	return &Service{}
}


// Merge godoc
// @Summary Merge two PDF files
// @Description Merge two PDF files into one output PDF
// @Tags pdf
// @Accept  json
// @Produce  json
// @Param file1 query string true "First PDF File"
// @Param file2 query string true "Second PDF File"
// @Success 200 {string} string "Merged PDF file created"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /merge [post]
func (a *Service) Merge(file1 string, file2 string) error {
	dir, _:= os.Getwd()

	inputFiles := []string{file1, file2}
	timestamp := time.Now().Format("02-01-06-15-04-05")
	outputFile := dir + "/pdf/output/" + fmt.Sprintf("merged-%s.pdf", timestamp)

	err := api.MergeCreateFile(inputFiles, outputFile, false, nil)
	if err != nil {
		return err
	}
	return nil
}

// Split godoc
// @Summary Split a PDF into smaller parts
// @Description Split a PDF file into smaller PDF files with a defined number of pages per split
// @Tags pdf
// @Accept  json
// @Produce  json
// @Param file query string true "PDF file to split"
// @Param pagesPerSplit query int true "Pages per split"
// @Success 200 {string} string "PDF file split"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /split [post]
func (a *Service) Split(file string, pagePerSplit int) error {
	conf := model.NewDefaultConfiguration()
	dir, _:= os.Getwd()

	err := api.SplitFile(file, dir + "/pdf/output/", pagePerSplit, conf)
	if err != nil {
		return err
	}
	return nil
}

// Compressed godoc
// @Summary Compress a PDF file
// @Description Compress a PDF file to reduce its size
// @Tags pdf
// @Accept  json
// @Produce  json
// @Param file query string true "PDF file to compress"
// @Success 200 {string} string "Compressed PDF file created"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /compress [post]
func (a *Service) Compressed(file string) error {
	dir, _:= os.Getwd()

	timestamp := time.Now().Format("02-01-06-15-04-05")
	outputFile := dir + "/pdf/output/" + fmt.Sprintf("compressed-%s.pdf", timestamp)

	err := api.OptimizeFile(file, outputFile, nil)
	if err != nil {
		return err
	}
	return nil
}