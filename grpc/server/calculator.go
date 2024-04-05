package server

import (
	"fmt"
	"time"

	"github.com/dengsgo/math-engine/engine"
	pb "github.com/developerc/finprojagent3/proto"
)

var RegisteredTaskMap map[int32]*pb.Task //хранилище зарегистрированных задач

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
