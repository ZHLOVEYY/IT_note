###### 1757
注意需要提取的是product_id

###### 584
如果查找的时候注意referee_id IS NULL;  NULL在比较大时候会认为是未知

###### 595 
注意select后面需要,分割

###### 1148
distinct 表示唯一    重命名列名用 as xx
正常select是表中的顺序，需要选出后order by xx 进行排序  （如果降序加desc）

###### 1683
运用：CHAR_LENGTH()：计算字符数（适用于多字节字符，如UTF-8）。
LENGTH()：计算字节数（可能和字符数不同，取决于编码）。


###### 1378
left就是以左边的为主   先执行from，然后left join （看on条件）最后select，所以
多的select要加,
如果用代指减少书写：
``` sql
SELECT 
    eu.unique_id, 
    e.name
FROM 
    EmployeeUNI eu
RIGHT JOIN 
    Employees e ON eu.id = e.id;
```


###### 1068 产品销售分析
题干中：获取 Sales 表中所有 sale_id 对应的 product_name
所以知道sales表是主表！！ （根据销售记录列产品，如果没有对应产品就是NULL）
如果是product作为主表就是：根据产品找销售记录，没有记录就是NULL
根据xx（主）找xx（别的）
这题没有要求如果没产品null表示，所以inner join或者join

###### 1581
COUNT() 是一个聚合函数，必须配合 GROUP BY 使用：正确写法是 COUNT(\*) AS count_no_trans，并用 GROUP BY 指定分组列。
LEFT JOIN + IS NULL​​：比 NOT IN 更高效
IS NULL不是 = NULL


###### 197 
datediff(today.recordDate,yesterday.recordDate) = 1
自己和自己join，同时不要求一定保留主表的记录，因为
查看join效果  （有就可以，所以join）
``` sql
select * 
from Weather today
join weather yesterday
on datediff(today.recordDate,yesterday.recordDate) = 1  #
```

###### 1661
也是自连接，方便计算
ROUND(..., 3) 保留 3 位小数
AVG计算平均数
- 通过这题看出sql书写思路：确定数据源from，join  然后筛选数据where  接着分组聚合group by（这样输出就聚为一个选项） 最后select要输出的数据
如果表都有同一个名字，需要a1.或者a2.

###### 577
自己搞定～

###### 1280
一对多，学生是一，课程是多    但是还多了一个考试的组合
group by 不止可以放一种，不然都聚到一块了！order by也是
思考的时候可以先想对其中的部分表操作比如先对exam操作，再和别的表合并，这也是个思路
比如这题：先考虑exam
官方题解还有先组合后就是grouped表这样 ()grouoped
``` sql
SELECT 
    student_id, subject_name, COUNT(*) AS attended_exams
FROM 
    Examinations
GROUP BY 
    student_id, subject_name

```
再考虑学生表和课程表，那么要和考试表配套，就需要交叉
``` sql
SELECT 
    *
FROM
    Students s
CROSS JOIN
    Subjects sub
```

使用LEFT JOIN​​：确保所有学生和科目的组合都会被保留，即使没有考试记录。
select中使用grouped.subject_name,  会存在NULL！

ai提供的代码中，COUNT(字段名)忽略NULL是SQL标准行为！！所以count（e.subject_name）就行
###### [570 至少有5名直接下属的经理](https://leetcode.cn/problems/managers-with-at-least-5-direct-reports/)
自己A出来了
- GROUP BY 先执行​​：将数据按指定字段分组，生成临时分组结果集（每个组包含多行原始数据）。
- HAVING 后执行​​：对分组后的结果进行过滤，仅保留满足条件的分组。

``` sql
select u.name
from Employee s
join Employee u
on s.managerId = u.id
group by u.name
having count(u.name) >= 5
```
看上去没问题，但是提交的时候就有问题了，因为可能有主管同名但是id不同!

###### [1934. 确认率](https://leetcode.cn/problems/confirmation-rate/)
主键的值在表中必须唯一，不能重复，同时不能设置为NULL（所以就是唯一滴～）
初步测试：
``` sql
select *
from Signups s
left join Confirmations c
on s.user_id = c.user_id
```
左边会有多的条数，同时也能统计到没有发请求的用户 ，会设置为NULL

