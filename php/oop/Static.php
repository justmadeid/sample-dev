<?php

require_once "helper/MathHelper.php";
use Helper\MathHelper;

echo MathHelper::$name;
MathHelper::$name = "Test Math".PHP_EOL;
echo MathHelper::$name;

$result = MathHelper::sum(20,20,20,40,50,90);
echo "Result : $result".PHP_EOL;



