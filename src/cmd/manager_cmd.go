package cmd

import (
	"github.com/Dataman-Cloud/swan/src/config"
	"github.com/Dataman-Cloud/swan/src/manager"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"golang.org/x/net/context"
)

func ManagerCmd() cli.Command {
	managerCmd := cli.Command{
		Name:        "manager",
		Usage:       "start a manager instance",
		Description: "start a swan manager",
		Action:      JoinAndStartManager,
	}

	managerCmd.Flags = append(managerCmd.Flags, FlagListenAddr())
	managerCmd.Flags = append(managerCmd.Flags, FlagAdvertiseAddr())
	managerCmd.Flags = append(managerCmd.Flags, FlagZkPath())
	managerCmd.Flags = append(managerCmd.Flags, FlagLogLevel())

	return managerCmd
}

func ManagerInitCmd() cli.Command {
	managerInitCmd := cli.Command{
		Name:        "init",
		Usage:       "init [ARG...]",
		Description: "start a manager",
		Flags:       []cli.Flag{},
		Action:      JoinAndStartManager,
	}

	return managerInitCmd
}

func JoinAndStartManager(c *cli.Context) error {
	conf := config.NewManagerConfig(c)
	setupLogger(conf.LogLevel)

	managerNode, err := manager.New(conf)
	if err != nil {
		logrus.Error("Node initialization failed")
		return err
	}

	if err := managerNode.InitAndStart(context.TODO()); err != nil {
		logrus.Errorf("start node failed. Error: %s", err.Error())
		return err
	}

	return nil
}
