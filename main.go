package main

import (
	"context"
	"time"

	"collectd.org/api"
	"collectd.org/exec"

	"fmt"
	"github.com/egidijus/go-libzfs"
	"log"
	"strconv"
)

func main() {

	var pname string
	pools, err := zfs.PoolOpenAll()
	if err != nil {
		fmt.Printf("terrible things happened")
		log.Print(err)
		return
	}

	// println("\tThere is/are ", len(pools), " ZFS pool/s.")

	for _, p := range pools {
		pname, err = p.Name()
		if err != nil {
			log.Print(err)
			p.Close()
			return
		}

		pstate, err := p.State()
		if err != nil {
			log.Print(err)
			p.Close()
			return
		}

		psize, err := p.GetProperty(zfs.PoolPropSize)
		if err != nil {
			log.Print(err)
			p.Close()
			return
		}

		isize, err := strconv.Atoi(psize.Value)
		if err != nil {
			log.Print(err)
			p.Close()
			return
		}

		pfree, err := p.GetProperty(zfs.PoolPropFree)
		if err != nil {
			log.Print(err)
			p.Close()
			return
		}

		vl := api.ValueList{
			Identifier: api.Identifier{
				Host:           exec.Hostname(),
				Plugin:         "gollectz",
				PluginInstance: pname,
				Type:           "gauge",
			},
			Time:     time.Now(),
			Interval: exec.Interval(),
			Values:   []api.Value{api.Gauge(isize)},
		}
		exec.Putval.Write(context.Background(), &vl)

		println("\tPool: ", pname, " state: ", pstate.String(), "Size", psize.Value, "Free", pfree.Value)
		p.Close()
	}
}
