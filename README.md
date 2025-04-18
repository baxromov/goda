### Testing Different Databases
#### For SQLite
- In `config.yaml`:
``` yaml
  database:
    driver: sqlite
    sqlite:
      dsn: test.db
```
#### For PostgreSQL
- In `config.yaml`:
``` yaml
  database:
    driver: postgres
    postgres:
      host: localhost
      port: 5432
      user: postgres
      password: postgres
      dbname: mydb
      sslmode: disable
```
#### For MySQL
- In `config.yaml`:
``` yaml
  database:
    driver: mysql
    mysql:
      host: localhost
      port: 3306
      user: mysql_user
      password: mysql_pass
      dbname: mydb
```
Restart the application, and it will automatically connect to the database defined in `config.yaml`.
