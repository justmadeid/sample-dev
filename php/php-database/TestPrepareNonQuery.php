<?php

require_once __DIR__ . '/GetConnection.php';

$connection = getConnection();

$username = "sanskara";
$password = "yoi";

//$sql = "INSERT INTO admin(username, password) VALUES (:username, :password)";
//$statement = $connection->prepare($sql);
//$statement->bindParam("username", $username);
//$statement->bindParam("password", $password);
//$statement->execute();

//recomended
$sql = "INSERT INTO admin(username, password) VALUES (?, ?)";
$statement = $connection->prepare($sql);
$statement->execute([$username,$password]);

$connection = null;