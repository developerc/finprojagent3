создал папку finprojagent3 для финального проекта
установил плагины
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
export PATH="$PATH:$(go env GOPATH)/bin"
C:\proj\Go\Exam\finprojorchestr2> go mod init github.com/developerc/finprojagent3
создал папку proto и grpc
mkdir proto
mkdir grpc
в папке proto создал файл grpc.proto
создал файлы 
grpc/client/clnt.go
grpc/server/srv.go
залил на github
////
Важно!! в файле grpc.proto строка 2
package finproj; // название пакета
должна совпадать со строкой в файле grpc.proto на оркестраторе, иначе запросы от агента не доходят до оркестратора даже при правильных IP и port
и строка 3
option go_package = "github.com/developerc/finprojagent3/proto";
должна соответствовать файлу на GitHub-e, когда даешь команду   go mod tidy   приложение подтягивает оттуда зависимости
////
сгенерировал файлы
protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     proto/grpc.proto  
go mod tidy
cd grpc
C:\proj\Go\Exam\finprojorchestr3\grpc> go mod init grpc
заполнил srv.go и подтянул зависимости
go mod tidy
в файле go.mod
replace grpc => ./grpc/
~/proj/Go/Exam/finprojagent3$ go get grpc
