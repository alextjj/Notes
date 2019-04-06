
<1:查看所有数据库>
	<command:1>	show databases;
	<command:2> show schemas;
	
<2:创建数据库>
	<command:1> create database 数据库名;
	<command:2> create schema	数据库名;
	
<3:删除数据库>
	<command:1> drop database 数据库名;
	

	
	
< --关键词-- >	
< if exists >  
	<例:> drop database if exists 数据库名;  if exists 条件判断 如果存在就XXX 不存在就显示警告信息(不报错误),	可以用 show warning  查看警告信息
	
< distinct > 去除重复 distinct必须放在开头
	<例:>select distinct roleId, rank from ranktable;<去除重复数据的查询>
	
< primary key >	主键
< engine >		数据库引擎
< charset >     字符编码
< AUTO_INCREMENT >自增长