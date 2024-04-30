package controller

import (
	// "encoding/json"
	//"fmt"
	"log"
	"net/http"

	//"github.com/cloud-barista/cb-webtool/src/model/tumblebug/mcis"

	// model "github.com/cloud-barista/cb-webtool/src/model"
	//"github.com/cloud-barista/cb-webtool/src/model"
	//"github.com/cloud-barista/cb-webtool/src/model/dragonfly"

	// spider "github.com/cloud-barista/cb-webtool/src/model/spider"
	// "github.com/cloud-barista/cb-webtool/src/model/tumblebug"
	// tbcommon "github.com/cloud-barista/cb-webtool/src/model/tumblebug/common"
	// tbmcir "github.com/cloud-barista/cb-webtool/src/model/tumblebug/mcir"
	tbmcis "github.com/cloud-barista/cb-webtool/src/model/tumblebug/mcis"

	//webtool "github.com/cloud-barista/cb-webtool/src/model/webtool"

	service "github.com/cloud-barista/cb-webtool/src/service"
	util "github.com/cloud-barista/cb-webtool/src/util"

	echotemplate "github.com/foolin/echo-template"
	"github.com/labstack/echo"
	// echosession "github.com/go-session/echo-session"
)

// type SecurityGroup struct {
// 	Id []string `form:"sg"`
// }

func WorkflowRegForm(c echo.Context) error {

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	// namespacelist 가져오기
	// nsList, _ := service.GetNameSpaceList()
	nsList, _ := service.GetStoredNameSpaceList(c)
	log.Println(" nsList  ", nsList)

	taskComponentList, _ := service.GetTaskComponentList()
	//workflowList, _ := service.GetWorkflowList(defaultNameSpaceID)

	return echotemplate.Render(c, http.StatusOK,
		"operation/migrations/workflowmng/WorkflowCreate", // 파일명
		map[string]interface{}{
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NameSpaceList":      nsList,

			"TaskComponentList": taskComponentList,
		})
}

// Workflow 목록 조회
func GetWorkflowList(c echo.Context) error {
	log.Println("GetMcisList : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	optionParam := c.QueryParam("option")
	filterKeyParam := c.QueryParam("filterKey")
	filterValParam := c.QueryParam("filterVal")

	workflowList, respStatus := service.GetWorkflowList(defaultNameSpaceID, optionParam, filterKeyParam, filterValParam)
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.JSON(respStatus.StatusCode, map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":            "success",
		"status":             respStatus.StatusCode,
		"DefaultNameSpaceID": defaultNameSpaceID,
		"WorkflowList":       workflowList,
	})

}

// Workflow 등록
func WorkflowRegProc(c echo.Context) error {
	log.Println("WorkflowRegProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	workflowReqInfo := &tbmcis.TbMcisReq{}
	if err := c.Bind(workflowReqInfo); err != nil {
		// if err := c.Bind(mCISInfoList); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "5001",
		})
	}
	log.Println(workflowReqInfo)

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	taskKey := defaultNameSpaceID + "||" + "workflow" + "||" + workflowReqInfo.Name

	service.StoreWebsocketMessage(util.TASK_TYPE_WORKFLOW, taskKey, util.WORKFLOW_LIFECYCLE_CREATE, util.TASK_STATUS_REQUEST, c) // session에 작업내용 저장

	// // go routin, channel
	go service.RegWorkflowByAsync(defaultNameSpaceID, workflowReqInfo, c)
	// 원래는 호출 결과를 return하나 go routine으로 바꾸면서 요청성공으로 return
	log.Println("before return")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  200,
	})

}

// // MCIS 등록
// func McisDynamicRegProc(c echo.Context) error {
// 	log.Println("McisDynamicRegProc : ")
// 	loginInfo := service.CallLoginInfo(c)
// 	if loginInfo.UserID == "" {
// 		return c.Redirect(http.StatusTemporaryRedirect, "/login")
// 	}

// 	// map[description:bb installMonAgent:yes name:aa vm:[map[connectionName:gcp-asia-east1 description:dd imageId:gcp-jsyoo-ubuntu name:cc provider:GCP securityGroupIds:[gcp-jsyoo-sg-01] specId:gcp-jsyoo-01 sshKeyId:gcp-jsyoo-sshkey subnetId:jsyoo-gcp-sub-01 vNetId:jsyoo-gcp-01 vm_add_cnt:0 vm_cnt:]]]
// 	log.Println("get info")

