# devgorant [![GoDoc](https://godoc.org/github.com/jayeshsolanki93/devgorant?status.svg)](https://godoc.org/github.com/jayeshsolanki93/devgorant) [![Build Status](https://travis-ci.org/jayeshsolanki93/devgorant.svg?branch=master)](https://travis-ci.org/jayeshsolanki93/devgorant)

Unofficial golang wrapper for the [devRant API](https://www.devrant.io/)

## Installation
```bash
go get github.com/jayeshsolanki93/devgorant
```

## Usage
Simple implementation to get the user score(points):

```go
package main

import (
  "fmt"
  "github.com/jayeshsolanki93/devgorant"
  "log"
)

func main() {
  devrant := devgorant.New()
  user, _, err := devrant.Profile("jayeshs")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(user.Score)
}
```


## API Reference

### `Rants` : Fetches rants

**Parameters:**

| Name | Type | Description | Default |
| ---- | ---- | -------- | ----------- | ------- |
| `sort` | string | Sort by `algo`, `top`, `recent` | `algo` |
| `limit` | integer | Number of rants required. Cannot be more than 50. | 50 |
| `skip` | integer | Number of rants to skip. | 0 |

**Example:**
```go
devrant := devgorant.New()
rants, err := devrant.Rants("algo", 20, 0)
```
---
### `Rant` : Fetches a rant and its comments given a valid rant id

**Parameters:**

| Name | Type | Description | Default |
| ---- | ---- | -------- | ----------- | ------- |
| `rantId` | integer | rant_id of a posted rant |  |

**Example:**
```go
devrant := devgorant.New()
rant, comments, err := devrant.Rant(27317)
```
---
### `Profile` : Fetches ranter's profile data

**Parameters:**

| Name | Type | Description | Default |
| ---- | ---- | -------- | ----------- | ------- |
| `username` | string | a valid username on devRant |  |

**Example:**
```go
devrant := devgorant.New()
user, content, err := devrant.Profile("jayeshs")
```
---
### `Search` : Search for rants matching the search term

**Parameters:**

| Name | Type | Description | Default |
| ---- | ---- | -------- | ----------- | ------- |
| `term` | string | any string to use as the search term |  |

**Example:**
```go
devrant := devgorant.New()
rants, err := devrant.Search("golang")
```
---
### `Surprise` : Returns a random rant

**Example:**
```go
devrant := devgorant.New()
rant, err := devrant.Surprise()
```
---
### `WeeklyRants` : Returns the rants tagged for 'weekly'

**Example:**
```go
devrant := devgorant.New()
rants, err := devrant.WeeklyRants()
```
---

## Tests
To run the tests locally:
```bash
go test -v 
```

## TODO
// TODO
