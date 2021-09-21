<?php

require_once "data/Programmer.php";

$company = new Company();
$company->programmer = new Programmer("Made");
var_dump($company);

$company->programmer = new BackendProgrammer("Garda");
var_dump($company);

$company->programmer = new FrontendProgrammer("Setiawan");
var_dump($company);

sayHelloProgrammer(new Programmer("Made"));
sayHelloProgrammer(new BackendProgrammer("Made"));
sayHelloProgrammer(new FrontendProgrammer("Made"));