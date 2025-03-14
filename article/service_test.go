package article

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"
)

// // Mock for the api.MergeCreateFile, api.SplitFile, and api.OptimizeFile functions
// type MockAPI struct {
// 	mock.Mock
// }

// func (m *MockAPI) MergeCreateFile(inputFiles []string, outputFile string, overwrite bool, conf interface{}) error {
// 	args := m.Called(inputFiles, outputFile, overwrite, conf)
// 	return args.Error(0)
// }

// func (m *MockAPI) SplitFile(file string, outputDir string, pagePerSplit int, conf interface{}) error {
// 	args := m.Called(file, outputDir, pagePerSplit, conf)
// 	return args.Error(0)
// }

// func (m *MockAPI) OptimizeFile(inputFile string, outputFile string, conf interface{}) error {
// 	args := m.Called(inputFile, outputFile, conf)
// 	return args.Error(0)
// }

func TestService_Merge(t *testing.T) {
	service := NewService()
	err := service.Merge("../pdf/input/file1.pdf", "../pdf/input/file2.pdf")
	assert.NoError(t, err)
}

func TestService_Split(t *testing.T) {
	service := NewService()
	err := service.Split("../pdf/input/2_pageinput.pdf", 1)
	assert.NoError(t, err)
}

func TestService_Compressed(t *testing.T) {
	service := NewService()
	err := service.Compressed("../pdf/input/file1.pdf")
	assert.NoError(t, err)
}