- NULLIF(expr1, expr2)​​：当 expr1 = expr2 时返回 NULL，否则返回 expr1 的值。
- 在分母中的应用​​：NULLIF(COUNT(c.action), 0) 会检查统计的行数是否为0：若 COUNT(c.action) = 0，返回 NULL → 整个除法变为 A / NULL，结果为 NULL（避免报错）。若 COUNT(c.action) ≠ 0，返回实际行数 → 正常计算除法
- 注意'confirmed' 用单引号是sql标准
- ROUND:之前用过，统计小数的，后面2代表位数
- THEN 1 END​​
​​作用​​：这是CASE WHEN表达式的一部分，用于条件计数或标记。THEN 1​​：当条件满足时，返回数值1（通常用于计数或布尔逻辑）。END​​：表示CASE语句的结束。

- \*1.0:强制将整数运算转换为浮点数运算，避免​**​整数除法截断​**​问题
- \*100.0:将比例转换为百分比形式

```sql
select s.user_id, ROUND(
    count(case when c.action = 'confirmed' THEN 1 END)*1.0/
    NULLIF(count(c.action),0),
    2) as confirmation_rate
from Signups s
left join Confirmations c
on s.user_id = c.user_id
group by s.user_id
```

 而NULL在最外层的count的时候会自动统计为0？并不会，结果来看返回NULL。
 而count（）会自动跳过中间NULL




直接求百分比的话用AVG就可以，然后判断如果是NULL （IFNULL和NULLIF不同）就令为0
``` sql
select s.user_id, ROUND(IFNULL(AVG(c.action = 'confirmed'),0),2) as confirmation_rate
from Signups s
left join Confirmations c
on s.user_id = c.user_id
group by s.user_id
```

###### [620. 有趣的电影](https://leetcode.cn/problems/not-boring-movies/)
- 第一版：
``` sql
select movie,description,rating
from cinema
where id%2=1 and description is not boring
order by rating desc
```
`IS NOT` 后面应该接 `NULL` 或布尔表达式
同时boring 字符串需要用单引号包裹
- 正解：
``` sql
select id,movie,description,rating
from cinema 
where id%2=1 and description != 'boring'
order by rating desc
```

###### [1251. 平均售价](https://leetcode.cn/problems/average-selling-price/)
- 第一版
	应该是拼一下，根据id和sold表中date在对应start-end之间来，然后通过price和unnits进行计算     语法不会
- 正解：
``` sql
select p.product_id, #不可为u，因为可能是NULL
    ROUND(IFNULL(SUM(u.units * p.price)/SUM(u.units),0),2) as average_price
from prices p
left join UnitsSold u 
on p.product_id = u.product_id 
    and u.purchase_date between p.start_date and p.end_date
group by product_id  
```

需要保留所有原始记录而不合并，应该去掉GROUP BY子句
需要left join 因为要求产品没有任何售出，则假设其平均售价为 0，需要保存商品记录


###### [1075. 项目员工 I](https://leetcode.cn/problems/project-employees-i/)
- 第一版
	主键为 (project_id, employee_id)并非说明每个员工只负责一个项目！而是说组合起来是主键
```sql
select p.project_id,ROUND(AVG(e.experience_years),2) as average_years
FROM Project p
join Employee e
where p.employee_id = e.employee_id
group by p.project_id
```
（运行排名真没用，每次运行都不一样的，算法的倒不太一样一点点）

###### [1633. 各赛事的用户注册率](https://leetcode.cn/problems/percentage-of-users-attended-a-contest/)
- 第一版
	需要根据总的user数来确定百分比，由于不是“根据user找比赛”，所以主表不是users。而是“根据比赛找user数量”.同时似乎不需要left join
	不知道怎么统计总人数用于除法
	尝试：
``` sql
select r.contest_id,Round(SUM(u.user_name)/SUM(distinct u.user_id)*100.0,2) as percentage
from Register r 
join users u 
on r.user_id = u.user_id
group by contest_id
order by percentage desc,contest_id
```

`SUM(u.user_name)` - 用户名是字符串不能求和
`SUM(distinct u.user_id)` - 用户ID求和没有意义,是统计条数
同时不需要left join因为没有用到Users表中别的信息
- 正解：
``` sql
select r.contest_id,
    Round(COUNT(distinct r.user_id)/(select Count(distinct user_id) from users)*100.0,2) as percentage
from Register r 
group by r.contest_id
order by percentage desc,contest_id
```


######  [1211. 查询结果的质量和占比](https://leetcode.cn/problems/queries-quality-and-percentage/)
- 第一版：
	不太清楚“四舍五入”怎么实现，以及怎么count结合where统计
