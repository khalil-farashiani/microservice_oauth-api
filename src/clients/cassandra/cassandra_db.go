package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	// Connect to cassandra cluster:
	cluster := gocql.NewCluster("127.0.0.1", "192.168.1.2", "192.168.1.3")
	cluster.Keyspace = "oauth"
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("connection to DB successfully")
	defer session.Close()
}

func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
