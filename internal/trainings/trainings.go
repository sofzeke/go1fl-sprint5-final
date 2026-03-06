package trainings

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps int // количество шагов, проделанных за тренировку
	TrainingType string // тип тренировки(бег или ходьба)
	Duration time.Duration // длительность тренировки
	personaldata.Personal // встроенная структура Personal из пакета personaldata, у которой есть метод Print()
}

func (t *Training) Parse(datastring string) (err error) {
	stepActivityDuration := strings.Split(datastring, ",")

	// Проверить, чтобы длина слайса была равна 3
	if len(stepActivityDuration) != 3 {
		return errors.New("the slice length is not equal to 3")
	}

	t.TrainingType = stepActivityDuration[1]

		steps, err := strconv.Atoi(stepActivityDuration[0])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}

	if steps <= 0 {
		return errors.New("steps not equal 0")
	}

	t.Steps = steps

	// Преобразовать третий элемент слайса в time.Duration.
	duration, err := time.ParseDuration(stepActivityDuration[2])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}

	if duration <= 0 {
		return errors.New("duration <= 0")
	}

	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// Вычислить дистанцию, используя функцию из пакета spentenergy
	distance:=spentenergy.Distance(t.Steps,t.Height)
	// Вычислить среднюю скорость, используя функцию из пакета spentenergy
	averageSpeed:=spentenergy.MeanSpeed(t.Steps,t.Height,t.Duration)
	caloriesBurned:= 0.0
	// Проверить, какой вид тренировки содержится в структуре Training
	switch t.TrainingType {
	case "Ходьба":
		calories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		caloriesBurned = calories
	case "Бег":
		calories, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		caloriesBurned = calories
	}
	// Сформируйте и верните строку по образцу
	str := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, averageSpeed, caloriesBurned)
	// Если был передан неизвестный тип тренировки, верните ошибку
	if t.TrainingType != "Ходьба" && t.TrainingType != "Бег" {
	errActivity := errors.New("неизвестный тип тренировки")
	log.Println(errActivity)
	return "", errActivity
	}

	return str, nil

}
