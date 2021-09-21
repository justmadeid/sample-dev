<?php
function createName()
{
    $name = "Made"; // local scope
}

createName();
echo $name . PHP_EOL;