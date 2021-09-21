<?php

class Person
{
    const AUTHOR = "JUSTMADE.ME";

    var string $name;
    var ?string $address = null;
    var string $country = "Indonesia";

    function __construct(string $name, ?string $address)
    {
        $this->name = $name;
        $this->address = $address;
    }

    function sayHello(?string $name)
    {
        if (is_null($name)) {
            echo "Hallo $this->name" . PHP_EOL;
        }else {
            echo "Hello my name is $name, nice to meet you $this->name".PHP_EOL;
        }
    }

    function info()
    {
        echo "Author : " . self::AUTHOR;
    }

    function __destruct()
    {
        echo "Object person $this->name is destroyed" . PHP_EOL;
    }

}
