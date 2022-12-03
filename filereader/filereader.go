package filereader

import (
	"fmt"
	"os"
)

func Reader(test bool, day int16) (string, error) {
	var filePath string = ""
	if test {
		filePath += "tests/"
	}
	dat, err := os.ReadFile(fmt.Sprintf("%sday%d.txt", filePath, day))
	if err != nil {
		fmt.Println("could not read file", err)
		return "", err
	}
	return string(dat), nil

}
