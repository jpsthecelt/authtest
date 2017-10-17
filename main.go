package main

import (
	"net/http"
	"net/http/cookiejar"
)

import (
	"encoding/base64"
	"log"
	"os"
	"fmt"
	"encoding/json"
	"flag"
	"strings"
	"io/ioutil"
	"io"
)

// Call this routine (for example) as:
//      ./authtest -auth ~/auth.txt -format json -computername A122758 -out savedOutput.txt

// What the config file looks like
type Config struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Serverurl string `json:"serverurl"`
}
var cfg Config

// Read username/password/url from config-file & return
func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
		}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

// Return base64-encoded authorization-string
func basicAuth(username, password string) string {
	  auth := username + ":" + password
  return base64.StdEncoding.EncodeToString([]byte(auth))
}

// handle server redirect & resend user/password
func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	  req.SetBasicAuth(cfg.Username,cfg.Password)
  return nil
}

func fileOutput(filepath, s string) error {
	fo, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fo.Close()

	_, err = io.Copy(fo, strings.NewReader(s))
	if err != nil {
		return err
	}

	return nil
}

func main() {

	jar, _ := cookiejar.New(nil)
	client := &http.Client {
	  Jar: jar,
	  CheckRedirect: redirectPolicyFunc,
	}

	// Assumes that the first argument is '-auth FQDN', no '~' and uses '/'s vs. '\'s
  	aPtr := flag.String("auth", ".", "an FQDN")
	fPtr := flag.String("format", "xml", "either xml or json")
	cPtr := flag.String("computername", "A123456", "name of desired computer")
	oPtr := flag.String("out", "savedfile.txt", "name of desired output-file")
	flag.Parse()

  	if len(*aPtr) > 0 {
  		fmt.Println("Using", *aPtr, "as the authorization file: ")
  	} else {
  		log.Fatal("\nERROR** - No auth file parameter on command-line>")
  	}

  	if len(*cPtr) > 0 {
  		fmt.Println("Finding history of", *cPtr, "...")
  	} else {
  		log.Fatal("\nERROR** - No indicated computer on command-line>")
  	}

  	sendOut := false
	if len(*oPtr) > 0 {
	   sendOut = true
	}

  	outputFormat := "text/xml"
	if (len(*fPtr) > 0) && (strings.Contains(*fPtr,"json")) {
		outputFormat = "application/json"
	}

	cfg = LoadConfiguration(*aPtr)
	//println(cfg.Username, cfg.Password, cfg.Serverurl)

	req, err := http.NewRequest("GET", cfg.Serverurl + "/JSSResource/computerhistory/name/"+*cPtr, nil)
    req.Header.Add("Authorization", "Basic " + basicAuth(cfg.Username, cfg.Password))
    req.Header.Add("accept",outputFormat)

    resp, err := client.Do(req)
	if err != nil {
	  	log.Fatal("Oh, crap; done screwed up on the Casper request...")
	} else {
		println("Get status was: ", resp.Status)
		if resp.StatusCode == 200 { // OK
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			bodyString := string(bodyBytes)
			println(bodyString)
			if sendOut {
				ok := fileOutput(*oPtr, bodyString )
				if ok != nil {
					log.Fatal("Cant write to specified output file")
				}
			}
		}
	}
}
