package service

import (
	"fmt"

	"github.com/cheyang/numa-utils/numa"
	pb "github.com/cheyang/numa-utils/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Server struct {
}

func (this Server) GetInfo(ctx context.Context, req *pb.Empty) (response *pb.InfoResponse, err error) {

	response = &pb.InfoResponse{Nodes: make([]*pb.InfoResponse_Node, 0)}

	nodes, err := numa.Nodes()
	if err != nil {
		return nil, err
	}
	fmt.Printf("available: %d nodes\n", len(nodes))

	for _, node := range nodes {
		cpus, err := numa.CPUsOfNode(node)

		cpusInt32 := make([]uint32, len(cpus))

		for i, cpu := range cpus {
			cpusInt32[i] = uint32(cpu)
		}

		if err != nil {
			return nil, grpc.Errorf(codes.Unavailable, "failed message: %v", err)
		}
		fmt.Printf("node %d cpus: %+v\n", node, cpus)
		total, free := numa.MemoryOfNode(node)
		fmt.Printf("node %d size: %d MB\n", node, numa.MemInMB(total))
		fmt.Printf("node %d free: %d MB\n", node, numa.MemInMB(free))
		pbNode := &pb.InfoResponse_Node{
			CpuSet: cpusInt32,
			Free:   numa.MemInMB(free),
			Total:  numa.MemInMB(total),
		}
		response.Nodes = append(response.Nodes, pbNode)
	}

	return response, nil

}

func (this Server) GetMetrics(ctx context.Context, req *pb.Empty) (response *pb.MetricsResponse, err error) {
	response = &pb.MetricsResponse{Distances: make([]*pb.MetricsResponse_Distance, 0)}

	distances := numa.GetDistances()

	for _, distance := range distances {
		response.Distances = append(response.Distances, &pb.MetricsResponse_Distance{
			Start:  distance.Start,
			End:    distance.End,
			Length: distance.Length,
		})
	}

	return response, nil
}
