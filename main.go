package main

// desired output example
// PUTVAL giddy-hp/zfs-fakepool_fakezfs2/zusedbydataset interval=10 N:8192
//

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"collectd.org/api"
	"collectd.org/exec"
	zfs "github.com/egidijus/go-libzfs"
)

var properties = []zfs.Prop{
	zfs.PoolPropSize,
	zfs.PoolPropFree,
	zfs.PoolPropAllocated,
}

func main() {

	pools, err := zfs.PoolOpenAll()
	if err != nil {
		fmt.Printf("terrible things happened")
		log.Print(err)
		return
	}

	for _, thepool := range pools {
		poolname, err := thepool.Name()
		if err != nil {
			log.Print(err)
			thepool.Close()
			return
		}
		stats := getstats(thepool)
		thepool.Close()
		sendmetric(poolname, stats)
	}

}

//this part fakes the zfs response
func getstats(thepool zfs.Pool) map[string]int {
	propertiesAndValues := make(map[string]int)

	for _, property := range properties {
		propertyAndValue, err := thepool.GetProperty(property)
		propertyName := zfs.PoolPropertyToName(property)
		propertyName = "zfs_" + propertyName
		if err != nil {
			log.Print(err)
			return propertiesAndValues
		}
		// print selected properties for fun
		// fmt.Println(propertyName, " ", propertyAndValue.Value)

		// formated pool size
		numericValue, err := strconv.Atoi(propertyAndValue.Value)
		if err != nil {
			log.Print(err)
			return propertiesAndValues
		}
		// add metrics to map
		propertiesAndValues[propertyName] = numericValue

	}

	return propertiesAndValues
}

func sendmetric(poolname string, stats map[string]int) {

	for key, value := range stats {
		vl := api.ValueList{
			Identifier: api.Identifier{
				Host:           exec.Hostname(),
				Plugin:         "gollectz",
				PluginInstance: poolname,
				Type:           key,
			},
			Time:     time.Now(),
			Interval: exec.Interval(),
			Values:   []api.Value{api.Gauge(value)},
		}
		exec.Putval.Write(context.Background(), &vl)
	}

}
