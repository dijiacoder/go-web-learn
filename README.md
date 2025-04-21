# HELLO
初学者入门 Go WEB 开发

go run cmd/gen/main.go


[ Service ]  
↓ 调用  
[ Repository ]  
↓ 封装 & 调用  
[ gorm/gen 生成的 Query ]  
↓ 实际操作  
[ 数据库 ]  


internal/  
├── model/  
│   └── gen/             # gorm/gen 自动生成（结构体 + query）  
├── repository/  
│   └── user_repository.go  # 我们手写的 repository，调用 gen 的 query  
├── service/  
│   └── user_service.go     # 业务逻辑层，调用 repository  

