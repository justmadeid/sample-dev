<?php

$host = "localhost";
$port = 3306;
$database = "php_database";
$username = "root";
$password = "";

try {
    $connection = new PDO("mysql:host=$host:$port;dbname=$database", $username, $password);
    echo "Connection Success" . PHP_EOL;

    //closed connection
    $connection = null;
}catch (PDOException $exception){
    echo "Can't COnnect to yout DATABASE error : ". $exception->getMessage().PHP_EOL;
}