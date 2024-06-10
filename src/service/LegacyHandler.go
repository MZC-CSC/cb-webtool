package service

import (
	"encoding/json"
	"fmt"
	"github.com/cloud-barista/cb-webtool/src/model"
	"github.com/cloud-barista/cb-webtool/src/model/honeybee/common"
	honeybeecommon "github.com/cloud-barista/cb-webtool/src/model/honeybee/common"
	honeybee "github.com/cloud-barista/cb-webtool/src/model/honeybee/sourcegroup"
	honeybeeinfra "github.com/cloud-barista/cb-webtool/src/model/honeybee/sourcegroup/infra"
	honeybeesoftware "github.com/cloud-barista/cb-webtool/src/model/honeybee/sourcegroup/software"
	"github.com/cloud-barista/cb-webtool/src/util"
	"log"
	"net/http"
)

/**
HoneyBee Api Manage Handler - 2024.06.05
*/

func GetConnectionInfoDataById(connectionId string) ([]honeybee.ConnectionInfo, model.WebStatus) {
	var originalUrl = "/connection_info/{connId}"

	var paramMapper = make(map[string]string)
	paramMapper["{connId}"] = connectionId
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.HONEYBEE + urlParam

	resp, err := util.CommonHttp(url, nil, http.MethodGet)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	connectionInfo := []honeybee.ConnectionInfo{}

	returnStatus := model.WebStatus{}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		failResultInfo := honeybeecommon.SimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		return nil, model.WebStatus{StatusCode: respStatus, Message: failResultInfo.Message}
	}
	decordErr := json.NewDecoder(respBody).Decode(&connectionInfo)
	if decordErr != nil {
		fmt.Println(decordErr)
	}
	fmt.Println(connectionInfo)

	returnStatus.StatusCode = respStatus
	log.Println(respBody)
	util.DisplayResponse(resp) // 수신내용 확인

	return connectionInfo, returnStatus
}

func GetConnectionInfoListBySourceId(sourceId string) (*[]honeybee.ConnectionInfo, model.WebStatus) {
	var originalUrl = "/source_group/{sgId}/connection_info"

	var paramMapper = make(map[string]string)
	paramMapper["{sgId}"] = sourceId
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.HONEYBEE + urlParam

	resp, err := util.CommonHttp(url, nil, http.MethodGet)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	connectionInfoList := &[]honeybee.ConnectionInfo{}

	returnStatus := model.WebStatus{}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		failResultInfo := honeybeecommon.SimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		return nil, model.WebStatus{StatusCode: respStatus, Message: failResultInfo.Message}
	}
	json.NewDecoder(respBody).Decode(&connectionInfoList)
	fmt.Println(connectionInfoList)

	returnStatus.StatusCode = respStatus
	log.Println(respBody)
	util.DisplayResponse(resp) // 수신내용 확인

	return connectionInfoList, returnStatus
}

func RegConnectionInfo(sourceGroupId string, connectionInfo *honeybee.ConnectionInfoRegReq) (*honeybee.ConnectionInfo, model.WebStatus) {
	var originalUrl = "/source_group/{sgId}/connection_info"

	var paramMapper = make(map[string]string)
	paramMapper["{sgId}"] = sourceGroupId
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.HONEYBEE + urlParam

	pbytes, _ := json.Marshal(connectionInfo)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	returnConnectionInfo := honeybee.ConnectionInfo{}
	returnStatus := model.WebStatus{}

	respBody := resp.Body
	respStatus := resp.StatusCode

	if err != nil {
		fmt.Println(err)
		return &returnConnectionInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		failResultInfo := honeybeecommon.SimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		return &returnConnectionInfo, model.WebStatus{StatusCode: respStatus, Message: failResultInfo.Message}
	}

	json.NewDecoder(respBody).Decode(&returnConnectionInfo)
	fmt.Println(returnConnectionInfo)

	returnStatus.StatusCode = respStatus

	return &returnConnectionInfo, returnStatus
}

