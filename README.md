# tinyurl  
tinyurl - моя реализация сервиса по созданию сокращённых URL-адресов.  
## Описание    
Сервис генерирует для URL сокращённый адрес - хэш длиной в 10 символов. Хэш состоит из символов латинского алфавита в врехнем и нижнем регистре, а также цифр. Из допустимого набора символов исключены 'I', 'l', 'O' и '0' (во избежание путаницы), а также буквы 'U' и 'u' (чтобы снизить вероятность появления в хэше нецензурных англоязычных слов).        
Доступ к API сервиса реализован по следующим протоколам:  
* HTTP (используется порт 8080)  
* GRPC (используется порт 50051)  

Предоставляется две опции для хранения данных (соответствие исходного URL и его хэша): напрямую в памяти или в БД Postgres. Выбор типа хранилища осуществляется через задание переменной окружения в файле .env:  
`STORAGE_TYPE=postgres` для использования БД  
`STORAGE_TYPE=inmemory` для использования памяти сервиса    

Также в файле .env задаются параметры подключения к БД и используемые порты для обращений HTTP и GRPC запросов. Пример параметров по уполчанию:  
`GRPC_PORT=50051`  
`HTTP_PORT=8080`  
`LOG_LEVEL=info`  
`DB_HOST=localhost`  
`DB_PORT=5432`  
`DB_USER= test`  
`DB_PASSWORD=test`  
`DB_NAME=urls_db`  

Для логирования используется библиотека zap, т.о. логи отображаются в формате JSON.    

## Использование сервиса  
### HTTP  
Для получения сокращённого адреса (хэш) с помощью метода POST передаётся JSON, где для поля longUrl указывается значение - исходный URL-адрес.  
Пример использования:       
```
curl -X POST http://localhost:8080/v1/set_url -d '{"longUrl":"https://wikipedia.org"}'  
```
В ответ от сервиса приходит JSON с полем shortUrl, где указан хэш отправленного URL-адреса.  
Пример ответа сервиса:  
```
{"shortUrl":"sTVmLmA66b"}
```  
  
Для получения полного (исходного) URL-адреса по хэшу используется метод GET.  
```
curl -X GET http://localhost:8080/sTVmLmA66b  
```  
В случае успеха приходит JSON с полем longUrl, где указан исходный (полный) URL-адрес.    
Пример ответа сервиса:  
```
{"longUrl":"https://wikipedia.org"}  
```  
Ответ в этом случае отправляется с кодом [301](https://ru.wikipedia.org/wiki/HTTP_301).  
Если запрошенный хэш имеет некорректный формат или не найден в базе, то в ответ будет отправлен код 400 или 404 соответственно.  

### Сборка сервиса    
Сервис может быть запущен в виде Docker-образа. В этом случае используется только inmemory хранилище.  
В случае использования БД до запуска сервиса нужно запустить соответствующий контейнер с БД postgres.    
Для удобства сборки и запуска docker и сервиса заданы соответствующие цели в Makefile.  

Сборка и запуск сервиса в виде docker:  
```bash
$ make docker-build  
$ make docker-run  
```  
Запуск docker с postgres:  
```bash
$ make compose-up  
``` 
Запуск сервиса без использования docker:  
```bash
$ make
```  
Запуск тестов:  
```bash
$ make test
```  
