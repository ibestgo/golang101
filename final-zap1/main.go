/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"flag"
	"os"
	"strconv"
	"net/http"
	"log"				// go default log
	"go.uber.org/zap" 
	
	"github.com/baechul/golang/v2/calc"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func init() {
	fmt.Println("Called from init()")

	// go default log setup
	logFile, _ := os.OpenFile("logfile.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	log.SetOutput(logFile)

	// zap logger setup
	logger, _ = zap.NewProduction()	// OR zap.NewDevelopment(), zap.New()
	sugarLogger = logger.Sugar()
}

func main() {
	var addr string
	var enableDebug bool

	// go run . --server-addr https://intuit.com  --enable-debug=false
	flag.StringVar(&addr, "server-addr", "localhost:8080", "The address the bacend server endpoint binds to.")
	flag.BoolVar(&enableDebug, "enable-debug", true, "Enable debug mode")
	flag.Parse()

	fmt.Println("addr: ", addr)
	fmt.Println("debug: ", strconv.FormatBool(enableDebug))
	
	// OR. seems slow though
	// strEnableDebug := fmt.Sprintf("%v", enableDebug) // or %t
	// fmt.Println("debug: ", strEnableDebug)

	// if newResult, err := calc.Add(3, 4, 3); err != nil {
	// 	fmt.Println("Error: => ", err)
	// 	os.Exit(1)
	// } else {
	// 	fmt.Println("calc.Add(3, 4, 3) => ", newResult)
	// }

	// Use Case 1
	newResult, err := calc.Add(3, 4, 3);
	if err != nil {
		fmt.Println("Error: => ", err)
		os.Exit(1)
	} 

	fmt.Println("calc.Add(3, 4, 3) => ", newResult)

	// Use Case 2
	// Using standard logger
	resp, err := http.Get(addr)
	if err != nil {
		log.Printf("Error fetching url %s : %s", addr, err.Error())
		os.Exit(2)
	} else {
		log.Printf("Status code for %s: %s", addr, resp.Status)
		resp.Body.Close()
	}		

	// Using uber-zap: Logger
	resp2, err2 := http.Get(addr)
	if err != nil {
		logger.Error("Error fetching url...",
			zap.String("url", addr),
			zap.Error(err2))
	} else {
		logger.Info("Success...", 
			zap.String("statusCode", resp2.Status),
			zap.String("url", addr))
		resp.Body.Close()	
	}

	// Using uber-zap: SugaredLogger
	defer sugarLogger.Sync()	// flushes buffer if any
	sugarLogger.Debugf("Trying to hit GET request for %s", addr)
	resp3, err3 := http.Get(addr)
	if err != nil {
		sugarLogger.Errorf("Error fetching url %s : Error = %s", addr, err3)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for url %s", resp3.Status, addr)
		resp.Body.Close()	
	}
}