// 	mcisReqInfo := &tbmcis.TbMcisDynamicReq{}
// 	if err := c.Bind(mcisReqInfo); err != nil {
// 		log.Println(err)
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": "fail",
// 			"status":  "5001",
// 		})
// 	}
// 	log.Println(mcisReqInfo) // 여러개일 수 있음.

// 	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
// 	// TODO : defaultNameSpaceID 가 없으면 설정화면으로 보낼 것

// 	// // socket의 key 생성 : ns + 구분 + id
// 	taskKey := defaultNameSpaceID + "||" + "mcis" + "||" + mcisReqInfo.Name // TODO : 공통 function으로 뺄 것.

// 	service.StoreWebsocketMessage(util.TASK_TYPE_MCIS, taskKey, util.MCIS_LIFECYCLE_CREATE, util.TASK_STATUS_REQUEST, c) // session에 작업내용 저장

// 	go service.RegMcisDynamicByAsync(defaultNameSpaceID, mcisReqInfo, c)
// 	// 원래는 호출 결과를 return하나 go routine으로 바꾸면서 요청성공으로 return
// 	log.Println("before return")
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success",
// 		"status":  200,
// 	})

// }

// 추천 vm spec 조회
// Recommend MCIS plan (filter and priority)
// func GetMcisRecommendVmSpecList(c echo.Context) error {

// 	log.Println("McisRegProc : ")
// 	loginInfo := service.CallLoginInfo(c)
// 	if loginInfo.UserID == "" {
// 		return c.Redirect(http.StatusTemporaryRedirect, "/login")
// 	}

// 	mcisDeploymentPlan := &tbmcis.DeploymentPlan{}
// 	if err := c.Bind(mcisDeploymentPlan); err != nil {
// 		// if err := c.Bind(mCISInfoList); err != nil {
// 		log.Println(err)
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": "fail",
// 			"status":  "fail",
// 		})
// 	}
// 	log.Println(mcisDeploymentPlan)

// 	vmSpecList, _ := service.GetMcisRecommendVmSpecList(mcisDeploymentPlan)

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message":    "success",
// 		"status":     200,
// 		"VmSpecList": vmSpecList,
// 	})
// }

// Workflow 삭제
func WorkflowDelProc(c echo.Context) error {
	log.Println("WorkflowDelProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	workflowID := c.Param("workflowID")
	//optionParam := c.QueryParam("option")
	log.Println("workflowID= " + workflowID)
	_, respStatus := service.DelWorkflow(defaultNameSpaceID, workflowID)
	log.Println("WorkflowDelProc service returned")
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.JSON(respStatus.StatusCode, map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  respStatus.StatusCode,
	})
}

// GetWorkflowInfoData
func GetWorkflowInfoData(c echo.Context) error {
	log.Println("GetWorkflowInfoData")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login") // 조회기능에서 바로 login화면으로 돌리지말고 return message로 하는게 낫지 않을까?
	}
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	workflowID := c.Param("workflowID")
	log.Println("workflowID= " + workflowID)
	optionParam := c.QueryParam("option")
	log.Println("optionParam= " + optionParam)

	resultWorkflowInfo, _ := service.GetWorkflowData(defaultNameSpaceID, workflowID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success",
		"status":       200,
		"WorkflowInfo": resultWorkflowInfo,
	})
}

func WorkflowDefaultMngForm(c echo.Context) error {
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	return echotemplate.Render(c, http.StatusOK,
		"operation/workflow/SequentialWorkflowDesigner",
		map[string]interface{}{
			"LoginInfo": loginInfo,
		})
}

func WorkflowFullscreenMngForm(c echo.Context) error {
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	return echotemplate.Render(c, http.StatusOK,
		"operation/workflow/SequentialWorkflowDesignerFullScreen",
		map[string]interface{}{
			"LoginInfo": loginInfo,
		})
}

func WorkflowDemoMngForm(c echo.Context) error {
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	targetFile := c.Param("targetFile")

	return echotemplate.Render(c, http.StatusOK,
		"operation/workflow/"+targetFile,
		map[string]interface{}{
			"LoginInfo": loginInfo,
		})
}
