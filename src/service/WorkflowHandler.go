package service

import (
	"encoding/json"
	"fmt"

	tbcommon "github.com/cloud-barista/cb-webtool/src/model/tumblebug/common"
	"github.com/labstack/echo"

	// "io"
	"log"
	"net/http"

	// "os"
	// model "github.com/cloud-barista/cb-webtool/src/model"
	"github.com/cloud-barista/cb-webtool/src/model"
	// spider "github.com/cloud-barista/cb-webtool/src/model/spider"	
	"github.com/cloud-barista/cb-webtool/src/model/cicada"
	// "github.com/cloud-barista/cb-webtool/src/model/tumblebug"

	util "github.com/cloud-barista/cb-webtool/src/util"
)

// 동기 방식 return 받을 때까지 대기.
func RegWorkflow(nameSpaceID string, workflowInfoReq interface{})(model.WebStatus){
	var originalUrl = "/ns/{namespace}/workflow"

	var paramMapper = make(map[string]string)
	paramMapper["{namespace}"] = nameSpaceID
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)
	url := util.CICADA + urlParam

	pbytes, _ := json.Marshal(workflowInfoReq)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	returnWorkflowInfo := cicada.WorkflowInfo{}
	returnStatus := model.WebStatus{}

	if err != nil {
		fmt.Println(err)
		return model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		errorInfo := model.ErrorInfo{}
		json.NewDecoder(respBody).Decode(&errorInfo)
		fmt.Println("respStatus != 200 reason ", errorInfo)
		returnStatus.Message = errorInfo.Message
	} else {
		json.NewDecoder(respBody).Decode(&returnWorkflowInfo)
		fmt.Println(returnWorkflowInfo)
	}
	returnStatus.StatusCode = respStatus

	return returnStatus
}

// 비동기 방식. 호출 후 결과를 socket에 저장
//func RegWorkflow(nameSpaceID string, workflowInfoReq interface{})(model.WebStatus){
func RegWorkflowByAsync(nameSpaceID string, workflowInfoReq interface{}, c echo.Context){
	var originalUrl = "/ns/{namespace}/workflow"

	var paramMapper = make(map[string]string)
	paramMapper["{namespace}"] = nameSpaceID
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)
	url := util.CICADA + urlParam

	taskKey := nameSpaceID + "||" + "workflow" + "||" 

	info, ok := workflowInfoReq.(map[string]interface{})
    if !ok {
        fmt.Println("RegWorkflowAsync wrong request")
		StoreWebsocketMessage(util.TASK_TYPE_WORKFLOW, taskKey, util.WORKFLOW_LIFECYCLE_CREATE, util.TASK_STATUS_FAIL, c) // session에 작업내용 저장
    }
    
    // "name" 필드에 접근
    name, ok := info["name"].(string)
	if !ok {        
		fmt.Println("RegWorkflowAsync there is no name field")
		StoreWebsocketMessage(util.TASK_TYPE_WORKFLOW, taskKey, util.WORKFLOW_LIFECYCLE_CREATE, util.TASK_STATUS_FAIL, c) // session에 작업내용 저장
    }

	pbytes, _ := json.Marshal(workflowInfoReq)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	taskKey = nameSpaceID + "||" + "workflow" + "||" + name

	if err != nil {
		fmt.Println(err)
		StoreWebsocketMessage(util.TASK_TYPE_WORKFLOW, taskKey, util.WORKFLOW_LIFECYCLE_CREATE, util.TASK_STATUS_FAIL, c) // session에 작업내용 저장
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		failResultInfo := tbcommon.TbSimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		log.Println("RegMcksByAsync ", failResultInfo)
		StoreWebsocketMessage(util.TASK_TYPE_WORKFLOW, taskKey, util.WORKFLOW_LIFECYCLE_CREATE, util.TASK_STATUS_FAIL, c) // session에 작업내용 저장

	} else {
		returnWorkflowInfo := cicada.WorkflowInfo{}
		json.NewDecoder(respBody).Decode(&returnWorkflowInfo)
		fmt.Println(returnWorkflowInfo)
		StoreWebsocketMessage(util.TASK_TYPE_WORKFLOW, taskKey, util.WORKFLOW_LIFECYCLE_CREATE, util.TASK_STATUS_COMPLETE, c) // session에 작업내용 저장
	}

}

