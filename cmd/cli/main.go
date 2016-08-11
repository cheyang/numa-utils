package main

import (
	"fmt"
	"os"

	"github.com/cheyang/numa-utils/numa"
	"github.com/spf13/cobra"
)

var mainCmd = &cobra.Command{
	Use:          os.Args[0],
	Short:        "Run numa display",
	SilenceUsage: false,
	RunE: func(cmd *cobra.Command, args []string) error {
		// log.SetOutput(os.Stderr)

		nodes, err := numa.Nodes()

		if err != nil {
			return err
		}

		fmt.Printf("available: %d nodes\n", len(nodes))

		for _, node := range nodes {
			cpus, err := numa.CPUsOfNode(node)
			if err != nil {
				return err
			}
			fmt.Printf("node %d cpus: %+v\n", node, cpus)
			all, free := numa.MemoryOfNode(node)
			fmt.Printf("node %d size: %d MB\n", node, numa.MemInMB(all))
			fmt.Printf("node %d free: %d MB\n", node, numa.MemInMB(free))
		}

		numa.PrintDistance()

		return nil
	},
}

func main() {
	if err := mainCmd.Execute(); err != nil {
		fmt.Printf("Err is %v", err)
		os.Exit(-1)
	}
}

func init() {
	mainCmd.Flags().BoolP("hardware", "H", true, "Display hardware and exit")
}