func GetConnectionInfoDataBysgIdAndconnId(connectionId string, sourceGroupId string) (*honeybee.ConnectionInfo, model.WebStatus) {
	var originalUrl = "/source_group/{sgId}/connection_info/{connId}"

	var paramMapper = make(map[string]string)
	paramMapper["{connId}"] = connectionId
	paramMapper["{sgId}"] = sourceGroupId
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	fmt.Println("GetConnectionInfoDataBysgIdAndconnId : URL")
	fmt.Println(urlParam)

	url := util.HONEYBEE + urlParam

	resp, err := util.CommonHttp(url, nil, http.MethodGet)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	connectionInfo := &honeybee.ConnectionInfo{}

	returnStatus := model.WebStatus{}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		failResultInfo := honeybeecommon.SimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		return nil, model.WebStatus{StatusCode: respStatus, Message: failResultInfo.Message}
	}
	json.NewDecoder(respBody).Decode(&connectionInfo)
	fmt.Println(connectionInfo)

	returnStatus.StatusCode = respStatus
	log.Println(respBody)
	util.DisplayResponse(resp) // 수신내용 확인

	return connectionInfo, returnStatus
}

// put /honeybee/source_group/{sgId}/connection_info/{connId}
func UpdateConnectionInfo(connectionId string, sourceGroupId string, connectionInfo honeybee.ConnectionInfoRegReq) (*honeybee.ConnectionInfo, model.WebStatus) {
	var originalUrl = "/source_group/{sgId}/connection_info/{connId}"

	var paramMapper = make(map[string]string)
	paramMapper["{connId}"] = connectionId
	paramMapper["{sgId}"] = sourceGroupId
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	fmt.Println("UpdateConnectionInfo : URL")
	fmt.Println(urlParam)

	url := util.HONEYBEE + urlParam
	pbytes, _ := json.Marshal(connectionInfo)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPut)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	returnConnectionInfo := &honeybee.ConnectionInfo{}

	returnStatus := model.WebStatus{}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		failResultInfo := honeybeecommon.SimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		return nil, model.WebStatus{StatusCode: respStatus, Message: failResultInfo.Message}
	}
	json.NewDecoder(respBody).Decode(&returnConnectionInfo)
	fmt.Println(returnConnectionInfo)

	returnStatus.StatusCode = respStatus
	log.Println(respBody)
	util.DisplayResponse(resp) // 수신내용 확인

	return returnConnectionInfo, returnStatus
}

// put /honeybee/source_group/{sgId}/connection_info/{connId}
func DeleteConnectionInfo(connectionId string, sourceGroupId string) (*common.SimpleMsg, model.WebStatus) {
	var originalUrl = "/source_group/{sgId}/connection_info/{connId}"

	var paramMapper = make(map[string]string)
	paramMapper["{connId}"] = connectionId
	paramMapper["{sgId}"] = sourceGroupId
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.HONEYBEE + urlParam

	resp, err := util.CommonHttp(url, nil, http.MethodDelete)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	message := &common.SimpleMsg{}

	returnStatus := model.WebStatus{}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		failResultInfo := honeybeecommon.SimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		return nil, model.WebStatus{StatusCode: respStatus, Message: failResultInfo.Message}
	}
	json.NewDecoder(respBody).Decode(&message)
	fmt.Println(message)

	returnStatus.StatusCode = respStatus
	log.Println(respBody)
	util.DisplayResponse(resp) // 수신내용 확인

	return message, returnStatus
}

// /honeybee/readyz
func CheckReadyHoneybee() (*common.SimpleMsg, model.WebStatus) {
	var originalUrl = "/readyz"

	url := util.HONEYBEE + originalUrl

	resp, err := util.CommonHttp(url, nil, http.MethodGet)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	returnStatus := model.WebStatus{}
	returnModel := common.SimpleMsg{}
	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		failResultInfo := honeybeecommon.SimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		return nil, model.WebStatus{StatusCode: respStatus, Message: failResultInfo.Message}
	}
	json.NewDecoder(respBody).Decode(&returnModel)
	fmt.Println(returnModel)

	returnStatus.StatusCode = respStatus
	log.Println(respBody)
	util.DisplayResponse(resp) // 수신내용 확인

	return &returnModel, returnStatus
}

