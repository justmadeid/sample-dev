<?php

namespace Jstmade\Test;

use PHPUnit\Framework\TestCase;

class MathTest extends TestCase
{
    public function testManual()
    {
        $this->assertEquals(10, Math::sum([5, 5]));
        $this->assertEquals(20, Math::sum([2, 18]));
    }

    /**
     * @dataProvider mathSumData
     */
    public function testDataProvider(array $values, int $expected)
    {
        self::assertEquals($expected, Math::sum($values));
    }

    public function mathSumData(): array
    {
        return [
            [[5, 5], 10],
            [[4, 4, 4, 4, 4], 20],
            [[], 0],
        ];
    }

    /**
     * @testWith [[5,5], 10]
     *[[4,4,4,4,4], 20]
     *[[], 0]
     */
    public function testWith(array $values, int $expected)
    {
        self::assertEquals($expected, Math::sum($values));
    }

}