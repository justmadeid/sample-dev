<?php

require_once __DIR__ . '/vendor/autoload.php';

use Jstmade\Data\People;

$people = new People("Made");

echo $people->sayHello("Budi") . PHP_EOL;