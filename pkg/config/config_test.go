package config

import (
	"openpitrix.io/logger"
	"os"
	"testing"
)

func TestLoadConf(t *testing.T) {
	logger.SetLevelByString("debug")


	os.Setenv("NF_APP_NAME", "NF_APP_NAME_test")
	os.Setenv("NF_APP_HOST", "NF_APP_HOST_test")
	os.Setenv("NF_APP_PORT", "NF_APP_PORT_test")
	os.Setenv("NF_APP_ENV", "NF_APP_ENV_test")
	os.Setenv("NF_APP_MAXTASKS", "9")
	os.Setenv("NF_APP_APPLOGMODE", "NF_APP_APPLOGMODE_INFO_test")




	os.Setenv("NF_DB_HOST", "NF_DB_HOST_test")
	os.Setenv("NF_DB_PORT", "NF_DB_PORT_Test")
	os.Setenv("NF_DB_USER", "NF_DB_USER_Test")
	os.Setenv("NF_DB_PASSWORD", "NF_DB_PASSWORD_test")
	os.Setenv("NF_DB_DBNAME", "NF_DB_DBNAME_test")
	os.Setenv("NF_DB_DISABLE", "false")
	os.Setenv("NF_DB_LOGMODE", "false")

	os.Setenv("NF_ETCD_ENDPOINTS", "NF_ETCD_ENDPOINTS_test")
	os.Setenv("NF_ETCD_PREFIX", "NF_ETCD_PREFIX_Test")
	os.Setenv("NF_ETCD_TOPIC", "NF_ETCD_TOPIC_Test")



	os.Setenv("NF_EMAIL_HOST", "NF_EMAIL_HOST_test")
	os.Setenv("NF_EMAIL_PORT", "11")
	os.Setenv("NF_EMAIL_USERNAME", "NF_EMAIL_USERNAME_Test")
	os.Setenv("NF_EMAIL_PASSWORD", "NF_EMAIL_PASSWORD_Test")


	mycfg:=GetInstance()
	mycfg.LoadConf()

	logger.Debugf(nil, "Appcfg=========================================")
	logger.Debugf(nil, "NF_APP_NAME : %+v",mycfg.App.Name)
	logger.Debugf(nil, "NF_APP_HOST : %+v",mycfg.App.Host)
	logger.Debugf(nil, "NF_APP_PORT : %+v", mycfg.App.Port)
	logger.Debugf(nil, "NF_APP_ENV : %+v", mycfg.App.Env)
	logger.Debugf(nil, "NF_APP_MAXTASKS : %+v", mycfg.App.Maxtasks)
	logger.Debugf(nil, "NF_APP_APPLOGMODE : %+v", mycfg.App.Applogmode)

	logger.Debugf(nil, "Dbcfg=========================================")
	logger.Debugf(nil, "NF_DB_HOST : %+v",mycfg.Db.Host)
	logger.Debugf(nil, "NF_DB_PORT : %+v",mycfg.Db.Port)
	logger.Debugf(nil, "NF_DB_USER : %+v",mycfg.Db.User)
	logger.Debugf(nil, "NF_DB_PASSWORD : %+v",mycfg.Db.Password)
	logger.Debugf(nil, "NF_DB_DBNAME : %+v",mycfg.Db.Dbname)
	logger.Debugf(nil, "NF_DB_DISABLE : %+v",mycfg.Db.Disable)
	logger.Debugf(nil, "NF_DB_DBLOGMODE : %+v",mycfg.Db.Logmode)

	logger.Debugf(nil, "Etcdcfg=========================================")
	logger.Debugf(nil, "NF_ETCD_ENDPOINTS : %+v",mycfg.Etcd.Endpoints)
	logger.Debugf(nil, "NF_ETCD_PREFIX : %+v",mycfg.Etcd.Prefix)
	logger.Debugf(nil, "NF_ETCD_TOPIC : %+v",mycfg.Etcd.Topic)


	logger.Debugf(nil, "Emailcfg=========================================")
	logger.Debugf(nil, "NF_EMAIL_HOST : %+v",mycfg.Email.Host)
	logger.Debugf(nil, "NF_EMAIL_PORT : %+v", mycfg.Email.Port)
	logger.Debugf(nil, "NF_EMAIL_USERNAME : %+v",mycfg.Email.Username)
	logger.Debugf(nil, "NF_EMAIL_PASSWORD : %+v",mycfg.Email.Password)


	mycfg.PrintUsage()


}

