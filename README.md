# gator

Part of my Go learning.

Gator is blog aggreGATOR that allows users to track, follow, unfollow RSS feeds.

## Required to install
- postgres v15+
- go 1.20+
- goose v3+
- SQLC v2+


## How to use

### install the package
simply run `go install github.com/chaasfr/gator`

### Creation your CONNECTION_URL
You will use this url to connect in the rest of the installation. It looks like this
`postgres://username:password@localhost:5432/gator` (localhost:5432 are default values, you may have something different)

### Start the postgresDB
Start the postgresdb `sudo service postgresql start`

Connect to postgres if you need to do manual stuff: `sudo -u postgres psql`

### create a configuration file
it expects that you have a configuration file in your home folder called `~/.gatorconfig.json` that contains:
```
{
    "db_url":"[CONNECTION_URL]?sslmode=disable",
    "current_user_name":"[ANY_VALUE]"
}
```
Current user does not matter, it will be updated when you log in.

### Run DB migrations
```
cd sql/schema
goose [CONNECTION_URL] up
```

### start the program
either run `go run . [command] [args]` if you are in the folder or `gator [command] [args]`if you did a go install.