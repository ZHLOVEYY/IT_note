## Mysql 
``` bash
mysql -u username -p password  -h host  -P  port    
```
 
 #### basic operation 
- View the database               `show databases;`  
 - Create a database               `create database xxx;`  
 - use database                 `USE xxx; `
 - Query the database currently in use  `SELECT DATABASE();
 - Delete the database                     `drop database xxx;``
 - Check the creation method of the table     `show create database db_name;`
demo of creating a table：
``` sql
 CREATE TABLE employee(
            id int primary key auto_increment ,
            name varchar(20),
            gender bit default 1,
            birthday date,
            department varchar(20),
            salary double(8,2) unsigned,
            resume text
          )character set=utf8;  
```
包括设置default    数据类型还有decimal  Enun  等

- show all the tables                    `show tables;`
- Check the creation method of a table                  `show create table xxx ;`
- View the structure of a certain table (specific and detailed)  `desc xxx ；`
- To delete a table with foreign keys: Delete the referenced table first and then the current table
#### Insert 
Insert a record
``` sql
INSERT employee (name,gender,birthday,salary,department) VALUES  ("yuan",1,"1985-12-12",8000,"教学部"), ("alvin",1,"1987-08-08",5000,"保安部")；
```
Additional addition methods
``` sql
INSERT INTO employee  SET name = '张三',  age = 25,  gender = '男';
```

#### Update
###### alter
- alter is mainly used for modifying the structure
operation：
- `alter table xxx add gender varchar(32) not null after age;`  
- change the field type: `ALTER TABLE employee MODIFY age VARCHAR(10) NOT NULL;`
- change the field name: `ALTER TABLE employee CHANGE age employee_age INT NOT NULL;`
- Delete the field:  `ALTER TABLE employee DROP employee_age;`
- change the table name :  `ALTER TABLE employee RENAME TO staff;`

###### update
- Update is used for modifying data 
``` sql
UPDATE employee SET salary = salary + 1000 , dep = 'Management';（where xx = xx）
update emp set salary=20000  where age > 30;
```
#### delete
- Delete the employee with the lowest salary in the Employee table: `DELETE FROM employee ORDER BY salary ASC LIMIT 1;`       （where ， ooerder by ， limit  ）
- Generally, pay attention to query the data first and then delete it.  `SELECT * FROM author WHERE name = 'bob';`
- Strings need to be wrapped in quotation marks：`delete from author where name = ‘bob’;`
- Clear the data：`DELETE FROM employee;`

####  Read
用于测试数据：
``` sql
CREATE TABLE emp(
    id       INT PRIMARY KEY AUTO_INCREMENT,
    name     VARCHAR(20),
    gender   ENUM("male","female","other"),
    age      TINYINT,
    dep      VARCHAR(20),
    city     VARCHAR(20),
   salary    DOUBLE(7,2)
)character set=utf8;


INSERT INTO emp (name,gender,age,dep,city,salary) VALUES
                ("yuan","male",24,"教学部","河北省",8000),
                ("eric","male",34,"销售部","山东省",8000),
                ("rain","male",28,"销售部","山东省",10000),
                ("alvin","female",22,"教学部","北京",9000),
                ("George", "male",24,"教学部","河北省",6000),
                ("danae", "male",32,"运营部","北京",12000),
                ("Sera", "male",38,"运营部","河北省",7000),
                ("Echo", "male",19,"运营部","河北省",9000),
                ("Abel", "female",24,"销售部","北京",9000);
```

###### select
- use of select 
``` sql
SELECT *|field1, field2 ...
FROM tab_name
WHERE 条件
GROUP BY field
HAVING 筛选
ORDER BY field
LIMIT 限制条数;
```

- Query the part of > <：`SELECT * FROM employee WHERE salary > 9000;`
		`SELECT * FROM emp WHERE dep="教学部" AND gender="male";`
- 查询平均工资大于 9000 的部门： having是通过分组后的进行筛选：
	``` sql
	SELECT dep, AVG(salary) AS avg_salary
	FROM emp
	GROUP BY dep
	HAVING avg_salary > 9000;
	```

- Query in descending order：`SELECT * FROM employee ORDER BY salary DESC;   降序排列（从高到低）`      `SELECT * FROM emp ORDER BY salary;`   
- Restrict the query： `SELECT * FROM employee LIMIT 3;`
- limit 2,2  :Take 2 items starting from index 2, so take 3 and 4

Range query
- between 80 and 100 

group by  分组语句
- 常用统计函数 max() avg()  sum() min() count()
- 按照部门划分然后查看平均工资：`select dep,avg(salary) from emp group by dep;`
- 搜索分组后将name 聚合展示`SELECT dep, GROUP_CONCAT(name) FROM emp GROUP BY dep;`

