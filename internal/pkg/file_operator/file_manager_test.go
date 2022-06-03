package file_operator_test

import (
	"github.com/hiteshpattanayak-tw/url_shortner/internal/pkg/file_operator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
)

func TestFileOperator_Create(t *testing.T) {
	wd, _ := os.Getwd()
	file1Path := path.Join(wd, "../../../", "test/files", "file1.txt")

	fo := file_operator.ProvideFileOperator(file1Path)
	err := fo.Create()
	require.NoError(t, err)

	err = fo.Close()
	require.NoError(t, err)
}

func TestFileOperator_Write(t *testing.T) {
	wd, _ := os.Getwd()
	file1Path := path.Join(wd, "../../../", "test/files", "file1.txt")

	fo := file_operator.ProvideFileOperator(file1Path)
	err := fo.Create()
	require.NoError(t, err)

	size, err := fo.Write("Hello world")
	require.NoError(t, err)
	assert.True(t, size > 0)

	err = fo.Close()
	require.NoError(t, err)
}

func TestFileOperator_Read(t *testing.T) {
	wd, _ := os.Getwd()
	file1Path := path.Join(wd, "../../../", "test/files", "file1.txt")

	fo := file_operator.ProvideFileOperator(file1Path)
	err := fo.Create()
	require.NoError(t, err)

	writeData := "Hello world"
	size, err := fo.Write(writeData)
	require.NoError(t, err)
	assert.True(t, size > 0)

	readData, err := fo.Read()
	require.NoError(t, err)
	assert.Equal(t, writeData, readData)

	err = fo.Close()
	require.NoError(t, err)
}
