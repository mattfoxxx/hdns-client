package hdnsclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

const apiToken = "secret_token"
const zoneID = "abcdef0123456789"
const records = "{\"records\":[{\"type\":\"A\",\"id\":\"83ccee3f0f1e32a485afccc344ec966b\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"@\",\"value\":\"138.201.246.71\",\"ttl\":0},{\"type\":\"A\",\"id\":\"263c1aa1d24397ec5e0fef82bee19ffe\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"salt.i\",\"value\":\"10.6.0.2\",\"ttl\":0},{\"type\":\"CNAME\",\"id\":\"05d81131824e77894034b3ae5538bf11\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"ci\",\"value\":\"node02\",\"ttl\":0},{\"type\":\"A\",\"id\":\"88d4b06028f0a37ed4d008385a09a0e4\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"mail\",\"value\":\"88.198.176.160\",\"ttl\":0},{\"type\":\"A\",\"id\":\"2d192f29a80cafc029d7d9f3df384c8b\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"*\",\"value\":\"138.201.246.71\",\"ttl\":0},{\"type\":\"A\",\"id\":\"310caaf7450a10ccae64f8f931c35a65\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"node02.i\",\"value\":\"10.6.0.3\",\"ttl\":0},{\"type\":\"A\",\"id\":\"4a092754adbfc2e988835ac68aae09e5\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"node01\",\"value\":\"138.201.246.71\",\"ttl\":0},{\"type\":\"CNAME\",\"id\":\"4e56937034526e869211305ed327b1a2\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"autoconfig\",\"value\":\"mail.your-server.de.\",\"ttl\":0},{\"type\":\"A\",\"id\":\"feb23da293034585e8e89a53a9d62d01\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"node02\",\"value\":\"94.130.173.77\",\"ttl\":0},{\"type\":\"A\",\"id\":\"20b123385cf2881754f583660f9d56b5\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"node01.i\",\"value\":\"10.6.0.1\",\"ttl\":0},{\"type\":\"CNAME\",\"id\":\"9578f6628d6d43357e4b682d78ae06ca\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"ftp\",\"value\":\"www\",\"ttl\":0},{\"type\":\"CNAME\",\"id\":\"3abe7ada2462ec4b699d99315c20822d\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"git\",\"value\":\"node01\",\"ttl\":0},{\"type\":\"CNAME\",\"id\":\"0423678e8ff141d158c77521b9f70cdd\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"icinga\",\"value\":\"node01\",\"ttl\":0},{\"type\":\"CNAME\",\"id\":\"a806a4de340449bc17f28e53f293fa9e\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"imap\",\"value\":\"mail\",\"ttl\":0},{\"type\":\"CNAME\",\"id\":\"fb856259f91561fe9141a063121544b2\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"monitor.ci\",\"value\":\"node02\",\"ttl\":0},{\"type\":\"CNAME\",\"id\":\"8c6aed4010aee438130e4bf4d3cf9871\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"oldno7\",\"value\":\"oldno7.dyndns.org.\",\"ttl\":0},{\"type\":\"CNAME\",\"id\":\"5962640404b4123adc5fc406b06624d3\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"pop\",\"value\":\"mail\",\"ttl\":0},{\"type\":\"CNAME\",\"id\":\"37384227c5b9552d32d42f0cd39e40f5\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"relay\",\"value\":\"mail\",\"ttl\":0},{\"type\":\"CNAME\",\"id\":\"888765dc3dcf947b4292b897f0978d62\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"smtp\",\"value\":\"mail\",\"ttl\":0},{\"type\":\"CNAME\",\"id\":\"bb1af01871a9c254954fc68513579c69\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"www\",\"value\":\"node01\",\"ttl\":0},{\"type\":\"TXT\",\"id\":\"936a04556787befc7c08904d278268bc\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"@\",\"value\":\"\\\"v=spf1 +a +mx ?all\\\"\",\"ttl\":0},{\"type\":\"MX\",\"id\":\"80ac6477bfd8d15f3c3debf3a4a58b76\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"@\",\"value\":\"10 mail\",\"ttl\":0},{\"type\":\"SRV\",\"id\":\"7d6542aaea885bd294fd4c106f6a4219\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"_autodiscover._tcp\",\"value\":\"0 100 443 mail.your-server.de.\",\"ttl\":0},{\"type\":\"NS\",\"id\":\"f2c7451ef794bbb82e19616bc36627b0\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"@\",\"value\":\"hydrogen.ns.hetzner.com.\",\"ttl\":0},{\"type\":\"NS\",\"id\":\"f0636a5e7c774c7ffeb4e19deebb7dcb\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"@\",\"value\":\"oxygen.ns.hetzner.com.\",\"ttl\":0},{\"type\":\"NS\",\"id\":\"8ce7818ffd302808066d3732ffbb26a4\",\"created\":\"2020-05-04 14:50:38 +0000 UTC\",\"modified\":\"2020-05-04 14:50:38 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"@\",\"value\":\"helium.ns.hetzner.de.\",\"ttl\":0},{\"type\":\"SOA\",\"id\":\"5ce1ba0484198e140307828385137855\",\"created\":\"2020-05-04 14:50:39 +0000 UTC\",\"modified\":\"2020-05-04 14:50:39 +0000 UTC\",\"zone_id\":\"5c95wmRRiFSdwNswRDcMuG\",\"name\":\"@\",\"value\":\"hydrogen.ns.hetzner.com. dns.hetzner.com. 2020050400 86400 10800 3600000 3600\",\"ttl\":0}],\"meta\":{\"pagination\":{\"page\":0,\"per_page\":0,\"last_page\":0,\"total_entries\":0}}}"
const recordID = "deadbeef1234567890"
const record = ""

func TestGetRecords(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		equals(t, fmt.Sprintf("/records?zone_id=%s&per_page=100&page=1", zoneID), r.URL.String())
		rw.WriteHeader(200)
		rw.Write([]byte(records))

	}))
	defer server.Close()

	TestClient := Client{
		server.URL,
		apiToken,
		server.Client(),
	}
	res, err := TestClient.GetRecords(zoneID, nil)
	if err != nil {
		fmt.Printf("An error occured: %s", err)
	}
	jsonResult, err := json.Marshal(res)
	equals(t, records, string(jsonResult))
}
