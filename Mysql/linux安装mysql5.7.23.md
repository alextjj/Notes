linux上安装mysql5.7 

1:下载tar包，这里使用wget从官网下载
	wget https://dev.mysql.com/get/Downloads/MySQL-5.7/mysql-5.7.23-linux-glibc2.12-x86_64.tar.gz
	

2: 解压tar包
	tar -zxvf mysql-5.7.20-linux-glibc2.12-x86_64.tar.gz.tar.gz
	mv mysql-5.7.20-linux-glibc2.12-x86_64/ /usr/local/mysql	//移动目录并重命名为mysql	
  	cd /usr/local/

	groupadd mysql  			//创建用户组mysql
	useradd -r -g mysql mysql 	//-r参数表示mysql用户是系统用户，不可用于登录系统，创建用户mysql并将其添加到用户组mysql中
  	chown -R mysql mysql/		//权限组　　
	chgrp -R mysql mysql/		//权限
	
	mkdir data			//datadir	
	chmod 777 data			//datadir用户权限	
	chown mysql:mysql data		//datadir用户组权限

	
3、创建配置文件

　	vi /etc/my.cnf

　	复制以下内容，wq保存，可以继续添加你需要的配置：

	[client]
	port = 3306
	socket = /tmp/mysql.sock	
	[mysqld]
	character_set_server=utf8
	init_connect='SET NAMES utf8'
	basedir=/usr/local/mysql
	datadir=/usr/local/mysql/data
	socket=/tmp/mysql.sock
	log-error=/var/log/mysqld.log
	pid-file=/var/run/mysqld/mysqld.pid
	#不区分大小写
	lower_case_table_names = 1
	
	sql_mode=STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION
	
	max_connections=5000			  
	default-time_zone = '+8:00'	

	
4、初始化数据库
	#先安装一下这个东东，要不然初始化有可能会报错	
	yum install libaio 或者 apt-get install libail
	
	#手动编辑一下日志文件，什么也不用写，直接保存退出	
	cd /var/log/	
	vi mysqld.log  //用:wq保存	
	chmod 777 mysqld.log
	chown mysql:mysql mysqld.log
	
	cd /var/run/	
	mkdir mysqld	
	chmod 777 mysqld	
	chown mysql:mysql mysqld
	
	cd mysqld	
	vi mysqld.pid	//:wq保存	
	chmod 777 mysqld.pid	
	chown mysql:mysql mysqld.pid 
	
	/usr/local/mysql/bin/mysqld --initialize --user=mysql --basedir=/usr/local/mysql --datadir=/usr/local/mysql/data --lc_messages_dir=/usr/local/mysql/share --lc_messages=en_US

	
5、查看初始密码
	cat /var/log/mysqld.log	
	执行后关注最后一点：root@localhost: 这里就是初始密码

	
6、启动服务，进入mysql，修改初始密码，运行远程连接（这里执行完后，密码将变成：123456）

	/usr/local/mysql/support-files/mysql.server start
	
	/usr/local/mysql/bin/mysql -uroot -p 你在上面看到的初始密码 //链接dbserver

	// 以下是进入数据库之后的sql语句
	use mysql;
	UPDATE `mysql`.`user` SET `Host`='%', `User`='root', `Select_priv`='Y', `Insert_priv`='Y', 
	`Update_priv`='Y', `Delete_priv`='Y', `Create_priv`='Y', `Drop_priv`='Y', `Reload_priv`='Y', 
	`Shutdown_priv`='Y', `Process_priv`='Y', `File_priv`='Y', `Grant_priv`='Y', 
	`References_priv`='Y', `Index_priv`='Y', `Alter_priv`='Y', `Show_db_priv`='Y', 
	`Super_priv`='Y', `Create_tmp_table_priv`='Y', `Lock_tables_priv`='Y', `Execute_priv`='Y',
	`Repl_slave_priv`='Y', `Repl_client_priv`='Y', `Create_view_priv`='Y', `Show_view_priv`='Y',
	`Create_routine_priv`='Y', `Alter_routine_priv`='Y', `Create_user_priv`='Y', `Event_priv`='Y',
	`Trigger_priv`='Y', `Create_tablespace_priv`='Y', `ssl_type`='', `ssl_cipher`='', `x509_issuer`='',
	`x509_subject`='', `max_questions`='0', `max_updates`='0', `max_connections`='0', 
	`max_user_connections`='0', `plugin`='mysql_native_password',
	`authentication_string`='*6BB4837EB74329105EE4568DDA7DC67ED2CA2AD9', 
	`password_expired`='N', `password_last_changed`='2017-11-20 12:41:07', 
	`password_lifetime`=NULL, `account_locked`='N' WHERE  (`User`='root');

	flush privileges;

	#如果提示必须要修改密码才可以进行操作的话则执行下面操作
	set password=password('新密码');
	flush privileges;
	UPDATE `mysql`.`user` SET `Host` = '%',  `User` = 'root'  WHERE (`Host` = 'localhost') AND (`User` = 'root');

	
7、开机自启;
	cd /usr/local/mysql/support-files;	
	cp mysql.server /etc/init.d/mysqld;	
	chkconfig --add mysqld   //个人测试时chkconfig 没能安装成功;
	
	
	
8、使用service mysqld命令启动/停止服务
	例如我的mysql：启动/停止/暂停：	
	su - mysql	
	service mysqld start/stop/restart
