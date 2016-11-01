# Brava Giulia

A go library to send IRCC commands to a Sony TV using a Pre-Shared-Key

## Setup the TV

First set the key in your TV under 

     Settings > Network > Home Network Setup > IP Control
     
Set Authentication to
    
    Normal and Pre-Shared-Key
    
And also set a key value.

Use your router to discover your TV IP address (and maybe give the TV
its own host name, say sonytv)


## Use the package

    package main

    import (
        "fmt"
        "github.com/baol/bravagiulia"
    )

    func main() {
        c := bravagiulia.NewClient("sonytv", "<YOUR-PSK-HERE>")
        commands = c.GetSupportedCommands()
        fmt.Println(commands) // will list the available commands
        c.SendIRCC(commands["PowerOff"]) // can call the commands by name
    }

