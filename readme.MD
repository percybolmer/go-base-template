# Example project
This is a example structure for a Go program.

`main.go` is simply a wrapper to start the COBRA Root command.  

This is project uses the Cobra CLI setup, which means that the `cmd` folder contains all the executable commands.

We can add helping information to the `root.go` that explains how to use the software.

Each CMD is then created seperatly and added Into the ROOTCMD.

Each CMD has its own input arguments and keyword to trigger it.

This approach makes it very easy to scale the software and add new features, as you can add a new CMD to easily test things.

The idea is simply this

`cmd` - Commands that executes from a CLI, can be scripts, long running processes etc.
`domain` - Here we store Libraries that are related to certain domains, etc modbus or mysql.
`services` - Combine domains into fully fledged "Services" that can do what we need to do.

CMD should not contain business logic, more act like a router and SETUP step and then Call on a Service OR services!

## Example usage
```bash
go run main.go
```

This will print the Help information located in `cmd/root.go`

```log
This project is structured with cobra

  The project uses only the Cobra Command tooling, not the Cobra CLI to generate code.

  The structure is Easy 

  the Folder cmd contains the Commads that are executable, such as a start command, or a small script
  Think of it as a bin folder.

  The domain/internal folder contains code/libraries that are used for certain domains
  such as the example ModBus Package

  Make sure to place all code that are related to ONE domain i the same package and dont share stuff across domain to avoid
  Coupling

Usage:
  base [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  start       Starts a long running scraper of data

Flags:
  -h, --help   help for base
```

As you can see, available Commands are commands that we ADD, we can see the `start` which is an Example Command.

To run it, simply add `start` after the ROot cmd.

```bash
go run main.go start
```

This will print any information about the command, and in this case we have a required `modbusIP` argument.

```log
Error: required flag(s) "modbusIP" not set
Usage:
  base start [flags]

Flags:
  -h, --help              help for start
  -m, --modbusIP string   The IP address of the ModBus unit

required flag(s) "modbusIP" not set
exit status 1
```

Add it and see it run ! 

