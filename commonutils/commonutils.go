package commonutils

import (
	"strings"
	"strconv"
	log "github.com/sirupsen/logrus"
)

func Add1HourToDateString(s string) string {
	log.WithFields(log.Fields{"package": "commonutils","function": "Add1HourToDateString",}).Debugf("Incoming string: %s", s)
	x := strings.SplitAfter(s, "T")
	y := strings.Split(x[1], ":")
	x1,_ := strconv.Atoi(y[0])
	y[0] =  strconv.Itoa(x1+1)
	x[1] = strings.Join(y, ":")
	z := strings.Join(x, "")
	log.WithFields(log.Fields{"package": "commonutils","function": "Add1HourToDateString",}).Debugf("Outgoing string: %s", z)
	return z
}

