https://ithelp.ithome.com.tw/articles/10234657

記得安裝mysql

建立一組mysql使用者和使用者密碼
。之後用程式登入時使用這組密碼
```
CREATE USER 'demo'@'%' IDENTIFIED BY 'demo123';
GRANT ALL PRIVILEGES ON demo.* TO 'demo'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;
```

1. GRANT ALL PRIVILEGES
GRANT: 用於授予權限給用戶。
ALL PRIVILEGES: 這是一個權限集合，授予用戶所有的權限。這包括了對資料庫的所有操作權限，如 SELECT、INSERT、UPDATE、DELETE、CREATE、DROP、GRANT 等。基本上，這個選項允許用戶進行對應資料庫的所有操作。
2. ON demo.*
ON: 指定授予權限的資料庫和表。
demo.*: 表示授權的範圍。demo 是資料庫名，* 是通配符，表示所有在 demo 資料庫中的表。也就是說，授予用戶對 demo 資料庫中所有表的權限。
3. TO 'demo'@'%'
TO: 指定授權的用戶和主機。
'demo'@'%':
'demo': 用戶名。
'%': 主機名的通配符，表示該用戶可以從任何主機連接到 MySQL 伺服器。這意味著這個用戶可以從任何 IP 地址進行連接。
4. WITH GRANT OPTION
WITH GRANT OPTION: 這個選項授予用戶的權限也包括授權其他用戶的權限。簡而言之，這意味著用戶 demo 除了可以執行所有操作外，還可以將這些權限授予其他用戶。
總結
這條 GRANT 語句的目的是：

授予用戶 demo 在資料庫 demo 上的所有操作權限。
允許用戶 demo 從任何主機（通過 '%'）連接到 MySQL 伺服器。
允許用戶 demo 將這些權限再授予其他用戶。