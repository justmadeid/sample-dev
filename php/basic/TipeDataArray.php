<?php

$value = array(10, 9, 8, 7.5);
var_dump($value);

$names = ["made", "garda", "setiawan"];
var_dump($names);

var_dump($names[0]);

$names[2] = "sanskara";
var_dump($names);
var_dump($names[2]);

unset($names[1]);
var_dump($names);

$names[] = "Mantap";
var_dump($names);

var_dump(count($names));

$made = array(
    "id" => "made",
    "name" => "garda",
    "age" => 20,
    "address" => array(
        "city" => "Jakarta",
        "country" => "Indonesia"
    )
);
var_dump($made);
var_dump($made["id"]);
var_dump($made["address"]["country"]);

$garda = [
    "id" => "made",
    "name" => "garda",
    "age" => 20,
    "address" => [
        "city" => "Jakarta",
        "country" => "Indonesia"
    ]
];
var_dump($garda);
var_dump($garda["id"]);