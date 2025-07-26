# encode
# Person CRUD API (Go + PostgreSQL)

REST API-сервис на Go  для сущности Person с использованием PostgreSQL.

---

## Структура проекта

person-crud/
├── internal/
│   ├── http/
│   ├── logic/
│   ├── models/
│   └── postgres/
├── main.go
├── .env
├── Dockerfile
├── docker-compose.yml
└── README.md

---

## Быстрый старт

### 1. Клонировать репозиторий
1.1 Склонировать проект 
```bash
git clone https://github.com/yourname/person-crud.git
```
1.2 Перейте в папку 
```bash
cd person-crud
```

3. Запустить сервисы
```bash
docker compose up -d --build
```
4. Вывести лог
```bash
 docker compose logs -f app

```
⸻

## Использование

### Создать Person
```bash
curl -X POST http://localhost:8080/person \
  -H "Content-Type: application/json" \
  -d '{
  "email": "egor@mail.com",
  "phone": "9969105118",
  "firstName": "Egor",
  "lastName": "Shapovalov"
}'
```
### Получить всех
```bash
curl http://localhost:8080/person
```
### Получить Person по ID
```bash
curl http://localhost:8080/person/1
```

### Удалить Person
```bash
curl -X DELETE http://localhost:8080/person/1
```

⸻

## Тестирование через Postman
1. Откройте Postman и создайте новую коллекцию, например, Person CRUD.
2. Добавьте следующие запросы:

### Create Person (POST)
URL: http://localhost:8080/person

Метод: POST	

Body → raw → JSON:
```bash
{
  "email": "john@example.com",
  "phone": "1234567890",
  "firstName": "John",
  "lastName": "Doe"
}
```

### Get All Persons (GET)

URL: http://localhost:8080/person
Метод: GET
	•	Get Person by ID (GET)
URL: http://localhost:8080/person/1
Метод: GET
	•	Update Person (PUT)
URL: http://localhost:8080/person/1

⸻

### 2. Очистка данных

Чтобы удалить все данные и тома PostgreSQL:

docker compose down -v
# Person CRUD API (Go + PostgreSQL)

REST API-сервис на Go  для сущности Person с использованием PostgreSQL.

---

## Структура проекта

person-crud/
├── internal/
│   ├── http/
│   ├── logic/
│   ├── models/
│   └── postgres/
├── main.go
├── .env
├── Dockerfile
├── docker-compose.yml
└── README.md

---

## Быстрый старт

### 1. Клонировать репозиторий
1.1 Склонировать проект 
```bash
git clone https://github.com/yourname/person-crud.git
```
1.2 Перейте в папку 
```bash
cd person-crud
```

3. Запустить сервисы
```bash
docker compose up -d --build
```
4. Вывести лог
```bash
 docker compose logs -f app 

```

⸻

## Использование

### Создать Person
```bash
curl -X POST http://localhost:8080/person \
  -H "Content-Type: application/json" \
  -d '{
  "email": "egor@mail.com",
  "phone": "9969105118",
  "firstName": "Egor",
  "lastName": "Shapovalov"
}'
```
Получить всех
```bash
curl http://localhost:8080/person
```
Получить Person по ID
```bash
curl http://localhost:8080/person/1
```

Удалить Person
```bash
curl -X DELETE http://localhost:8080/person/1
```

⸻

## Тестирование через Postman
	
1.	Откройте Postman и создайте новую коллекцию, например, Person CRUD.	
2.	Добавьте следующие запросы:

### Cate Person (POST)
URL: http://localhost:8080/person

Метод: POST

Body → raw → JSON:
```bash
{
  "email": "john@example.com",
  "phone": "1234567890",
  "firstName": "John",
  "lastName": "Doe"
}
```

###	Get All Persons (GET)
URL: http://localhost:8080/person

Метод: GET
	•	Get Person by ID (GET)

URL: http://localhost:8080/person/1

Метод: GET
	•	Update Person (PUT)
URL: http://localhost:8080/person/1

⸻

### 2. Очистка данных

Чтобы удалить все данные и тома PostgreSQL:
 ```bash
docker compose down -v
```