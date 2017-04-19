package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "MyDatabase"
	app.Usage = "To query some data from the database adn execute binary file"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		// NOTE : Command to count rows
		{
			Name:    "count-rows",
			Aliases: []string{"cr"},
			Usage:   "Count total rows from table",
			Action: func(c *cli.Context) error {
				createdAt := c.Args().Get(0)
				if createdAt == "" {
					fmt.Println("Please use input the date on input  = Y-m-d HH:mm:ss")
					return nil
				}

				db := connect()
				defer db.Close()

				stm := countRows(db)
				defer stm.Close()

				// begin execute query
				result := stm.QueryRow(createdAt)

				var countResult int64
				result.Scan(&countResult)

				fmt.Println("count rows = ", countResult)

				return nil
			},
		},

		// NOTE : count all mo_process
		{
			Name:    "mo-process",
			Aliases: []string{"mp"},
			Usage:   "Count total mo that receive but and not yet processed",
			Action: func(c *cli.Context) error {
				db := connect()
				defer db.Close()

				stm := moReceive(db)
				defer stm.Close()

				// begin execute query
				result := stm.QueryRow()

				var countResult int64
				result.Scan(&countResult)

				fmt.Println("Total Mo Receive = ", countResult)

				return nil
			},
		},

		// NOTE : remove all mo_process
		{
			Name:    "remove-mo-process",
			Aliases: []string{"rmp"},
			Usage:   "Remove all mo_process !!! WARNING this will delete all datas",
			Action: func(c *cli.Context) error {
				db := connect()
				defer db.Close()

				stm := moRemove(db)
				defer stm.Close()

				// begin execute query
				_, err := stm.Exec()
				if err != nil {
					logger.Println(err)
					return nil
				}

				fmt.Println("Success delte data")

				return nil
			},
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}
