package server

import (
	"net"

	"github.com/azmodb/azmo/pb"
	"github.com/azmodb/db"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
	db *db.DB
}

func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.Event, error) {
	rec, err := s.db.Get(req.Key, req.Rev, req.Versions)
	if err != nil {
		rec.Close()
		return nil, err
	}

	ev := &pb.Event{Record: rec.Record, Duration: 0}
	return ev, nil
}

func (s *server) Batch(req *pb.BatchRequest, srv pb.DB_BatchServer) (err error) {
	batch := s.db.Next()
	ev := &pb.Event{}

	for _, arg := range req.GetArgs() {
		var rec *db.Record

		switch t := arg.Command.(type) {
		case *pb.Argument_Put:
			r := t.Put
			if !r.Tombstone {
				rec, err = batch.Insert(r.Key, r.Value, r.Prev)
			} else {
				rec, err = batch.Put(r.Key, r.Value, r.Prev)
			}
		case *pb.Argument_Delete:
			r := t.Delete
			rec, err = batch.Delete(r.Key, r.Prev)
		case *pb.Argument_Increment:
			r := t.Increment
			rec, err = batch.Increment(r.Key, r.Value, r.Prev)
		case *pb.Argument_Decrement:
			r := t.Decrement
			rec, err = batch.Decrement(r.Key, r.Value, r.Prev)
		}
		if err != nil { // rec is already released
			break
		}

		ev.Record = rec.Record
		err = srv.Send(ev)
		rec.Close()
		if err != nil {
			break
		}
	}
	return err
}

func (s *server) Range(req *pb.RangeRequest, srv pb.DB_RangeServer) error {
	return nil
}

func (s *server) Watch(req *pb.WatchRequest, srv pb.DB_WatchServer) error {
	return nil
}

type Option func(*server) error

func Listen(listener net.Listener, options ...Option) error {
	server := &server{db: db.New()}
	for _, opt := range options {
		if err := opt(server); err != nil {
			return err
		}
	}

	s := grpc.NewServer()
	pb.RegisterDBServer(s, server)
	return s.Serve(listener)
}
