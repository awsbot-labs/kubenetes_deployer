package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	app := cli.NewApp()
	app.Name = "Kubenetes Deployer"
	app.Usage = "creates a kubenetes cluster in a Amazon Web Services"
	app.Version = "1"
	app.Flags = []cli.Flag{}

	app.Commands = cli.Commands{
		{
			Name:  "aws",
			Usage: "Deploys to Amazon Web Services",
			Flags: ,
			Action: func(c *cli.Context) error {
				fmt.Println("Deploying to Amazon Web Services")
				sess, err := session.NewSession(&aws.Config{
					Region: aws.String("us-west-2")},
				)
				fmt.Println(sess)
				if err != nil {
					log.Fatal(err)
				}
				return nil
			},
		},
	}

	// start our application
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}