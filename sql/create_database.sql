create database gamereviews default character set utf8;

grant all privileges on gamereviews.* to 'gamereviews'@'%' identified by 'gamereviews';
grant all privileges on gamereviews.* to 'gamereviews'@'localhost' identified by 'gamereviews';