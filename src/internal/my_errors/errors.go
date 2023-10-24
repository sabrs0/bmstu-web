package my_errors

import "errors"

const (
	Err                     = "ОШИБКА\n"
	ErrNotExists            = Err + "не найдено"
	ErrAlreadyExists        = Err + "Конфликтная ситуация. Возможно сущность с введенными данными уже существует.\nвведите другие данные"
	ErrCountry              = Err + "введенная страна не поддерживается"
	ErrMoney                = Err + "неверный формат ввода денежной суммы.\n формат: хххххх.хх"
	ErrNotExistsFoundrising = Err + "У данного фонда нет сборов"
	//SupportedCountries = "Поддерживаемы страны:\nРоссия, США, Великобритания, Канада, Франция, Германия, Китай, Италия, Испания"
)

var ErrorNotExists = errors.New(ErrNotExists)
var ErrorConflict = errors.New(ErrNotExists)
