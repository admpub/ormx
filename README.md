#ormx

Ormx is a package for golang. 
It allows load balancing in a round robin style between master and replica databases.

Read queries are executed by replica.  
Write queries are executed by the master. 
Use .MasterCanRead(true) to perform WRITE queries with the master and READ queries with the master and replica.  

## Usage

```
	func main(){
		dsns := "root:root@tcp(master:3306)/masterpwd;"
		dsns += "root:root@tcp(replica:3306)/replicapwd"
		b, err := ormx.NewBalancer("mysql", gorp.MySQLDialect{"InnoDB", "utf8mb4"}	, dsns)
		if err != nil{
			panic(err)
		}
		err = b.Ping()
		if err != nil{
			panic(err)
		}
		
		// Master is used only for write queries, this is the default value
        b.MasterCanRead(false) 

        // Master is used for write and read queries
        b.MasterCanRead(true)
		
		count, err := b.SelectInt("SELECT COUNT(*) FROM mytable")
		if err != nil{
			panic(err)
		}
		fmt.Println(count)
		
	}
```
## License

Ormx is licensed under the [MIT License](./LICENSE).

