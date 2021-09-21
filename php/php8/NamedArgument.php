<?php

function sayHello(string $first, string $middle = "", string $last):void
{
    echo "Hello, $first $middle $last".PHP_EOL;
}

sayHello("Made","Garda","Setiawan");
//sayHello("Made","Setiawan");

sayHello(last: "Setiawan", first: "Made", middle: "Garda");
sayHello(first: "Made", last: "Setiawan");