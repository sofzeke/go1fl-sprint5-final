package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps int // количество шагов
	Duration time.Duration // длительность прогулки
	personaldata.Personal // встроенная структура Personal из пакета personaldata, у которой есть метод Print()
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// Проверить пустую строку
	if datastring == "" {
		return errors.New("less string")
	}

	duratironStep := strings.Split(datastring, ",")

	if len(duratironStep) != 2 {
		return errors.New("incorrect format")
	}
	// Проверить на преобразование
	steps, err := strconv.Atoi(duratironStep[0])
	if err != nil {
		return fmt.Errorf("transformation error: %w", err)
	}
	// Проверка на отрицательное или нулево значени
	if steps <= 0 {
		return errors.New("steps is less than or equal to zero")
	}

	ds.Steps = steps

	duration, err := time.ParseDuration(duratironStep[1])
	if err != nil {
		return fmt.Errorf("transformation error: %w", err)
	}

	if duration <= 0 {
		return errors.New("distance is less than or equal to zero")
	}

	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// Вычислите дистанцию
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	// Вычислите количество сожжённых калорий. При возникновении ошибки верните пустую строку и ошибку
	caloriesBurned:= 0.0

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	caloriesBurned = calories

	str := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, caloriesBurned)
	return str, nil
}

