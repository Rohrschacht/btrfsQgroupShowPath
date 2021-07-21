package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/rohrschacht/checkroot"
)

func getSubvolumePathsByID(fspath string) (map[string]string, error) {
	pathsByID := make(map[string]string)

	cmd := exec.Command("btrfs", "subvolume", "list", fspath)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		cols := strings.Fields(line)

		id := cols[1]
		path := cols[len(cols)-1]

		pathsByID[id] = path
	}

	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	return pathsByID, nil
}

func main() {
	checkroot.RootOrExit()

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Please specify path to btrfs filesystem!\n")
		os.Exit(1)
	}

	pathsByID, err := getSubvolumePathsByID(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := exec.Command("btrfs", "qgroup", "show", os.Args[1])
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(stdout)

	// skip first two lines of output
	scanner.Scan()
	fmt.Print(scanner.Text())
	fmt.Println(" path")
	scanner.Scan()
	fmt.Print(scanner.Text())
	fmt.Println(" ----")

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Print(scanner.Text())

		cols := strings.Fields(line)
		qgroupid := cols[0]

		// find correct path and print it in additional column
		if qgroupid == "0/5" {
			fmt.Println(" /")
		} else if strings.HasPrefix(qgroupid, "0/") {
			subvolid := strings.TrimPrefix(qgroupid, "0/")
			path := pathsByID[subvolid]
			fmt.Println(" " + path)
		} else {
			fmt.Println()
		}
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
