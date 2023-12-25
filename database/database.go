package database

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"

	"github.com/bwmarrin/snowflake"
	"go.etcd.io/bbolt"
)

var DB *bbolt.DB
var SnowflakeNode *snowflake.Node

type Customer struct {
	Sequence  int
	Snowflake string
}

func CreateCustomer(c *Customer) error {
	return DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("customers"))

		id, err := b.NextSequence()
		if err != nil {
			return err
		}

		c.Snowflake = SnowflakeNode.Generate().String()
		c.Sequence = int(id)

		buf, err := json.Marshal(c)
		if err != nil {
			return err
		}

		return b.Put([]byte(itob(c.Sequence)), buf)
	})
}

func ListCustomers(limit int) (list []*Customer, err error) {
	err = DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("customers"))
		c := b.Cursor()

		for k, v := c.First(); k != nil && len(list) < limit; k, v = c.Next() {
			var customer Customer
			json.Unmarshal(v, &customer)
			list = append(list, &customer)
		}

		return nil
	})
	return
}

func Init() {
	fmt.Println("[DB] was requested, initializing...")

	// Setup snowflake node for id generation
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	SnowflakeNode = node

	database_path, ok := os.LookupEnv("USE_DATABASE")
	if !ok {
		fmt.Println("[CONFIG-ERROR] Database file not specified.")
		os.Exit(1)
	}

	fmt.Println("[CONFIG] Choosen DATABASE:", database_path)
	db, err := bbolt.Open(database_path, 0600, nil)
	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *bbolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("customers"))
		return nil
	})

	if err != nil {
		panic(err)
	}

	DB = db
	CreateCustomer(&Customer{})
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
