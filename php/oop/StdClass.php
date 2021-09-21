<?php
require_once "data/Person.php";

$array = [
    "firstName" => "Made",
    "middleName" => "Garda",
    "lastName" => "Setiawan"
];

$object = (object) $array;
var_dump($object);

echo "First Name $object->firstName".PHP_EOL;
echo "Middle Name $object->middleName".PHP_EOL;
echo "Last Name $object->lastName".PHP_EOL;

$arrayLagi = (array) $object;
var_dump($arrayLagi);

$person = new Person("Made","Semarang");
var_dump($person);

$arrayPerson = (array) $person;
var_dump($arrayPerson);