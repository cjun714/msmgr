package eureka

// {
//     "instance": {
//         "hostName": "WKS-SOF-L011",
//         "app": "com.automationrhapsody.eureka.app",
//         "vipAddress": "com.automationrhapsody.eureka.app",
//         "secureVipAddress": "com.automationrhapsody.eureka.app",
//         "ipAddr": "10.0.0.10",
//         "status": "STARTING",
//         "port": {"$": "8080", "@enabled": "true"},
//         "securePort": {"$": "8443", "@enabled": "true"},
//         "healthCheckUrl": "http://WKS-SOF-L011:8080/healthcheck",
//         "statusPageUrl": "http://WKS-SOF-L011:8080/status",
//         "homePageUrl": "http://WKS-SOF-L011:8080",
//         "dataCenterInfo": {
//             "@class": "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo",
//             "name": "MyOwn"
//         }
//     }
// }

// RegInfo is POJO mapped JSON used to register into Eureka
type RegInfo struct {
	Instance struct {
		HostName         string `json:"hostName"`
		App              string `json:"app"`
		VipAddress       string `json:"vipAddress"`
		SecureVipAddress string `json:"secureVipAddress"`
		IPAddr           string `json:"ipAddr"`
		Status           string `json:"status"`
		Port             struct {
			Pt     string `json:"$"`
			Enable string `json:"@enabled"`
		} `json:"port"`
		SecurePort struct {
			Pt     string `json:"$"`
			Enable string `json:"@enabled"`
		} `json:"securePort"`
		HealthCheckURL string `json:"healthCheckUrl"`
		StatusPageURL  string `json:"statusPageUrl"`
		HomePageURL    string `json:"homePageUrl"`
		DataCenterInfo struct {
			Class string `json:"@class"`
			Name  string `json:"name"`
		} `json:"dataCenterInfo"`
	} `json:"instance"`
}

type APP struct {
	Application struct {
		Name     string `json:"name"`
		Instance []struct {
			HostName         string `json:"hostName"`
			App              string `json:"app"`
			VipAddress       string `json:"vipAddress"`
			SecureVipAddress string `json:"secureVipAddress"`
			IPAddr           string `json:"ipAddr"`
			Status           string `json:"status"`
			Port             struct {
				Pt     int    `json:"$"`
				Enable string `json:"@enabled"`
			} `json:"port"`
			SecurePort struct {
				Pt     int    `json:"$"`
				Enable string `json:"@enabled"`
			} `json:"securePort"`
			HealthCheckURL string `json:"healthCheckUrl"`
			StatusPageURL  string `json:"statusPageUrl"`
			HomePageURL    string `json:"homePageUrl"`
			DataCenterInfo struct {
				Class string `json:"@class"`
				Name  string `json:"name"`
			} `json:"dataCenterInfo"`
		} `json:"instance"`
	} `json:"application"`
}
