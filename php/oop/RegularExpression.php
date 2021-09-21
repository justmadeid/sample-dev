<?php

$matches = [];
$result = preg_match_all("/mad|awan|sety/i", "Made Garda Setiawan", $matches);

var_dump($result);
var_dump($matches);

$result = preg_replace("/anjing|bangsat/i", "***", "dasar lu ANJING dan BANGSAT!");

var_dump($result);

$result = preg_split("/[\s,-]/", "Made Garda Setiawan,Programmer,Jstmade");

var_dump($result);