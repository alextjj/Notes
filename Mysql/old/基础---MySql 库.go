
<1:�鿴�������ݿ�>
	<command:1>	show databases;
	<command:2> show schemas;
	
<2:�������ݿ�>
	<command:1> create database ���ݿ���;
	<command:2> create schema	���ݿ���;
	
<3:ɾ�����ݿ�>
	<command:1> drop database ���ݿ���;
	

	
	
< --�ؼ���-- >	
< if exists >  
	<��:> drop database if exists ���ݿ���;  if exists �����ж� ������ھ�XXX �����ھ���ʾ������Ϣ(��������),	������ show warning  �鿴������Ϣ
	
< distinct > ȥ���ظ� distinct������ڿ�ͷ
	<��:>select distinct roleId, rank from ranktable;<ȥ���ظ����ݵĲ�ѯ>
	
< primary key >	����
< engine >		���ݿ�����
< charset >     �ַ�����
< AUTO_INCREMENT >������