补充：
- 查总人数SELECT COUNT(\*) 员工总人数 FROM emp;
- distinct 去重： SELECT distinct salary from emp order by salary;

#### 关联查询
###### 多对多
用于测试数据： （含有一对多）
``` sql
--作者表
CREATE TABLE book
   (
       id     INT PRIMARY KEY AUTO_INCREMENT,
       title  VARCHAR(32),
       price  DOUBLE(5, 2),
       pub_id INT NOT NULL
       --这里没有设置外键不过其实就是“外键”
   ) ENGINE = INNODB
     CHARSET = utf8;

--发布
CREATE TABLE publisher
(
   id    INT PRIMARY KEY AUTO_INCREMENT,
   name  VARCHAR(32),
   email VARCHAR(32),
   addr  VARCHAR(32)
) ENGINE = INNODB 
 CHARSET = utf8;
   
INSERT INTO book(title, price, pub_id)
VALUES ('西游记', 15, 1),
	  ('三国演义', 45, 2),
	  ('红楼梦', 66, 3),
	  ('水浒传', 21, 2),
	  ('红与黑', 67, 3),
	  ('乱世佳人', 44, 6),
	  ('飘', 56, 1),
	  ('放风筝的人', 78, 3);
   
INSERT INTO publisher(id, name, email, addr)
VALUES (1, '清华出版社', "123", "bj"),
	  (2, '北大出版社', "234", "bj"),
	  (3, '机械工业出版社', "345", "nj"),
	  (4, '邮电出版社', "456", "nj"),
	  (5, '电子工业出版社', "567", "bj"),
	  (6, '人民大学出版社', "678", "bj");
	
CREATE TABLE author
(
   id   INT PRIMARY KEY AUTO_INCREMENT,
   NAME VARCHAR(32) NOT NULL
) ENGINE = INNODB
 CHARSET = utf8;
   
-- 关系表 用于建立多对多  建立书本到作者的多对多
CREATE TABLE book2author
(
   id        INT NOT NULL UNIQUE AUTO_INCREMENT,
   author_id INT NOT NULL,
   book_id   INT NOT NULL
) ENGINE = INNODB
 CHARSET = utf8;
   
INSERT INTO author(NAME)
VALUES ('yuan'),
	  ('rain'),
	  ('alvin'),
	  ('eric');
   
-- 插入关系数据  用于建立多对多
INSERT INTO book2author(author_id, book_id)
VALUES (1, 1),
	  (1, 2),
	  (2, 1),
	  (3, 3),
	  (3, 4),
	  (1, 3);
   
-- 作者详细数据
CREATE TABLE authorDetail
(
   id        INT PRIMARY KEY AUTO_INCREMENT, --id 自动增长
   tel       VARCHAR(32),
   addr      VARCHAR(32),
   author_id INT NOT NULL unique -- 也可以给author添加一个关联字段： alter table author add authorDetail_id INT NOT NULL
   ) ENGINE = INNODB CHARSET = utf8;
   
INSERT INTO authorDetail(tel, addr, author_id)
VALUES ("110", "北京", 1),
	  ("911", "成都", 2),
	  ("119", "上海", 3),
	  ("111", "广州", 4);
   
   -- 一对多的关联查询  一个发行商可以对应多本书，在多的一侧：书的表的简历外键
   select *
   from book,
        publisher
   where book.pub_id = publisher.id;
        
   -- 关联然后设置条件查询
   select *
   from book   
		inner join publisher on book.pub_id = publisher.id
   where publisher.name = "清华出版社";  
   
   -- 多对多的关联查询   利用中间表简历关联
   select *
   from book
		inner join book2author on book.id = book2author.book_id
		inner join author on book2author.author_id = author.id
   
   -- 选出yuan出版的书籍的名字（book.title）
   select book.title
   from book
		inner join book2author on book.id = book2author.book_id
		inner join author on book2author.author_id = author.id
   where author.NAME = "yuan";
   
   -- 每一本书的作者的个数   优先book.title分组后  查看统计结果（每本书的作者数）
   select book.title,count(*) as author_count
   from book
            inner join book2author on book.id = book2author.book_id
            inner join author on book2author.author_id = author.id
   group by book.title;
   
   -- left join 以left join后面的表名为基础在左侧添加
   select book.title, publisher.name
   from book
            left join publisher on book.pub_id = publisher.id
   
   -- yuan出版的每一本书的名字和对应出版社的名字
   select book.title,publisher.name
   from book
            inner join publisher on book.pub_id=publisher.id --先拼接这个表
            inner join book2author on book.id = book2author.book_id
            inner join author on book2author.author_id = author.id
   where author.NAME="yuan";
   
   --delete from publisher where id=6; 一般不会一起删除，没有设置级联
   -- 添加外键约束   一对多     
   -- show create table 可以看到外键情况
   ALTER TABLE book add CONSTRAINT dep_fk
       FOREIGN KEY(pub_id) REFERENCES publisher(id) ON DELETE CASCADE;
   -- 会删除
   delete from publisher where id=3

	-- 添加外键约束
	ALTER TABLE <数据表名> （ADD CONSTRAINT <外键名>）
	FOREIGN KEY(<列名>) REFERENCES <主表名> (<列名>);
		
	-- 删除外键约束
	ALTER TABLE <表名> DROP FOREIGN KEY <外键约束名>; -- 注意有约束名字的
		ALTER TABLE book DROP FOREIGN KEY dep_fk;
	--查看约束
	show index from book;
	drop index 外键约束名 on<表名>; -- 同时将索引删除
	drop index dep_fk on book;
	show create table book;
   
```


