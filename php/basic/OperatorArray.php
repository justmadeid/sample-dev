<?php

$first = [
    "first_name" => "Made"
];

$last = [
    "first_name" => "Garda",
    "last_name" => "Setiawan"
];

$full = $first + $last;
var_dump($full);

$a = [
    "first_name" => "Made",
    "last_name" => "Garda"
];

$b = [
    "last_name" => "Garda",
    "first_name" => "Made"
];

var_dump($a == $b);
var_dump($a === $b);