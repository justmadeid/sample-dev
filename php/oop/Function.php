<?php

require_once "data/Person.php";

$person = new Person("Made",null);
$person ->name = "Made";
$person->sayHello("Garda");
$person->info();

$person2 = new Person("Sanskara", "Bali");
$person2 ->name = "Sanskara";
$person2->sayHello(null);