package client

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	pb "github.com/developerc/finprojagent3/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure" // для упрощения не будем использовать SSL/TLS аутентификация
)

type Agent struct {
	Id   int    `json:"id"`
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}

type Orchestrator struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}

var agent Agent
var orchestrator Orchestrator

// настройки агента и оркестратора читаем из настроечного файла
func readSettings(settingsFile string) {
	var numLine = 0
	f, err := os.OpenFile(settingsFile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println("error opening setting file!")
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		str := sc.Text() // GET the line string
		//в нулевой линии записаны настройки оркестратора, в первой агента
		switch numLine {
		case 0:
			//применяем настройки агента, записанные в файле
			err := json.Unmarshal([]byte(str), &orchestrator)
			if err != nil {
				fmt.Println("error unmarshal settings!")
				return
			}
			//fmt.Println(orchestrator)
		case 1:
			//применяем настройки агента, записанные в файле
			err := json.Unmarshal([]byte(str), &agent)
			if err != nil {
				fmt.Println("error unmarshal settings!")
				return
			}
			//fmt.Println(agent)
		}
		numLine++
	}
	if err := sc.Err(); err != nil {
		fmt.Println("error scanning setting file!")
		return
	}
}

func RegisterAgent() {
	readSettings("settings.txt")
	host := orchestrator.Ip                  //"localhost"
	port := orchestrator.Port                //"5000"
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
	idAgent, err := grpcClient.RegisterNewAgent(context.TODO(), &pb.AgentParams{Id: -1, Ip: agent.Ip /*"localhost"*/, Port: int32(agent.Port) /*5001*/})
	if err != nil {
		log.Println("failed invoking idAgent: ", err)
	}
	agent.Id = int(idAgent.GetId())
	fmt.Println("idAgent: ", agent.Id)
}

func GetAgentId() int {
	return agent.Id
}

func GetAgent() Agent {
	return agent
}

func GetOrchestrator() Orchestrator {
	return orchestrator
}
