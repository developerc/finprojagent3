package server

import (
	"context"
	"errors"
	"fmt"
	"grpc/client"
	"log"
	"net"
	"os"
	"sync"

	pb "github.com/developerc/finprojagent3/proto"
	"google.golang.org/grpc"
)

var RegisteredAgentMap map[int]Agent //хранилище зарегистрированных агентов
var IdAgent int                      //счетчик ID агента
var MaxTasks int                     //максимальное количество одновременно выполняемых задач
var mutex sync.Mutex

type Agent struct {
	Id   int    `json:"id"`
	Ip   string `json:"ip"`
	Port int    `json:"port"`
	//mutex sync.Mutex
}

type Server struct {
	pb.OrchServerServiceServer // сервис из сгенерированного пакета
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SendTask(ctx context.Context, task *pb.Task) (*pb.Task, error) {
	fmt.Println("Get Task:  ", task)
	task.Status = "in_progress"
	task.Agentid = int32(client.GetAgentId())
	if ok := HandleSendTask(task); ok {
		return task, nil
	}

	return nil, errors.New("agent can't to receive task")
}
func (s *Server) PushFinishTask(ctx context.Context, task *pb.Task) (*pb.Task, error) {

	return task, nil
}

func (s *Server) HBreq(ctx context.Context, heartBit *pb.HeartBit) (*pb.HeartBitResp, error) {
	//hbr := pb.HeartBitResp
	return &pb.HeartBitResp{}, nil
}

func CreateOrchGRPCserver() {
	MaxTasks = 2
	RegisteredAgentMap = make(map[int]Agent)
	RegisteredTaskMap = make(map[int32]*pb.Task)
	go CheckPushSovedTask()

	//изменить значения host port !!
	host := "localhost"
	port := "5001"

	addr := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("tcp", addr) // будем ждать запросы по этому адресу

	if err != nil {
		log.Println("error starting tcp listener: ", err)
		os.Exit(1)
	}

	log.Println("tcp listener started at port: ", port)
	// создадим сервер grpc
	grpcServer := grpc.NewServer()
	// объект структуры, которая содержит реализацию серверной части OrchServerServiceServer
	orchserverServiceServer := NewServer()
	// зарегистрируем нашу реализацию сервера
	pb.RegisterOrchServerServiceServer(grpcServer, orchserverServiceServer)
	// запустим grpc сервер
	if err := grpcServer.Serve(lis); err != nil {
		log.Println("error serving grpc: ", err)
		os.Exit(1)
	}
}
