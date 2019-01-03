<?php
set_error_handler(function ($code, $error) {
    die("$error\n");
}, E_ALL | E_NOTICE | E_WARNING);

array_shift($argv);

if (!$argv) {
    trigger_error('i need at least 1 to pick');
}

print($argv[array_rand($argv)] . "\n");
