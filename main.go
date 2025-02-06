package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync/atomic"
	"time"
)

var globalID int64

type Excutor struct {
	nameExcutor string
}

type Date struct {
	startDate time.Time
	endDate   time.Time
}

type Task struct {
	ID              int64
	headingTask     string
	descriptionTask string
	deadLineTask    Date
	executorTask    Excutor
	completeTask    bool
}

type ConsoleInput struct {
	reader *bufio.Reader
}

func NewConsoleInput() ConsoleInput {
	return ConsoleInput{reader: bufio.NewReader(os.Stdin)}
}

// enterString считывает строку ввода и удаляет лишние символы.
func (ci ConsoleInput) enterString() string {
	input, err := ci.reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	// Убираем символы переноса строки и возврата каретки
	return strings.TrimSpace(input)
}

// enterDate считывает дату в формате YYYY-MM-DD.
func (ci ConsoleInput) enterDate() (time.Time, error) {
	dateStr := ci.enterString()
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

func (task *Task) createIdTask() {
	task.ID = atomic.AddInt64(&globalID, 1)
}

func (task *Task) addTask(input ConsoleInput) {
	fmt.Println("Введите название задачи:")
	task.headingTask = input.enterString()

	fmt.Println("Введите описание задачи:")
	task.descriptionTask = input.enterString()

	fmt.Println("Введите дату начала задачи в формате YYYY-MM-DD:")
	startDate, err := input.enterDate()
	if err != nil {
		log.Fatal("Ошибка ввода даты начала:", err)
	}
	task.deadLineTask.startDate = startDate

	fmt.Println("Введите дату завершения задачи в формате YYYY-MM-DD:")
	endDate, err := input.enterDate()
	if err != nil {
		log.Fatal("Ошибка ввода даты завершения:", err)
	}
	task.deadLineTask.endDate = endDate

	fmt.Println("Введите работника, отвечающего за данную задачу:")
	task.executorTask.nameExcutor = input.enterString()

	task.createIdTask()
	fmt.Println("Создана задача с ID:", task.ID)

	task.completeTask = false
}

func (task *Task) viewingTask() {
	fmt.Println("========================================")
	fmt.Printf("ID задачи: %d\n", task.ID)
	fmt.Println("Название:", task.headingTask)
	fmt.Println("Описание:", task.descriptionTask)
	fmt.Printf("Дата начала: %s, Дата завершения: %s\n",
		task.deadLineTask.startDate.Format("2006-01-02"),
		task.deadLineTask.endDate.Format("2006-01-02"))
	fmt.Println("Ответственный:", task.executorTask.nameExcutor)
	if task.completeTask {
		fmt.Println("Статус: Выполнено")
	} else {
		fmt.Println("Статус: Не выполнено")
	}
	fmt.Println("========================================")
}

func (task *Task) removeTask(taskSlice *[]Task, input ConsoleInput) {
	fmt.Println("Введите ID удаляемой задачи:")
	idStr := input.enterString()
	var id int64
	_, err := fmt.Sscan(idStr, &id)
	if err != nil {
		log.Fatal(err)
	}

	index := -1
	for i, t := range *taskSlice {
		if t.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Printf("Задача с ID %d не найдена.\n", id)
		return
	}

	*taskSlice = append((*taskSlice)[:index], (*taskSlice)[index+1:]...)
	fmt.Printf("Задача с ID %d удалена.\n", id)
}

func menu(input ConsoleInput) int {
	fmt.Println("\nЗдравствуйте! Выберите действие:")
	fmt.Println("1 - Добавить задачу")
	fmt.Println("2 - Просмотреть все задачи")
	fmt.Println("3 - Отметить задачу как выполненную")
	fmt.Println("4 - Удалить задачу")
	fmt.Println("5 - Выход")
	fmt.Print("Ваш выбор: ")

	choiceStr := input.enterString()
	var choice int
	_, err := fmt.Sscan(choiceStr, &choice)
	if err != nil {
		log.Println("Ошибка ввода:", err)
		return 0
	}
	return choice
}

func main() {
	var tasks []Task
	input := NewConsoleInput()

	for {
		choice := menu(input)
		switch choice {
		case 1:
			var newTask Task
			newTask.addTask(input)
			tasks = append(tasks, newTask)
		case 2:
			if len(tasks) == 0 {
				fmt.Println("Список задач пуст.")
			} else {
				for i := range tasks {
					tasks[i].viewingTask()
				}
			}
		case 3:
			if len(tasks) == 0 {
				fmt.Println("Список задач пуст.")
			} else {
				fmt.Println("Введите ID задачи, которую нужно отметить как выполненную:")
				idStr := input.enterString()
				var id int64
				_, err := fmt.Sscan(idStr, &id)
				if err != nil {
					log.Println("Ошибка ввода:", err)
					continue
				}
				found := false
				for i := range tasks {
					if tasks[i].ID == id {
						tasks[i].completeTask = true
						fmt.Printf("Задача с ID %d отмечена как выполненная.\n", id)
						found = true
						break
					}
				}
				if !found {
					fmt.Printf("Задача с ID %d не найдена.\n", id)
				}
			}
		case 4:
			if len(tasks) == 0 {
				fmt.Println("Список задач пуст.")
			} else {
				tasks[0].removeTask(&tasks, input)
			}
		case 5:
			fmt.Println("Выход из программы...")
			return
		default:
			fmt.Println("Некорректный выбор. Пожалуйста, попробуйте снова.")
		}
	}
}
