package exporter

import (
	"bytes"
	"fmt"
	pcg "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"io"
	"net/http"
	"time"
)

func getNebulaMetrics(ipAddress string, port int32) (map[string]*pcg.MetricFamily, error) {
	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	resp, err := httpClient.Get(fmt.Sprintf("http://%s:%d/metrics?format=prometheus", ipAddress, port))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	parser := expfmt.TextParser{}
	return parser.TextToMetricFamilies(bytes.NewReader(body))
}
