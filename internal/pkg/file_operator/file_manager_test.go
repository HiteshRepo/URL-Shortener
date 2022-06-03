package file_operator_test

import (
	"github.com/hiteshpattanayak-tw/url_shortner/internal/pkg/file_operator"
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
