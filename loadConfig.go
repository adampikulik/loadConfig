package loadConfig

import (
 	"bufio"
 	"io"
 	"os"
 	"strings"
)

type Config map[string]string

func ReadConfig(filename string) (Config, error) {
     // init with some bogus data
 	config := Config{
 		"PRIVATE_KEY":"11111111-1111-1111-1111-111111111111",
 		"APP_NAME":"default_app",
 		"SUB_NAME":"default_sub",
 		"ES_KEY":"11111111-1111-1111-1111-111111111111",
 		"QAPP_NAME":"default_app",
       		"QSUB_NAME":"default_sub",
 		"TS_FIELD":"timestamp",
 		"LAYOUT":"2006-01-02 15:04:05.0000",
 		"TZ":"Asia/Jerusalem",
 	}
 	if len(filename) == 0 {
 		return config, nil
 	}
 	file, err := os.Open(filename)
 	if err != nil {
 		return nil, err
 	}
 	defer file.Close()
 	
 	reader := bufio.NewReader(file)
 	
 	for {
 		line, err := reader.ReadString('\n')
 		
 		// check if the line has = sign
             // and process the line. Ignore the rest.
 		if equal := strings.Index(line, "="); equal >= 0 {
 			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
 				value := ""
 				if len(line) > equal {
 					value = strings.TrimSpace(line[equal+1:])
 				}
                // assign the config map
 				config[key] = value
 			}
 		}
 		if err == io.EOF {
 			break
 		}
 		if err != nil {
 			return nil, err
 		}
 	}
 	if os.Getenv("PRIVATE_KEY") != "" {
 		config["PRIVATE_KEY"] = os.Getenv("ES_KEY")
 		}
  	if os.Getenv("ES_KEY") != "" {
 		config["ES_KEY"] = os.Getenv("ES_KEY")
 		}

 	return config, nil
}
