package logic

import (
	"context"
	"database-service/internal/svc"
	"git.ickey.com.cn/infra/utils/log"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	das20200116 "github.com/alibabacloud-go/das-20200116/v2/client"
	dbs20190306 "github.com/alibabacloud-go/dbs-20190306/client"
	dms_enterprise20181101 "github.com/alibabacloud-go/dms-enterprise-20181101/client"
	dts20200101 "github.com/alibabacloud-go/dts-20200101/client"
	r_kvstore20150101 "github.com/alibabacloud-go/r-kvstore-20150101/v2/client"
	rds20140815 "github.com/alibabacloud-go/rds-20140815/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

type ALiClientLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func (l *ALiClientLogic) GetAliRDSToken() (_result *rds20140815.Client, _err error) {
	accessKeyId := l.svcCtx.ApolloConfig["aliAccessKeyId"]
	accessKeySecret := l.svcCtx.ApolloConfig["aliAccessKeySecret"]
	config := &openapi.Config{
		AccessKeyId:     &accessKeyId,
		AccessKeySecret: &accessKeySecret,
	}
	config.Endpoint = tea.String("rds.aliyuncs.com")
	_result = &rds20140815.Client{}
	_result, _err = rds20140815.NewClient(config)
	if _err != nil {
		_ = log.Error(_err)
		return nil, _err
	}
	return _result, _err
}

func (l *ALiClientLogic) GetAliRedisToken() (_result *r_kvstore20150101.Client, _err error) {
	accessKeyId := l.svcCtx.ApolloConfig["aliAccessKeyId"]
	accessKeySecret := l.svcCtx.ApolloConfig["aliAccessKeySecret"]
	config := &openapi.Config{
		AccessKeyId:     &accessKeyId,
		AccessKeySecret: &accessKeySecret,
	}
	config.Endpoint = tea.String("r-kvstore.aliyuncs.com")
	_result = &r_kvstore20150101.Client{}
	_result, _err = r_kvstore20150101.NewClient(config)
	if _err != nil {
		_ = log.Error(_err)
		return nil, _err
	}
	return _result, _err
}

func (l *ALiClientLogic) GetAliDMSToken() (_result *dms_enterprise20181101.Client, _err error) {
	accessKeyId := l.svcCtx.ApolloConfig["aliAccessKeyId"]
	accessKeySecret := l.svcCtx.ApolloConfig["aliAccessKeySecret"]
	config := &openapi.Config{
		AccessKeyId:     &accessKeyId,
		AccessKeySecret: &accessKeySecret,
	}
	config.Endpoint = tea.String("dms-enterprise.aliyuncs.com")
	_result = &dms_enterprise20181101.Client{}
	_result, _err = dms_enterprise20181101.NewClient(config)
	if _err != nil {
		_ = log.Error(_err)
		return nil, _err
	}
	return _result, _err
}

func (l *ALiClientLogic) GetAliDBSToken() (_result *dbs20190306.Client, _err error) {
	accessKeyId := l.svcCtx.ApolloConfig["aliAccessKeyId"]
	accessKeySecret := l.svcCtx.ApolloConfig["aliAccessKeySecret"]
	config := &openapi.Config{
		AccessKeyId:     &accessKeyId,
		AccessKeySecret: &accessKeySecret,
	}
	config.Endpoint = tea.String("dbs-api.cn-hangzhou.aliyuncs.com")
	_result = &dbs20190306.Client{}
	_result, _err = dbs20190306.NewClient(config)
	if _err != nil {
		_ = log.Error(_err)
		return nil, _err
	}
	return _result, _err
}

func (l *ALiClientLogic) GetAliDTSToken() (_result *dts20200101.Client, _err error) {
	accessKeyId := l.svcCtx.ApolloConfig["aliAccessKeyId"]
	accessKeySecret := l.svcCtx.ApolloConfig["aliAccessKeySecret"]
	config := &openapi.Config{
		AccessKeyId:     &accessKeyId,
		AccessKeySecret: &accessKeySecret,
	}
	config.Endpoint = tea.String("dts.cn-shanghai.aliyuncs.com")
	_result = &dts20200101.Client{}
	_result, _err = dts20200101.NewClient(config)
	if _err != nil {
		_ = log.Error(_err)
		return nil, _err
	}
	return _result, _err
}

func (l *ALiClientLogic) GetAliDASToken() (_result *das20200116.Client, _err error) {
	accessKeyId := l.svcCtx.ApolloConfig["aliAccessKeyId"]
	accessKeySecret := l.svcCtx.ApolloConfig["aliAccessKeySecret"]
	config := &openapi.Config{
		AccessKeyId:     &accessKeyId,
		AccessKeySecret: &accessKeySecret,
	}
	config.Endpoint = tea.String("das.cn-shanghai.aliyuncs.com")
	_result = &das20200116.Client{}
	_result, _err = das20200116.NewClient(config)
	if _err != nil {
		_ = log.Error(_err)
		return nil, _err
	}
	return _result, _err
}
