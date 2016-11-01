// Package bravagiulia controls a SONY BRAVIA TV model using a Pre-Shared-Key
package bravagiulia

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/parnurzeal/gorequest"
)

// Client is the bravagiulia client
type Client struct {
	IP  string
	PSK string
}

var irccBody = `<?xml version="1.0"?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" 
            s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
  <s:Body>
    <u:X_SendIRCC xmlns:u="urn:schemas-sony-com:service:IRCC:1">
      <IRCCCode>{IRCC}</IRCCCode>
    </u:X_SendIRCC>
  </s:Body>
</s:Envelope>`

// NewClient creates a new bravagiulia client
func NewClient(ip string, psk string) Client {
	return Client{IP: ip, PSK: psk}
}

// RemoteControllerInfo is part of the getRemoteControllerInfo json response parser
type RemoteControllerInfo struct {
	Result [][]RemoteControllerCommand `json:"result"`
}

// RemoteControllerCommand is part of the getRemoteControllerInfo json response parser
type RemoteControllerCommand struct {
	Name  string
	Value string
}

// GetSupportedCommands returns a map from command names to command
// codes for the SendIRCC method
func (c Client) GetSupportedCommands() map[string]string {
	commands := make(map[string]string)
	url := strings.Replace("http://{ip}/sony/system", "{ip}", c.IP, 1)
	resp, body, _ := gorequest.New().Post(url).
		Type("json").
		Send(`{"method":"getRemoteControllerInfo","params":[""],"id":1,"version":"1.0"}`).
		End()
	if resp.StatusCode == 200 {
		var jb RemoteControllerInfo
		json.Unmarshal([]byte(body), &jb)
		log.Println("JB", jb)
		jresult := jb.Result[1]
		log.Println("JResult", jresult)
		for _, c := range jresult {
			commands[c.Name] = c.Value
			log.Println(c.Name, c.Value)
		}
	}
	return commands
}

// SendIRCC Sends an IRCC command string to the specified client
func (c Client) SendIRCC(command string) []error {
	url := strings.Replace("http://{ip}/sony/IRCC", "{ip}", c.IP, 1)
	body := strings.Replace(irccBody, "{IRCC}", command, 1)
	_, _, errs := gorequest.New().
		Post(url).
		Set("X-Auth-PSK", c.PSK).
		Type("xml").
		Set("SOAPACTION", "urn:schemas-sony-com:service:IRCC:1#X_SendIRCC").
		Send(body).
		End()
	return errs
}
