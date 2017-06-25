package main

import (
	"config"
	"eureka"
	"fmt"
	"io/ioutil"
	"k8"
	"net/http"
	"os"
	"time"
	"unsafe"
	"util/file"
	"util/json"
	"util/log"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// /list
func list(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	list, e := file.GetFileInfoList(config.YamlPath)
	if e != nil {
		log.E(e.Error())
		fmt.Fprint(w, "Error!\n")
		return
	}

	bytes := json.ToJSON(list)
	writeJSON(w, bytes)
}

// /read/{id}
func read(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	svcInfo, e := loadServiceInfo(id)
	if e != nil {
		log.E(e.Error())
		fmt.Fprint(w, "Error!\n")
		return
	}
	bytes := json.ToJSON(*svcInfo)
	writeJSON(w, bytes)
}

// /save/{id}
func save(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jsonBytes, e := getBodyFromReq(r) // get json from request's body
	if e != nil {
		log.E(e)
		writeJSON(w, json.ToJSON(false))
	}

	svcInfo := k8.ServiceInfo{}
	e = json.ToObj(jsonBytes, &svcInfo)
	if e != nil {
		log.E(e)
		writeJSON(w, json.ToJSON(false))
		return
	}

	e = k8.CreateYaml(svcInfo)
	if e != nil {
		log.E(e)
		writeJSON(w, json.ToJSON(false))
		return
	}
	writeJSON(w, json.ToJSON(true))
	return
}

// /stop/{id}
func stop(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	if !k8.Stop(config.YamlPath + id) {
		writeJSON(w, json.ToJSON(false))
		return
	}
	bytes := json.ToJSON(true)
	writeJSON(w, bytes)
}

// /run/{id}
func run(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	if !k8.Create(config.YamlPath + id) {
		writeJSON(w, json.ToJSON(false))
		return
	}
	bytes := json.ToJSON(true)
	writeJSON(w, bytes)
}

// /delete/{id}
func delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	k8.Stop(config.YamlPath + id)
	e := file.Delete(config.YamlPath + id) // ../yamls/<id>
	if e != nil {
		log.E(e)
		writeJSON(w, json.ToJSON(false))
		return
	}
	bytes := json.ToJSON(true)
	writeJSON(w, bytes)
}

// /find-user/:name
func findUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	user := struct{ Name, Email string }{name, name + "@hpe.com"}

	bytes := json.ToJSON(user)
	writeJSON(w, bytes)
}

func query(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("service")
	result, e := eureka.Query(config.EurekaURL, name)
	if e != nil {
		fmt.Fprint(w, "Error to query:"+name)
		return
	}
	if len(result) == 0 {
		fmt.Fprint(w, "No result:", name)
		return 
	}

	bytes := json.ToJSON(result)
	writeJSON(w, bytes)
}

// /quit/
func quit(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Close server at "+time.Now().Format("15:04:05")+"!\n")
	go func() {
		log.H("Close server")
		time.Sleep(1000 * time.Millisecond)
		eureka.Cancel(config.EurekaURL, config.AppName)
		os.Exit(0)
	}()
}

func loadServiceInfo(id string) (*k8.ServiceInfo, error) {
	svcInfo := new(k8.ServiceInfo)

	bts, e := file.Read(config.YamlPath + id + "/svc.yaml")
	if e != nil {
		return nil, e
	}
	svc := k8.NewSVC(bts) //k8.SVCTmpl

	bts, e = file.Read(config.YamlPath + id + "/rc.yaml")
	if e != nil {
		return nil, e
	}
	rc := k8.NewRC(bts)

	svcInfo.Name = svc.Metadata.Name
	svcInfo.ImgURL = rc.Spec.Template.Spec.Containers[0].Image
	svcInfo.CPUQuota = rc.Spec.Template.Spec.Containers[0].Resources.Limits.CPU
	svcInfo.RAMQuota = rc.Spec.Template.Spec.Containers[0].Resources.Limits.Memory
	svcInfo.ClusterPort = svc.Spec.Ports[0].Port
	svcInfo.NodePort = svc.Spec.Ports[0].NodePort
	svcInfo.TargetPort = svc.Spec.Ports[0].TargetPort
	return svcInfo, nil
}

func writeJSON(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	fmt.Fprintf(w, *(*string)(unsafe.Pointer(&bytes)))
}

func getBodyFromReq(req *http.Request) ([]byte, error) {
	result, e := ioutil.ReadAll(req.Body)
	if e != nil {
		return nil, e
	}
	return result, nil
}
