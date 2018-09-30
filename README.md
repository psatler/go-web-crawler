# A simple Go Web Crawler

> A go web crawler with MySQL persistency

A simple go web crawler that finds the 10 most valuable companies according to the [Fundamentus](https://www.fundamentus.com.br/detalhes.php) website. It searches for stock name, company name, average daily rate of stock, and company's market value, storing them at a _MySQL_ dabatase.

# How to Run

You can either run it locally via a shell script or in a container via Docker Compose. Also, you need to provide an **_.env_** file to access the database. An example is shown in [.env.example]().

## Using a shell script and running locally

Do the following:

```
go get github.com/psatler/go-web-crawler
cd go-web-crawler
sh init-db.sh
```

The script might ask you about _sudo_ passwords to make sure _MySQL_ service is up. Also, it opens the _MySQL_ db as a root user.

## Using Docker Compose

The project also comes with a Docker Compose file to create containers for each service (go application and _MySQL_). So, to run using it, do:

```
git clone https://github.com/psatler/go-web-crawler.git
cd go-web-crawler
sudo docker-compose up
```

**NOTE:** If the _MySQL_ service is up, it might rise some port conflicts/errors, so you might have to do `sudo service mysql stop` first to be able to run docker compose.

# Main Dependencies

- [GoQuery](https://godoc.org/github.com/PuerkitoBio/goquery): implements features similar to jQuery, including the chainable syntax, to manipulate and query an HTML document.
- [Go-sql-driver](https://godoc.org/github.com/go-sql-driver/mysql): package mysql provides a MySQL driver for Go's database/sql package.
- [StrConv](https://godoc.org/strconv): implements conversions to and from string representations of basic data types.
- [Sort](https://godoc.org/sort#example-Slice): provides primitives for sorting slices and user-defined collections.

# License

# Acknowledgments

The app first searches a list of links to be queried afterwards. Then, it pulls some information of these links, like stock price, market value, etc. This second search (for details of each link) is the process which takes longer.

The first implementation, without any concurrency used, took about **4min30s** to **5min** to be completed. Then, another approach was dividing the slice into _go routines_, where each _go routine_ would take care of a part of the slice, appending the result to a final slice of structs. With that approach, the time spent dropped down to **1min09s** ish.

It was used a _WaitGroup_. A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls _Add_ to set the number of goroutines to wait for. Then each of the goroutines runs and calls _Done_ when finished. At the same time, _Wait_ can be used to block until all goroutines have finished.

## Useful/Basic MySQL Database Commands

```
show databases;
use <DBName>;
show tables;
describe <TableName>;
drop database <DBName>;
drop table <TableName>;
select * from <TableName>;
```

## Opening MySQL db from inside the container

It was used the command below (as a root user), where `db-mysql-container` is the container name defined on Docker Compose file.

```
sudo docker exec -it db-mysql-container mysql -uroot -proot
```

## Docker Compose

When using volumes

---

#0 - Company: CVRD ON N1
Market Value: 317121000000.000000
#1 - Company: PETROBRAS ON
Market Value: 304719000000.000000
#2 - Company: AMBEV S/A ON NM
Market Value: 294004000000.000000
#3 - Company: ITAUUNIBANCO PN N1
Market Value: 283143000000.000000
#4 - Company: AMBEV ON
Market Value: 272509000000.000000
#5 - Company: AMBEV PN
Market Value: 271883000000.000000
#6 - Company: PETROBRAS PN
Market Value: 263368000000.000000
#7 - Company: ITAUUNIBANCO ON N1
Market Value: 247456000000.000000
#8 - Company: BANCO BRADESCO S.A. PN N1
Market Value: 191986000000.000000
#9 - Company: CVRD PNA N1
Market Value: 176290000000.000000

Time without any concurrency (go routines) used
real 4m33,679s
user 0m0,000s
sys 0m0,015s

dividing the slice of urls and running them using go routines:
With 5 go routines
real 1m14,511s
user 0m0,000s
sys 0m0,015s

with 10 go routines
real 1m18,038s
user 0m0,000s
sys 0m0,030s

with 20 go routines
real 1m17,083s
user 0m0,015s
sys 0m0,000s

Compiling the code as an exe (with 6 go routines)
real 1m9,395s
user 0m0,000s
sys 0m0,031s

explanation:
A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.

https://nathanleclaire.com/blog/2014/02/15/how-to-wait-for-all-goroutines-to-finish-executing-before-continuing/

https://www.ardanlabs.com/blog/2014/01/concurrency-goroutines-and-gomaxprocs.html

Maximizing throughput is about getting rid of bottlenecks. First of all find where is most time is lost.

Sometimes running too many goroutines doesn’t make things faster but makes them slower, because goroutines start to compete for resources. Usually the fastest configuration is one that uses resources in the most effective way: if things are limited by computation then it's best to set the number of running goroutines equal to the number of CPU cores. But if it is limited by IO then choose the optimal number by measuring its performance.

# installed Docker

See links used as base to install docker on Ubuntu 16:04 LTS
