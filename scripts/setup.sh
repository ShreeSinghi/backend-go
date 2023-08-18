#!/bin/bash

if [[ ! $(command -v migrate) ]]; then
    echo "Migrate not found, please install it to continue"
    exit

elif [[ ! $(command -v go) ]]; then
    echo "Golang not found, please install it to continue"
    exit

elif [[ ! $(command -v mysql) ]]; then
    echo "MySQL not found, please install it to continue"
    exit

else
    echo "Prerequisites ready"

fi

echo "Making config.yaml, enter database credentials"
read -p "Username: " dbUser && echo "dbUser: $dbUser" > config.yaml
read -s -p "Password: " password && echo "password: $password" >> config.yaml && echo ""
read -p "Database Name: " dbName && echo "dbName: $dbName" >> config.yaml
read -p "Host: " host && echo "host: $host" >> config.yaml

echo "Applying migration..."
migrate -path database/migration/ -database "mysql://$dbUser:$password@tcp($host)/$dbName" up
echo "Done"

go mod tidy
go mod vendor