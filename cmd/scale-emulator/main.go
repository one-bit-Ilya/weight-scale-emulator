package main

import (
	"fmt"

	"github.com/one-bit-Ilya/weight-scale-emulator/internal/config"
	"github.com/one-bit-Ilya/weight-scale-emulator/internal/logger"
)

func main() {
	logger.Init()
	c := config.LoadConfig()
	fmt.Printf("%+v\n", c)
}
