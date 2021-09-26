<?php

namespace Jstmade\Mvc\Controller;

use Jstmade\Mvc\App\View;


class HomeController
{
    function index(): void
    {
        $model = [
            "title" => "First MVC",
            "content" => "LOREM IPSUM DOLOR"
        ];

        View::render('Home/index', $model);

    }

    function hello(): void
    {
        echo "HomeController.hello()";
    }

    function world(): void
    {
        echo "HomeController.world()";
    }

    function about(): void
    {
        echo "made";
    }

    function login(): void
    {
        $request = [
            "username" => $_POST['username'],
            "password" => $_POST['password'],
        ];

        $user = [

        ];

        $reponse = [
            "message" => "login success"
        ];
        //ke view
    }
}