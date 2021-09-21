<?php

$data = [1,2,3,4,5,6,7,8,9,10];

$dataResult = array_map(fn(int $value) => $value * 10, $data);
var_dump($dataResult);

var_dump(rsort($data));
var_dump(array_values($data));
var_dump(array_keys($data));

$person = [
    "firstname" => "Made",
    "lastname" => "Garda"
];

var_dump(array_values($person));
var_dump(array_keys($person));

