<?php
require_once "../View/ViewAddTodoList.php";
require_once "../BusinessLogic/ShowTodoList.php";
require_once "../BusinessLogic/AddTodoList.php";

addTodoList("Made");
addTodoList("Garda");
addTodoList("Setiawan");

viewAddTodoList();

showTodoList();

viewAddTodoList();

showTodoList();