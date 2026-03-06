package actioninfo

import "log"


type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		// Парсим текущее значение с помощью метода Parse()
		err := dp.Parse(data)
		if err != nil {
			// При ошибке логируем и переходим к следующей итерации
			log.Printf("Error parsing '%s': %v", data, err)
			continue
	}
	// Получаем строку с информацией об активности с помощью метода ActionInfo()
		info, err := dp.ActionInfo()
		if err != nil {
			// При ошибке получения информации логируем её
			log.Printf("Error info: %v", err)
			continue
	}

		// Выводим строку с информацией об активности
		log.Println(info)
	}
}