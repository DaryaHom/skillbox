1. Создайте директорию attestation в директории src вашего окружения golang
2. Склонируйте этот репозиторий в директорию attestation (git clone ...)
3. Запустите программу (go run main.go). Симулятор данных запустится автоматически.
4. По умолчанию программа использует порт :8585, симулятор - порт :8383. 
Упорядоченный набор сгенерированных данных можно получить get-запросом по адресу http://localhost:8585/struct