// Workflow 목록 조회
func GetWorkflowList(nameSpaceID string, optionParam string, filterKeyParam string, filterValParam string) ([]cicada.WorkflowInfo, model.WebStatus) {
	var originalUrl = "/ns/{namespace}/workflow"

	var paramMapper = make(map[string]string)
	paramMapper["{namespace}"] = nameSpaceID
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.CICADA + urlParam
	// url := util.LADYBUG + "/ns/" + nameSpaceID + "/clusters"
	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	// resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	// 원래는 items 와 kind 가 들어오는데
	// kind에는 clusterlist 라는 것만 있고 실제로는 items 에 cluster 정보들이 있음.
	// 그래서 굳이 kind까지 처리하지 않고 item만 return
	workflowList := map[string][]cicada.WorkflowInfo{}
	json.NewDecoder(respBody).Decode(&workflowList)
	fmt.Println(workflowList["items"])
	log.Println(respBody)
	// util.DisplayResponse(resp) // 수신내용 확인

	return workflowList["items"], model.WebStatus{StatusCode: respStatus}
}

// 특정 Cluster 조회
func GetWorkflowData(nameSpaceID string, workflowID string) (*cicada.WorkflowInfo, model.WebStatus) {
	var originalUrl = "/ns/{namespace}/workflow/{workflow}"

	var paramMapper = make(map[string]string)
	paramMapper["{namespace}"] = nameSpaceID
	paramMapper["{workflow}"] = workflowID
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.CICADA + urlParam

	// resp, err := util.CommonHttp(url, nil, http.MethodGet)
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	// defer body.Close()
	workflowInfo := cicada.WorkflowInfo{}
	if err != nil {
		fmt.Println(err)
		return &workflowInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// util.DisplayResponse(resp) // 수신내용 확인

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&workflowInfo)
	fmt.Println(workflowInfo)

	return &workflowInfo, model.WebStatus{StatusCode: respStatus}
}


// Workflow 삭제
func DelWorkflow(nameSpaceID string, workflowId string) (*cicada.WorkflowStatusInfo, model.WebStatus) {
	var originalUrl = "/ns/{namespace}/workflow/{workflow}"

	var paramMapper = make(map[string]string)
	paramMapper["{namespace}"] = nameSpaceID
	paramMapper["{workflow}"] = workflowId
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)
	url := util.CICADA + urlParam

	if workflowId == "" {
		return nil, model.WebStatus{StatusCode: 500, Message: "workflow is required"}
	}

	// 경로안에 parameter가 있어 추가 param없이 호출 함.
	resp, err := util.CommonHttp(url, nil, http.MethodDelete)
	statusInfo := cicada.WorkflowStatusInfo{}
	if err != nil {
		fmt.Println("delCluster ", err)
		return &statusInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&statusInfo)
	fmt.Println(statusInfo)

	if respStatus != 200 && respStatus != 201 {
		fmt.Println(respBody)
		return &statusInfo, model.WebStatus{StatusCode: respStatus, Message: statusInfo.Message}
	}
	return &statusInfo, model.WebStatus{StatusCode: respStatus}
}

// Cluster 삭제 비동기 처리
func DelWorkflowByAsync(nameSpaceID string, workflowID string, c echo.Context) {
	var originalUrl = "/ns/{namespace}/workflow/{workflow}"

	var paramMapper = make(map[string]string)
	paramMapper["{namespace}"] = nameSpaceID
	paramMapper["{workflow}"] = workflowID
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)
	url := util.CICADA + urlParam

	// 경로안에 parameter가 있어 추가 param없이 호출 함.
	resp, err := util.CommonHttp(url, nil, http.MethodDelete)


	taskKey := nameSpaceID + "||" + "workflow" + "||" + workflowID

	if err != nil {
		fmt.Println(err)
		StoreWebsocketMessage(util.TASK_TYPE_WORKFLOW, taskKey, util.WORKFLOW_LIFECYCLE_DELETE, util.TASK_STATUS_FAIL, c) // session에 작업내용 저장
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		
		failResultInfo := tbcommon.TbSimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		log.Println("DelMcksByAsync ", failResultInfo)
		StoreWebsocketMessage(util.TASK_TYPE_WORKFLOW, taskKey, util.WORKFLOW_LIFECYCLE_DELETE, util.TASK_STATUS_FAIL, c) // session에 작업내용 저장

	} else {
		returnWorkflowInfo := cicada.WorkflowInfo{}
		json.NewDecoder(respBody).Decode(&returnWorkflowInfo)
		fmt.Println(returnWorkflowInfo)
		StoreWebsocketMessage(util.TASK_TYPE_WORKFLOW, taskKey, util.WORKFLOW_LIFECYCLE_DELETE, util.TASK_STATUS_COMPLETE, c) // session에 작업내용 저장
	}
}


// custom task component 목록 조회
func GetTaskComponentList()([]cicada.TaskComponentInfo, model.WebStatus){
	
	return nil, model.WebStatus{}
}