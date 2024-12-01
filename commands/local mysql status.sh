#!/bin/bash

# Define the credentials
USERNAME="your_username"
PASSWORD="your_password"
HOST="your_host"  
PORT="your_port"
DATABASE="your_database"

# Ping the MySQL database
mysql --user="$USERNAME" --password="$PASSWORD" --host="$HOST" --port="$PORT" --execute="SELECT 1;" "$DATABASE" &>/dev/null

if [ $? -eq 0 ]; then
    echo "online"
else
    echo "offline"
fi
