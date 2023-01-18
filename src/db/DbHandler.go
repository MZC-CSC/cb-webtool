package db

import (
	"fmt"

	"github.com/cloud-barista/cb-webtool/src/db/dbmodel"
	"github.com/cloud-barista/cb-webtool/src/model"
	tbcommon "github.com/cloud-barista/cb-webtool/src/model/tumblebug/common"
	_ "github.com/go-sql-driver/mysql"

	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getMySQLConnectionString() string {
	DBUser := os.Getenv("db_user")
	DBPassword := os.Getenv("db_password")
	DBName := os.Getenv("db_name")
	DBHost := os.Getenv("db_host")
	DBPort := os.Getenv("db_port")
	//DBtype     := "mysql"

	dataBase := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUser,
		DBPassword,
		DBHost,
		DBPort,
		DBName)

	return dataBase
}

func NewMysqlDBConnection(params ...string) *gorm.DB {
	//var err error
	//conString := getMySQLConnectionString()
	//log.Print(conString)
	//
	////DB, err = gorm.Open(config.GetDBType(), conString)
	////
	////if err != nil {
	////	log.Panic(err)
	////}

	return DB
}

func NewPostresqlDBConnnection(params ...string) *gorm.DB {
	var err error

	DBUser := os.Getenv("db_user")
	DBPassword := os.Getenv("db_password")
	DBName := os.Getenv("db_name")
	DBHost := os.Getenv("db_host")
	DBPort := os.Getenv("db_port")

	//dbUrl := "postgres://" + DBUser + ":" + DBPassword + "@" + DBHost + ":" + DBPort + "/" + DBName
	dbUrl := "host=" + DBHost + " port=" + DBPort + " user=" + DBUser + " dbname=" + DBName + " password=" + DBPassword + " sslmode=disable"
	//DB, err = gorm.Open("postgres", dbUrl)
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	//db, err = gorm.Open( "postgres", "host=db port=5432 user=postgres dbname=postgres sslmode=disable password=postgres")
	//DB, err = gorm.Open(postgres.Open(dbUrl)), &gorm.Config{}
	if err != nil {
		log.Panic(err)
	}

	return DB
}

func GetDBInstance(dbType string) *gorm.DB {
	if DB == nil {
		switch dbType {
		case "mysql":
			NewMysqlDBConnection()
			break
		case "postgres":
			NewPostresqlDBConnnection()
			break
		}

	}
	return DB
}

// TABEL Handling
// Find, First 등은 primary key로 조회
// 그외에는 db.where로

func AddUserInfo(reqUserInfo dbmodel.UserInfo) (dbmodel.UserInfo, error) {
	if DB == nil {
		GetDBInstance("postgres")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(reqUserInfo.Passwd), bcrypt.DefaultCost)
	if err != nil {
		log.Println(" GenerateFromPassword ", err)
		return reqUserInfo, err
	}
	sendUserInfo := dbmodel.UserInfo{}
	sendUserInfo.Id = reqUserInfo.Id
	sendUserInfo.UserName = reqUserInfo.UserName
	sendUserInfo.Passwd = string(passwordHash)
	log.Println(" sendUserInfo ", sendUserInfo)
	err = DB.Create(&sendUserInfo).Error
	if err != nil {
		log.Println(" Create userInfo ", err)
		return sendUserInfo, err
	}
	return sendUserInfo, nil
}

// user 정보조회 : memberNo가 serial인데... 그러면 long인가?
func GetUserInfoByPk(memberNo int) (dbmodel.UserInfo, error) {
	if DB == nil {
		GetDBInstance("postgres")
	}

	userInfo := dbmodel.UserInfo{}

	if result := DB.First(&userInfo, memberNo); result.Error != nil {
		fmt.Println(result.Error)
		return userInfo, result.Error
	}

	return userInfo, nil
}

func GetUserInfoById(id string) (dbmodel.UserInfo, error) {
	if DB == nil {
		GetDBInstance("postgres")
	}

	userInfo := dbmodel.UserInfo{}

	if result := DB.Where("id = ? ", id).First(&userInfo); result.Error != nil {
		fmt.Println(result.Error)
		return userInfo, result.Error
	}

	return userInfo, nil
}

// 등록된 사용자 조회 후 hashPassword 비교
func GetUserInfoByIdPass(id string, passwd string) (dbmodel.UserInfo, error) {
	if DB == nil {
		GetDBInstance("postgres")
	}

	userInfo := dbmodel.UserInfo{}

	if result := DB.Where("id = ?", id).First(&userInfo); result.Error != nil {
		fmt.Println(result.Error)
		return userInfo, result.Error
	}
	err := bcrypt.CompareHashAndPassword([]byte(userInfo.Passwd), []byte(passwd))
	if err != nil {
		return userInfo, err
	}
	return userInfo, nil
}

func GetUserInfoList() ([]dbmodel.UserInfo, error) {
	if DB == nil {
		GetDBInstance("postgres")
	}

	userInfoList := []dbmodel.UserInfo{}

	if result := DB.Find(&userInfoList); result.Error != nil {
		fmt.Println(result.Error)
		return userInfoList, result.Error
	}

	return userInfoList, nil
}

