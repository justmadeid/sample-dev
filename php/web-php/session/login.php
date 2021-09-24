<?php
session_start();

if ($_SESSION['login'] == true){
    header('Location: member.php');
    exit();
}

if ($_SERVER['REQUEST_METHOD'] == "POST"){
    if ($_POST['username'] == "made" AND $_POST['password'] == "made"){
        $_SESSION['login'] = true;
        $_SESSION['username'] = "made";
        header('Location: member.php');
        exit();
    }else{
        $error = "Faied Login";
    }
}

?>

<html>
<body>
<?php if($error != ""){ ?>
    <h2><?= $error ?></h2>
<?php } ?>
<h1>Login</h1>
<form action="/session/login.php" method="POST">
    <label>Username :
        <input type="text" name="username">
    </label>
    <br/>
    <label>Password :
        <input type="password" name="password">
    </label>
    <br/>
    <input type="submit" value="Login">
</form>
</body>
</html>
