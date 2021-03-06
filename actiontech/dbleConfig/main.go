package main

import (
	"encoding/json"
	"fmt"
)

type ComponentGroupMeta struct {
	// @field
	// @description=component group id
	// @example=uproxy-g1
	GroupId string `json:"group_id" gorm:"primary_key"`
	// @field
	// @description=component group type
	// @example=uproxy
	Type string `json:"type"`
	// @field
	// @description=component group admin user
	// @example=admin
	AdminUser string `json:"admin_user"`
	// @field
	// @description=component group password
	// @example=aeN8MV7d2Ae6x1
	AdminPassword string `json:"admin_password"`
	// @field
	// @description=component group port(ushard, uproxy)
	// @example=13665
	Port string `json:"port"`
	// @field
	// @description=if component should be monitored
	// @default=true
	MonitorEnable bool `json:"monitor_enable"`
	// @field
	// @description=sip of component group
	// @example=172.30.10.163
	Sip string `json:"sip"`
	// @field
	// @description=sip of component group
	// @example=172.30.10.163
	SipNetSegment string `json:"sip_net_segment"`
	// @field
	// @description=if auto commit (uproxy)
	IsAutoCommit bool `json:"is_auto_commit"`
	// @field
	// @description=ignore update route error
	IgnoreUpdateRouteError bool `json:"ignore_update_route_error"`
	// @field
	// @description=ushard params
	// @example={"charset":"utf8","fakeMySQLVersion":"5.6.23","managerPort":"13310","serverPort":"13306"}
	UshardParams string `json:"ushard_params"`
	// @field
	// @description=ushard xml config files
	// @example={"cacheservice.properties":"#\\n# Copyright (C) 2016-2019 ActionTech.\\n# License: http://www.gnu.org/licenses/gpl.html...
	UshardXmlConfigFiles string `json:"ushard_xml_config_files"`
	// @field
	// @description=ushard json config
	// @example={}
	UshardJsonConfig string `json:"ushard_json_config"`
	// @field
	// @description=ushard reload error
	// @example=err
	UshardLastReloadError string `json:"ushard_last_reload_error"`
	// @field
	// @description=ushard used-work-ids
	// @example=0,1,2,3
	UshardUsedWorkIds string `json:"ushard_used_work_ids"`
	// @field
	// @description=ushard slow query alert limited-time
	// @example=60(s)
	UshardSlowQueryTime int `json:"ushard_slow_query_time"`
}

