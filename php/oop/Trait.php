<?php

require_once "data/SayGoodBye.php";

use Data\Traits\{Person,SayHello,SayGoodBye};

$person = new Person();
$person->goodBye("made");
$person->hello("garda");

$person->name = "Sanskara";
var_dump($person);

$person->run();