// //////////////////////////////////////SourceGroup
// /honeybee/source_group get
func GetSourceGroupList() ([]honeybee.SourceGroupInfo, model.WebStatus) {
	var originalUrl = "/source_group"

	url := util.HONEYBEE + originalUrl

	resp, err := util.CommonHttp(url, nil, http.MethodGet)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	sourceGroupInfoList := []honeybee.SourceGroupInfo{}

	returnStatus := model.WebStatus{}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		failResultInfo := honeybeecommon.SimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		return nil, model.WebStatus{StatusCode: respStatus, Message: failResultInfo.Message}
	}
	fmt.Println(respBody)
	decordErr := json.NewDecoder(respBody).Decode(&sourceGroupInfoList)
	if decordErr != nil {
		fmt.Println("Decode Error : ", decordErr)
	}

	fmt.Println("sourceGroupInfoList")
	fmt.Println(sourceGroupInfoList)
	fmt.Println("sourceGroupInfoList end")

	returnStatus.StatusCode = respStatus
	log.Println(respBody)
	util.DisplayResponse(resp) // 수신내용 확인

	return sourceGroupInfoList, returnStatus
}

// /honeybee/source_group post
func RegSourceGroup(regSourceGroupInfo honeybee.SourceGroupRegReq) (*honeybee.SourceGroupInfo, model.WebStatus) {
	var originalUrl = "/source_group"

	url := util.HONEYBEE + originalUrl
	pbytes, _ := json.Marshal(regSourceGroupInfo)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	sourceGroupInfo := &honeybee.SourceGroupInfo{}

	returnStatus := model.WebStatus{}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		failResultInfo := honeybeecommon.SimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		return nil, model.WebStatus{StatusCode: respStatus, Message: failResultInfo.Message}
	}
	json.NewDecoder(respBody).Decode(&sourceGroupInfo)
	fmt.Println(sourceGroupInfo)

	returnStatus.StatusCode = respStatus
	log.Println(respBody)
	util.DisplayResponse(resp) // 수신내용 확인

	return sourceGroupInfo, returnStatus
}

// /honeybee/source_group/{sgId} get
func GetSourceGroupDataById(sourceGroupId string) (*honeybee.SourceGroupInfo, model.WebStatus) {
	var originalUrl = "/source_group/{sgId}"

	var paramMapper = make(map[string]string)
	paramMapper["{sgId}"] = sourceGroupId
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.HONEYBEE + urlParam

	resp, err := util.CommonHttp(url, nil, http.MethodGet)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	sourceGroupInfo := &honeybee.SourceGroupInfo{}

	returnStatus := model.WebStatus{}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		failResultInfo := honeybeecommon.SimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		return nil, model.WebStatus{StatusCode: respStatus, Message: failResultInfo.Message}
	}
	json.NewDecoder(respBody).Decode(&sourceGroupInfo)
	fmt.Println(sourceGroupInfo)

	returnStatus.StatusCode = respStatus
	log.Println(respBody)
	util.DisplayResponse(resp) // 수신내용 확인

	return sourceGroupInfo, returnStatus
}

// /honeybee/source_group/{sgId} update
// Update Naming Rule 물어볼것
func UpdateSourceGroupData(sourceGroupId string, updateSourceGroupInfo honeybee.SourceGroupRegReq) (*honeybee.SourceGroupInfo, model.WebStatus) {
	var originalUrl = "/source_group/{sgId}"

	var paramMapper = make(map[string]string)
	paramMapper["{sgId}"] = sourceGroupId
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.HONEYBEE + urlParam

	pbytes, _ := json.Marshal(updateSourceGroupInfo)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPut)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	sourceGroupInfo := &honeybee.SourceGroupInfo{}

	returnStatus := model.WebStatus{}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		failResultInfo := honeybeecommon.SimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		return nil, model.WebStatus{StatusCode: respStatus, Message: failResultInfo.Message}
	}
	json.NewDecoder(respBody).Decode(&sourceGroupInfo)
	fmt.Println(sourceGroupInfo)

	returnStatus.StatusCode = respStatus
	log.Println(respBody)
	util.DisplayResponse(resp) // 수신내용 확인

	return sourceGroupInfo, returnStatus
}

