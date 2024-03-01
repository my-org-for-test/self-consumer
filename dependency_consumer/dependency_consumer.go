package dependency_consumer

import (
	"fmt"
	"github.com/my-org-org/self-consumer/internal"
)

func PrintDependencyMessage() {
	fmt.Println("Dependency Consumer")
    internal.InternalFunction()
}
