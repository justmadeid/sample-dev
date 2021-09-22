<?php

require_once __DIR__ . '/GetConnection.php';

$connetion = getConnection();

$sql = <<<SQL
    INSERT INTO customers(id,name,email) 
    VALUES ("2","Garda","garda@gmail.com");
SQL;

$connetion->exec($sql);

$connetion = null;