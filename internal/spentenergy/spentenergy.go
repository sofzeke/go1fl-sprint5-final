package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	//Проверить входные параметры на корректность
	if steps <0 || weight <=0 || height <=0 || duration <=0 {
		return 0, errors.New("invalid input parameters: steps, weight, height, and duration")
	} 
	//Рассчитать среднюю скорость с помощью meanSpeed()
	averageSpeed:=MeanSpeed(steps,height,duration)	
	// Если средняя скорость равна 0 
    if averageSpeed == 0 {
        return 0, nil
    }
	// Перевод продолжительности в минуты
	durationInMinutes:=duration.Minutes()
	//Рассчитать количество калорий
	caloriesSpent:=(weight * averageSpeed * durationInMinutes) / minInH
	 // Умножить на корректирующий коэффициент для ходьбы
    caloriesBurned := caloriesSpent * walkingCaloriesCoefficient

	return caloriesBurned,nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	//Проверить входные параметры на корректность
	if steps <0 || weight <=0 || height <=0 || duration <=0 {
		return 0, errors.New("invalid input parameters: steps, weight, height, and duration")
	} 
	//Рассчитать среднюю скорость с помощью meanSpeed()
	averageSpeed:=MeanSpeed(steps,height,duration)
	// Если средняя скорость равна 0 
    if averageSpeed == 0 {
        return 0, nil
    }
	// Перевод продолжительности в минуты
	durationInMinutes:=duration.Minutes()
	//Рассчитать количество калорий
	caloriesBurned:=(weight * averageSpeed * durationInMinutes) / minInH 

	return caloriesBurned,nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	//добавить проверки отрицательных шагов (должно возвращать 0)
	if steps < 0 {
		return 0
	}
	//Проверить, что продолжительность duration больше 0. Если это не так, вернуть 0.
	if duration <= 0 {
		return 0
	}
	//Вычислить дистанцию с помощью Distance()
	distance:= Distance(steps,height)
	// Перевод продолжительности в часы
	durationInHours := duration.Hours()
	//Вычислить и вернуть среднюю скорость
	averageSpeed := distance / durationInHours

	return averageSpeed
}

func Distance(steps int, height float64) float64 {
	//рассчитываем длину шага
	lengthStep := height*stepLengthCoefficient
	 // Умножаем количество шагов на длину шага 
	 totalDistanceMeters := float64(steps) * lengthStep
	 // Переводим метры в километры: делим на число метров в километре
    distanceKm := totalDistanceMeters / mInKm

	return distanceKm
}
