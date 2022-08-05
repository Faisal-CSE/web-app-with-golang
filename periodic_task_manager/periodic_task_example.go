package periodic_task_manager

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

func myTask() {
	fmt.Println("This task will run periodically")
}
func ExecuteACronJob() {
	gocron.Every(6).Hours().Do(myTask)
	<-gocron.Start()
}
