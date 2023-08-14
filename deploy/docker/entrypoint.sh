#!/bin/bash
if [ ! -d "/var/lib/mysql/ecovacs" ]; then
    mysqld --initialize-insecure --user=mysql --datadir=/var/lib/mysql
    mysqld --daemonize --user=mysql
    sleep 5s
    mysql -uroot -e "create database ecovacs default charset 'utf8' collate 'utf8_bin'; grant all on ecovacs.* to 'root'@'127.0.0.1' identified by '123456'; flush privileges;"
else
    mysqld --daemonize --user=mysql
fi
redis-server &
if [ "$1" = "actions" ]; then
    cd /opt/ecovacs/server && go run main.go &
    cd /opt/ecovacs/web/ && yarn serve &
else 
    /usr/sbin/nginx &
    cd /usr/share/nginx/html/ && ./server &
fi
echo "ecovacs ALL start!!!"
tail -f /dev/null