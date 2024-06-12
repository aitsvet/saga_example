package domain

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/aitsvet/saga_example/internal/db"
	"github.com/aitsvet/saga_example/internal/kafka"
	"github.com/aitsvet/saga_example/internal/model"
	"github.com/aitsvet/saga_example/internal/rabbit"
)

type Service struct {
	Name   string
	DB     *sql.DB
	Rabbit *rabbit.Client
	Kafka  *kafka.Client
}

func NewService(name string, migration string, listening ...string) *Service {
	var err error
	s := &Service{Name: name}
	if os.Getenv("DB_NAME") != "" {
		s.DB, err = db.ConnectDB()
		if err != nil {
			log.Fatal(err)
		}
		_, err = s.DB.Exec(migration)
		if err != nil {
			log.Fatal(err)
		}
	}
	s.Rabbit, err = rabbit.Connect(name)
	if err != nil {
		log.Fatal("Failed to create RabbitMQ client:", err)
	}
	s.Kafka, err = kafka.Connect(name, listening...)
	if err != nil {
		log.Fatal("Failed to create Kafka client:", err)
	}
	fmt.Println(name + " service started")
	return s
}

func (s *Service) Close() {
	if s.DB != nil {
		_ = s.DB.Close()
	}
	if s.Rabbit != nil {
		s.Rabbit.Close()
	}
	if s.Kafka != nil {
		s.Kafka.Close()
	}
}

func (s *Service) LogError(m *model.Message, o any, err error) {
	_, filename, line, _ := runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "%s:%d %s: tx %d from %s with %v got error %v\n",
		filename, line, s.Name, m.TxId, m.Queue, o, err)
	os.Stderr.Sync()
}
