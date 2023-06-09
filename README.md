# gometr
Простой http сервер, который отдает go метрики

### Зависимости
Для установки зависимостей перейдите в корень проекта и выполните следующую команду
```
$ make install
```

### Запуск
Чтобы запустить сервис, из корня прокета выполните команду  
```
$ make run
```

### Линтер
Для запуска линтера, перейдите в корень проекта и выполните следующую команду 
```
$ make lint
```

### Структура папок
* ./config - папка с конфигами приложения
    * gometr.conf.yaml - в этом файле можно сконфигурировать порт, на котором запустится приложение

* ./api - в этой папке может быть сваггер спицификация   

* ./internal/handlers - api хэндлеры нашего сервиса

* ./internal/app - папка с описанием и стартом сервиса

Для более подробной информации о том, как строить структуру файлов и папок для проекта на GO читайте по [ссылке](https://github.com/golang-standards/project-layout)

### http Методы
Приложение стартует на порту 8000. После запуска будет доступен по url localhost:8000
* /metrics - отдает go метрики (согласно протоколу prometheus)
* /health - сокращенный формат ответа стандарта хелсчека ([полный формат](https://tools.ietf.org/id/draft-inadarei-api-health-check-01.html))
```json
{
     "status": "pass",
     "service_id": "gometr",
     "checks": {
         "ping_mysql": {
             "component_id": "mysql",
             "component_type": "db",
             "status": "pass"
         }
     }
 }
```