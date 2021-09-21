<?php

class Product
{
     protected string $name;
     private int $price;

     public function __construct(string $name, int $price){
         $this->name = $name;
         $this->price = $price;
     }

    /**
     * @return string
     */
    public function getName(): string
    {
        return $this->name;
    }

    /**
     * @return int
     */
    public function getPrice(): int
    {
        return $this->price;
    }
}

class ProductDummy extends Product
{
    public function info()
    {
        echo "name $this->name".PHP_EOL;
    }
}
