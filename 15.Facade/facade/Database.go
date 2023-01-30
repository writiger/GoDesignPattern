package facade

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type database struct {
}

func (d database) GetProperty(email string) (string, error) {
	path := "database.txt"
	file, err := os.OpenFile(path, os.O_RDONLY, 0666)

	if err != nil {
		return "", err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			fmt.Println("File Read Over.")
			break
		}
		info := strings.Split(string(line), "=")
		if email == info[0] {
			return info[1], nil
		}
	}
	return "", errors.New("查询邮箱为空")
}
