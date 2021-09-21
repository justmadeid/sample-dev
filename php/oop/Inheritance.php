<?php

require_once "data/Manager.php";

$manager = new Manager();
$manager->name = "Made";
$manager->sayHello("Sanskara");

$vp = new Manager();
$vp->name = "Garda";
$vp->sayHello("Setiawan");