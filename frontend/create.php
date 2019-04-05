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
                <div class='firstline'><span class='title'> <input type="text" name="title" placeholder="Title"> </span> - <span class='date'> <input type="date" name="date"> </span></div>
                <span class='category'><input type="text" name="category" placeholder="Category"></span>
                <div class='description'><textarea name="description">Description</textarea></div>
                </div>
                <input type="submit" value="Create">
            </form>
        </div>
    </body>
</html>