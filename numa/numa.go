package numa

import (
	"errors"
	"fmt"
)

// #cgo LDFLAGS: -lnuma
// #include <numa.h>
import "C"

var NumaNotAvailable = errors.New("Numa not Available")

/* Is Numa available? */
func IsNumaAvailable() int {
	return int(C.numa_available())
}

/* Get max available node */
func MaxNode() int         { return int(C.numa_max_node()) }
func MaxPossibleNode() int { return int(C.numa_max_possible_node()) }

func NumConfiguredCPUs() int { return int(C.numa_num_configured_cpus()) }

/**
* Get Memory of the Numa node
**/
func MemoryOfNode(node int) (inAll, free uint64) {

	cFree := C.longlong(0)
	cInAll := C.numa_node_size64(C.int(node), &cFree)
	return uint64(cInAll), uint64(cFree)
}

func MemInMB(mem uint64) string {
	return fmt.Sprintf("%d", mem)
}

/**
* Get Numa Nodes
**/
func Nodes() (nodes []int, err error) {

	if IsNumaAvailable() < 0 {
		return nodes, NumaNotAvailable
	}

	mask := C.numa_allocate_nodemask()
	defer C.numa_free_nodemask(mask)

	maxnode := C.numa_max_node()
	for i := 0; i < maxnode; i++ {
		if C.numa_bitmask_isbitset(mask, C.uint(i)) > 0 {
			nodes = append(nodes, i)
		}
	}

	return nodes, nil
}

/**
* Get CPU slice from the specified Node
**/
func CPUsOfNode(node int) (cpus []int, err error) {

	if IsNumaAvailable() < 0 {
		return cpus, NumaNotAvailable
	}

	mask := C.numa_allocate_cpumask()
	defer C.numa_free_cpumask(mask)

	rc := C.numa_node_to_cpus(C.int(node), mask)
	maxCpus := NumConfiguredCPUs()
	if rc >= 0 {
		for i := 0; i < maxCpus; i++ {
			if C.numa_bitmask_isbitset(mask, C.uint(i)) > 0 {
				cpus = append(cpus, i)
			}
		}
	}

	return cpus, nil

}
