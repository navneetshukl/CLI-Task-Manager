package cmd

import (
	"fmt"
	"strings"

	"github.com/navneetshukl/task/db"
	"github.com/spf13/cobra"
)

// addCmd function will add the task to our todo list via the terminal
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("add called")
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong : ", err.Error())
			return
		}
		fmt.Printf("Added \"%s\"to your task list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
