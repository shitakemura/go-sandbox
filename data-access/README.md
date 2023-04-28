## Tutorial: Accessing a relational database

https://go.dev/doc/tutorial/database-access


#### login mysql
- docker compose up
- mysql -h 127.0.0.1 -u docker sampledb -p
- show tables;

#### setup db
- docker compose up
- mysql -h 127.0.0.1 -u docker sampledb -p < create-tables.sql
