# gator

Part of my Go learning.

Gator is blog aggreGATOR that allows users to track, follow, unfollow RSS feeds.

## Required to install
- postgres v15+
- go 1.20+
- goose v3+
- SQLC v2+


## How to

### Start the postgresDB
Start the postgresdb `sudo service postgresql start`

Connect to postgres if you need to do manual stuff: `sudo -u postgres psql`


### Run DB migrations
```
cd sql/schema
goose [CONNECTION_URL] up
```
