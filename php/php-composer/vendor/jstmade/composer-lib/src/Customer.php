<?php

namespace Jstmade\belajar;

class Customer {

    public function __construct(private string $name)
    {
    }

    public function sayHello(string $name="guest"):string {

        return "Hello $name my name is $this->name".PHP_EOL;
    }
}
