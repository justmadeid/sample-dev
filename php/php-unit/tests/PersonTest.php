<?php

namespace Jstmade\Test;

use PHPUnit\Framework\TestCase;

class PersonTest extends TestCase
{

    private Person $person;

    protected function setUp(): void
    {
//        $this->person = new Person("made");
    }

    /**
     * @before
     */
    public function createPerson()
    {
        $this->person = new Person("made");
    }

    public function testSuccess()
    {
        self::assertEquals("Hello Garda, my name is made", $this->person->sayHello("Garda"));
    }

    public function testException()
    {
        $this->expectException(\Exception::class);
        $this->person->sayHello(null);
    }

    public function testOutput()
    {
        $this->expectOutputString("Good Bye Made" . PHP_EOL);
        $this->person->sayGoodbye("Made");
    }

}

