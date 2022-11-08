package homematicutils

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type HomematicInfo struct {
	Hostname string
	User     string
	Password string
}

func SetIntVar(client http.Client, homematic HomematicInfo, varname string, value int) {
	var rega_cmd string = fmt.Sprintf("dom.GetObject('%s').State('%d');", varname, value)
	req, err := http.NewRequest("POST", homematic.Hostname+"/Test.exe", strings.NewReader(rega_cmd))
	if err != nil {
		log.Fatalln(err)
	}
	req.SetBasicAuth(homematic.User, homematic.Password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
}
