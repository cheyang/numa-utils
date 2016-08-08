package numa

import (
	"errors"
	"fmt"
)

// #cgo LDFLAGS: -lnuma
// #include <numa.h>
import "C"

var NumaNotAvailable = errors.New("No NUMA support available on this system")

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

func MemInMB(mem uint64) uint64 {
	return mem >> 20
}

/**
* Get Numa Nodes
**/
func Nodes() (nodes []int, err error) {

	if IsNumaAvailable() < 0 {
		return nodes, NumaNotAvailable
	}

	maxnode := int(C.numa_max_node())
	for i := 0; i <= maxnode; i++ {
		if C.numa_bitmask_isbitset(C.numa_nodes_ptr, C.uint(i)) > 0 {
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
	} else {
		return cpus, NumaNotAvailable
	}

	return cpus, nil

}

func PrintDistance() {

	maxnode := MaxNode()
	fst := 0

	for i := 0; i < maxnode; i++ {
		if C.numa_bitmask_isbitset(C.numa_nodes_ptr, C.uint(i)) > 0 {
			fst = i
			break
		}
	}

	if C.numa_distance(C.int(maxnode), C.int(fst)) == 0 {
		fmt.Println("No distance information available.")
		return
	}

	fmt.Printf("node distances:\n")
	fmt.Printf("node ")
	for i := 0; i <= maxnode; i++ {
		if int(C.numa_bitmask_isbitset(C.numa_nodes_ptr, C.uint(i))) > 0 {

			fmt.Printf("% 3d ", i)
		}
	}
	fmt.Printf("\n")
	for i := 0; i <= maxnode; i++ {
		if int(C.numa_bitmask_isbitset(C.numa_nodes_ptr, C.uint(i))) == 0 {
			continue
		}
		fmt.Printf("% 3d: ", i)
		for k := 0; k <= maxnode; k++ {
			if C.numa_bitmask_isbitset(C.numa_nodes_ptr, C.uint(i)) > 0 &&
				C.numa_bitmask_isbitset(C.numa_nodes_ptr, C.uint(k)) > 0 {
				fmt.Printf("% 3d ", C.numa_distance(C.int(i), C.int(k)))
			}
		}
		fmt.Printf("\n")
	}
}