// /honeybee/source_group/{sgId} delete
func DeleteSourceGroupList(sourceGroupId string) (*common.SimpleMsg, model.WebStatus) {
	var originalUrl = "/source_group/{sgId}"

	var paramMapper = make(map[string]string)
	paramMapper["{sgId}"] = sourceGroupId
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.HONEYBEE + urlParam

	resp, err := util.CommonHttp(url, nil, http.MethodDelete)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	message := &common.SimpleMsg{}

	returnStatus := model.WebStatus{}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		failResultInfo := honeybeecommon.SimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		return nil, model.WebStatus{StatusCode: respStatus, Message: failResultInfo.Message}
	}
	json.NewDecoder(respBody).Decode(&message)
	fmt.Println(message)

	returnStatus.StatusCode = respStatus
	log.Println(respBody)
	util.DisplayResponse(resp) // 수신내용 확인

	return message, returnStatus
}

// /honeybee/source_group/{sgId}/connection_check get
func GetCheckConnectionSourceGroupData(sourceGroupId string) ([]honeybee.ConnectionInfo, model.WebStatus) {
	var originalUrl = "/source_group/{sgId}/connection_check"

	var paramMapper = make(map[string]string)
	paramMapper["{sgId}"] = sourceGroupId
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.HONEYBEE + urlParam
	resp, err := util.CommonHttp(url, nil, http.MethodGet)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	sourceGroupConnectionInfo := []honeybee.ConnectionInfo{}

	returnStatus := model.WebStatus{}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		failResultInfo := honeybeecommon.SimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		return nil, model.WebStatus{StatusCode: respStatus, Message: failResultInfo.Message}
	}
	json.NewDecoder(respBody).Decode(&sourceGroupConnectionInfo)
	fmt.Println(sourceGroupConnectionInfo)

	returnStatus.StatusCode = respStatus
	log.Println(respBody)
	util.DisplayResponse(resp) // 수신내용 확인

	return sourceGroupConnectionInfo, returnStatus
}

//[Import] Import source info

// /honeybee/source_group/{sgId}/connection_info/{connId}/import/infra Get
func GetImportInfraInfoBySourceIdAndConnId(connectionId string, sourceGroupId string) (*honeybee.SaveInfraInfo, model.WebStatus) {
	var originalUrl = "/source_group/{sgId}/connection_info/{connId}/import/infra"

	var paramMapper = make(map[string]string)
	paramMapper["{connId}"] = connectionId
	paramMapper["{sgId}"] = sourceGroupId
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	fmt.Println("GetImportInfraInfoBySourceIdAndConnId : URL")
	fmt.Println(urlParam)

	url := util.HONEYBEE + urlParam

	resp, err := util.CommonHttp(url, nil, http.MethodGet)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	saveInfraInfo := &honeybee.SaveInfraInfo{}

	returnStatus := model.WebStatus{}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		failResultInfo := honeybeecommon.SimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		return nil, model.WebStatus{StatusCode: respStatus, Message: failResultInfo.Message}
	}
	decodeErr := json.NewDecoder(respBody).Decode(&saveInfraInfo)
	if decodeErr != nil {
		fmt.Println(decodeErr)
		fmt.Println(respBody)
	}
	fmt.Println(saveInfraInfo)

	returnStatus.StatusCode = respStatus
	log.Println(respBody)
	util.DisplayResponse(resp) // 수신내용 확인

	return saveInfraInfo, returnStatus
}

