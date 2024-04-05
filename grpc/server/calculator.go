package server

import (
	"context"
	"fmt"
	"grpc/client"
	"log"
	"os"
	"time"

	"github.com/dengsgo/math-engine/engine"
	pb "github.com/developerc/finprojagent3/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var RegisteredTaskMap map[int32]*pb.Task //хранилище зарегистрированных задач

func CheckPushSovedTask() {
	for {
		time.Sleep(1 * time.Second)
		for k, task := range RegisteredTaskMap {
			if task.Status != "in_progress" {
				fmt.Println("finished task: ", k, task)
				if successPushTask(task) {
					mutex.Lock()
					delete(RegisteredTaskMap, k)
					mutex.Unlock()
				}
			}
		}
	}
}

func successPushTask(task *pb.Task) bool {
	orch := client.GetOrchestrator()
	host := orch.Ip                          //"localhost"
	port := orch.Port                        //"5000"
	addr := fmt.Sprintf("%s:%d", host, port) // используем адрес сервера
	// установим соединение
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("could not connect to grpc server: ", err)
		os.Exit(1)
	}
	// закроем соединение, когда выйдем из функции
	defer conn.Close()

	grpcClient := pb.NewOrchServerServiceClient(conn)
	//idAgent, err := grpcClient.RegisterNewAgent(context.TODO(), &pb.AgentParams{Id: -1, Ip: agent.Ip /*"localhost"*/, Port: int32(agent.Port) /*5001*/})
	_, err = grpcClient.PushFinishTask(context.TODO(), task)
	if err != nil {
		return false
	}
	return true
}

func HandleSendTask(task *pb.Task) bool {
	if len(RegisteredTaskMap) >= MaxTasks { //если количество обрабатываемых задач больше максимального
		fmt.Println("too many handling task!")
		return false
	}

	mutex.Lock()
	RegisteredTaskMap[task.Id] = task
	mutex.Unlock()
	go solveTask(task)
	return true
}

func solveTask(task *pb.Task) {
	time.Sleep(30 * time.Second) //имитируем длительную задачу
	toks, err := engine.Parse(task.Expr)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		task.Enddate = time.Now().Unix()
		task.Status = "error"
		RegisteredTaskMap[task.Id] = task
		fmt.Println(RegisteredTaskMap)
		return
	}
	// []token -> AST Tree
	ast := engine.NewAST(toks, task.Expr)
	if ast.Err != nil {
		fmt.Println("ERROR: " + ast.Err.Error())
		task.Enddate = time.Now().Unix()
		task.Status = "error"
		RegisteredTaskMap[task.Id] = task
		fmt.Println(RegisteredTaskMap)
		return
	}
	// AST builder
	ar := ast.ParseExpression()
	if ast.Err != nil {
		fmt.Println("ERROR: " + ast.Err.Error())
		task.Enddate = time.Now().Unix()
		task.Status = "error"
		RegisteredTaskMap[task.Id] = task
		fmt.Println(RegisteredTaskMap)
		return
	}
	fmt.Printf("ExprAST: %+v\n", ar)
	// AST traversal -> result
	r := engine.ExprASTResult(ar)
	fmt.Println("progressing ...\t", r)
	fmt.Printf("%s = %v\n", task.Expr, r)
	task.Result = float32(r)
	task.Enddate = time.Now().Unix()
	task.Status = "finish"
	RegisteredTaskMap[task.Id] = task
	fmt.Println(RegisteredTaskMap)
}
