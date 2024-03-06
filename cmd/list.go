package cmd

import (
	"fmt"
	"os"

	"github.com/navneetshukl/task/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your task",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("list called")

		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong : ", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete!")
			return
		}

		fmt.Println("You have the following tasks:")

		for i, task := range tasks {
			fmt.Printf("%d, %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

}
