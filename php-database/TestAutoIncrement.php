<?php

require_once __DIR__ . '/GetConnection.php';

$connetion = getConnection();

$connetion->exec("INSERT INTO comments(email,comment) VALUES ('made@gmail','Haloo whats up')");
$id = $connetion->lastInsertId();

echo "$id".PHP_EOL;

$connetion = null;
