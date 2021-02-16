#!/bin/bash

sudo apt-get update
sudo apt install -y apache2 apache2-utils 
sudo systemctl enable apache2
sudo systemctl start apache2
sudo apt install -y  mysql-client mysql-server
sudo mysql_secure_installation --y
sudo apt install -y php php-mysql libapache2-mod-php php-cli php-cgi php-gd
sudo a2enmod rewrite
sudo systemctl restart apache2
wget -c http://wordpress.org/latest.tar.gz
tar -xzvf latest.tar.gz
sudo rsync -av wordpress/* /var/www/html/
sudo chown -R www-data:www-data /var/www/html/
sudo chmod -R 755 /var/www/html/
sudo rm /var/www/html/index.html
sudo mysql -u root -p <<EOF
CREATE DATABASE wp_testbase;
CREATE USER 'wp_user'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON wp_testbase.* TO 'wp_user'@'localhost';
FLUSH PRIVILEGES;
EXIT;
EOF
