package task

import (
	"openpitrix.io/logger"
	"openpitrix.io/notification/pkg/config"
	"openpitrix.io/notification/pkg/models"
	"openpitrix.io/notification/pkg/util/emailutil"

	//	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"openpitrix.io/notification/pkg/util/etcdutil"

	//	"openpitrix.io/openpitrix/pkg/models"
	//	"openpitrix.io/openpitrix/pkg/pi"
	//	"openpitrix.io/openpitrix/pkg/plugins"
	//	"openpitrix.io/openpitrix/pkg/util/ctxutil"
	"time"
)

//Contains all of the logic for the User model.
type taskService struct {
	db             *gorm.DB
	queue          *etcdutil.Queue
	runningTaskIds chan string
}

func NewService(db *gorm.DB, queue *etcdutil.Queue) Service {
	tasksc := &taskService{db: db, queue: queue}
	MaxTasks:=config.GetInstance().App.MaxWorkingTasks
	tasksc.runningTaskIds = make(chan string, MaxTasks)
	return tasksc
}

func (sc *taskService) ExtractTasks() error {
	for {
		taskId, err := sc.queue.Dequeue()
		//taskId := time.Now().Format("2006-01-02 15:04:05")
		//time.Sleep(1 * time.Second)
		if err != nil {
			logger.Criticalf(nil, "Failed to dequeue job from etcd queue: %+v", err)
			time.Sleep(3 * time.Second)
			continue
		}

		logger.Infof(nil,"%+v","Dequeue from etcd queue success  " + taskId)
		sc.runningTaskIds <- taskId
	}
	return nil
}

func (sc *taskService) HandleTask(handlerNum string) error {
	for {
		taskId := <-sc.runningTaskIds
		logger.Debugf(nil,time.Now().Format("2006-01-02 15:04:05")+" handlerNum:"+handlerNum+"  Receive:", taskId)

		taskWNfInfo, err := sc.getTaskwithNfContentbyID(taskId)
		if err != nil {
			logger.Criticalf(nil, "Error-Get Task from DB withNfContent byID : %+v", err)
			return err
		}
		addrsStr := taskWNfInfo.AddrsStr
		titel := taskWNfInfo.Title
		content := taskWNfInfo.Content
		emailutil.SendMail(addrsStr, titel, content)
	}
	return nil
}

func (sc *taskService) getTaskbyID(taskID string) (*models.Task, error) {
	task := &models.Task{}
	err := sc.db.
		Where("task_id = ?", taskID).
		First(task).Error

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, nil
	}
	return task, nil
}

func (sc *taskService) getTaskwithNfContentbyID(taskID string) (*models.TaskWNfInfo, error) {
	task := &models.TaskWNfInfo{}
	sc.db.Raw("SELECT  t3.title,t3.short_content,  t3.content,t1.task_id,t1.addrs_str "+
		"	FROM task t1,job t2,notification_center_post t3 where t1.job_id=t2.job_id and t2.nf_post_id=t3.nf_post_id  and t1.task_id=? ", taskID).Scan(&task)

	return task, nil
}