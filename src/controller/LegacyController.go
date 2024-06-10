package controller

import (
	"fmt"
	honeybee "github.com/cloud-barista/cb-webtool/src/model/honeybee/sourcegroup"
	"github.com/cloud-barista/cb-webtool/src/service"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func GetConnectionInfoDataById(c echo.Context) error {

	connectionId := c.Param("connId")

	connectionInfo, respStatus := service.GetConnectionInfoDataById(connectionId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":        "success",
		"status":         respStatus,
		"connectionInfo": connectionInfo,
	})
}

func GetConnectionInfoListBySourceId(c echo.Context) error {

	sourceId := c.Param("sgId")

	connectionInfo, respStatus := service.GetConnectionInfoListBySourceId(sourceId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":        "success",
		"status":         respStatus,
		"connectionInfo": connectionInfo,
	})
}

func RegConnectionInfo(c echo.Context) error {

	sourceGroupId := c.Param("sgId")

	connectionInfoReq := &honeybee.ConnectionInfoRegReq{}
	if err := c.Bind(connectionInfoReq); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "5001",
		})
	}
	log.Println(connectionInfoReq)

	connectionInfo, respStatus := service.RegConnectionInfo(sourceGroupId, connectionInfoReq)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":        "success",
		"status":         respStatus,
		"connectionInfo": connectionInfo,
	})
}

func GetConnectionInfoDataBysgIdAndconnId(c echo.Context) error {

	sourceGroupId := c.Param("sgId")
	connectionId := c.Param("connId")

	connectionInfo, respStatus := service.GetConnectionInfoDataBysgIdAndconnId(connectionId, sourceGroupId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":        "success",
		"status":         respStatus,
		"connectionInfo": connectionInfo,
	})
}

func UpdateConnectionInfo(c echo.Context) error {

	sourceGroupId := c.Param("sgId")
	connectionId := c.Param("connId")

	connectionInfoReq := &honeybee.ConnectionInfoRegReq{}
	if err := c.Bind(connectionInfoReq); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "5001",
		})
	}

	connectionInfo, respStatus := service.UpdateConnectionInfo(connectionId, sourceGroupId, *connectionInfoReq)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":        "success",
		"status":         respStatus,
		"connectionInfo": connectionInfo,
	})
}

func DeleteConnectionInfo(c echo.Context) error {

	sourceGroupId := c.Param("sgId")
	connectionId := c.Param("connId")

	connectionInfo, respStatus := service.DeleteConnectionInfo(connectionId, sourceGroupId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":        "success",
		"status":         respStatus,
		"connectionInfo": connectionInfo,
	})
}

func CheckReadyHoneybee(c echo.Context) error {
	message, respStatus := service.CheckReadyHoneybee()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success",
		"status":          respStatus,
		"honeybeeMessage": message,
	})
}

func GetSourceGroupList(c echo.Context) error {

	sourceGroupInfo, respStatus := service.GetSourceGroupList()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success",
		"status":          respStatus,
		"sourceGroupInfo": sourceGroupInfo,
	})
}

func RegSourceGroup(c echo.Context) error {

	sourceGroupRegReq := &honeybee.SourceGroupRegReq{}
	if err := c.Bind(sourceGroupRegReq); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "5001",
		})
	}

	connectionInfo, respStatus := service.RegSourceGroup(*sourceGroupRegReq)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":        "success",
		"status":         respStatus,
		"connectionInfo": connectionInfo,
	})
}

func GetSourceGroupDataById(c echo.Context) error {

	sourceGroupId := c.Param("sgId")
	fmt.Print("sgID : %s", sourceGroupId)
	connectionInfo, respStatus := service.GetSourceGroupDataById(sourceGroupId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":        "success",
		"status":         respStatus,
		"connectionInfo": connectionInfo,
	})
}

func UpdateSourceGroupData(c echo.Context) error {
	sourceGroupId := c.Param("sgId")

	sourceGroupRegReq := &honeybee.SourceGroupRegReq{}
	if err := c.Bind(sourceGroupRegReq); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "5001",
		})
	}

	connectionInfo, respStatus := service.UpdateSourceGroupData(sourceGroupId, *sourceGroupRegReq)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":        "success",
		"status":         respStatus,
		"connectionInfo": connectionInfo,
	})
}

func DeleteSourceGroupList(c echo.Context) error {
	sourceGroupId := c.Param("sgId")

	honeybeeMessage, respStatus := service.DeleteSourceGroupList(sourceGroupId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success",
		"status":          respStatus,
		"honeybeeMessage": honeybeeMessage,
	})
}

func GetCheckConnectionSourceGroupData(c echo.Context) error {
	sourceGroupId := c.Param("sgId")

	connectionInfo, respStatus := service.GetCheckConnectionSourceGroupData(sourceGroupId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":        "success",
		"status":         respStatus,
		"connectionInfo": connectionInfo,
	})
}

func GetImportInfraInfoBySourceIdAndConnId(c echo.Context) error {
	connectionId := c.Param("connId")
	sourceGroupId := c.Param("sgId")

	saveInfraInfo, respStatus := service.GetImportInfraInfoBySourceIdAndConnId(connectionId, sourceGroupId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success",
		"status":        respStatus,
		"saveInfraInfo": saveInfraInfo,
	})
}

func GetSoftwareInfoBySourceIdAndConnId(c echo.Context) error {
	connectionId := c.Param("connId")
	sourceGroupId := c.Param("sgId")

	softwareInfo, respStatus := service.GetSoftwareInfoBySourceIdAndConnId(connectionId, sourceGroupId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success",
		"status":       respStatus,
		"softwareInfo": softwareInfo,
	})
}

func GetLegacyInfraInfoBySourceIdAndConnId(c echo.Context) error {
	connectionId := c.Param("connId")
	sourceGroupId := c.Param("sgId")

	infraInfo, respStatus := service.GetLegacyInfraInfoBySourceIdAndConnId(connectionId, sourceGroupId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "success",
		"status":    respStatus,
		"infraInfo": infraInfo,
	})
}

func GetLegacySoftwareInfoBySourceIdAndConnId(c echo.Context) error {
	connectionId := c.Param("connId")
	sourceGroupId := c.Param("sgId")

	softwareInfo, respStatus := service.GetLegacySoftwareInfoBySourceIdAndConnId(connectionId, sourceGroupId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success",
		"status":       respStatus,
		"softwareInfo": softwareInfo,
	})
}
