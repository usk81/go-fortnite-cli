# go-fortnite-cli
example of https://github.com/usk81/go-fortnite

## Install 

```
git clone git@github.com:usk81/go-fortnite-cli.git
```

## Usage

```
// Get Fornite BR Player Stats
go run main.go -t your-token stats gamepad player-nick-name

// Get Fornite Match History
go run main.go -t your-token history player-account-id

// Get Current Fortnite Store
go run main.go -t your-token store

// Get Current Active Challenges
//   note: This API is inactive, maybe.
go run main.go -t your-token challenges
```

