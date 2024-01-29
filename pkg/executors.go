package pkg

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
MVEMathExpressionExecutor исполнитель представляющий функционал
по реализации вычисления математический выражений. Исполнитель
принимает по своему маршруту сообщение с выражение в виде строки,
и массив со значениями для переменных в выражении. Значения
должны быть в формате словаря: map[имя переменной]значение.
Обмен даннями происходит в формате JSON
*/
type MVEMathExpressionExecutor struct{}

/*
MVEMathExpressionMessage структура сообщения для данного исполнителя
*/
type MVEMathExpressionMessage struct {
	MathExpression string                   `json:"mathExpression"`
	Variables      []map[string]interface{} `json:"variables"`
}

/*
MVEMathExpressionMessage структура ответа для данного исполнителя
*/
type MVEMathExpressionResponse struct {
	MathExpression string        `json:"mathExpression"`
	Results        []interface{} `json:"results"`
}

func NewMVEMathExpressionExecutor() *MVEMathExpressionExecutor {
	return &MVEMathExpressionExecutor{}
}

func (m *MVEMathExpressionExecutor) getExecutorName() string {
	return "Multy Var Evaluate Math Expression Executor"
}

func (m *MVEMathExpressionExecutor) getExecutorRoute() string {
	return "/MathExpressionEvaluate/MultyVar"
}

func (m *MVEMathExpressionExecutor) getExecutorHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Сообщение
		var message MVEMathExpressionMessage

		// Декодируем тело запроса в сообщение
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&message)
		if err != nil {
			http.Error(w, "[ERROR]: Decoding JSON was failed: "+err.Error(), http.StatusBadRequest)
			log.Println("[ERROR]: Decoding JSON was failed: " + err.Error())
			return
		}

		// Находим значение выражения для всех входных данных
		result, err := MultyVarEvaluateMathExpression(message.MathExpression, message.Variables)
		if err != nil {
			http.Error(w, "[ERROR]: Expression evaluate was failed: "+err.Error(), http.StatusBadRequest)
			log.Println("[ERROR]: Expression evaluate was failed: " + err.Error())
			return
		}

		// Создаем отклик
		response := MVEMathExpressionResponse{
			MathExpression: message.MathExpression,
			Results:        result,
		}

		// Конвертируем отклик в json-отклик
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "[ERROR]: Can not encoding to JSON"+err.Error(), http.StatusInternalServerError)
			log.Println("[ERROR]: Can not encoding to JSON" + err.Error())
			return
		}

		// Заполняем тело запроса и заголовки
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)

		log.Printf("[OK]: Successful work. Executor: %v, Expression: %v, Result: %v",
			m.getExecutorName(), message.MathExpression, result)
	}
}
