# gometr
Простой http сервер, который отдает go метрики

### Зависимости
Для установки зависимостей перейдите в корень проекта и выполните следующую команду
```
$ make install
```

### Linter
Для запуска линтера, перейдите в корень проекта и выполните следующую команду 
```
$ make lint
```

### Folders
* ./config - папка с конфигами приложения
    * gometr.conf.yaml - в этом файле можно сконфигурировать порт, на котором запустится приложение

* ./api - в этой папке может быть сваггер спицификация   

* ./internal/handlers - api хэндлеры нашего сервиса

* ./internal/app - папка с описанием и стартом сервиса

Для более подробной информации о том, как строить структуру файлов и папок для проекта на GO читайте по [ссылке](https://github.com/golang-standards/project-layout)

### http Методы
* /metrics - отдает go метрики (согласно протоколу prometheus)
* /health - сокращенный формат ответа ([полный формат](https://tools.ietf.org/id/draft-inadarei-api-health-check-01.html))
```json
{
     "status": "pass",
     "service_id": "MBPadmincity101",
     "checks": {
         "ping_mysql": {
             "component_id": "mysql",
             "component_type": "db",
             "status": "pass"
         }
     }
 }
```