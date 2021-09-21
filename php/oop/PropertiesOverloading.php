<?php

class Zero
{
    private array $properties = [];

    public function __get($name)
    {
//        echo "access Property $name".PHP_EOL;
//        return "contoh";
        return $this->properties[$name];
    }

    public function __set($name, $value)
    {
//        echo "Set Properties with value $value".PHP_EOL;
        $this->properties[$name] = $value;
    }

    public function __isset($name): bool
    {
//        echo "isset $name".PHP_EOL;
        return isset($this->properties[$name]);
    }

    public function __unset($name)
    {
//        echo "unset $name".PHP_EOL;
        unset($this->properties[$name]);
    }

    public function __call($name, $arguments)
    {
        $join = join(",", $arguments);
        echo "Call function $name with arguments $join". PHP_EOL;
    }

    public static function __callStatic($name, $arguments)
    {
        $join = join(",", $arguments);
        echo "Call static function $name with arguments $join". PHP_EOL;
    }

}

$zero = new Zero();
$zero->firstName = "Made";
$zero->middleName = "GA";
$zero->lastName = "Sans";

echo "First Name : $zero->firstName" . PHP_EOL;
echo "Middle Name : $zero->middleName" . PHP_EOL;
echo "Last Name : $zero->lastName" . PHP_EOL;

$zero->sayHello("Made", "Yoi");
Zero::sayHello("Garda", "Ok");