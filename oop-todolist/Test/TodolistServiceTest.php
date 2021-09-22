<?php

require_once __DIR__."/../Entity/TodoList.php";
require_once __DIR__."/../Repository/TodolistRepository.php";
require_once __DIR__."/../Service/TodolistService.php";

use Entity\TodoList;
use Service\TodolistServiceImpl;
use Repository\TodolistRepositoryImpl;

function testShowTodolist():void{

    $todolistRepository = new TodolistRepositoryImpl();
    $todolistRepository->todolist[1] = new TodoList("AAAA");
    $todolistRepository->todolist[2] = new TodoList("BBBB");
    $todolistRepository->todolist[3] = new TodoList("CCCC");

    $todolistService = new TodolistServiceImpl($todolistRepository);

    $todolistService->showTodolist();

}

function testAddTodolist():void{

    $todolistRepository = new TodolistRepositoryImpl();

    $todolistService = new TodolistServiceImpl($todolistRepository);
    $todolistService->addTodolist("BBB");
    $todolistService->addTodolist("AAA");

    $todolistService->showTodolist();
}

function testRemoveTodolist():void{

    $todolistRepository = new TodolistRepositoryImpl();

    $todolistService = new TodolistServiceImpl($todolistRepository);
    $todolistService->addTodolist("BBB");
    $todolistService->addTodolist("AAA");

    $todolistService->showTodolist();

    $todolistService->removeTodolist(2);
    $todolistService->showTodolist();

    $todolistService->removeTodolist(4);
    $todolistService->showTodolist();

    $todolistService->removeTodolist(1);
    $todolistService->showTodolist();
}

testRemoveTodolist();