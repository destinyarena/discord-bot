package faceit

import (
    "fmt"
    "google.golang.org/grpc"
    pb "github.com/arturoguerra/d2arena/pkg/faceit/proto"
)

func New(host, port string) (pb.FaceitClient, error) {
    address := fmt.Sprintf("%s:%s", host, port)
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        return nil, err
    }

    defer conn.Close()

    c := pb.NewFaceitClient(conn)

    return c, nil
}
