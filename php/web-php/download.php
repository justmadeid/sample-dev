<?php

if (isset($_GET['key']) && $_GET['key'] == 'RAHASIA') {
    header('Content-Disposition: attachment; filename="wozniak.jpg"');
    header('Content-Type: image/jpeg');
    readfile(__DIR__ . '/file/wozniak.jpg');
    exit();
} else {
    echo "Invalid Key";
}