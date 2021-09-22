<?php

require_once __DIR__ . '/GetConnection.php';

$connetion = getConnection();

$sql = "SELECT * FROM customers";
$statement = $connetion->query($sql);
foreach ($statement as $row) {
    $id = $row['id'];
    $name = $row['name'];
    $email = $row['email'];

    echo "id : $id".PHP_EOL;
    echo "name : $name".PHP_EOL;
    echo "email : $email".PHP_EOL;
}

$connetion = null;
