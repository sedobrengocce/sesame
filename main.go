/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/sedobrengocce/sesame/cmd"
	"github.com/sedobrengocce/sesame/obj/configDefaults"
)

func main() {
    configdefaults.CheckConfigPath()
	cmd.Execute()
}
