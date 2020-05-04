/*
 *
 * 	Copyright (C) 2020  <Semchenko Aleksandr>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.If not, see <http://www.gnu.org/licenses/>.
 * /
 */

package kaspersky

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

//VServers Class Reference
//Virtual servers processing.
//
//Detailed Description
//
//Virtual servers processing.
//
//Allows to create and destroy virtual servers, acquire and modify their attributes.
type VServers struct {
	client *Client
}

//Acquire virtual servers for the specified group.
//
//Returns array of virtual servers for the specified group
//
//Parameters:
//	- lParentGroup	(int) id of parent group, -1 means 'from all groups'
//Returns:
//	- (array) array, each element is a container KLPAR::ParamsPtr containing attributes "KLVSRV_*"
//(see List of virtual server attributes).
func (vs *VServers) GetVServers(ctx context.Context, lParentGroup int) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lParentGroup": %d}`, lParentGroup))

	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.GetVServers", bytes.NewBuffer(postData))
	raw, err := vs.client.Do(ctx, request, nil)
	return raw, err
}

//Register new virtual server.
//
//Registers new virtual server
//
// Parameters:
//	- strDisplayName	(string) virtual server display name, if display name is non-unique,
//	it will be modified to become unique
//	- lParentGroup	(int) virtual server parent group
//
//	Returns:
//	(params) a container KLPAR::ParamsPtr containing attributes "KLVSRV_ID" and "KLVSRV_DN" (
//	see List of virtual server attributes).
//
//Example Result
//{
//  "PxgRetVal" : {
//    "KLVSRV_CREATED" : {
//      "type" : "datetime",
//      "value" : "2020-05-03T00:59:09Z"
//    },
//    "KLVSRV_DN" : "vservx",
//    "KLVSRV_ENABLED" : true,
//    "KLVSRV_GROUPS" : 167,
//    "KLVSRV_GRP" : 0,
//    "KLVSRV_HST_UID" : "VSRVa80a675f-40d1-4f50-aec8-ff79bd8793d4",
//    "KLVSRV_ID" : 1,
//    "KLVSRV_LIC_ENABLED" : true,
//    "KLVSRV_SUPER" : 166,
//    "KLVSRV_UID" : "VSRVa80a675f-40d1-4f50-aec8-ff79bd8793d4",
//    "KLVSRV_UNASSIGNED" : 170
//  }
//}
func (vs *VServers) AddVServerInfo(ctx context.Context, strDisplayName string, lParentGroup int) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lParentGroup": %d, "strDisplayName" : "%s"}`, lParentGroup, strDisplayName))
	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.AddVServerInfo", bytes.NewBuffer(postData))
	raw, err := vs.client.Do(ctx, request, nil)
	return raw, err
}

//Unregister specified Virtual Server.
//
//Unregisters specified Virtual Server
//
// Parameters:
//	- lVServer	(int) Virtual Server id
//	- [out]	strActionGuid	(string) id of asynchronous operation,
//	to get status use AsyncActionStateChecker.CheckActionState
func (vs *VServers) DelVServer(ctx context.Context, lVServer int) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lVServer": %d}`, lVServer))

	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.DelVServer", bytes.NewBuffer(postData))
	raw, err := vs.client.Do(ctx, request, nil)
	return raw, err
}

func (vs *VServers) GetPermissions(ctx context.Context, lVServer int) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lVServer": %d}`, lVServer))

	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.GetPermissions", bytes.NewBuffer(postData))
	raw, err := vs.client.Do(ctx, request, nil)
	return raw, err
}

//HostInfoParams struct
type VServerParams struct {
	LVServer       int64    `json:"lVServer"`
	PFields2Return []string `json:"pFields2Return"`
}

//Acquire info on virtual server.
//
//Returns info about the specified virtual server
//
//	Parameters:
//	- lVServer	(int) virtual server id
//	- pFields2Return	(array) attributes "KLVSRV_*" to acquire (see List of virtual server attributes).
//
//	Returns:
//	- (params) a container containing attributes "KLVSRV_*" (see List of virtual server attributes)
func (vs *VServers) GetVServerInfo(ctx context.Context, lVServer int64) ([]byte, error) {
	v := VServerParams{LVServer: lVServer, PFields2Return: []string{
		"KLVSRV_CUSTOM_INFO",
		"KLVSRV_ID",
		"KLVSRV_UID",
		"KLVSRV_GRP",
		"KLVSRV_DN",
		"KLVSRV_GROUPS",
		"KLVSRV_SUPER",
		"KLVSRV_UNASSIGNED",
		"KLVSRV_ENABLED",
		"KLVSRV_LIC_ENABLED",
		"KLVSRV_HST_UID",
		"KLVSRV_CREATED",
	}}
	postData, _ := json.Marshal(v) //[]byte(fmt.Sprintf(`{"lVServer": %d}`, lVServer))

	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.GetVServerInfo", bytes.NewBuffer(postData))
	raw, err := vs.client.Do(ctx, request, nil)
	return raw, err
}

//TODO ↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓

func (vs *VServers) UpdateVServerInfo(ctx context.Context, lVServer int, v interface{}) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lVServer": %d}`, lVServer))

	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.UpdateVServerInfo", bytes.NewBuffer(postData))
	raw, err := vs.client.Do(ctx, request, nil)
	return raw, err
}

func (vs *VServers) MoveVServer(ctx context.Context, lVServer int) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lVServer": %d}`, lVServer))

	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.MoveVServer", bytes.NewBuffer(postData))
	raw, err := vs.client.Do(ctx, request, nil)
	return raw, err
}

func (vs *VServers) RecallCertAndCloseConnections(ctx context.Context, lVServer int) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lVServer": %d}`, lVServer))

	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.RecallCertAndCloseConnections", bytes.NewBuffer(postData))
	raw, err := vs.client.Do(ctx, request, nil)
	return raw, err
}

func (vs *VServers) SetPermissions(ctx context.Context, lVServer int) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lVServer": %d}`, lVServer))

	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.SetPermissions", bytes.NewBuffer(postData))
	raw, err := vs.client.Do(ctx, request, nil)
	return raw, err
}