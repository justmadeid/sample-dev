<?php

$sayHello = function (string $name){
    echo "Hello $name". PHP_EOL;
};

$sayHello("Made");
$sayHello("Garda");

function sayGoodBye(string $name, $filter){
    $finalName = $filter($name);
    echo "Good bye $finalName" . PHP_EOL;
}

sayGoodBye("Made", function (string $name): string {
 return strtoupper($name);
});

$filterFunction = function (string $name): string {
  return strtolower($name);
};

sayGoodBye("Garda", $filterFunction);

$firstName = "made";
$lastName = "garda";

$sayHelloNew = function () use ($firstName,$lastName){
    echo "Hello $firstName $lastName" . PHP_EOL;
};

$sayHelloNew();