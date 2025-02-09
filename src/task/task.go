package task

import (
	"fmt"
	"log"
	"sync/atomic"
	input "taskManager/src/input"
	"taskManager/src/menu"
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

func (task *Task) SetCompletedTask(answer bool) {
	task.completeTask = answer
}

func (task *Task) SetDescriptionTask(description string) {
	task.descriptionTask = description
}

func (task *Task) SetHeadingTask(title string) {
	task.headingTask = title
}

func (task *Task) SetDeadlineTask(startDate time.Time, endDate time.Time) {
	task.deadLineTask.startDate = startDate
	task.deadLineTask.endDate = endDate
}

func (task *Task) SetExecutorTask(executor string) {
	task.executorTask.nameExcutor = executor
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
		log.Println("Ошибка ввода")
	}

	fmt.Println("Введите дату завершения задачи в формате YYYY-MM-DD:")
	endDate, err := input.EnterDate()
	if err != nil {
		log.Println("Ошибка ввода")
	}
	task.SetDeadlineTask(startDate, endDate)

	fmt.Println("Введите работника, отвечающего за данную задачу:")
	task.SetExecutorTask(input.EnterString())

	task.createIdTask()
	fmt.Println("Создана задача с ID:", task.ID)

	task.completeTask = false
}

func (task *Task) ViewingTask() {
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
	}

	tempTask := (*taskSlice)[index]

	for {
		expr := menu.MenuPatchPrint(input)
		switch expr {
		case 1:
			fmt.Println("Введите новый загаловок задачи")
			tempTask.SetHeadingTask(input.EnterString())
		case 2:
			fmt.Println("Введите новое описание задачи")
			tempTask.SetDescriptionTask(input.EnterString())
		case 3:
			fmt.Println("Введите новый день начала таски")
			startDate, err := input.EnterDate()
			if err != nil {
				log.Fatal("Ошибка ввода даты начала:", err)
			}

			fmt.Println("Введите дедлайн таски")
			endDate, err := input.EnterDate()
			if err != nil {
				log.Fatal("Ошибка ввода даты конца", err)
			}
			tempTask.SetDeadlineTask(startDate, endDate)
		case 4:
			fmt.Println("Введите нового исполнителя таски")
			tempTask.SetExecutorTask(input.EnterString())
		case 5:
			fmt.Println("Выход из программы...")
			(*taskSlice)[index] = tempTask
			return
		default:
			fmt.Println("Некорректный выбор. Пожалуйста, попробуйте снова.")
		}
	}

}
