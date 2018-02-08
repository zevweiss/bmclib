package discover

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/ncode/bmc/dell"
	"github.com/ncode/bmc/devices"
	"github.com/ncode/bmc/hp"
	"github.com/ncode/bmc/httpclient"
	"github.com/ncode/bmc/ilo"
	"github.com/ncode/bmc/supermicrox10"
)

// ScanAndConnect will scan the bmc trying to learn the device type and return a working connection
func ScanAndConnect(host string, username string, password string) (bmcConnection interface{}, err error) {
	log.WithFields(log.Fields{"step": "ScanAndConnect", "host": host}).Debug("detecting vendor")
	var vendor string

	client, err := httpclient.Build()
	if err != nil {
		return bmcConnection, err
	}

	resp, err := client.Get(fmt.Sprintf("http://%s/res/ok.png", host))
	if err != nil {
		return bmcConnection, err
	}
	io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		log.WithFields(log.Fields{"step": "ScanAndConnect", "host": host, "vendor": devices.Cloudline}).Debug("it's a discrete")
		return bmcConnection, httpclient.ErrVendorNotSupported
	}

	resp, err = client.Get(fmt.Sprintf("https://%s/xmldata?item=all", host))
	if err != nil {
		return bmcConnection, err
	}
	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return bmcConnection, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		iloXMLC := &hp.Rimp{}
		err = xml.Unmarshal(payload, iloXMLC)
		if err != nil {
			return bmcConnection, err
		}

		if iloXMLC.Infra2 != nil {
			log.WithFields(log.Fields{"step": "ScanAndConnect", "host": host, "vendor": devices.HP}).Debug("it's a chassis")
			// TODO: Return chassis here
			return bmcConnection, err
		}

		iloXML := &hp.RimpBlade{}
		err = xml.Unmarshal(payload, iloXML)
		if err != nil {
			return bmcConnection, err
		}

		if iloXML.HSI != nil {
			var isKnow bool
			if strings.HasPrefix(iloXML.MP.Pn, "Integrated Lights-Out") {
				return ilo.New(host, username, password)
			}

			return bmcConnection, fmt.Errorf("it's an HP, but I cound't not identify the hardware type. Please verify")
		}
	}

	resp, err = client.Get(fmt.Sprintf("https://%s/session?aimGetProp=hostname,gui_str_title_bar,OEMHostName,fwVersion,sysDesc", host))
	if err != nil {
		return bmcConnection, err
	}

	payload, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return bmcConnection, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		dellJSON := &dell.HwDetection{}
		err = json.Unmarshal(payload, dellJSON)
		if err != nil {
			return bmcConnection, err
		}

		return bmcConnection, err
	}
	resp, err = client.Get(fmt.Sprintf("https://%s/cgi-bin/webcgi/login", host))
	if err != nil {
		return bmcConnection, err
	}
	io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		log.WithFields(log.Fields{"step": "connection", "host": host, "vendor": devices.Dell}).Debug("it's a chassis")
		// TODO: Return a chassis here
		return bmcConnection, err
	}

	resp, err = client.Get(fmt.Sprintf("https://%s/cgi/login.cgi", host))
	if err != nil {
		return bmcConnection, err
	}
	io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return supermicrox10.New(host, username, password)
	}

	return bmcConnection, httpclient.ErrVendorUnknown
}
