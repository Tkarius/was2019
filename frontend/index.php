<!DOCTYPE HTML>
<html>
    <head>
        <title>View Business Announcements</title>
        <link rel="stylesheet" type="text/css" href="css\main.css">
    </head>
    <body>
        <?php require ('menu_bar.php'); ?>
        <div class="contentarea">
            <?php 
                for($i = 0; $i<3; $i++){
                    echo 
                    "<div class='bnsannouncement'>
                    <div class='firstline'><span class='title'> Title </span> - <span class='date'> 00.00.0000 </span></div>
                    <span class='category'>Category</span>
                    <div class='description'>Description</div>
                    </div>";
                }
            ?>
        </div>
    </body>
</html>