package environment

import "os"

func SetUpEnvironmentVariables() {
	os.Setenv("mysqlLogin", "root:Chaugn@rs2@/mini_golang_project")
	os.Setenv("redisHost", "localhost:6379")
	os.Setenv("port", "127.0.0.1:8080")
}