func AddNamespaceInfo(userInfo dbmodel.UserInfo, reqNamespaceInfo dbmodel.NamespaceInfo) (dbmodel.NamespaceInfo, error) {
	if DB == nil {
		GetDBInstance("postgres")
	}

	//reqNamespaceInfo.OwnerNo = userInfo.MemberNo

	// namespace 추가
	err := DB.Create(reqNamespaceInfo).Error
	if err != nil {
		return reqNamespaceInfo, err
	}

	// namespaceShare 추가
	reqNamespaceShare := dbmodel.NamespaceShare{}
	reqNamespaceShare.NsID = reqNamespaceInfo.NsId
	//reqNamespaceShare.MemberNo = userInfo.MemberNo
	err = DB.Create(reqNamespaceShare).Error
	if err != nil {
		return reqNamespaceInfo, err
	}

	return reqNamespaceInfo, nil
}

// namespace 정보조회
func GetNamespaceInfo(nsId string) (dbmodel.NamespaceInfo, error) {
	if DB == nil {
		GetDBInstance("postgres")
	}

	nsInfo := dbmodel.NamespaceInfo{}

	if result := DB.First(&nsInfo, nsId); result.Error != nil {
		fmt.Println(result.Error)
		return nsInfo, result.Error
	}

	return nsInfo, nil
}

func GetNamespaceInfoList() ([]dbmodel.NamespaceInfo, error) {
	if DB == nil {
		GetDBInstance("postgres")
	}

	nsInfoList := []dbmodel.NamespaceInfo{}

	if result := DB.Find(&nsInfoList); result.Error != nil {
		fmt.Println(result.Error)
		return nsInfoList, result.Error
	}

	return nsInfoList, nil
}

// 해당 member가 사용하는 namespace 조회
func GetNamespaceShareByMemberNo(memberNo int) (dbmodel.NamespaceShare, error) {
	if DB == nil {
		GetDBInstance("postgresql")
	}

	namespaceShareInfo := dbmodel.NamespaceShare{}

	if result := DB.Where("memberNo = ? ", memberNo).First(&namespaceShareInfo); result.Error != nil {
		fmt.Println(result.Error)
		return namespaceShareInfo, result.Error
	}

	return namespaceShareInfo, nil
}

func ListNamespaceByUserId(userId string) ([]dbmodel.NamespaceInfo, model.WebStatus) {
	if DB == nil {
		GetDBInstance("postgresql")
	}

	//namespaceShareInfo := dbmodel.NamespaceShare{}
	//
	//if result := DB.Where("id = ? ", userId).First(&namespaceShareInfo); result.Error != nil {
	//	fmt.Println(result.Error)
	//	return namespaceShareInfo, result.Error
	//}
	//
	//return namespaceShareInfo, nil

	//dbNsInfo := dbmodel.NamespaceInfo{}
	namespaceList := []dbmodel.NamespaceInfo{}
	//err := DB.Model(&dbNsInfo).Select("namespace_infos.ns_id, namespace_infos.ns_name").Joins("INNER JOIN namespace_shares nss ON namespace_infos.ns_id = nss.ns_id").Joins("INNER JOIN USER_INFOS ui ON nss.member_no = ui.member_no").Scan(namespaceList).Error
	//err := DB.Find(&namespaceList).Joins("INNER JOIN namespace_shares nss ON namespace_infos.ns_id = nss.ns_id").Joins("INNER JOIN USER_INFOS ui ON nss.member_no = ui.member_no").Where("ui.id = ?", "abc").Error
	err := DB.Joins("INNER JOIN namespace_shares nss ON namespace_infos.ns_id = nss.ns_id").Joins("INNER JOIN USER_INFOS ui ON nss.member_no = ui.member_no").Where("ui.id = ?", userId).Find(&namespaceList).Error
	fmt.Println("query result")
	fmt.Println(namespaceList)
	if err != nil {
		return namespaceList, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	return namespaceList, model.WebStatus{StatusCode: 200, Message: ""}
}

// Namespace를 사용하는 User 조회
func GetNamespaceShareByNsId(nsId string) (dbmodel.NamespaceShare, error) {
	if DB == nil {
		GetDBInstance("postgresql")
	}

	namespaceShareInfo := dbmodel.NamespaceShare{}

	if result := DB.Where("id = ? ", nsId).First(&namespaceShareInfo); result.Error != nil {
		fmt.Println(result.Error)
		return namespaceShareInfo, result.Error
	}

	return namespaceShareInfo, nil
}

// User
func GetUserFromDB(userId string, passwd string) (model.User, error) {
	log.Println(" getUserFromDB ")

	dbUserInfo, err := GetUserInfoByIdPass(userId, passwd)
	if err != nil {
		return model.User{}, err
	}

	userInfo := model.User{}
	userInfo.UserID = dbUserInfo.Id
	//db := db.GetDBInstance()
	//log.Println(" db ", db)

	//var userInfo = db.UserDB{}
	//
	//err := db.Find(&userInfo, "email = ", userId)
	//if err != nil {
	//	return model.User{}, err
	//}
	//err1 := bcrypt.CompareHashAndPassword([]byte(userInfo.PasswordHash), []byte(password))
	//if err1 != nil {
	//	return model.User{}, err1
	//}

	return userInfo, nil
}

func GetUserNamespaceList(userId string) ([]tbcommon.TbNsInfo, model.WebStatus) {
	nsList := []tbcommon.TbNsInfo{}
	dbNsList, nsStatus := ListNamespaceByUserId(userId)
	if nsStatus.StatusCode == 200 {
		for _, dbNs := range dbNsList {
			tbns := tbcommon.TbNsInfo{}
			tbns.Name = dbNs.NsName
			tbns.ID = dbNs.NsId
			nsList = append(nsList, tbns)
		}
	}
	return nsList, nsStatus
}
