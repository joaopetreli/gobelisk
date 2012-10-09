package main

import (
	"flag"
	"fmt"
	"gobelisk/manager"
	"gobelisk/protocol/action"
	"time"
)

var (
	host             = flag.String("host", "", "host to connect with.")
	port             = flag.String("port", "", "port to connect with.")
	username         = flag.String("username", "", "valid username.")
	secret           = flag.String("secret", "", "valid password.")
	keepAlive        = flag.Bool("keep-alive", false, "keep the program alive by sending ping action.")
	keepAlivetimeout = flag.Uint("keep-alive-timeout", 0, "Keep alive timeout in seconds.")
)

func main() {
	flag.Parse()
	if *host == "" || *port == "" || *username == "" || *secret == "" {
		flag.Usage()
		return
	}

	login := action.NewLogin(*username, *secret)
	conn, fullyBooted, err := manager.Connect(*host, *port, &login)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer manager.Logoff(conn)
	fullyBooted.Callback()

	if *keepAlive {
		stop := false
		if *keepAlivetimeout > 0 {
			go func() {
				time.Sleep(time.Duration(*keepAlivetimeout) * time.Second)
				stop = true
			}()
		}

		ping := action.NewPing()
		for {
			if stop {
				break
			}
			if err = manager.SendQuery(conn, &ping); err != nil {
				break
			}
			time.Sleep(1 * time.Second)
		}
	}

}
