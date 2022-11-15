# calculator http service
## Calculator service

edit config.go in the utils package to edit configurations details

Start postgres and redis server with docker-compose

and

go run main.go

Visit: http://localhost:3330/calculate?first={a}&second{b}

or http://localhost:3330/calculate/history -> to get list of recent calculations

where {a} and {b} is any integer











