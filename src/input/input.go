package input

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

type ConsoleInput struct {
	reader *bufio.Reader
}

func NewConsoleInput() ConsoleInput {
	return ConsoleInput{reader: bufio.NewReader(os.Stdin)}
}

// enterString считывает строку ввода и удаляет лишние символы.
func (ci ConsoleInput) EnterString() string {
	input, err := ci.reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	// Убираем символы переноса строки и возврата каретки
	return strings.TrimSpace(input)
}

// enterDate считывает дату в формате YYYY-MM-DD.
func (ci ConsoleInput) EnterDate() (time.Time, error) {
	dateStr := ci.EnterString()
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}
