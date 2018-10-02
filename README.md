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

# Acknowledgments

The app first searches a list of links to be queried afterwards. Then, it pulls some information of these links, like stock price, market value, etc. This second search (for details of each link) is the process which takes longer.

The first implementation, without any concurrency used, took about **9min30s** to **10min** to be completed. Then, another approach was dividing the slice into _go routines_, where each _go routine_ would take care of a part of the slice, appending the result to a final slice of structs. With that approach, the time spent dropped down to **4mins** ish.

It was used a _WaitGroup_. A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls _Add_ to set the number of goroutines to wait for. Then each of the goroutines runs and calls _Done_ when finished. At the same time, _Wait_ can be used to block until all goroutines have finished.

### Useful/Basic MySQL Database Commands

```
show databases;
use <DBName>;
show tables;
describe <TableName>;
drop database <DBName>;
drop table <TableName>;
select * from <TableName>;
```

### Opening MySQL db from inside the container

It was used the command below (as a root user), where `db-mysql-container` is the container name defined on Docker Compose file.

```
sudo docker exec -it db-mysql-container mysql -uroot -proot
```

### Load a SQL file using Docker Compose

It's done via volumes, as shown in the `docker-compose.yaml` file in this project and as shown [here](https://stackoverflow.com/questions/44533534/docker-how-to-use-sql-file-in-directory) and [here](https://gist.github.com/onjin/2dd3cc52ef79069de1faa2dfd456c945).

```
db:
     volumes:
        /path-to-sql-files-on-your-host:/docker-entrypoint-initdb.d
```

then run `docker-compose down -v` to destroy containers and volumes and run `docker-compose up` to recreate them.

### Linking one container to another

To reach a service on another container, take [this docker tutorial](https://docs.docker.com/compose/networking/) as reference.

For example, in this project, _DB_HOST_ env var is defined as `db`, the name given to the mysql service. And _DB_PORT_ is set with the same number exposed in **ports** inside the mysql.
