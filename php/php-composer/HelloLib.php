<?php

require_once __DIR__ . '/vendor/autoload.php';

$customer = new \Jstmade\belajar\Customer("made");
echo $customer->sayHello() . PHP_EOL;