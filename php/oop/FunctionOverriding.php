<?php

require_once "data/Manager.php";

$manager = new Manager();
$manager->name = "Made";
$manager->sayHello("G");

$vp = new VicePresident();
$vp->name = "Sanskara";
$vp->sayHello("A");