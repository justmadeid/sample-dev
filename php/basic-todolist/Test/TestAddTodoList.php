<?php

require_once "../Model/TodoList.php";
require_once "../BusinessLogic/AddTodoList.php";

addTodoList("Made");
addTodoList("Garda");
addTodoList("Setiawan");

var_dump($todoList);