package customhttpexporter

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"go.opentelemetry.io/collector/pdata/pmetric"
)

// No default function for this. It must be implemented
// Note: You can change the function name if you like
func pushMetrics(
	ctx context.Context,
	td pmetric.Metrics,
) (err error) {
	fmt.Println("THIS IS A CUSTOM EXPORTER!")
	const myurl = "http://localhost:6379"
	xmlbody := createRawMetricsPayload()

	resp, err := http.Post(myurl, "text/xml", strings.NewReader(xmlbody))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Do something with resp.Body
	return nil

}

func createRawMetricsPayload() string {
	payload := `
	<ns:save-raw-metrics-request
	xmlns:ns="http://services.www.up.com/monitorable-asset/save-raw-metrics-1_0"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xmlns:id="http://schemas.www.up.com/monitorable-asset/common/identifier-1_0"
	xmlns:env="http://schemas.www.up.com/change-management/common/environments-1_1">
<ns:metric-list>
	<ns:metric>
		<id:name>omhqp14e32</id:name>
		<id:type>CHK_PTC_MDM_DOC</id:type>
		<env:environment>TEST</env:environment>
		<id:metric-name>status</id:metric-name>
		<ns:value>NORMAL</ns:value>
<!--            <ns:occurrence-time>2023-03-17T20:37:29Z</ns:occurrence-time>-->
<!--            <ns:collection-time>2023-03-17T20:37:29Z</ns:collection-time>-->
<!--            <ns:correlation-id>corrId</ns:correlation-id>-->
		<ns:details>testing</ns:details>
<!--            <ns:physical-host>physHost1</ns:physical-host>-->
	</ns:metric>
</ns:metric-list>
</ns:save-raw-metrics-request>`

	return payload
}