// /honeybee/source_group/{sgId}/connection_info/{connId}/import/software Get
func GetSoftwareInfoBySourceIdAndConnId(connectionId string, sourceGroupId string) (*honeybee.SaveSoftwareInfo, model.WebStatus) {
	var originalUrl = "/source_group/{sgId}/connection_info/{connId}/import/infra"

	var paramMapper = make(map[string]string)
	paramMapper["{connId}"] = connectionId
	paramMapper["{sgId}"] = sourceGroupId
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	fmt.Println("GetSoftwareInfoBySourceIdAndConnId : URL")
	fmt.Println(urlParam)

	url := util.HONEYBEE + urlParam

	resp, err := util.CommonHttp(url, nil, http.MethodGet)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	softwareInfraInfo := &honeybee.SaveSoftwareInfo{}

	returnStatus := model.WebStatus{}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		failResultInfo := honeybeecommon.SimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		return nil, model.WebStatus{StatusCode: respStatus, Message: failResultInfo.Message}
	}
	json.NewDecoder(respBody).Decode(&softwareInfraInfo)
	fmt.Println(softwareInfraInfo)

	returnStatus.StatusCode = respStatus
	log.Println(respBody)
	util.DisplayResponse(resp) // 수신내용 확인

	return softwareInfraInfo, returnStatus
}

// [Get] Get source info
// /honeybee/source_group/{sgId}/connection_info/{connId}/infra
func GetLegacyInfraInfoBySourceIdAndConnId(connectionId string, sourceGroupId string) (honeybeeinfra.InfraInfo, model.WebStatus) {
	var originalUrl = "/source_group/{sgId}/connection_info/{connId}/import/infra"

	var paramMapper = make(map[string]string)
	paramMapper["{connId}"] = connectionId
	paramMapper["{sgId}"] = sourceGroupId
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	fmt.Println("GetInfraInfoBySourceIdAndConnId : URL")
	fmt.Println(urlParam)

	url := util.HONEYBEE + urlParam

	resp, err := util.CommonHttp(url, nil, http.MethodGet)

	if err != nil {
		fmt.Println(err)
		return honeybeeinfra.InfraInfo{}, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	infraInfo := honeybeeinfra.InfraInfo{}

	returnStatus := model.WebStatus{}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		failResultInfo := honeybeecommon.SimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		return honeybeeinfra.InfraInfo{}, model.WebStatus{StatusCode: respStatus, Message: failResultInfo.Message}
	}
	fmt.Println("respBody")
	fmt.Println(respBody)
	fmt.Println(&respBody)
	fmt.Println("respBody end")

	decodeErr := json.NewDecoder(respBody).Decode(&infraInfo)
	if decodeErr != nil {
		fmt.Println(respBody)
		fmt.Println(decodeErr)
	}
	fmt.Println(infraInfo)

	returnStatus.StatusCode = respStatus
	log.Println(respBody)
	util.DisplayResponse(resp) // 수신내용 확인

	return infraInfo, returnStatus
}

// /honeybee/source_group/{sgId}/connection_info/{connId}/software
func GetLegacySoftwareInfoBySourceIdAndConnId(connectionId string, sourceGroupId string) (*honeybeesoftware.SoftwareInfo, model.WebStatus) {
	var originalUrl = "/source_group/{sgId}/connection_info/{connId}/software"

	var paramMapper = make(map[string]string)
	paramMapper["{connId}"] = connectionId
	paramMapper["{sgId}"] = sourceGroupId
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	fmt.Println("GetLegacySoftwareInfoBySourceIdAndConnId : URL")
	fmt.Println(urlParam)

	url := util.HONEYBEE + urlParam

	resp, err := util.CommonHttp(url, nil, http.MethodGet)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	softwareInfo := &honeybeesoftware.SoftwareInfo{}
	fmt.Println("Body")
	fmt.Println(respBody)
	fmt.Println("Body")
	returnStatus := model.WebStatus{}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		failResultInfo := honeybeecommon.SimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		return nil, model.WebStatus{StatusCode: respStatus, Message: failResultInfo.Message}
	}
	decodeErr := json.NewDecoder(respBody).Decode(&softwareInfo)
	if decodeErr != nil {
		fmt.Println(respBody)
		fmt.Println(decodeErr)
	}
	fmt.Println("softwareInfo")
	fmt.Println(softwareInfo)
	fmt.Println("softwareInfo end")

	returnStatus.StatusCode = respStatus
	log.Println(respBody)
	util.DisplayResponse(resp) // 수신내용 확인

	return softwareInfo, returnStatus
}
