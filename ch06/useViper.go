package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func aliasNormalizeFunc(f *pflag.FlagSet, n string) pflag.NormalizedName {
	switch n {
	case "pass":
		n = "password"
		break
	case "ps":
		n = "password"
		break
	}
	return pflag.NormalizedName(n)
}

func main() {
	pflag.StringP("name", "n", "Mike", "Name parameter")
	pflag.StringP("password", "p", "hardToGuess", "Password")
	pflag.CommandLine.SetNormalizeFunc(aliasNormalizeFunc)

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	name := viper.GetString("name")
	password := viper.GetString("password")

	fmt.Println(name, password)

	// Reading an Environment variable
	viper.BindEnv("GOMAXPROCS")
	val := viper.Get("GOMAXPROCS")
	if val != nil {
		fmt.Println("GOMAXPROCS:", val)
	}

	// Setting an Environment variable
	viper.Set("GOMAXPROCS", 16)
	val = viper.Get("GOMAXPROCS")
	fmt.Println("GOMAXPROCS:", val)
}
