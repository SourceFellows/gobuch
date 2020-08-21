/*Package main zeigt einen einfachen HTTP Service, der getestet werden soll.

Das Abschalten des Power-Safe Modus unter Linux hilft folgendes Script:

	for i in /sys/devices/system/cpu/cpu[0-7]
	do
    	echo performance > $i/cpufreq/scaling_governor
	done

Die Laufzeitgeschwindigkeit kann zusätzlich über 'ab' (Apache HTTP server benchmarking tool -
https://httpd.apache.org/docs/2.4/programs/ab.html) betestet werden.
*/
package main

import (
	"encoding/json"
	"net/http"
	_ "net/http/pprof"
	"regexp"
)

type Result struct {
	Name string
}

func myHandler(writer http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	re := regexp.MustCompile("^/json/(.*)$")

	if !re.MatchString(path) {
		writer.WriteHeader(401)
		return
	}

	parts := re.FindStringSubmatch(path)
	res := &Result{Name: parts[1]}
	bites, err := json.Marshal(*res)
	if err != nil {
		writer.WriteHeader(500)
		return
	}
	writer.Header().Add("Content-Type", "application/json")
	writer.Write(bites)
}

func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":8080", nil)
}

