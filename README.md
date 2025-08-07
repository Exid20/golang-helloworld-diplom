# Golang HelloWorld Diplom

Простой HTTP-сервер на Go, который возвращает "Hello, World!" в ответ на запросы.

## Описание

Этот проект представляет собой минималистичный HTTP-сервер, написанный на языке Go. Сервер запускается на порту `8080` и отвечает строкой `Hello, World!` на любой GET-запрос к корневому пути (`/`).


##Technologies

Go

Docker

Оркестрация через docker-compose

terraform

yandex cloud

Github action CI/CD

Prometheus + Grafana

##Project Structure

├── docker-compose.yaml

├── Dockerfile

├── go.mod

├── go.sum

├── LICENSE

├── main.go

├── prometheus.yml

├── provisioning

│   ├── dashboards

│   │   ├── dashboards.yml

│   │   └── prometheus-stats.json

│   └── datasources

│       └── datasources.yml

├── README.md

├── renovate.json

└── terraform

    ├── main.tf
    
    ├── outputs.tf
    
    ├── terraform.tfstate
    
    ├── terraform.tfstate.backup
    
    ├── terraform.tfvars
    
    └── variables.tf
    
##CI/CD

Проект использует GitHub Actions для автоматической сборки, тестирования и деплоя приложения в Yandex Cloud.
Весь процесс запускается при пуше в ветку main или создании pull request.

## Этапы CI/CD
Build & Test (build-test):
Установка Go 1.21

Проверка форматирования кода (go fmt)

Запуск тестов (go test)

Сборка бинарного файла (go build)

Сборка Docker-образа с уникальным тегом (на основе SHA коммита)

Сохранение артефакта (бинарный файл)

Deploy (deploy):

Инициализация Terraform (terraform init)

Применение инфраструктуры в Yandex Cloud (terraform apply)

Получение внешнего IP виртуальной машины

Ожидание доступности VM по SSH

Деплой приложения на VM:

Установка Docker и Docker Compose (если отсутствуют)

Копирование проекта через SCP

Запуск через docker compose up --build -d

Уведомления:

Отправка статуса выполнения в Telegram-чат (успех/ошибка)

## Требуемые Secrets

Для работы пайплайна необходимы следующие секреты в GitHub:

SSH_PUBLIC_KEY — публичный SSH-ключ для доступа к VM

SSH_KEY — приватный SSH-ключ

SSH_USERNAME — имя пользователя для SSH

YC_TOKEN, YC_CLOUD_ID, YC_FOLDER_ID — для аутентификации в Yandex Cloud

TELEGRAM_TOKEN, TELEGRAM_CHAT_ID — для уведомлений


## Особенности

Инфраструктура управляется через Terraform (см. папку terraform/)

Используется Docker Compose для оркестрации сервисов

Мониторинг через Prometheus + Grafana (конфиги в provisioning/)

## Технологии CI/CD

GitHub Actions

Terraform (Yandex Cloud Provider)

Docker + Docker Compose

Telegram Bot API (уведомления)

### Запуск сервера

1. Убедитесь, что у вас установлен Go (версия 1.20 или выше).
2. Клонируйте репозиторий:
   git clone https://github.com/Exid20/golang-helloworld-diplom.git


