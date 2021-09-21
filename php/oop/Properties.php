<?php

require_once "data/Person.php";

$person = new Person("Made", "Bali");
$person ->name = "Made";
$person ->address = null;
$person ->country = "USA";

var_dump($person);

echo "Name : {$person->name}" . PHP_EOL;
echo "Address : {$person->address}". PHP_EOL;
echo "Country : {$person->country}". PHP_EOL;