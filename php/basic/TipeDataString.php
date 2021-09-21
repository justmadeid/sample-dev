<?php

echo 'Name : ';
echo 'Made Garda Setiawan';
echo "\n";

echo "Name : ";
echo "Made\t Garda\t Setiawan\n";

echo <<<Made
Selamat belajar PHP
sekarang, kita belajar tipe data string
ini adalah cara ke-3 membuat string
bisa menggunakan heredoc

Made;

echo <<<'Made'
Selamat belajar PHP
sekarang, kita belajar tipe data string
ini adalah cara ke-3 membuat string
bisa menggunakan heredoc
Made;
