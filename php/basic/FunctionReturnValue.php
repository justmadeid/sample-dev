<?php

function sum(int $first, int $second) : int
{
    $total = $first + $second;
    return $total;
}

$result = sum(90,50);
echo $result;
var_dump($result);

function getFinalValue(int $value): string
{
    if ($value >= 80) {
        return "A";
    } else if ($value >= 70) {
        return "B";
    } else if ($value >= 60) {
        return "C";
    } else if ($value >= 50) {
        return "D";
    } else {
        return "E";
    }

    //tidak akan pernah di eksekusi
    echo "Ups" . PHP_EOL;
}

$score = getFinalValue(90);
var_dump($score);

$score = getFinalValue(30);
var_dump($score);