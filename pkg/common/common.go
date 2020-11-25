package common

import (
	"time"

	"k8s.io/apimachinery/pkg/util/wait"

)

const (
	BindPrefix         = "external-dns.providers.%s.%s"
)

var (
	Debug = false
	Backoff = wait.Backoff{
		Duration: 30 * time.Second,
		Factor:   1,
		Steps:    3,
	} // retry 3 times, total 90 seconds.
)