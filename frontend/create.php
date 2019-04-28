<!DOCTYPE HTML>
<html>
    <head>
        <title>Create Business Announcement</title>
        <link rel="stylesheet" type="text/css" href="css\main.css">
    </head>
    <body>
        <?php require ('menu_bar.php'); ?>

        <div class="contentarea">
        <h2>Create a Business Announcement</h2>
            <form method="POST" action="/wassivu">
                <div class='bnsannouncement'>
                <div class='firstline'><span class='title'> <input type="text" name="title" placeholder="Title" required> </span> - <span class='date'> <input type="date" name="date" required> - <input type="email" name="email" placeholder="email" required> - <input type="text" name="fullname" placeholder="Full name" required></div>
                <span class='category'>
                <select name="category">
                    <option value="buying">Buying</option>
                    <option value="selling">Selling</option>
                </select>
                </span>
                <div class='description'><textarea name="description">Description</textarea></div>
                </div>
                <input type="submit" value="Create">
            </form>
        </div>
    </body>
</html>