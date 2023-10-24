package my_errors

import "errors"

const (
	Err                     = "ОШИБКА\n"
	ErrAuth                 = "Неверные данные авторизации\n"
	ErrAccess               = "Отказано в доступе\n"
	ErrNotExists            = Err + "не найдено"
	ErrAlreadyExists        = Err + "Конфликтная ситуация. Возможно сущность с введенными данными уже существует.\nвведите другие данные"
	ErrCountry              = Err + "введенная страна не поддерживается"
	ErrMoney                = Err + "неверный формат ввода денежной суммы.\n формат: хххххх.хх"
	ErrNotExistsFoundrising = Err + "У данного фонда нет сборов"
	//SupportedCountries = "Поддерживаемы страны:\nРоссия, США, Великобритания, Канада, Франция, Германия, Китай, Италия, Испания"
)

/*
type ErrorNotExists struct{

}

	func (err ErrorNotExists) Error() string{
		return ErrNotExists
	}
*/
var ErrorNotExists = errors.New(ErrNotExists)
var ErrorConflict = errors.New(ErrAlreadyExists)
var ErrorAccess = errors.New(ErrAccess)
var ErrorAuth = errors.New(ErrAuth)
