<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>防止多次提交</title>
</head>
<body>
    <form action="http://127.0.0.1:8080/login" method="post">
        用户名:<input type="text" name="username"><br>
        密&nbsp&nbsp&nbsp码:<input type="password" name="password"><br>
        <input type="hidden" name="token" value="{{.}}">
        <input type="submit" value="登陆">
    </form>
</body>
</html>