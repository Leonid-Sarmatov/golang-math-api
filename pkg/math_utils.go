package pkg

import (
	"github.com/Knetic/govaluate"
)

/*
parseMathExpression Функция для парсинга выражения из строки

Parameters:

	string: Входное выражение в строке

Returns:

	*govaluate.EvaluableExpression: Выражение, готовое к вычислению
	error: Ошибки
*/
func parseMathExpression(expr string) (*govaluate.EvaluableExpression, error) {
	parsed, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

/*
evaluateMathExpression Функция вычисления выражения

Parameters:

	*govaluate.EvaluableExpression: Входное выражение
	map[string]interface{}: Словарь со значениями переменных

Returns:

	interface{}: Результат вычисления
	error: Ошибки
*/
func evaluateMathExpression(expr *govaluate.EvaluableExpression,
	vars map[string]interface{}) (interface{}, error) {
	result, err := expr.Evaluate(vars)
	if err != nil {
		return nil, err
	}

	return result, nil
}

/*
evaluateMathExpression Функция вычисления значений выражения при разных переменных

Parameters:

	*govaluate.EvaluableExpression: Входное выражение
	[]map[string]interface{}: Срез со словарями значений

Returns:

	[]interface{}: Массив с результатами вычислений
	error: Ошибки
*/
func MultyVarEvaluateMathExpression(expr string,
	varsArray []map[string]interface{}) ([]interface{}, error) {
	parsed, err := parseMathExpression(expr)
	if err != nil {
		return nil, err
	}

	resultArray := make([]interface{}, 0)
	for _, varMap := range varsArray {
		res, err := evaluateMathExpression(parsed, varMap)
		if err != nil {
			return nil, err
		}
		resultArray = append(resultArray, res)
	}

	return resultArray, nil
}
