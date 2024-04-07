package server

import (
	"fmt"
	"testing"

	pb "github.com/developerc/finprojagent3/proto"
)

func TestSolveTask(t *testing.T) {
	t.Run("testSolveTaskValid", testSolveTaskValid)
	t.Run("testSolveTaskValid", testSolveTaskInValid)
}

func testSolveTaskInValid(t *testing.T) {
	RegisteredTaskMap = make(map[int32]*pb.Task)
	var task pb.Task
	task.Id = 1
	task.Agentid = 1
	task.Status = "in_progress"
	task.Expr = "2+3"
	solveTask(&task)
	fmt.Println("RegisteredTaskMap: ", RegisteredTaskMap)
	res := RegisteredTaskMap[1]
	if res.Result == float32(15) {
		t.Error("not valid result!")
	}
}

func testSolveTaskValid(t *testing.T) {
	RegisteredTaskMap = make(map[int32]*pb.Task)
	var task pb.Task
	task.Id = 1
	task.Agentid = 1
	task.Status = "in_progress"
	task.Expr = "2+3"
	solveTask(&task)
	//time.Sleep(35 * time.Second)
	fmt.Println("RegisteredTaskMap: ", RegisteredTaskMap)
	res := RegisteredTaskMap[1]
	if res.Result != float32(5) {
		t.Error("not valid result!")
	}
}

func TestHandleSendTask(t *testing.T) {
	t.Run("testHandleSendTaskValid", testHandleSendTaskValid)
	t.Run("testHandleSendTaskInValid", testHandleSendTaskInValid)
}

func testHandleSendTaskInValid(t *testing.T) {
	MaxTasks = 2
	RegisteredTaskMap = make(map[int32]*pb.Task)
	var task pb.Task
	task.Id = 1
	task.Agentid = 1
	task.Status = "in_progress"
	task.Expr = "2+3"
	RegisteredTaskMap[0] = &task
	RegisteredTaskMap[1] = &task
	got := HandleSendTask(&task)
	if got == true {
		t.Error("not valid result!")
	}
}

func testHandleSendTaskValid(t *testing.T) {
	MaxTasks = 2
	RegisteredTaskMap = make(map[int32]*pb.Task)
	var task pb.Task
	task.Id = 1
	task.Agentid = 1
	task.Status = "in_progress"
	task.Expr = "2+3"
	got := HandleSendTask(&task)
	if got == false {
		t.Error("not valid result!")
	}
	/*fmt.Println("RegisteredTaskMap: ", RegisteredTaskMap)
	solveTask(&task)
	time.Sleep(35 * time.Second)
	fmt.Println("RegisteredTaskMap: ", RegisteredTaskMap)
	res := RegisteredTaskMap[1]
	if res.Result != float32(5) {
		t.Error("not valid result!")
	}*/
}
