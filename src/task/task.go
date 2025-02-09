package task

import (
	"fmt"
	"log"
	"sync/atomic"
	"time"

	input "taskManager/src/input"
	menu "taskManager/src/menu"
)

var globalID int64

// Тип для исполнителя задачи.
type Executor struct {
	Name string `json:"name"`
}

// Тип для дат задачи.
type Date struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

// Экспортируемая структура задачи.
type Task struct {
	ID              int64    `json:"id"`
	HeadingTask     string   `json:"heading_task"`
	DescriptionTask string   `json:"description_task"`
	DeadLineTask    Date     `json:"dead_line_task"`
	ExecutorTask    Executor `json:"executor_task"`
	CompleteTask    bool     `json:"complete_task"`
}

func (task *Task) SetCompletedTask(answer bool) {
	task.CompleteTask = answer
}

func (task *Task) SetDescriptionTask(description string) {
	task.DescriptionTask = description
}

func (task *Task) SetHeadingTask(title string) {
	task.HeadingTask = title
}

func (task *Task) SetDeadlineTask(startDate time.Time, endDate time.Time) {
	task.DeadLineTask.StartDate = startDate
	task.DeadLineTask.EndDate = endDate
}

func (task *Task) SetExecutorTask(executor string) {
	task.ExecutorTask.Name = executor
}

func (task *Task) createIdTask() {
	task.ID = atomic.AddInt64(&globalID, 1)
}

func (task *Task) AddTask(input input.ConsoleInput) {
	fmt.Println("Введите название задачи:")
	task.SetHeadingTask(input.EnterString())

	fmt.Println("Введите описание задачи:")
	task.SetDescriptionTask(input.EnterString())

	fmt.Println("Введите дату начала задачи в формате YYYY-MM-DD:")
	startDate, err := input.EnterDate()
	if err != nil {
		log.Println("Ошибка ввода даты начала:", err)
	}

	fmt.Println("Введите дату завершения задачи в формате YYYY-MM-DD:")
	endDate, err := input.EnterDate()
	if err != nil {
		log.Println("Ошибка ввода даты завершения:", err)
	}
	task.SetDeadlineTask(startDate, endDate)

	fmt.Println("Введите работника, отвечающего за данную задачу:")
	task.SetExecutorTask(input.EnterString())

	task.createIdTask()
	fmt.Println("Создана задача с ID:", task.ID)

	task.CompleteTask = false
}

func (task *Task) ViewingTask() {
	fmt.Println("========================================")
	fmt.Printf("ID задачи: %d\n", task.ID)
	fmt.Println("Название:", task.HeadingTask)
	fmt.Println("Описание:", task.DescriptionTask)
	fmt.Printf("Дата начала: %s, Дата завершения: %s\n",
		task.DeadLineTask.StartDate.Format("2006-01-02"),
		task.DeadLineTask.EndDate.Format("2006-01-02"))
	fmt.Println("Ответственный:", task.ExecutorTask.Name)
	if task.CompleteTask {
		fmt.Println("Статус: Выполнено")
	} else {
		fmt.Println("Статус: Не выполнено")
	}
	fmt.Println("========================================")
}

func (task *Task) RemoveTask(taskSlice *[]Task, input input.ConsoleInput) {
	fmt.Println("Введите ID удаляемой задачи:")
	idStr := input.EnterString()
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

func (task *Task) PatchTask(taskSlice *[]Task, input input.ConsoleInput) {
	fmt.Println("Введите ID изменяемой задачи:")
	idStr := input.EnterString()
	var id int64
	_, err := fmt.Sscan(idStr, &id)
	if err != nil {
		log.Fatal(err)
	}

	index := -1
	for i, t := range *taskSlice {
		if t.ID == id {
			index = i
		}
	}
	if index == -1 {
		fmt.Println("Данный ID не найден в системе")
		return
	}

	tempTask := (*taskSlice)[index]

	for {
		expr := menu.MenuPatchPrint(input)
		switch expr {
		case 1:
			fmt.Println("Введите новый заголовок задачи")
			tempTask.SetHeadingTask(input.EnterString())
		case 2:
			fmt.Println("Введите новое описание задачи")
			tempTask.SetDescriptionTask(input.EnterString())
		case 3:
			fmt.Println("Введите новый день начала задачи")
			startDate, err := input.EnterDate()
			if err != nil {
				log.Fatal("Ошибка ввода даты начала:", err)
			}

			fmt.Println("Введите новый дедлайн задачи")
			endDate, err := input.EnterDate()
			if err != nil {
				log.Fatal("Ошибка ввода даты завершения", err)
			}
			tempTask.SetDeadlineTask(startDate, endDate)
		case 4:
			fmt.Println("Введите нового исполнителя задачи")
			tempTask.SetExecutorTask(input.EnterString())
		case 5:
			fmt.Println("Выход из редактирования задачи...")
			(*taskSlice)[index] = tempTask
			return
		default:
			fmt.Println("Некорректный выбор. Пожалуйста, попробуйте снова.")
		}
	}
}
