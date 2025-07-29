package main

import (
    "os"

    "k8s.io/component-base/cli"
    "k8s.io/kubernetes/cmd/kube-scheduler/app"

    "stackelberg-scheduler-plugin/pkg/plugin"
)

func main() {
    command := app.NewSchedulerCommand(
        app.WithPlugin(plugin.StackelbergPluginName, plugin.New),
    )

    code := cli.Run(command)
    os.Exit(code)
}