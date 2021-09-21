<?php

require_once "../Model/TodoList.php";
require_once "../BusinessLogic/AddTodoList.php";
require_once "../BusinessLogic/ShowTodoList.php";
require_once "../BusinessLogic/RemoveTodoList.php";

addTodoList("Made");
addTodoList("Garda");
addTodoList("Setiawan");
addTodoList("Sanskara");
addTodoList("JuanDeJanon");

showTodoList();

removeTodoList(3);
showTodoList();

removeTodoList(2);
showTodoList();

$success = removeTodoList(4);
var_dump($success);

showTodoList();