<?php

require_once "../View/ViewShowTodoList.php";
require_once "../BusinessLogic/AddTodoList.php";

addTodoList("Made");
addTodoList("Garda");
addTodoList("Setiawan");
addTodoList("Sanskara");
addTodoList("JuanDeJanon");

viewShowTodoList();