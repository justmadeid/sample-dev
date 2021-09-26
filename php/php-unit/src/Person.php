<?php

namespace Jstmade\Test;

class Person
{

    public function __construct(private string $name)
    {
    }

    public function sayHello(?string $name)
    {
        if ($name == null) throw new \Exception("Name is Null");

        return "Hello $name, my name is $this->name";

    }

    public function sayGoodbye(?string $name)
    {
        echo "Good Bye $name". PHP_EOL;
    }

}