// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package notification

import (
	"context"
	"testing"
	"time"

	"openpitrix.io/logger"
	notification "openpitrix.io/notification/pkg/pb"
	"openpitrix.io/notification/pkg/util/pbutil"
)

func TestNewServer(t *testing.T) {
	logger.SetLevelByString("debug")
	s, _ := NewServer()
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	logger.Infof(nil, "[%+v]", s)
}

func TestCreateNfWithAddrs(t *testing.T) {
	s, _ := NewServer()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	testAddrsStr := "huojiao2006@163.com;513590612@qq.com"
	var req = &notification.CreateNfWithAddrsRequest{
		ContentType:  pbutil.ToProtoString("Information"),
		SentType:     pbutil.ToProtoString("Email"),
		AddrsStr:     pbutil.ToProtoString(testAddrsStr),
		Title:        pbutil.ToProtoString("Run case"),
		Content:      pbutil.ToProtoString("Run case content"),
		ShortContent: pbutil.ToProtoString("Run case ShortContent"),
		ExpiredDays:  pbutil.ToProtoString("7"),
		Owner:        pbutil.ToProtoString("HuoJiao"),
		Status:       pbutil.ToProtoString("New"),
	}
	s.CreateNfWithAddrs(ctx, req)
}