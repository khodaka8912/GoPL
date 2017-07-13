#!/bin/bash
go run clock.go -port 8010 -tz Asia/Tokyo &
go run clock.go -port 8020 -tz US/Eastern &
go run clockwall.go  Japan=localhost:8010 US=localhost:8020