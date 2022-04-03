go run ./client/client_grpc.go -create id 1231
go run ./client/client_grpc.go -create id 1232
go run ./client/client_grpc.go -create id 1233
go run ./client/client_grpc.go -create id 1234
go run ./client/client_grpc.go -create id 1235
go run ./client/client_grpc.go -create id 1236
go run ./client/client_grpc.go -create id 1237
go run ./client/client_grpc.go -create id 1238
go run ./client/client_grpc.go -create id 1239
go run ./client/client_grpc.go -create id 1240

go run ./client/client_grpc.go -write -id 1231 -name abc -low 0 -mid 10 -high 100 localhost port 50051
go run ./client/client_grpc.go -write -id 1232 -name def -low 0 -mid 10 -high 100 localhost port 50051
go run ./client/client_grpc.go -write -id 1233 -name ghi -low 0 -mid 10 -high 100 localhost port 50051
go run ./client/client_grpc.go -write -id 1234 -name jkl -low 0 -mid 10 -high 100 localhost port 50051
go run ./client/client_grpc.go -write -id 1235 -name mno -low 0 -mid 10 -high 100 localhost port 50051
go run ./client/client_grpc.go -write -id 1236 -name pqr -low 0 -mid 10 -high 100 localhost port 50051
go run ./client/client_grpc.go -write -id 1237 -name stu -low 0 -mid 10 -high 100 localhost port 50051
go run ./client/client_grpc.go -write -id 1238 -name vwx -low 0 -mid 10 -high 100 localhost port 50051
go run ./client/client_grpc.go -write -id 1239 -name yza -low 0 -mid 10 -high 100 localhost port 50051
go run ./client/client_grpc.go -write -id 1240 -name yzb -low 0 -mid 10 -high 100 localhost port 50051

go run ./client/client_grpc.go -read -id 1231 localhost port 50051
go run ./client/client_grpc.go -read -id 1232 localhost port 50051
go run ./client/client_grpc.go -read -id 1233 localhost port 50051
go run ./client/client_grpc.go -read -id 1234 localhost port 50051
go run ./client/client_grpc.go -read -id 1237 localhost port 50051
go run ./client/client_grpc.go -read -id 1238 localhost port 50051
go run ./client/client_grpc.go -read -id 1237 localhost port 50051
go run ./client/client_grpc.go -read -id 1238 localhost port 50051
go run ./client/client_grpc.go -read -id 1239 localhost port 50051
go run ./client/client_grpc.go -read -id 1240 localhost port 50051

go run ./client/client_grpc.go -drop 1231 localhost port 50051
go run ./client/client_grpc.go -drop 1232 localhost port 50051
go run ./client/client_grpc.go -drop 1233 localhost port 50051
go run ./client/client_grpc.go -drop 1234 localhost port 50051
go run ./client/client_grpc.go -drop 1235 localhost port 50051
go run ./client/client_grpc.go -drop 1236 localhost port 50051
go run ./client/client_grpc.go -drop 1237 localhost port 50051
go run ./client/client_grpc.go -drop 1238 localhost port 50051
go run ./client/client_grpc.go -drop 1239 localhost port 50051
go run ./client/client_grpc.go -drop 1240 localhost port 50051

