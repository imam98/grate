package xls

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestManyFiles(t *testing.T) {
	err := filepath.Walk("../testdata", func(p string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(info.Name(), ".xls") {
			return nil
		}
		wb, err := Open(p)
		if err != nil {
			return nil //err
		}

		sheets, err := wb.List()
		if err != nil {
			return err
		}
		for _, s := range sheets {
			//log.Println(s)
			sheet, err := wb.Get(s)
			if err != nil {
				return err
			}

			for sheet.Next() {
				sheet.Strings()
			}
		}

		return wb.Close()
	})
	if err != nil {
		t.Fatal(err)
	}
}
