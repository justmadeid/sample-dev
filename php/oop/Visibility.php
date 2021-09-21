<?php

require_once "data/Product.php";

$product = new Product("Apple", 20000);

echo $product->getName().PHP_EOL;
echo $product->getPrice().PHP_EOL;

$dummy = new ProductDummy("Melon",30000);
echo $dummy->getName();