func main() {
	jsonString := `{
    "group_id": "ushard-shard_group_1",
    "type": "ushard",
    "admin_user": "root",
    "admin_password": "C1yDa/NDw1Xork9ATMZlpnPxhYNGlacH7KBconXvAT+/g+5+YI5oEN1+Z9GL9fsq8Tz/X+C08518t+3f+yWRpGmoT2MWaArr+48hAjoXflphf+aJ+dLIwBOpl4DV7fMULvW5M42qxMuLsxGW8LwvDoSu2uqH/4umVAAvez+qtLo=",
    "port": "8066",
    "monitor_enable": true,
    "sip": "",
    "sip_net_segment": "",
    "is_auto_commit": false,
    "ignore_update_route_error": false,
    "ushard_params": "{\"charset\":\"utf8\",\"fakeMySQLVersion\":\"5.6.23\",\"managerPort\":\"9066\",\"serverPort\":\"8066\"}",
    "ushard_xml_config_files": "{\"autoGeneratedMapFile_0.txt\":\"//auto generated by rule default_enum_int\\n1=0\\n2=1\\n3=2\\n4=3\",\"rule.xml\":\"\\u003c?xml version=\\\"1.0\\\"?\\u003e\\n\\u003c!DOCTYPE dble:rule SYSTEM \\\"rule.dtd\\\"\\u003e\\n\\u003cdble:rule xmlns:dble=\\\"http://dble.cloud/\\\" version=\\\"1.0\\\"\\u003e\\n    \\u003cfunction name=\\\"default_mod_two\\\" class=\\\"Hash\\\"\\u003e\\n        \\u003cproperty name=\\\"partitionCount\\\"\\u003e2\\u003c/property\\u003e\\n        \\u003cproperty name=\\\"partitionLength\\\"\\u003e1\\u003c/property\\u003e\\n    \\u003c/function\\u003e\\n    \\u003cfunction name=\\\"default_enum_int\\\" class=\\\"Enum\\\"\\u003e\\n        \\u003cproperty name=\\\"defaultNode\\\"\\u003e0\\u003c/property\\u003e\\n        \\u003cproperty name=\\\"type\\\"\\u003e0\\u003c/property\\u003e\\n        \\u003cproperty name=\\\"mapFile\\\"\\u003eautoGeneratedMapFile_0.txt\\u003c/property\\u003e\\n    \\u003c/function\\u003e\\n    \\u003cfunction name=\\\"default_enum_int2\\\" class=\\\"Enum\\\"\\u003e\\n\\u003cproperty name=\\\"mapFile\\\"\\u003etest.enum\\u003c/property\\u003e\\n\\u003cproperty name=\\\"defaultNode\\\"\\u003e0\\u003c/property\\u003e\\n\\u003cproperty name=\\\"type\\\"\\u003e0\\u003c/property\\u003e\\n \\u003c/function\\u003e\\n\\u003c/dble:rule\\u003e\",\"schema.xml\":\"\\u003c?xml version=\\\"1.0\\\"?\\u003e\\n\\u003c!DOCTYPE dble:schema SYSTEM \\\"schema.dtd\\\"\\u003e\\n\\u003cdble:schema xmlns:dble=\\\"http://dble.cloud/\\\" \\u003e\\n\\u003c!--\\u003cschema name=\\\"testdb\\\" sqlMaxLimit=\\\"100\\\" \\u003e--\\u003e\\n\\u003c!--\\u003ctable name=\\\"sharding_two_node\\\" dataNode=\\\"dn1,dn2\\\" rule=\\\"two_node_hash\\\" primaryKey=\\\"id\\\"/\\u003e--\\u003e\\n\\u003c!--\\u003c/schema\\u003e--\\u003e\\n\\n\\u003c!--\\u003cdataNode name=\\\"dn1\\\" dataHost=\\\"dh1\\\" database=\\\"test1\\\" /\\u003e--\\u003e\\n\\u003c!--\\u003cdataNode name=\\\"dn2\\\" dataHost=\\\"dh1\\\" database=\\\"test2\\\" /\\u003e--\\u003e\\n\\n\\u003c!--\\u003cdataHost name=\\\"dh1\\\" maxCon=\\\"10\\\" minCon=\\\"1\\\" balance=\\\"3\\\"--\\u003e\\n\\u003c!--switchType=\\\"-1\\\" slaveThreshold=\\\"100\\\" \\u003e--\\u003e\\n\\u003c!--\\u003cheartbeat\\u003eshow slave status\\u003c/heartbeat\\u003e--\\u003e\\n\\u003c!--\\u003cwriteHost host=\\\"hostM1\\\" url=\\\"xxx.xxx.xxx.a:3306\\\" user=\\\"user1\\\" password=\\\"***\\\"\\u003e--\\u003e\\n\\u003c!--\\u003creadHost host=\\\"hostS1\\\" url=\\\"xxx.xxx.xx.b:3306\\\" user=\\\"user2\\\" password=\\\"***\\\"/\\u003e--\\u003e\\n\\u003c!--\\u003c/writeHost\\u003e--\\u003e\\n\\u003c!--\\u003c/dataHost\\u003e--\\u003e\\n\\u003c/dble:schema\\u003e\",\"server.xml\":\"\\u003c?xml version=\\\"1.0\\\"?\\u003e\\n\\u003c!DOCTYPE dble:server SYSTEM \\\"server.dtd\\\"\\u003e\\n\\u003cdble:server xmlns:dble=\\\"http://dble.cloud/\\\" version=\\\"1.0\\\"\\u003e\\n    \\u003csystem\\u003e\\n        \\u003cproperty name=\\\"serverPort\\\"\\u003e8066\\u003c/property\\u003e\\n        \\u003cproperty name=\\\"managerPort\\\"\\u003e9066\\u003c/property\\u003e\\n        \\u003cproperty name=\\\"fakeMySQLVersion\\\"\\u003e5.6.23\\u003c/property\\u003e\\n        \\u003cproperty name=\\\"sequnceHandlerType\\\"\\u003e2\\u003c/property\\u003e\\n        \\u003cproperty name=\\\"charset\\\"\\u003eutf8\\u003c/property\\u003e\\n        \\u003cproperty name=\\\"txIsolation\\\"\\u003e2\\u003c/property\\u003e\\n        \\u003cproperty name=\\\"sqlExecuteTimeout\\\"\\u003e3600\\u003c/property\\u003e\\n    \\u003c/system\\u003e\\n    \\u003cuser name=\\\"root\\\"\\u003e\\n        \\u003cproperty name=\\\"password\\\"\\u003e1\\u003c/property\\u003e\\n        \\u003cproperty name=\\\"readOnly\\\"\\u003etrue\\u003c/property\\u003e\\n        \\u003cproperty name=\\\"manager\\\"\\u003etrue\\u003c/property\\u003e\\n        \\u003cproperty name=\\\"usingDecrypt\\\"\\u003e0\\u003c/property\\u003e\\n    \\u003c/user\\u003e\\n\\u003c/dble:server\\u003e\"}",
    "ushard_json_config": "",
    "ushard_last_reload_error": "",
    "ushard_used_work_ids": ""
}`
	compGroupMeta := &ComponentGroupMeta{}
	if err := json.Unmarshal([]byte(jsonString), compGroupMeta); err != nil {
		panic(err)
	}
	vals := map[string]string{}
	if err := json.Unmarshal([]byte(compGroupMeta.UshardXmlConfigFiles), &vals); err != nil {
		panic(err)
	}
	fmt.Println(vals["test.enum"])
}
