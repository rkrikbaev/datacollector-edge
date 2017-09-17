package websocket

import (
	"github.com/streamsets/datacollector-edge/container/common"
	"github.com/streamsets/datacollector-edge/container/creation"
	"testing"
)

func getStageContext(
	resourceUrl string,
	headers []interface{},
	parameters map[string]interface{},
) *common.StageContextImpl {
	stageConfig := common.StageConfiguration{}
	stageConfig.Library = LIBRARY
	stageConfig.StageName = STAGE_NAME
	stageConfig.Configuration = []common.Config{
		{
			Name:  "conf.resourceUrl",
			Value: resourceUrl,
		},
		{
			Name:  "conf.headers",
			Value: headers,
		},
		{
			Name:  "conf.dataFormat",
			Value: "JSON",
		},
	}
	return &common.StageContextImpl{
		StageConfig: stageConfig,
		Parameters:  parameters,
	}
}

func TestWebSocketClientDestination_Init(t *testing.T) {
	resourceUrl := "http://test:9000"
	headers := make([]interface{}, 2)
	headers[0] = map[string]interface{}{
		"key":   "X-SDC-APPLICATION-ID",
		"value": "SDCe",
	}
	headers[1] = map[string]interface{}{
		"key":   "DUMMY-HEADER",
		"value": "DUMMY",
	}

	stageContext := getStageContext(resourceUrl, headers, nil)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters)
	if err != nil {
		t.Error(err)
		return
	}

	stageInstance := stageBean.Stage
	if stageInstance == nil {
		t.Error("Failed to create stage instance")
	}

	if stageInstance.(*WebSocketClientDestination).Conf.ResourceUrl != resourceUrl {
		t.Error("Failed to inject config value for resourceUrl")
	}

	if stageInstance.(*WebSocketClientDestination).Conf.Headers == nil {
		t.Error("Failed to inject config value for Headers")
		return
	}

	err = stageInstance.Init(stageContext)
	if err != nil {
		t.Error(err)
	}

	if stageInstance.(*WebSocketClientDestination).Conf.DataGeneratorFormatConfig.RecordWriterFactory == nil {
		t.Error("Failed to initialize RecordWriterFactory")
	}
}
