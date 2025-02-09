package files

import (
	"encoding/json"
	"log"
	"os"
	"taskManager/src/task"
)

func SaveTasksToFile(filename string, tasks []task.Task) {
	data, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		log.Fatal("Ошибка сериализации задач:", err)
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatal("Ошибка записи в файл:", err)
	}
}

func LoadTasksFromFile(filename string) []task.Task {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Ошибка чтения файла:", err)
	}
	var tasks []task.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Fatal("Ошибка десериализации JSON:", err)
	}
	return tasks
}
