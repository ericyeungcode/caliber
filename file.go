package caliber

import (
	"os"
	"path"
	"strings"
	"time"
)

func EnsureDir(fpath string) error {
	baseDir := path.Dir(fpath)
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		return os.MkdirAll(baseDir, 0700) // Create your file
	}
	return nil
}

func WriteFileEnsureDir(fpath string, data []byte) error {
	err := EnsureDir(fpath)
	if err != nil {
		return err
	}
	return os.WriteFile(fpath, data, 0666)
}

func GetMegaBytes(v float64) float64 {
	return v / 1e6
}

func GetDirFileList(aDir string) ([]string, error) {
	entries, err := os.ReadDir(aDir)
	if err != nil {
		return nil, err
	}

	list := make([]string, 0, len(entries))
	for _, v := range entries {
		if strings.HasSuffix(v.Name(), ".csv") {
			list = append(list, v.Name())
		}
	}

	return list, nil
}

func GetFileSize(fileName string) (int64, error) {
	stat, err := os.Stat(fileName)
	if err != nil {
		return 0, err
	}
	return stat.Size(), nil
}

func RemoveWithRetry(fileName string) error {
	var err error
	for i := 0; i < 5; i++ {
		err = os.Remove(fileName)
		if err == nil {
			return nil
		}
		time.Sleep(time.Millisecond * 10)
	}

	return err
}
