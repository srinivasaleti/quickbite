package bloomfilters

import (
	"compress/gzip"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper function to write coupons to a gzip file
func writeGzipFile(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := gzip.NewWriter(file)
	defer writer.Close()

	for _, line := range lines {
		_, err := writer.Write([]byte(line + "\n"))
		if err != nil {
			return err
		}
	}
	return nil
}

func TestGzipFileToBloomFilter(t *testing.T) {
	tempDir := t.TempDir()

	file := "large_file.tz"
	var coupons []string
	for i := 1; i <= 1000000; i++ {
		coupons = append(coupons, fmt.Sprintf("COUPON%d", i))
	}

	filePath := filepath.Join(tempDir, file)
	err := writeGzipFile(filePath, coupons)
	assert.NoError(t, err)

	filter, err := fileToBloomFilter(filePath)
	assert.NoError(t, err)

	assert.True(t, filter.Test([]byte("COUPON1")))
	assert.True(t, filter.Test([]byte("COUPON50000")))
	assert.True(t, filter.Test([]byte("COUPON99999")))
	assert.True(t, filter.Test([]byte("COUPON123456")))

	assert.False(t, filter.Test([]byte("INVALID-COUPON")))
}

func TestBloomFiltersWithMultipleFiles(t *testing.T) {
	tempDir := t.TempDir()

	fileCoupons := map[string][]string{
		"file1.tz": {"C1", "C2", "C3", "C4", "C5", "C6", "C7", "C8", "C9", "C10"},
		"file2.tz": {"C5", "C12", "C13", "C14", "C15", "C16", "C17", "C18", "C19", "C20"},
		"file3.tz": {"C5", "C22", "C23", "C24", "C25", "C26", "C27", "C28", "C29", "C30"},
	}

	var filePaths []string

	for fname, coupons := range fileCoupons {
		path := filepath.Join(tempDir, fname)
		err := writeGzipFile(path, coupons)
		assert.NoError(t, err)
		filePaths = append(filePaths, path)
	}

	filter := NewGzipBloomFilter()
	err := filter.Load(filePaths)
	assert.NoError(t, err)

	// Get full paths
	file1 := filepath.Join(tempDir, "file1.tz")
	file2 := filepath.Join(tempDir, "file2.tz")
	file3 := filepath.Join(tempDir, "file3.tz")

	assert.ElementsMatch(t, filter.ElmentExistsInWhichFiles("C1"), []string{file1})
	assert.ElementsMatch(t, filter.ElmentExistsInWhichFiles("C15"), []string{file2})
	assert.ElementsMatch(t, filter.ElmentExistsInWhichFiles("C28"), []string{file3})
	assert.ElementsMatch(t, filter.ElmentExistsInWhichFiles("C5"), []string{file1, file2, file3})
	assert.Equal(t, len(filter.ElmentExistsInWhichFiles("C5")), 3)
	assert.Empty(t, filter.ElmentExistsInWhichFiles("INVALID-COUPON"))
}
