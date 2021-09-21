<?php

require_once "../Model/TodoList.php";
require_once "../View/ViewRemoveTodoList.php";
require_once "../BusinessLogic/AddTodoList.php";
require_once "../BusinessLogic/ShowTodoList.php";

addTodoList("Made");
addTodoList("Garda");
addTodoList("Setiawan");
addTodoList("Sanskara");
addTodoList("JuanDeJanon");

showTodoList();

viewRemoveTodoList();

showTodoList();