package store

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/scylladb/gocqlx/v2"
	"github.com/szymon676/ogauth-grpc/proto"
)

type ScyllaStore struct {
	cluster *gocql.ClusterConfig
}

func NewScyllaStore() *ScyllaStore {
	cluster := gocql.NewCluster("127.0.0.1:9042")
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatalf("Failed to create scylladb session: %v", err)
	}
	defer session.Close()

	err = createKeyspace(&session)
	if err != nil {
		log.Fatalf("Failed to create keyspace: %v", err)
	}

	err = createTable(&session)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	fmt.Println("userStorage cluster successfully created")

	return &ScyllaStore{cluster: cluster}
}

func (s *ScyllaStore) SaveUser(user *proto.RegisterRequest) error {
	session, err := gocqlx.WrapSession(s.cluster.CreateSession())
	if err != nil {
		return err
	}
	userid := uuid.New()
	q := fmt.Sprintf("INSERT INTO users.users (id, username, password, email) VALUES ('%s', '%s', '%s', '%s')", userid.String(), user.Username, user.Password, user.Email)
	err = session.Query(q, nil).Exec()
	if err != nil {
		return err
	}

	return nil
}

func createKeyspace(session *gocqlx.Session) error {
	q := "CREATE KEYSPACE IF NOT EXISTS users WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}"
	err := session.Query(q, nil).Exec()
	if err != nil {
		return err
	}
	return nil
}

func createTable(session *gocqlx.Session) error {
	q := "CREATE TABLE IF NOT EXISTS users.users (id text, username text, password text, email text, PRIMARY KEY (id))"
	err := session.Query(q, nil).Exec()
	if err != nil {
		return err
	}
	return nil
}
