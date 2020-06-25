package conf

import (
	"os"

	"github.com/joho/godotenv"
	"server/database"
)

// 初始化配置
func Init() {
	// 从本地读取环境变量
	_ = godotenv.Load()

	// 连接数据库
	database.InitDB(os.Getenv("MYSQL_DSN"))
}
