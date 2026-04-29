package main

import ("fmt"
		"log"
		"os"
		"log/slog")

func PrintLogger () {
	fmt.Println("*******************Printing logger*************")
	
	// To control how logs look.
	log.SetPrefix("APP: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) 

	log.Println("log with new line")
	log.Print("log without new line")

	// It may be useful to create a custom logger and pass it around. 
	// When creating a new logger, we can set a prefix to distinguish its output from other loggers.
	newLogger := log.New(os.Stdout, "newInstanceLogger :", log.LstdFlags ) // Here newInstanceLogger can be used for log level like INFO, ERROR
	newLogger.Print("testing new instance")

	// The slog package provides structured log output, recommended to use in prod
	slog.Info("I am info log", "arg1Key", "arg1value", "arg2key", "arg2value") // arg
	slog.Error("I am error log")
	customLog := slog.With("defaultKey", "defaultvalue") // this is a way to always log something i.e. logging with context
	customLog.Info("test", "sdfd", "sdfdf")



	log.Fatal("Log and exit")

}