``` sql
select query_name,
    ROUND(SUM(rating/position)*100.0/count(*),2) as quality
    ROUND(select(count(rating where rating < 3))*100.0/count(*),2) as poor_query_percentage
from Queries
Group by query_name
```
count是已经分组过的所以count(\*)是没问题的
理解SUM上又一些问题     
当前的表中的情况可以用case来处理，不像上一题中需要从user表中distint，获取新的统计数据，这里的数据直接是分组后的结果中的统计，用case
- 正解
``` sql
select query_name,
    ROUND(AVG(rating/position),2) as quality,
    ROUND(SUM(CASE WHEN rating < 3 THEN 1 ELSE 0 END)*100.0/count(*),2) as poor_query_percentage
from Queries
Group by query_name
```
###### [1193. 每月交易 I](https://leetcode.cn/problems/monthly-transactions-i/)
- 第一版
需要select出月份，同时需要分组需要按照国家和日期按照月分组，算approve的总数和请求的总数

- 正解

``` sql
select DATE_FORMAT(trans_date,'%Y-%m') as month,
    country,
    COUNT(*) AS trans_count,
    SUM(case when state = 'approved' THEN 1 ELSE 0 END) as approved_count,
    SUM(amount) as trans_total_amount,
    SUM(case when state = 'approved' THEN amount ELSE 0 END) AS approved_total_amount
from Transactions
group by country,DATE_FORMAT(trans_date,'%Y-%m')
```

###### [1174. 即时食物配送 II](https://leetcode.cn/problems/immediate-food-delivery-ii/)

- 第一版
	根据所有用户统计，但是每个用户需要根据自己的数据单独统计，需group
	然后应该是需要统计第一条订单（需要排序）并判断条件为1或者0
	然后再除所有的distinct用户总数
``` sql
select ROUND(SUM(select(1)case when order_date = customer_pref_delivery_date THEN 1 ElSE 0 END)/select(count(distinct customer_id from Delivery)),2) as immediate_percentage
from Delivery
GROUP by customer_id
order by delivery_id
```
比较乱写的属于是，不知道语法

- 正解

使用 `WHERE` 子句筛选出每个客户的首次订单
子查询的作用是找出每个客户的首次订单！
``` sql
select ROUND(SUM(case when order_date = customer_pref_delivery_date THEN 1 ElSE 0 END)*100.0/count(*),2) as immediate_percentage
# 先搞好表
from Delivery
where (customer_id, order_date) in(
    select customer_id,min(order_date)
    from Delivery
    group by customer_id
)
```

######  [550. 游戏玩法分析 IV](https://leetcode.cn/problems/game-play-analysis-iv/)
次日留存类问题
- 第一版
	类似上一题，不过需要找到首次和第二次登陆进行比较
	自己尝试写但是发现判断条件不太会
```sql
select ROUND(SUM(case when  )*1.0/select count(distinct player_id) from Activity,2)
from Activity
where (player_id,event_date) in (
    select player_id,event_date
    from Activity
    order by event_date 
    limit 2
)
group by player_id
```
limit 2 只会返回前两个，还是不太行的，没有分组等
同时自己的where返回两条数据后面还要处理，不太聪明
- 正解1
``` sql
select round(count(*)/(select count(distinct player_id) from activity),2) as fraction
from activity
where (player_id,date_sub(event_date,interval 1 day)) in (
    select player_id,min(event_date)
    from activity
    group by player_id
)
```
满足event_date减去1后在where的子查询（最早一天），其实就是第二天。

- 正解2
利用join
```sql
SELECT ROUND(
    COUNT(a.player_id) / 
    (SELECT COUNT(DISTINCT player_id) FROM activity),
    2
) AS fraction
FROM activity a
JOIN (
    SELECT player_id, MIN(event_date) AS first_login
    FROM activity
    GROUP BY player_id
) b ON a.player_id = b.player_id 
    AND a.event_date = DATE_ADD(b.first_login, INTERVAL 1 DAY)
    #或者datediff(a.event_date, b.first_login) = 1
```

 
######  [2356. 每位教师所教授的科目种类的数量](https://leetcode.cn/problems/number-of-unique-subjects-taught-by-each-teacher/)
第一次搞出来
- 正解
``` sql
select teacher_id,count(distinct subject_id) as cnt
from Teacher
group by teacher_id
```


###### [1141. 查询近30天活跃用户数](https://leetcode.cn/problems/user-activity-for-the-past-30-days-i/)

