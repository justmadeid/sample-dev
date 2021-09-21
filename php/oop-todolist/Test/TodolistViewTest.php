<?php
require_once __DIR__ . "/../Entity/TodoList.php";
require_once __DIR__ . "/../Repository/TodolistRepository.php";
require_once __DIR__ . "/../Service/TodolistService.php";
require_once __DIR__ . "/../View/TodolistView.php";
require_once __DIR__ . "/../Helper/InputHelper.php";

use Entity\TodoList;
use Repository\TodolistRepositoryImpl;
use Service\TodolistServiceImpl;
use View\TodolistView;

function testViewShowTodolist(): void
{
    $todolistRepository = new TodolistRepositoryImpl();
    $todolistService = new TodolistServiceImpl($todolistRepository);
    $todolistView = new TodolistView($todolistService);

    $todolistService->addTodolist("Made");
    $todolistService->addTodolist("Made1");
    $todolistService->addTodolist("Made2");

    $todolistView->showTodolist();
}

function testViewAddTodolist(): void
{
    $todolistRepository = new TodolistRepositoryImpl();
    $todolistService = new TodolistServiceImpl($todolistRepository);
    $todolistView = new TodolistView($todolistService);

    $todolistService->addTodolist("Made");
    $todolistService->addTodolist("Made1");
    $todolistService->addTodolist("Made2");

    $todolistService->showTodolist();

    $todolistView->addTodolist();
    $todolistService->showTodolist();

    $todolistView->addTodolist();
    $todolistService->showTodolist();
}

function testViewRemoveTodolist(): void
{
    $todolistRepository = new TodolistRepositoryImpl();
    $todolistService = new TodolistServiceImpl($todolistRepository);
    $todolistView = new TodolistView($todolistService);

    $todolistService->addTodolist("Made");
    $todolistService->addTodolist("Made1");
    $todolistService->addTodolist("Made2");

    $todolistService->showTodolist();

    $todolistView->removeTodolist();
    $todolistService->showTodolist();
}

testViewRemoveTodolist();