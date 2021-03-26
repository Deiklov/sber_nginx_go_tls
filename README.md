# sber_nginx_go_tls
### Задача:  
Сделать авторизацию в nginx с использованием сертификатов,  
клиент go https ,  
nginx  отвечает либо 200 ok, либо проксирует запрос дальше  
### Исполнитель: 
Романов Андрей(дивизион "Риски розничного бизнеса")
### Описание работы:  
Статья на основе которой делал https://medium.com/rahasak/tls-mutual-authentication-with-golang-and-nginx-937f0da22a0e  
1.  Генерим клиентские и серверные сертификаты c помощью openssl, лежат в папке certs
2.  Запускаем nginx накидываем в /etc/certs сертификаты
3.  Запускаем: docker-compose up -d nginx
4.  Запускаем: bash netcat.sh(фейковый сервер)
5.  Запускаем go клиент: go run main.go

 

