<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>Go Web 表单输入</title>
    </head>
    <body>
        <form action="http://127.0.0.1:9090/login" method="post">
            用户名：<input type="text" name="username"/>
            密  码：<input type="password" name="password"/>
            <input type="submit" value="Login" />
        </form>
    </body>
</html>