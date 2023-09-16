# Database

### transaction

### Pessimistic lock

lock while start to access even select, slower

### Optimisc lock

lock while update, more fast

### isolation lv

- RU => dirty read, unrepeatable read, phantom read
- RC => unrepeatable read, phantom read
- RR => phantom read
- Serialization => all solved

### sql tuning

- index aka secondary key
- primary key
- join

check quantity of query data set

### performance tuning under large-scale

- db sharding
- master-slave for read/write segregatioin
- use memory
- batch process
- limit selection
- regional backup

### further

re-index?
