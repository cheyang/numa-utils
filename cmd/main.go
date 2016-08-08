package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/cheyang/numa-utils/numa"
	"github.com/spf13/cobra"
)

var mainCmd = &cobra.Command{
	Use:          os.Args[0],
	Short:        "Run numa display",
	SilenceUsage: false,
	RunE: func(cmd *cobra.Command, args []string) error {
		log.SetOutput(os.Stderr)

		nodes, err := numa.Nodes()

		if err != nil {
			return error
		}

		log.Infof("available: %d nodes\n", len(nodes))

		for _, node := range nodes {
			cpus := numa.CPUsOfNode(node)
			log.Infof("node %d cpus: %+v\n", node, cpus)
			all, free := MemoryOfNode(node)
			log.Infof("node %d size: %d MB\n", node, numa.MemInMB(all))
			log.Infof("node %d free: %d MB\n", node, numa.MemInMB(free))
		}

	},
}

func main() {
	if err := mainCmd.Execute(); err != nil {
		log.Fatalf("Err is %v", err)
	}
}

func init() {
	mainCmd.Flags().BoolP("hardware", "h", true, "Display hardware and exit")
}
