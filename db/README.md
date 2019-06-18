## Setup Config Database where alembic.ini
### Example
```bash
sqlalchemy.url = driver://user:pass@localhost/dbname
# to
sqlalchemy.url = mysql://user:pass@localhost/dbname
```

### Generate Migration
```bash
alembic revision -m "create account table"
```

### Run migrate to head
```bash
alembic upgrade head
```

### Run rollback one step
```bash
alembic downgrade -1
alembic downgrade base 
```
