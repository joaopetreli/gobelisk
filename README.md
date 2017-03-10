gobelisk
--------

An easy way to interact with Asterisk Manager API in golang.

Documentation
-------------

### How do I connect to a server running Asterisk Manager API?

It is possible to connect to a server running Asterisk Manager API through the function `Connect` inside the `manager` package, like depicted below:

    package main

    import (
      "fmt"
      "gobelisk/manager"
      "gobelisk/protocol/action"
    )

    func main() {
      login := action.NewLogin("username", "password")
      conn, fullyBooted, err := manager.Connect("host", "port", &login)
      if err != nil {
        fmt.Println(err)
        return
      }
      defer manager.Logoff(conn)
      fullyBooted.Callback()

      ping := action.NewPing()
      if err = manager.SendQuery(conn, &ping); err != nil {
        fmt.Println(err)
        return
      }
	}

For more information, take a look at the `main.go` program in the project root.
