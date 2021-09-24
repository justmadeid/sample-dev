<?php
session_start();

if ($_SESSION['login'] !== true){
    header('Location: member.php');
    exit();
}

$say = "Hello ". $_SESSION['username'];

?>

<html>
<header>
    <title>

    </title>
</header>
<body>

<h1><?= $say ?></h1>
<a href="/session/logout.php">Logout</a>

</body>
</html>

