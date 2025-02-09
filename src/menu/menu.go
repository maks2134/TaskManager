package menu

import (
	"fmt"
	"log"
	"taskManager/src/input"
)

func MenuPrint(input input.ConsoleInput) int {
	fmt.Println("\nЗдравствуйте! Выберите действие:")
	fmt.Println("1 - Добавить задачу")
	fmt.Println("2 - Просмотреть все задачи")
	fmt.Println("3 - Отметить задачу как выполненную")
	fmt.Println("4 - Удалить задачу")
	fmt.Println("5 - Изменить задачу по ID")
	fmt.Println("6 - Выход")
	fmt.Print("Ваш выбор: ")

	choiceStr := input.EnterString()
	var choice int
	_, err := fmt.Sscan(choiceStr, &choice)
	if err != nil {
		log.Println("Ошибка ввода:", err)
		return 0
	}
	return choice
}

func MenuPatchPrint(input input.ConsoleInput) int {
	fmt.Println("\nЗдравствуйте! Выберите, что изменить:")
	fmt.Println("1 - Изменить название")
	fmt.Println("2 - Изменить описание")
	fmt.Println("3 - Изменить дедлайн")
	fmt.Println("4 - Изменить исполнителя")
	fmt.Println("5 - Выход")
	fmt.Print("Ваш выбор: ")

	choiceStr := input.EnterString()
	var choice int
	_, err := fmt.Sscan(choiceStr, &choice)
	if err != nil {
		log.Println("Ошибка ввода:", err)
		return 0
	}
	return choice
}
