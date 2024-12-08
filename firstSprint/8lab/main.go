package main

import (
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Логируем информацию о запросе
		log.Printf("Запрос: %s %s", r.Method, r.URL.Path)

		// Передаём управление следующему обработчику
		next.ServeHTTP(w, r)

		// Вычисляем время выполнения запроса
		duration := time.Since(start)
		log.Printf("Время выполнения запроса: %s", duration)
	})
}

func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет, мир!"))
}

func main() {
	//mux := http.NewServeMux()
	//
	//// Создаём обработчик для маршрута "/"
	//hello := http.HandlerFunc(helloHandler)
	//
	//// Применяем logging middleware к обработчику "/"
	//mux.Handle("/", PanicMiddleware(hello))
	//
	//// Запускаем сервер на порту 8080
	//if err := http.ListenAndServe(":8080", mux); err != nil {
	//
	//	log.Fatal(err)
	//}

	StartServer3(time.Second)
}