- 一对多：一个出版社可以对应多本书，这样理解也很顺
	- 接着在多的地方建立关联字段 pubish_id进行关联
		``` sql
		select book.title,publisher.email from book inner join publisher on book.pub_id = publisher.id where book.title = "xxx"
		--即使没有外键约束，这条查询仍然可以正常运行 不过外键会保持一致性
		```

###### 一对一
一对一的关系表其实用的并不多，一对一关联案例：
``` sql
--上面创建过
CREATE TABLE author(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(32) NOT NULL
) ENGINE=INNODB CHARSET=utf8;

CREATE TABLE authorDetail(
    id INT PRIMARY KEY AUTO_INCREMENT,
    tel VARCHAR(32),
    addr VARCHAR(32),
    author_id INT NOT NULL UNIQUE,
    FOREIGN KEY (author_id) REFERENCES author(id)
) ENGINE=INNODB CHARSET=utf8;

-- 插入 author 数据,上面插入过

-- 插入 authorDetail 数据
INSERT INTO authorDetail (tel, addr, author_id) VALUES ('123456', 'Beijing', 1);
INSERT INTO authorDetail (tel, addr, author_id) VALUES ('654321', 'Shanghai', 2);

-- 查询 author 及其 authorDetail （利用别名）
SELECT a.id, a.name, ad.tel, ad.addr
FROM author a
JOIN authorDetail ad ON a.id = ad.author_id;
```

笛卡尔积查询：​   
`select * from author,authordetail;`  ->  全部都会两两组合然后呈现

-  join  （应该不需要利用foreign key 也可以实现
	- 利用条件从笛卡尔积结果中筛选出正确结果
		```sql
		select * from author,authordetail where author_id = author.id;       
		```
	- left join      左边的表的为主都保存 ，（也会显示 NULL
		```sql
		select * from authordetail left join author on author_id = author.id;
		```
	- right join       右边的表为主都保存，（也会显示 NULL
	- inner join     拼接能拼上的


flowchart LR
Start --> Stop
#### 约束
###### 主键
- 增加和删除主键
	``` sql
	 alter table t2 add primary key(id);
	 alter table t2 drop primary key;
	 alter table t2 add primary key(name);
	```

- 联合主键 比如学生1选课1 只能出现一次，联合主键
	- 如果用学生编号做主键，那么一个学生就只能选择一门课程。如果用课程编号做主键，那么一门课程只能有一个学生来选
	``` sql
	-- ①创建时：
	create table sc (
		studentid int,
		courseid int,
		score int,
	primary key (studentid,courseid)
	);        
	-- ②修改时：
	alter table sc add primary key (字段1,字段2,字段3); --修改为三个字段作为主键
			```
- 主键自增长约束
	``` sql
	CREATE TABLE t1(
		id INT(4) PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(25) 
		);
	```
###### unique约束等
- unique 约束关系 唯一约束 （确保表中某一列或一组列中的数据具有唯一性）
	```sql
	CREATE TABLE user(
		id INT(11) PRIMARY KEY,
		name VARCHAR(22) UNIQUE
		);  
	INSERT user (id,name) values (1,"yuan"),(2,"rain");   
	INSERT user (id,name) values (3,"alvin"),(4,"alvin");
	-- ERROR 1062 (23000): Duplicate entry 'alvin' for key 'name'
	```
	对应修改唯一约束
	```sql
	ALTER TABLE <表名> DROP INDEX <唯一约束名>;  -- 删除唯一约束
	ALTER TABLE <数据表名> ADD CONSTRAINT <唯一约束名> UNIQUE(<列名>);  -- 添加唯一约束
	ALTER TABLE user  DROP INDEX name;
	ALTER TABLE user ADD CONSTRAINT NAME_INDEX UNIQUE(name);
		```
- 非空约束 not null 也可以修改和设置
- 默认值default 约束

###### 外键约束
- 外键约束foreign key 建立约束关系！！！ （见上面demo）
