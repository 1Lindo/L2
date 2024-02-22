/*
Данный код реализует HTTP-сервер для работы с календарем с использованием стандартной HTTP-библиотеки.

Реализовный методы:

POST /create_event
POST /update_event
POST /delete_event
GET /events_for_day
GET /events_for_week
GET /events_for_month

Логика работы с базой данных НЕ зависит от кода HTTP сервера, она вынесена в отдельные функции: CreateEventInDb, UpdateEventInDb,
DeleteEventFromDb, GetDayEventFromDb, GetWeekEventFromDb, GetMonthEventFromDb, которые в дальнейшем могут быть вынесены
в другой пакет.

Логика работы методов для сериализации вынесена в отдельные фукнции: serializeObject и deserializeObject.

Логика работы проверки параметров, а также парсинга queryParams из POST запросов вынесена в отдельные фукнции:
checkEventParams, checkGetEventParams, parseQueryStringParams.

За логирование запросов отвечает logMiddleware.

Интерпретация:
1) В случае ошибки бизнес-логики сервер возвращает HTTP 503;
2) В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400;
3) В случае остальных ошибок сервер должен возвращать HTTP 500;
5) Web-сервер запускается на порту, который указан в конфиге и выводит в лог каждый обработанный запрос.

Для тестирования предлагается использовать Postman:
1) Для тестов POST запросов используйте JSON из константы EVENT1 и поместите его в body запроса;
2) Для тестов GET запросов запустите программу на локальном сервере, параметры указаны в queryParams.
http://localhost:8086/create_event?user_id=3&?date=2019-09-09


*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

const EVENT1 = `{
	"user_id":"2",
	"date":"2019-09-09"
}`

// Event представляет собой структуру для события в календаре.
type Event struct {
	UserID int    `json:"user_id" db:"userId"`
	Date   string `json:"date" db:"date"`
}

// CalendarService предоставляет методы для работы с календарем.
type CalendarService struct {
	// Здесь могут быть поля для хранения данных или зависимостей.
}

func main() {
	// Создание экземпляра CalendarService
	calendarService := &CalendarService{}

	// Регистрация обработчиков
	http.HandleFunc("/create_event", calendarService.createEventHandler)
	http.HandleFunc("/update_event", calendarService.updateEventHandler)
	http.HandleFunc("/delete_event", calendarService.deleteEventHandler)
	http.HandleFunc("/events_for_day", calendarService.getEventForDayHandler)
	http.HandleFunc("/events_for_week", calendarService.getEventForWeekHandler)
	http.HandleFunc("/events_for_month", calendarService.getEventForMonthHandler)

	// Применение middleware для логирования
	loggedRouter := logMiddleware(http.DefaultServeMux)

	// Запуск сервера на порту 8086
	port := 8086
	fmt.Printf("Server is running on :%d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), loggedRouter))
}

// Ниже представлены тестовые фукнции CRUD и методы для запросов в базу данных
func (cs *CalendarService) CreateEventInDb(e *Event) (string, error) {
	fmt.Printf("Writing this: \"%v\" to db", e)
	return "success!", nil
}

func (cs *CalendarService) UpdateEventInDb(e *Event) (string, error) {
	fmt.Printf("Updating this: \"%v\" in db", e)
	return "success!", nil
}

func (cs *CalendarService) DeleteEventFromDb(e *Event) (string, error) {
	fmt.Printf("Deleted this: \"%v\" from db", e)
	return "success!", nil
}

func (cs *CalendarService) GetDayEventFromDb(e *Event) (Event, error) {
	fmt.Printf("Got this day event: \"%v\" from db", e)
	res := *e
	return res, nil
}

func (cs *CalendarService) GetWeekEventFromDb(e *Event) (Event, error) {
	fmt.Printf("Got this week event: \"%v\" from db", e)
	res := *e
	return res, nil
}

func (cs *CalendarService) GetMonthEventFromDb(e *Event) (Event, error) {
	fmt.Printf("Got this month: \"%v\" from db", e)
	res := *e
	return res, nil
}

// Сериализация объекта в JSON
func serializeObject(obj Event) ([]byte, error) {
	const x string = "Сериализация объекта в JSON"
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return nil, fmt.Errorf("%s:%v", x, err)
	}
	return jsonData, nil
}

// Десериализация JSON в объект
func deserializeObject(jsonStr string, obj *Event) error {
	const x string = "Десериализация объекта из JSON"
	err := json.Unmarshal([]byte(jsonStr), obj)
	if err != nil {
		return fmt.Errorf("%s:%v", x, err)
	}
	return nil
}

// Вспомогательная функция для парсинга и валидации параметров создания
func checkEventParams(e *Event) error {
	// Валидация параметров
	if e.UserID < 0 {
		return fmt.Errorf("неверное значение ID: %d", e.UserID)
	}
	_, err := time.Parse("2006-01-02", e.Date)
	if err != nil {
		return fmt.Errorf("неверный формат даты: %v", err)
	}
	return nil
}

func checkGetEventParams(params map[string]string) (*Event, error) {
	userId, err := strconv.Atoi(params["user_id"])
	if err != nil {
		return nil, fmt.Errorf("ошибка преобразования user_id: %v", err)
	}
	parsedParams := &Event{
		UserID: userId,
		Date:   params["date"],
	}
	return parsedParams, nil
}

// Парсинг url в map с параментрами event
func parseQueryStringParams(r *http.Request) (map[string]string, error) {
	// Получение строки запроса из URL
	queryString := r.URL.Query()

	// Преобразование параметров в map[string]string
	paramsMap := make(map[string]string)
	for key, values := range queryString {
		// Берем только первое значение, игнорируем остальные (если они есть)
		paramsMap[key] = values[0]
	}

	return paramsMap, nil
}

// createEventHandler обрабатывает запросы на создание события.
func (cs *CalendarService) createEventHandler(w http.ResponseWriter, r *http.Request) {

	var event Event
	//процесс десериализации:
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		// Обработка прочих ошибок
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"error": "Внутренняя ошибка сервера"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Парсинг и валидация параметров:
	err = checkEventParams(&event)
	if err != nil {
		// Обработка ошибок входных данных
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"error": "Ошибка входных данных"}
		json.NewEncoder(w).Encode(response)
		return

	}

	//Создание event в DB календаря
	_, err = cs.CreateEventInDb(&event)
	if err != nil {
		// Обработка ошибок бизнес-логики
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"error": "Ошибка в бизнес-логике"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Обработка успешного выполнения метода
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{"result": "Успешный результат"}
	json.NewEncoder(w).Encode(response)
	return
}

// updateEventHandler обрабатывает запросы на обновление события.
func (cs *CalendarService) updateEventHandler(w http.ResponseWriter, r *http.Request) {

	var event Event
	//процесс десериализации:
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		// Обработка прочих ошибок
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"error": "Внутренняя ошибка сервера"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Парсинг и валидация параметров:
	err = checkEventParams(&event)
	if err != nil {
		// Обработка ошибок входных данных
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"error": "Ошибка входных данных"}
		json.NewEncoder(w).Encode(response)
		return
	}

	//Создание event в DB календаря
	_, err = cs.UpdateEventInDb(&event)
	if err != nil {
		// Обработка ошибок бизнес-логики
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"error": "Ошибка в бизнес-логике"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Обработка успешного выполнения метода
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{"result": "Успешный результат"}
	json.NewEncoder(w).Encode(response)
	return
}

// deleteEventHandler обрабатывает запросы на удаление события.
func (cs *CalendarService) deleteEventHandler(w http.ResponseWriter, r *http.Request) {

	var event Event
	//процесс десериализации:
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		// Обработка прочих ошибок
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"error": "Внутренняя ошибка сервера"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Парсинг и валидация параметров:
	err = checkEventParams(&event)
	if err != nil {
		// Обработка ошибок входных данных
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"error": "Ошибка входных данных"}
		json.NewEncoder(w).Encode(response)
		return
	}

	//Создание event в DB календаря
	_, err = cs.DeleteEventFromDb(&event)
	if err != nil {
		// Обработка ошибок бизнес-логики
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"error": "Ошибка в бизнес-логике"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Обработка успешного выполнения метода
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{"result": "Успешный результат"}
	json.NewEncoder(w).Encode(response)
	return
}

// eventsForDayHandler обрабатывает запросы на получение событий за день.
func (cs *CalendarService) getEventForDayHandler(w http.ResponseWriter, r *http.Request) {
	// Парсинг url и преобразование полученных данных в map[string]string
	params, err := parseQueryStringParams(r)
	if err != nil {
		// Обработка прочих ошибок
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"error": "Внутренняя ошибка сервера"}
		json.NewEncoder(w).Encode(response)
		return
	}
	// Преобразуем map в с новую структуру Event и возвращаем указатель на нее
	parsedParams, err := checkGetEventParams(params)
	if err != nil {
		// Обработка ошибок входных данных
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"error": "Ошибка входных данных"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Передаем указатель на структуру Event в метод DB для поиска там соответсвий, получаем необходимый Event
	data, err := cs.GetDayEventFromDb(parsedParams)
	if err != nil {
		// Обработка ошибок бизнес-логики
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"error": "Ошибка в бизнес-логике"}
		json.NewEncoder(w).Encode(response)
		return
	}
	//Маршлим data в JSON
	jsonData, err := serializeObject(data)
	if err != nil {
		// Обработка ошибок сервера
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"error": "Внутренняя ошибка сервера"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Обработка успешного выполнения метода и передача Client данных в JSON
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	response := map[string]interface{}{"result": "Успешный результат"}
	json.NewEncoder(w).Encode(response)
	return
}

// eventsForWeekHandler обрабатывает запросы на получение событий за неделю.
func (cs *CalendarService) getEventForWeekHandler(w http.ResponseWriter, r *http.Request) {
	// Парсинг url и преобразование полученных данных в map[string]string
	params, err := parseQueryStringParams(r)
	if err != nil {
		// Обработка прочих ошибок
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"error": "Внутренняя ошибка сервера"}
		json.NewEncoder(w).Encode(response)
		return
	}
	// Преобразуем map в с новую структуру Event и возвращаем указатель на нее
	parsedParams, err := checkGetEventParams(params)
	if err != nil {
		// Обработка ошибок входных данных
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"error": "Ошибка входных данных"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Передаем указатель на структуру Event в метод DB для поиска там соответсвий, получаем необходимый Event
	data, err := cs.GetWeekEventFromDb(parsedParams)
	if err != nil {
		// Обработка ошибок бизнес-логики
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"error": "Ошибка в бизнес-логике"}
		json.NewEncoder(w).Encode(response)
		return
	}
	//Маршлим data в JSON
	jsonData, err := serializeObject(data)
	if err != nil {
		// Обработка ошибок сервера
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"error": "Внутренняя ошибка сервера"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Обработка успешного выполнения метода и передача Client данных в JSON
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	response := map[string]interface{}{"result": "Успешный результат"}
	json.NewEncoder(w).Encode(response)
	return

}

// eventsForMonthHandler обрабатывает запросы на получение событий за месяц.
func (cs *CalendarService) getEventForMonthHandler(w http.ResponseWriter, r *http.Request) {
	// Парсинг url и преобразование полученных данных в map[string]string
	params, err := parseQueryStringParams(r)
	if err != nil {
		// Обработка прочих ошибок
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"error": "Внутренняя ошибка сервера"}
		json.NewEncoder(w).Encode(response)
		return
	}
	// Преобразуем map в с новую структуру Event и возвращаем указатель на нее
	parsedParams, err := checkGetEventParams(params)
	if err != nil {
		// Обработка ошибок входных данных
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"error": "Ошибка входных данных"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Передаем указатель на структуру Event в метод DB для поиска там соответсвий, получаем необходимый Event
	data, err := cs.GetMonthEventFromDb(parsedParams)
	if err != nil {
		// Обработка ошибок бизнес-логики
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"error": "Ошибка в бизнес-логике"}
		json.NewEncoder(w).Encode(response)
		return
	}
	//Маршлим data в JSON
	jsonData, err := serializeObject(data)
	if err != nil {
		// Обработка ошибок сервера
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"error": "Внутренняя ошибка сервера"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Обработка успешного выполнения метода и передача Client данных в JSON
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	response := map[string]interface{}{"result": "Успешный результат"}
	json.NewEncoder(w).Encode(response)
	return
}

// logMiddleware выполняет логирование каждого обработанного запроса.
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
