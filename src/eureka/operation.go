package eureka

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"util/json"
	"util/log"
	"util/nic"
)

// Register itself to Eureka
func Register(eurekaURL string, appName string) error {
	url := eurekaURL + "/eureka/apps/" + appName
	log.H("register: ", url)

	ipAddr, _ := nic.GetIPByNicName("WLAN")

	regInfo := RegInfo{}
	regInfo.Instance.HostName = nic.GetHostName() // be used as instanceID in Eureka
	regInfo.Instance.App = appName
	regInfo.Instance.VipAddress = appName
	regInfo.Instance.SecureVipAddress = appName
	regInfo.Instance.IPAddr = ipAddr
	regInfo.Instance.Status = "UP"
	regInfo.Instance.Port.Pt = "8080"
	regInfo.Instance.Port.Enable = "true"
	regInfo.Instance.SecurePort.Pt = "8080"
	regInfo.Instance.SecurePort.Enable = "true"
	regInfo.Instance.HealthCheckURL = "http://" + appName + "/status"
	regInfo.Instance.StatusPageURL = "http://" + appName + "/status"
	regInfo.Instance.HomePageURL = "http://" + appName + "/health"
	regInfo.Instance.DataCenterInfo.Class = "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo"
	regInfo.Instance.DataCenterInfo.Name = "MyOwn" // must set this value

	bts := json.ToJSON(regInfo)

	bodybuf := bytes.NewBuffer([]byte(bts))
	// "http://20.26.33.122:32010/eureka/apps/com.automationrhapsody.eureka.app"

	log.I(string(bts))
	resp, e := http.Post(url, "application/json", bodybuf)
	if e != nil {
		return e
	}
	defer resp.Body.Close()

	log.I("resp: ", resp.Status)

	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return e
	}

	if len(body) != 0 {
		log.E(string(body))
	}

	return nil
}

// Cancel is used to un-register from Eureka
func Cancel(eurekaURL string, appName string) error {
	url := eurekaURL + "/eureka/apps/" + appName + "/" + nic.GetHostName()
	log.H("unregister: ", url)

	client := http.Client{}
	request, e := http.NewRequest("DELETE", url, nil)
	if e != nil {
		panic(e)
	}
	resp, e := client.Do(request)
	if e != nil {
		return e
	}
	defer resp.Body.Close()

	log.I("resp: ", resp.Status)

	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return e
	}
	if len(body) != 0 {
		log.E(string(body))
	}

	return nil
}

// Renew is used to renew register
func Renew(eurekaURL string, appName string) error {
	// url := "http://20.26.33.122:32010/eureka/apps/com.automationrhapsody.eureka.app/WKS-SOF-L011"

	url := eurekaURL + "/eureka/apps/" + appName + "/" + nic.GetHostName()
	log.I("renew:", url)

	client := &http.Client{}
	request, e := http.NewRequest("PUT", url, nil)

	if e != nil {
		log.E(e)
		return e
	}

	resp, e := client.Do(request)
	if e != nil {
		return e
	}
	defer resp.Body.Close()

	log.I("resp: ", resp.Status)

	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return e
	}
	if len(body) != 0 {
		log.E(string(body))
	}

	return nil
}

// Query is used to query service
func Query(eurekaURL string, appName string) ([]string, error) {
	// url := "http://20.26.33.122:32010/eureka/apps/com.automationrhapsody.eureka.app/WKS-SOF-L011"

	url := eurekaURL + "/eureka/apps/" + appName
	log.I("Query:", url)

	client := &http.Client{}
	request, e := http.NewRequest("GET", url, nil)
	request.Header.Add("Accept", "text/html,application/json;q=0.9,*/*;q=0.8")

	if e != nil {
		log.E(e)
		return nil, e
	}

	resp, e := client.Do(request)
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()

	log.I("resp: ", resp.Status)

	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return nil, e
	}
	if len(body) != 0 {
		log.E(string(body))
	}

	app := APP{}
	e = json.ToObj(body, &app)
	if e != nil {
		log.E(e)
	}
	// log.E("Show IP:", app.Application.Name)

	count := len(app.Application.Instance)
	result := make([]string, count)

	for i := 0; i < count; i++ {
		result[i] = app.Application.Instance[i].HostName + ":" + strconv.Itoa(app.Application.Instance[i].Port.Pt)
	}

	return result, nil
}
