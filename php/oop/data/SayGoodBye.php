<?php

namespace Data\Traits;

trait SayGoodBye
{
    public function goodBye(?string $name): void
    {
        if (is_null($name)) {
            echo "Good bye" . PHP_EOL;
        } else {
            echo "Good Bye $name" . PHP_EOL;
        }
    }
}

trait SayHello
{
    public function hello(?string $name): void
    {
        if (is_null($name)) {
            echo "Hello " . PHP_EOL;
        } else {
            echo "Hello $name" . PHP_EOL;
        }
    }
}

trait HasName
{
    public string $name;
}

trait CanRun
{
    public abstract function run(): void;
}

class ParentPerson
{
    public function goodBye(?string $name): void
    {
        echo "good bye in person" . PHP_EOL;
    }

    public function hello(?string $name): void
    {
        echo "hello in person" . PHP_EOL;
    }
}

trait All
{
    use SayGoodBye, SayHello, HasName, CanRun {
        // Bisa di override
        // hello as private;
        // goodBye as private;
    }

}

class Person extends ParentPerson
{

    use All;

    public function run(): void
    {
        echo "person $this->name is running" . PHP_EOL;
    }

}