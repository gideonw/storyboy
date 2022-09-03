package main

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Repo struct {
	driver  neo4j.Driver
	session neo4j.Session
}

func NewRepo() Repo {
	driver, err := neo4j.NewDriver("neo4j://127.0.0.1:7687", neo4j.BasicAuth("neo4j", "password", ""))
	if err != nil {
		panic(err)
	}

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	return Repo{
		driver:  driver,
		session: session,
	}
}

func (r Repo) Close() {
	r.session.Close()
	r.driver.Close()
}

func (r Repo) CreateEntry(t, name, race string) string {
	greeting, err := r.session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (a:Entry{type:$type, name:$name, race:$race}) RETURN id(a) + a.name",
			map[string]interface{}{
				"type": t,
				"name": name,
				"race": race,
			})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		panic(err)
	}

	return greeting.(string)
}
