package main

import (
	"fmt"
	"log"
	input "taskManager/src/input"
	menu "taskManager/src/menu"
	task "taskManager/src/task"
)

func main() {
	var tasks []task.Task
	input := input.NewConsoleInput()

	for {
		choice := menu.MenuPrint(input)
		switch choice {
		case 1:
			var newTask task.Task
			newTask.AddTask(input)
			tasks = append(tasks, newTask)
		case 2:
			if len(tasks) == 0 {
				fmt.Println("Список задач пуст.")
			} else {
				for i := range tasks {
					tasks[i].ViewingTask()
				}
			}
		case 3:
			if len(tasks) == 0 {
				fmt.Println("Список задач пуст.")
			} else {
				fmt.Println("Введите ID задачи, которую нужно отметить как выполненную:")
				idStr := input.EnterString()
				var id int64
				_, err := fmt.Sscan(idStr, &id)
				if err != nil {
					log.Println("Ошибка ввода:", err)
					continue
				}
				found := false
				for i := range tasks {
					if tasks[i].ID == id {
						tasks[i].SetCompletedTask(true)
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
				tasks[0].RemoveTask(&tasks, input)
			}
		case 5:
			if len(tasks) == 0 {
				fmt.Println("Список задач пуст.")
			} else {
				tasks[0].PatchTask(&tasks, input)
			}
		case 6:
			fmt.Println("Выход из программы...")
			return
		default:
			fmt.Println("Некорректный выбор. Пожалуйста, попробуйте снова.")
		}
	}
}
