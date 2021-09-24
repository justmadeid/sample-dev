<?php

header('Application: Belajar PHP Web');
header('Author: Made G');

$client = $_SERVER['HTTP_CLIENT_NAME'];

echo "Hello $client";