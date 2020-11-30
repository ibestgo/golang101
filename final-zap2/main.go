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
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore" 
	"github.com/natefinch/lumberjack"	// for logfile rotation that zap doesn't support
)

var sugarLogger *zap.SugaredLogger

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)		// OR NewJSONEncoder
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger {
		Filename: "./test.log",			
		MaxSize: 10,				// max log file size in MB before rotated		
		MaxBackups: 5,				// max # of old files to retain 
		MaxAge: 30,					// max # of days to retain old files
		Compress: false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func InitLogger() {
	encoder := getEncoder()
	writerSyncer := getLogWriter()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func main() {
	InitLogger()
	defer sugarLogger.Sync()	// flushes buffer if any
	SimpleHttpGet("www.google.com")		
	SimpleHttpGet("http://www.google.com")
}

func SimpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)		
	resp, err := http.Get(url)		
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)			
		resp.Body.Close()
	}
}