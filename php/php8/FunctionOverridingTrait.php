<?php

trait SampleTrait
{
    public abstract function simpleFunction(string $name): string;
}

class SampleClass {
    use SampleTrait;

    public function simpleFunction(int $name): string
    {

    }
}