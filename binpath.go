package binpath

import (
	"os"
	"path/filepath"
)

var bindir binDirectory

func init() {
	bin, err := filepath.Abs(os.Args[0])
	if err != nil {
		panic(err)
	}

	bindir = binDirectory(filepath.Dir(bin))
}

type binDirectory string

func (bd binDirectory) join(elements []string) string {
	elements = append([]string{string(bd)}, elements...)
	return filepath.Join(elements...)
}

func Join(elem ...string) string {
	return bindir.join(elem)
}
