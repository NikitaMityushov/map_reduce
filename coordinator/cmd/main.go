package main

import (
	"fmt"

	"github.com/NikitaMityushov/map_reduce/coordinator/internal/config"
)

func main() {
	fmt.Println("Coordinator is started")

	// 1) config init
	cfg := config.MustLoad()

	fmt.Println(cfg)
	// 2) logger init
	// 3) init app
	// 4) start grpc server
}
