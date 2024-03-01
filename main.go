package main

import (
	"fmt"
	"github.com/my-org-org/self-producer/dependency"
	"github.com/my-org-org/self-consumer/dependency_consumer"
)

func main() {
	fmt.Println("Dependency Demo:")
	dependency.PrintDependencyMessage()

	dependency_consumer.PrintDependencyMessage()
}
