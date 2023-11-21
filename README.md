# Итоговый проект 2023
### Задание
Разработать систему заказа такси. Сервисы должны удовлетворять контрактам из репозитория. Разработать схемы для хранения данных, написать миграции и заполнить базы данных тестовыми данными.
#### Обеспечить наблюдаемость работы сервисов: 
- сервисы должны писать структурированный логи в stdout,stderr
- писать метрики в prometheus 
- обеспечивать возможности сквозной трассировки (обрабатывать данные трассировки в заголовках запросов/сообщениях kafka, писать данные о трассировке в jaeger). 
#### Обеспечить возможность развертывания сервисов и их зависимостей в docker с использованием docker compose. 
- баз данных
- kafka
- jaeger
- prometheus
- grafana (опционально)
#### Разработать тесты 
Покрыть сервисы unit и интеграционными тестами. 
#### Студенты ВШЭ реализуют сервисы
- Driver
- Location
#### Студенты МФТИ реализуют сервисы
- Client
- Offering
- Trip
#### Требования
- команда работает в монорепозитории (все разрабатываемые сервисы должны находиться в одном репозитории)
- для тестирования сервисов разрабатываемых другими командами необходимо разработать моки на основании контрактов
- использовать подход чистая архитектура
- использовать https://github.com/golang-standards/project-layout
- все настраиваемые параметры должны передаваться через env, создать файл .env.dev с настройками для работы сервисов в окружении для разработки. 
#### Система оценки
Реализация всех требований основного задания - **7 баллов**

Покрытие тестами > 60% **+ 1 балл**

Уведомление клиентов и водителей об изменении состояния поездки/предложении поездки через Websocket (масштабируемое решение с использованием pub/sub) **+ 1 балл**

Предложить ментору улучшение системы и реализовать после согласования **+ 1 балл**
### Общая схема системы
![schema](img/schema.png)
### Client Service
Отвечает за управление заказом со стороны клиента (создание, отмена, получение обновлений).
Реагирует на события об изменении состояния поездки.
### Trip Service 
Отвечает за управление состоянием поездки. Публикует события об изменении состояния поздки в отдельный топик. Реагирует на команды в специальном топике.
###  Driver Service 
Отвечает за управление заказом со стороны водителя (принятие, отмена, завершение, получение заказов)
Отвечает за поиск исполнителя заказа.
Реагирует на события об изменении состояния поездки.
### Location Service 
Отвечает за хранение и предоставление данных о местонахождении водителей
### Offering Service 
Отвечает за создание предложения. Рассчитвает стоимость поездки (методика рассчета на усмотрение студента).
### Состояния поздки
![states](img/states.png)
### Процесс успешной поездки
![process1](image.png)