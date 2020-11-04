package hdnsclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

const recordID = "deadbeef1234567890"
const record = ""

func TestGetRecords(t *testing.T) {
	t.Parallel()

	Convey("Given a test server", t, func() {
		fakeRecordTime, _ := time.Parse(ctLayout, "2020-05-04 14:50:38 +0000 UTC")
		recordsResponse := ListResp{
			[]Record{Record{
				Type:     "A",
				ID:       "83ccee3f0f1e32a485afccc344ec966b",
				Created:  CustomTime{fakeRecordTime},
				Modified: CustomTime{fakeRecordTime},
				ZoneID:   "5c95wmRRiFSdwNswRDcMuG",
				Name:     "@",
				Value:    "1.1.1.1",
				TTL:      0,
			},
				Record{
					Type:     "A",
					ID:       "263c1aa1d24397ec5e0fef82bee19ffe",
					Created:  CustomTime{fakeRecordTime},
					Modified: CustomTime{fakeRecordTime},
					ZoneID:   "5c95wmRRiFSdwNswRDcMuG",
					Name:     "node01",
					Value:    "2.2.2.2",
					TTL:      0,
				},
				Record{
					Type:     "CNAME",
					ID:       "05d81131824e77894034b3ae5538bf11",
					Created:  CustomTime{fakeRecordTime},
					Modified: CustomTime{fakeRecordTime},
					ZoneID:   "5c95wmRRiFSdwNswRDcMuG",
					Name:     "node02",
					Value:    "node01",
					TTL:      0,
				}},
			Meta{
				Pagination{
					Page:         0,
					PerPage:      0,
					LastPage:     0,
					TotalEntries: 0,
				},
			},
		}
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.WriteHeader(200)
			response, _ := json.Marshal(recordsResponse)
			rw.Write(response)

		}))
		defer server.Close()

		Convey("and a test client", func() {
			const apiToken = "secret_token"
			TestClient := Client{
				server.URL,
				apiToken,
				server.Client(),
			}
			Convey("given a zone id", func() {
				zoneID := "abcdef0123456789"

				Convey("a request should return all records", func() {
					res, err := TestClient.GetRecords(zoneID, nil)
					if err != nil {
						fmt.Printf("An error occured: %s", err)
					}
					jsonResult, _ := json.Marshal(res)
					expectedResult, _ := json.Marshal(recordsResponse)
					So(string(jsonResult), ShouldEqual, string(expectedResult))
					So(len(res.Records), ShouldEqual, 3)
				})
			})
		})
	})
}
