#!/bin/bash
# /usr/bin/mysqld_safe --skip-grant-tables &

printf "Making sure mysql service is up. Enter SUDO password to start mysql service \n"
sudo service mysql start

printf "Creating db structure from init.sql file. Insert your mysql root password below. \n"
# mysql -u root -e "CREATE DATABASE mydb"
mysql -u root -p < init.sql #you'll be asked for your root password

printf "Starting the go application... \n"
sleep 1
go run *.go