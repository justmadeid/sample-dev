<?php

//$path = '/index';
//
//if (isset($_SERVER['PATH_INFO'])) {
//    $path = $_SERVER['PATH_INFO'];
//}
//
//require __DIR__ . '/../app/View' . $path . '.php';

require_once __DIR__ . '/../vendor/autoload.php';

use Jstmade\Mvc\App\Router;
use Jstmade\Mvc\Controller\HomeController;
use Jstmade\Mvc\Controller\ProductController;
use Jstmade\Mvc\Middleware\AuthMiddleware;

Router::add('GET', '/products/([0-9a-zA-Z]*)/categories/([0-9a-zA-Z]*)', ProductController::class, 'categories');

Router::add('GET', '/', HomeController::class, 'index');
Router::add('GET', '/hello', HomeController::class, 'hello', [AuthMiddleware::class]);
Router::add('GET', '/world', HomeController::class, 'world', [AuthMiddleware::class]);
Router::add('GET', '/about', HomeController::class, 'about', [AuthMiddleware::class]);

Router::run();