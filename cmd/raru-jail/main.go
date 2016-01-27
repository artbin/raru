package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/ArtemKulyabin/raru"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "raru-jail"
	app.Version = "0.0.1"
	app.Usage = "run program as random user"
	app.Commands = []cli.Command{
		{
			Name: "run",
			Usage: `raru-jail run [options] command [arguments...]

   The raru-jail utility changes its current and root directories to the
   supplied directory --chroot and then exec's command, if supplied, or an
   interactive copy of the user's login shell.

   The following environment variable is referenced by raru-jail:

   SHELL  If set, the string specified by SHELL is interpreted as the name
          of the shell to exec. If the variable SHELL is not set, /bin/sh
          is used.
`,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "chroot, c",
					Usage: "Chroot before execute",
				},
				cli.BoolFlag{
					Name:  "fork, f",
					Usage: "Fork and execute",
				},
			},
			Action: run,
		},
		{
			Name:   "build",
			Usage:  "raru-jail build [options] jail-name [programs...]",
			Action: build,
		},
	}
	app.Run(os.Args)
}

func run(c *cli.Context) {
	name := c.Args().First()
	tail := c.Args().Tail()

	if name == "" {
		if name = os.Getenv("SHELL"); name == "" {
			name = "sh"
		}
		tail = []string{"-i"}
	}

	exer, err := raru.NewExecutor()
	if err != nil {
		log.Printf("raru critical failure: %s", err.Error())
		os.Exit(1)
	}

	if c.IsSet("chroot") {
		exer.SetChrootDir(c.String("chroot"))
	}

	if c.Bool("fork") {
		cmd := exec.Command(name, tail...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := exer.Spawn(cmd); err != nil {
			log.Printf("raru critical failure: %s", err.Error())
			os.Exit(1)
		}
	} else {
		if err := exer.Exec(name, tail...); err != nil {
			log.Printf("raru critical failure: %s", err.Error())
			os.Exit(1)
		}
	}
}

func build(c *cli.Context) {
	name := c.Args().First()
	programs := c.Args().Tail()
	err := raru.MkJail(name, programs)
	if err != nil {
		log.Println("Build failed", err)
	}
}
