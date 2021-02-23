/*
 * MIT License
 *
 * Copyright (c) [2020] [Semchenko Aleksandr]
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package kaspersky

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// HostGroup service allow to Hosts and management groups processing.
type HostGroup service

// AddDomain Add a new domain to the database.
func (hg *HostGroup) AddDomain(ctx context.Context, strDomain string, nType int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strDomain": "%s", "nType" : %d }`, strDomain, nType))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.AddDomain", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

// AddGroupParams struct
type AddGroupParams struct {
	PInfo *GroupPInfo `json:"pInfo,omitempty"`
}

type GroupPInfo struct {
	Name     *string `json:"name,omitempty"`
	ParentID *int64  `json:"parentId,omitempty"`
}

// AddGroup Creates new group with the specified attributes and returns its Id.
// If such group already exists returns Id of existing group.
func (hg *HostGroup) AddGroup(ctx context.Context, params AddGroupParams) (*PxgValInt, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.AddGroup", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := hg.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

// AddGroupHostsForSync Add hosts from specified group to synchronization.
func (hg *HostGroup) AddGroupHostsForSync(ctx context.Context, nGroupId int64, strSSType string) (*WActionGUID, []byte,
	error) {
	postData := []byte(fmt.Sprintf(` {"nGroupId": %d , "strSSType": "%s" }`, nGroupId, strSSType))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.AddGroupHostsForSync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	wActionGUID := new(WActionGUID)
	raw, err := hg.client.Do(ctx, request, &wActionGUID)
	return wActionGUID, raw, err
}

// AddHost Create new host record.
func (hg *HostGroup) AddHost(ctx context.Context, params interface{}) (*PxgValStr, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.AddHost", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := hg.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// HostsForSyncParams struct using in HostGroup.AddHostsForSync
type HostsForSyncParams struct {
	// PHostNames array of host names
	PHostNames []string `json:"pHostNames"`

	// StrSSType setting storage identity (empty string means synchronization of all setting storages)
	StrSSType string `json:"strSSType,omitempty"`
}

// AddHostsForSync Performs synchronization of settings between server and host.
func (hg *HostGroup) AddHostsForSync(ctx context.Context, params HostsForSyncParams) (*WActionGUID, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.AddHostsForSync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	wActionGUID := new(WActionGUID)
	raw, err := hg.client.Do(ctx, request, &wActionGUID)
	return wActionGUID, raw, err
}

// AddIncidentsParams struct
type AddIncidentsParams struct {
	PData PData `json:"pData"`
}

// PData struct
type PData struct {
	KlincdtSeverity  int64     `json:"KLINCDT_SEVERITY"`
	KlincdtAdded     *DateTime `json:"KLINCDT_ADDED"`
	KlincdtBody      string    `json:"KLINCDT_BODY"`
	KlhstWksHostname string    `json:"KLHST_WKS_HOSTNAME"`
	KlhstuserID      int64     `json:"KLHSTUSER_ID"`
}

// AddIncident Create new incident.
func (hg *HostGroup) AddIncident(ctx context.Context, params AddIncidentsParams) (*PxgValStr, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.AddIncident", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := hg.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// DelDomain Removes a domain from the database.
func (hg *HostGroup) DelDomain(ctx context.Context, strDomain string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strDomain": "%s"}`, strDomain))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.DelDomain", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

// DeleteIncident Delete incident.
func (hg *HostGroup) DeleteIncident(ctx context.Context, nId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nId": %d}`, nId))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.DeleteIncident", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

// HGParams struct
type HGParams struct {
	WstrFilter        string          `json:"wstrFilter"`
	VecFieldsToReturn []string        `json:"vecFieldsToReturn"`
	VecFieldsToOrder  []FieldsToOrder `json:"vecFieldsToOrder"`
	PParams           PParams         `json:"pParams"`
	LMaxLifeTime      int64           `json:"lMaxLifeTime"`
}

type PParams struct {
	KlsrvhSlaveRecDepth    int64 `json:"KLSRVH_SLAVE_REC_DEPTH"`
	KlgrpFindFromCurVsOnly bool  `json:"KLGRP_FIND_FROM_CUR_VS_ONLY"`
}

type HostGroupAttributes struct {
	HostGroupPChunk *HostGroupPChunk `json:"pChunk,omitempty"`
	PxgRetVal       *int64           `json:"PxgRetVal,omitempty"`
}

type HostGroupIteratorArray struct {
	Type           *string         `json:"type,omitempty"`
	HostGroupValue *HostGroupValue `json:"value,omitempty"`
}

type HostGroupValue struct {
	HostGroupID   int64  `json:"id,omitempty"`
	HostGroupName string `json:"name,omitempty"`
}

type HostGroupPChunk struct {
	HostGroupIteratorArray []HostGroupIteratorArray `json:"KLCSP_ITERATOR_ARRAY"`
}

// FindGroups Finds groups that satisfy conditions from filter pParams, and creates a server-side collection of found groups.
// Search is performed over the hierarchy
func (hg *HostGroup) FindGroups(ctx context.Context, params HGParams) (*Accessor, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.FindGroups", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	accessor := new(Accessor)
	raw, err := hg.client.Do(ctx, request, &accessor)
	return accessor, raw, err
}

type HostAttributes struct {
	HostPChunk *HostPChunk `json:"pChunk,omitempty"`
	PxgRetVal  *int64      `json:"PxgRetVal,omitempty"`
}

type HostIteratorArray struct {
	Type      *string    `json:"type,omitempty"`
	HostValue *HostValue `json:"value,omitempty"`
}

type HostValue struct {
	HostDN                  string   `json:"KLHST_WKS_DN,omitempty"`
	HostGroupID             uint64   `json:"KLHST_WKS_GROUPID,omitempty"`
	HostCreated             DateTime `json:"KLHST_WKS_CREATED,omitempty"`
	HostLastVisible         DateTime `json:"KLHST_WKS_LAST_VISIBLE,omitempty"`
	HostLastInfoUpdate      DateTime `json:"KLHST_WKS_LAST_INFOUDATE,omitempty"`
	HostStatus              uint8    `json:"KLHST_WKS_STATUS,omitempty"`
	HostLastUpdate          DateTime `json:"KLHST_WKS_LAST_UPDATE,omitempty"`
	HostLastNAgentConnected DateTime `json:"KLHST_WKS_LAST_NAGENT_CONNECTED,omitempty"`
	HostKeepConnection      bool     `json:"KLHST_WKS_KEEP_CONNECTION,omitempty"`
	HostName                string   `json:"KLHST_WKS_HOSTNAME,omitempty"`
	HostInstanceID          string   `json:"KLHST_INSTANCEID,omitempty"`
	HostWinHostname         string   `json:"KLHST_WKS_WINHOSTNAME,omitempty"`
	HostWinDomain           string   `json:"KLHST_WKS_WINDOMAIN,omitempty"`
	HostDNSDomain           string   `json:"KLHST_WKS_DNSDOMAIN,omitempty"`
	HostDNSName             string   `json:"KLHST_WKS_DNSNAME,omitempty"`
	HostFQDN                string   `json:"KLHST_WKS_FQDN,omitempty"`
	HostCType               uint8    `json:"KLHST_WKS_CTYPE,omitempty"`
	HostPType               uint8    `json:"KLHST_WKS_PTYPE,omitpermit"`
	HostOSName              string   `json:"KLHST_WKS_OS_NAME,omitpermit"`
	HostOSVerMajor          uint8    `json:"KLHST_WKS_OS_VER_MAJOR,omitpermit"`
	HostOSVerMinor          uint8    `json:"KLHST_WKS_OS_VER_MINOR,omitpermit"`
	HostLastFullScan        DateTime `json:"KLHST_WKS_LAST_FULLSCAN,omitpermit"`
	HostVirusCount          uint64   `json:"KLHST_WKS_VIRUS_COUNT,omitpermit"`
	HostRTPState            uint8    `json:"KLHST_WKS_RTP_STATE,omitpermit"`
	HostRTPErrorCode        uint64   `json:"KLHST_WKS_RTP_ERROR_CODE,omitpermit"`
	HostComment             string   `json:"KLHST_WKS_COMMENT,omitpermit"`
	HostIPLong              Long     `json:"KLHST_WKS_IP_LONG,omitpermit"`
	HostConnectIPLong       Long     `json:"KLHST_WKS_CONNECT_IP_LONG,omitpermit"`
	HostFromUnassigned      bool     `json:"KLHST_WKS_FROM_UNASSIGNED,omitpermit"`
	HostUncuredCount        Long     `json:"KLHST_WKS_UNCURED_COUNT,omitpermit"`
}

type HostPChunk struct {
	HostIteratorArray []HostIteratorArray `json:"KLCSP_ITERATOR_ARRAY"`
}

// FindHosts Finds hosts that satisfy conditions from filter string wstrFilter, and creates a server-side collection of found hosts.
// Search is performed over the hierarchy
func (hg *HostGroup) FindHosts(ctx context.Context, params HGParams) (*Accessor, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.FindHosts", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	accessor := new(Accessor)
	raw, err := hg.client.Do(ctx, request, &accessor)
	return accessor, raw, err
}

// FindHostsAsync Find host asynchronously by filter string.
// Finds hosts asynchronously that satisfy conditions from filter string wstrFilter,
// and creates a server-side collection of found hosts. Search is performed over the hierarchy
// to get status use AsyncActionStateChecker.CheckActionState, lStateCode "1" means OK and "0" means fail
// to get accessor id call HostGroup.FindHostsAsyncGetAccessor
// to cancel operation call HostGroup.FindHostsAsyncCancel
func (hg *HostGroup) FindHostsAsync(ctx context.Context, params HGParams) (*RequestID, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.FindHostsAsync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	requestID := new(RequestID)
	raw, err := hg.client.Do(ctx, request, &requestID)
	return requestID, raw, err
}

// FindHostsAsyncCancel Cancels asynchronous operation HostGroup.FindHostsAsync
func (hg *HostGroup) FindHostsAsyncCancel(ctx context.Context, strRequestId string) error {
	postData := []byte(fmt.Sprintf(`{"strRequestId": "%s"}`, strRequestId))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.FindHostsAsyncCancel", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = hg.client.Do(ctx, request, nil)
	if err != nil {
		return err
	}

	return nil
}

// FindHostsAsyncGetAccessor Gets result of asynchronous operation HostGroup.FindHostsAsync
func (hg *HostGroup) FindHostsAsyncGetAccessor(ctx context.Context, strRequestId string) (*AsyncAccessor, []byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"strRequestId" : "%s" }`, strRequestId))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.FindHostsAsyncGetAccessor", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	asyncAccessor := new(AsyncAccessor)
	raw, err := hg.client.Do(ctx, request, &asyncAccessor)
	return asyncAccessor, raw, err
}

// FindIncidentsParams struct
type FindIncidentsParams struct {
	StrFilter       string          `json:"strFilter,omitempty"`
	PFieldsToReturn []string        `json:"pFieldsToReturn,omitempty"`
	PFieldsToOrder  []FieldsToOrder `json:"pFieldsToOrder,omitempty"`
	LMaxLifeTime    int64           `json:"lMaxLifeTime,omitempty"`
}

// FindIncidents Find incident by filter string. Finds incidents that satisfy conditions from filter string strFilter.
func (hg *HostGroup) FindIncidents(ctx context.Context, params FindIncidentsParams) (*Accessor, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.FindIncidents", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	accessor := new(Accessor)
	raw, err := hg.client.Do(ctx, request, &accessor)
	return accessor, raw, err
}

// FindUsers Finds existing users. Finds users that satisfy conditions from filter string strFilter.
func (hg *HostGroup) FindUsers(ctx context.Context, params PFindParams) (*Accessor, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.FindUsers", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	accessor := new(Accessor)
	raw, err := hg.client.Do(ctx, request, &accessor)
	return accessor, raw, err
}

// HostFixes struct
type HostFixes struct {
	PxgRetVal *HostFixesVal `json:"PxgRetVal,omitempty"`
}

type HostFixesVal struct {
	KlhstHFData     []KlhstHFDatum   `json:"KLHST_HF_DATA"`
	KlhstHFProducts []KlhstHFProduct `json:"KLHST_HF_PRODUCTS"`
}

type KlhstHFDatum struct {
	Type  *string            `json:"type,omitempty"`
	Value *KLHSTHFDATUMValue `json:"value,omitempty"`
}

type KLHSTHFDATUMValue struct {
	KlhstHFDN     *string `json:"KLHST_HF_DN,omitempty"`
	KlhstHFID     *string `json:"KLHST_HF_ID,omitempty"`
	KlhstHFProdid *int64  `json:"KLHST_HF_PRODID,omitempty"`
}

type KlhstHFProduct struct {
	Type  *string              `json:"type,omitempty"`
	Value *KLHSTHFPRODUCTValue `json:"value,omitempty"`
}

type KLHSTHFPRODUCTValue struct {
	KlhstHFProdid          *int64  `json:"KLHST_HF_PRODID,omitempty"`
	KlhstWksProductID      *string `json:"KLHST_WKS_PRODUCT_ID,omitempty"`
	KlhstWksProductName    *string `json:"KLHST_WKS_PRODUCT_NAME,omitempty"`
	KlhstWksProductVersion *string `json:"KLHST_WKS_PRODUCT_VERSION,omitempty"`
}

// GetAllHostFixes Returns all hotfixes installed in the network.
func (hg *HostGroup) GetAllHostFixes(ctx context.Context) (*HostFixes, error) {
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetAllHostfixes", nil)
	if err != nil {
		return nil, err
	}
	hostFixes := new(HostFixes)
	_, err = hg.client.Do(ctx, request, &hostFixes)
	return hostFixes, err
}

//	ProductComponents is returned by GetComponentsForProductOnHost
type ProductComponents struct {
	ProductComponentsArray []ProductComponentsArray `json:"PxgRetVal"`
}

type ProductComponentsArray struct {
	Type             string            `json:"type,omitempty"`
	ProductComponent *ProductComponent `json:"value,omitempty"`
}

type ProductComponent struct {
	KlhstPrcstComponentDN      string                      `json:"KLHST_PRCST_COMPONENT_DN,omitempty"`
	KlhstPrcstComponentID      *KlhstPrcstComponentID      `json:"KLHST_PRCST_COMPONENT_ID,omitempty"`
	KlhstPrcstComponentStatus  int64                       `json:"KLHST_PRCST_COMPONENT_STATUS,omitempty"`
	KlhstPrcstComponentVersion *KlhstPrcstComponentVersion `json:"KLHST_PRCST_COMPONENT_VERSION,omitempty"`
}

type KlhstPrcstComponentID struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type KlhstPrcstComponentVersion struct {
	Type  string `json:"type,omitempty"`
	Value int64  `json:"value,omitempty"`
}

// GetComponentsForProductOnHost Return array of product components for specified host and product.
func (hg *HostGroup) GetComponentsForProductOnHost(ctx context.Context, strHostName, strProductName,
	strProductVersion string) (*ProductComponents, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"strHostName": "%s","strProductName": "%s","strProductVersion": "%s"}`,
		strHostName, strProductName, strProductVersion))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetComponentsForProductOnHost", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	var productComponents *ProductComponents
	raw, err := hg.client.Do(ctx, request, &productComponents)
	return productComponents, raw, err
}

// GetDomainHosts Return a list of workstation names in the domain.
//
// The information is obtained from the domain controller.
// This call returns the full list of workstations in the domain, even if the workstation is now turned off.
//
// Deprecated: use either HostGroup.FindHostsAsync or HostGroup.FindHosts instead.
func (hg *HostGroup) GetDomainHosts(ctx context.Context, domain string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"domain": "%s"}`, domain))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetDomainHosts", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

// Domains struct
type Domains struct {
	PxgRetVal []DomainsValue `json:"PxgRetVal"`
}

type DomainsValue struct {
	Type  *string      `json:"type,omitempty"`
	Value *DomainValue `json:"value,omitempty"`
}

type DomainValue struct {
	KlhstWksWindomain     *string `json:"KLHST_WKS_WINDOMAIN,omitempty"`
	KlhstWksWindomainType *int64  `json:"KLHST_WKS_WINDOMAIN_TYPE,omitempty"`
}

// GetDomains List of Windows domain in the network.
func (hg *HostGroup) GetDomains(ctx context.Context) (*Domains, error) {
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetDomains", nil)
	if err != nil {
		return nil, err
	}

	domains := new(Domains)
	_, err = hg.client.Do(ctx, request, &domains)
	return domains, err
}

// GetGroupId Acquire administration group id by its name and id of parent group.
func (hg *HostGroup) GetGroupId(ctx context.Context, nParent int64, strName string) (*PxgValInt, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nParent": %d, "strName": "%s"}`, nParent, strName))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetGroupId", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := hg.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

// GroupInfo struct
type GroupInfo struct {
	PxgRetVal *GroupInfoVal `json:"PxgRetVal,omitempty"`
}

type GroupInfoVal struct {
	KlgrpChldgrpCnt          *int64        `json:"KLGRP_CHLDGRP_CNT,omitempty"`
	KlgrpChldhstCnt          *int64        `json:"KLGRP_CHLDHST_CNT,omitempty"`
	KlgrpChldhstCntCRT       *int64        `json:"KLGRP_CHLDHST_CNT_CRT,omitempty"`
	KlgrpChldhstCntOk        *int64        `json:"KLGRP_CHLDHST_CNT_OK,omitempty"`
	KlgrpChldhstCntWrn       *int64        `json:"KLGRP_CHLDHST_CNT_WRN,omitempty"`
	KLGRPHlfForceChildren    *bool         `json:"KLGRP_HlfForceChildren,omitempty"`
	KLGRPHlfForced           *bool         `json:"KLGRP_HlfForced,omitempty"`
	KlsrvHststatCritical     *KlsrvHststat `json:"KLSRV_HSTSTAT_CRITICAL,omitempty"`
	KlsrvHststatWarning      *KlsrvHststat `json:"KLSRV_HSTSTAT_WARNING,omitempty"`
	AutoRemovePeriod         *int64        `json:"autoRemovePeriod,omitempty"`
	CGrpAutoInstallPackageID []interface{} `json:"c_grp_autoInstallPackageId"`
	ChildGroupsNum           *int64        `json:"childGroupsNum,omitempty"`
	CreationDate             *DateTime     `json:"creationDate,omitempty"`
	GrpEnableFscan           *bool         `json:"grp_enable_fscan,omitempty"`
	GrpFromUnassigned        *bool         `json:"grp_from_unassigned,omitempty"`
	GrpFullName              *string       `json:"grp_full_name,omitempty"`
	HostsNum                 *int64        `json:"hostsNum,omitempty"`
	ID                       *int64        `json:"id,omitempty"`
	Level                    *int64        `json:"level,omitempty"`
	Name                     *string       `json:"name,omitempty"`
	NotifyPeriod             *int64        `json:"notifyPeriod,omitempty"`
	LastUpdate               *DateTime     `json:"lastUpdate,omitempty"`
	ParentID                 *int64        `json:"parentId,omitempty"`
}

type KlsrvHststat struct {
	Type  *string            `json:"type,omitempty"`
	Value *KlsrvHststatValue `json:"value,omitempty"`
}

type KlsrvHststatValue struct {
	KlsrvHststatInheritable     *bool  `json:"KLSRV_HSTSTAT_INHERITABLE,omitempty"`
	KlsrvHststatInherited       *bool  `json:"KLSRV_HSTSTAT_INHERITED,omitempty"`
	KlsrvHststatLocked          *bool  `json:"KLSRV_HSTSTAT_LOCKED,omitempty"`
	KlsrvHststatMask            *int64 `json:"KLSRV_HSTSTAT_MASK,omitempty"`
	KlsrvHststatOldAVBases      *int64 `json:"KLSRV_HSTSTAT_OLD_AV_BASES,omitempty"`
	KlsrvHststatOldFscan        *int64 `json:"KLSRV_HSTSTAT_OLD_FSCAN,omitempty"`
	KlsrvHststatOldLastConnect  *int64 `json:"KLSRV_HSTSTAT_OLD_LAST_CONNECT,omitempty"`
	KlsrvHststatOldLicense      *int64 `json:"KLSRV_HSTSTAT_OLD_LICENSE,omitempty"`
	KlsrvHststatSPBootExpired   *int64 `json:"KLSRV_HSTSTAT_SP_BOOT_EXPIRED,omitempty"`
	KlsrvHststatSPBootReasons   *int64 `json:"KLSRV_HSTSTAT_SP_BOOT_REASONS,omitempty"`
	KlsrvHststatSPDiskspace     *int64 `json:"KLSRV_HSTSTAT_SP_DISKSPACE,omitempty"`
	KlsrvHststatSPEncryption    *int64 `json:"KLSRV_HSTSTAT_SP_ENCRYPTION,omitempty"`
	KlsrvHststatSPRptDiffers    *int64 `json:"KLSRV_HSTSTAT_SP_RPT_DIFFERS,omitempty"`
	KlsrvHststatSPRptInactive   *int64 `json:"KLSRV_HSTSTAT_SP_RPT_INACTIVE,omitempty"`
	KlsrvHststatSPUasearch      *int64 `json:"KLSRV_HSTSTAT_SP_UASEARCH,omitempty"`
	KlsrvHststatSPUncured       *int64 `json:"KLSRV_HSTSTAT_SP_UNCURED,omitempty"`
	KlsrvHststatSPViruses       *int64 `json:"KLSRV_HSTSTAT_SP_VIRUSES,omitempty"`
	KlsrvHststatSPVulnerability *int64 `json:"KLSRV_HSTSTAT_SP_VULNERABILITY,omitempty"`
	KlsrvHststatSPVulnFlags     *int64 `json:"KLSRV_HSTSTAT_SP_VULN_FLAGS,omitempty"`
}

// GetGroupInfo Acquire administration group attributes.
//
// Deprecated: Use HostGroup.GetGroupInfoEx instead
func (hg *HostGroup) GetGroupInfo(ctx context.Context, nGroupId int64) (*GroupInfo, error) {
	postData := []byte(fmt.Sprintf(`{"nGroupId": %d}`, nGroupId))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetGroupInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	groupInfo := new(GroupInfo)
	_, err = hg.client.Do(ctx, request, &groupInfo)
	return groupInfo, err
}

// GroupInfoExParams struct
type GroupInfoExParams struct {
	// NGroupID Id of existing group
	NGroupID int64 `json:"nGroupId,omitempty"`

	// PArrAttributes Array of up to 100 strings. Each entry is an attrbute name (see List of group attributes).
	PArrAttributes []string `json:"pArrAttributes"`
}

// GetGroupInfoEx Acquire administration group attributes.
//
// Remark: not working on KSC 10
func (hg *HostGroup) GetGroupInfoEx(ctx context.Context, params GroupInfoExParams) (*GroupInfo, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetGroupInfoEx", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	groupInfo := new(GroupInfo)
	raw, err := hg.client.Do(ctx, request, &groupInfo)
	return groupInfo, raw, err
}

// ProductFixes struct
type ProductFixes struct {
	Fixes []Fixes `json:"PxgRetVal"`
}

type Fixes struct {
	Type       string     `json:"type"`
	FixesValue FixesValue `json:"value"`
}

type FixesValue struct {
	KlhstHFDN string `json:"KLHST_HF_DN"`
	KlhstHFID string `json:"KLHST_HF_ID"`
}

// GetHostfixesForProductOnHost Return array of hotfixes for specified host and product.
// Array is ordered according hotfix installation order.
func (hg *HostGroup) GetHostfixesForProductOnHost(ctx context.Context, strHostName, strProductName, strProductVersion string) (*ProductFixes, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"strHostName": "%s","strProductName": "%s","strProductVersion": "%s"}`, strHostName, strProductName, strProductVersion))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetHostfixesForProductOnHost", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	productFixes := new(ProductFixes)
	raw, err := hg.client.Do(ctx, request, &productFixes)
	return productFixes, raw, err
}

// GetHostInfo Acquire specified host attributes.
func (hg *HostGroup) GetHostInfo(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetHostInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

// GetHostProducts Return information about installed products on the host.
func (hg *HostGroup) GetHostProducts(ctx context.Context, strHostName string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strHostName": "%s"}`, strHostName))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetHostProducts", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

// GetHostTasks Return server specific identity to acquire and manage host tasks.
func (hg *HostGroup) GetHostTasks(ctx context.Context, hostId string) (*PxgValStr, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"strHostName": "%s"}`, hostId))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetHostTasks", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := hg.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// InstanceStatisticsParams struct
type InstanceStatisticsParams struct {
	// VecFilterFields Array of needed attributes
	// Remark: can be empty, but not nil
	VecFilterFields []string `json:"vecFilterFields"`
}

// GetInstanceStatistics Acquire Server statistics info.
//
// Remark: not working on KSC 10
func (hg *HostGroup) GetInstanceStatistics(ctx context.Context, params InstanceStatisticsParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetInstanceStatistics", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

// StaticInfoParams struct
type StaticInfoParams struct {
	PValues []string `json:"pValues"`
}

// GetRunTimeInfo Return server run-time info.
func (hg *HostGroup) GetRunTimeInfo(ctx context.Context, params StaticInfoParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetRunTimeInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

// GetStaticInfo Return server static info.
func (hg *HostGroup) GetStaticInfo(ctx context.Context, params StaticInfoParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetStaticInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

// SubGroups struct
type SubGroups struct {
	PxgRetVal []SubGroup `json:"PxgRetVal"`
}

type SubGroup struct {
	Type  *string        `json:"type,omitempty"`
	Value *SubGroupValue `json:"value,omitempty"`
}

type SubGroupValue struct {
	GrpPartOfAdViewByRule *bool   `json:"grp_part_of_ad_view_by_rule,omitempty"`
	ID                    *int64  `json:"id,omitempty"`
	Name                  *string `json:"name,omitempty"`
}

// GetSubgroups Acquire administration group subgroups tree.
func (hg *HostGroup) GetSubgroups(ctx context.Context, nGroupId int64, nDepth int64) (*SubGroups, error) {
	postData := []byte(fmt.Sprintf(`{"nParent": %d, "nDepth": %d }`, nGroupId, nDepth))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GetSubgroups", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	subGroups := new(SubGroups)
	_, err = hg.client.Do(ctx, request, &subGroups)
	return subGroups, err
}

// GroupIdGroups Id of predefined root group "Managed computers".
func (hg *HostGroup) GroupIdGroups(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GroupIdGroups", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := hg.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

// GroupIdSuper Id of predefined group "Master server".
func (hg *HostGroup) GroupIdSuper(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GroupIdSuper", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := hg.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

// GroupIdUnassigned Id of predefined group "Unassigned computers".
func (hg *HostGroup) GroupIdUnassigned(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.GroupIdUnassigned", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := hg.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

// MoveHostsFromGroupToGroup Moves hosts from root of source group to root of destination group. Operation is asynchronous.
func (hg *HostGroup) MoveHostsFromGroupToGroup(ctx context.Context, nSrcGroupId int64,
	nDstGroupId int64) (*WActionGUID, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nSrcGroupId": %d, "nDstGroupId": %d}`, nSrcGroupId, nDstGroupId))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.MoveHostsFromGroupToGroup", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	wActionGUID := new(WActionGUID)
	raw, err := hg.client.Do(ctx, request, &wActionGUID)
	return wActionGUID, raw, err
}

// HostsToGroupParams struct
type HostsToGroupParams struct {
	NGroup     int64    `json:"nGroup"`
	PHostNames []string `json:"pHostNames"`
}

// MoveHostsToGroup Move multiple hosts into specified administration group.
func (hg *HostGroup) MoveHostsToGroup(ctx context.Context, params HostsToGroupParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.MoveHostsToGroup",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

// RemoveGroup Delete administration group.
func (hg *HostGroup) RemoveGroup(ctx context.Context, nGroup, nFlags int64) (*WActionGUID, []byte, error) {
	postData := []byte(fmt.Sprintf(`{ "nGroup": %d, "nFlags": %d }`, nGroup, nFlags))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.RemoveGroup", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	wActionGUID := new(WActionGUID)
	raw, err := hg.client.Do(ctx, request, &wActionGUID)
	return wActionGUID, raw, err
}

// RemoveHost Removes host record.
func (hg *HostGroup) RemoveHost(ctx context.Context, strHostName string) error {
	postData := []byte(fmt.Sprintf(`{ "strHostName": "%s" }`, strHostName))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.RemoveHost", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = hg.client.Do(ctx, request, nil)
	if err != nil {
		return err
	}
	return nil
}

// RemoveHostsParams struct
type RemoveHostsParams struct {
	//Array of host names
	PHostNames []string `json:"pHostNames"`
	//Whether to force deleting hosts records
	BForceDestroy bool `json:"bForceDestroy"`
}

// RemoveHosts Remove multiple hosts.
//
// Removes multiple hosts. Function behavior depends on bForceDestroy flag.
//
//If bForceDestroy is true then hosts records are deleted.
//
//If bForceDestroy is false hosts records will be deleted only for hosts located in group "Unassigned computers"
// or its subgroups, others will be moved into corresponding subgroups of group "Unassigned computers".
func (hg *HostGroup) RemoveHosts(ctx context.Context, params RemoveHostsParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.RemoveHosts", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

// KlhstWksResults struct
type KlhstWksResults struct {
	PResults PResults `json:"pResults"`
}

type PResults struct {
	KlhstWksResults []bool `json:"KLHST_WKS_RESULTS"`
}

// PInfoRaM struct
type PInfoRaM struct {
	PInfo PInfo `json:"pInfo"`
}

type PInfo struct {
	KlhstWksAnyname []string `json:"KLHST_WKS_ANYNAME"`
	KlhstWksGroupid int64    `json:"KLHST_WKS_GROUPID"`
}

// ResolveAndMoveToGroup Moves hosts into a group by name or ip-address.
//
// If the entered name corresponds to the ip-address format, then the server tries to find in the database a host with the indicated ip-address.
// Otherwise, the server tries to interpret the host as follows:
//
// 1. Name (KLHST_WKS_HOSTNAME)
//
//2. Display name (KLHST_WKS_DN)
//
//3. NetBIOS name (KLHST_WKS_WINHOSTNAME)
//
//4. DNS name (KLHST_WKS_DNSNAME)
func (hg *HostGroup) ResolveAndMoveToGroup(ctx context.Context, params PInfoRaM) (*KlhstWksResults, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.ResolveAndMoveToGroup", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	klhstWksResults := new(KlhstWksResults)
	raw, err := hg.client.Do(ctx, request, &klhstWksResults)
	return klhstWksResults, raw, err
}

// RestartNetworkScanning Restarts specified network scanning type.
func (hg *HostGroup) RestartNetworkScanning(ctx context.Context, nType int64) (*PxgRetError, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nType": %d	}`, nType))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.RestartNetworkScanning", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgRetError := new(PxgRetError)
	raw, err := hg.client.Do(ctx, request, &pxgRetError)
	return pxgRetError, raw, err
}

// SetLocInfo Allows to set server localization information.
func (hg *HostGroup) SetLocInfo(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.SetLocInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

//SectionParams struct using in HostGroup.SS_CreateSection | HostGroup.SS_DeleteSection
type SectionParams struct {
	//host name (unique server-generated string)
	StrHostName string `json:"strHostName,omitempty"`

	//type of storage (for example: "SS_SETTINGS")
	StrType string `json:"strType,omitempty"`

	//product name string, non-empty string, not longer than 31 character, and cannot contain characters /\:*?"<>.
	StrProduct string `json:"strProduct,omitempty"`

	//version string, non-empty string, not longer than 31 character, and cannot contain characters /\:*?"<>.
	StrVersion string `json:"strVersion,omitempty"`

	//section name string, non-empty string, not longer than 31 character, and cannot contain characters /\:*?"<>.
	StrSection string `json:"strSection,omitempty"`

	//write option, values:
	//	1 - "Update", updates existing variables in the specified section. If a variable does not exist an error occurs.
	//	2 - "Add", adds new variables to the specified section. If a variable already exists an error occurs.
	//	3 - "Replace", replaces variables in the specified section. If a variable already exists it will be updates, if a variable does not exist it will be added.
	//	4 - "Delete", deletes variables specified in pData from the specified section.
	//	7 - "Clear", replaces existing section contents with pData, i.e. existing section contents will deleted and variables from pData will be written to the section.
	NOption   int64       `json:"nOption,omitempty"`
	PSettings interface{} `json:"pSettings,omitempty"`
}

// SSCreateSection Create section in host settings storage.
func (hg *HostGroup) SSCreateSection(ctx context.Context, params SectionParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.SS_CreateSection", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

// SSWrite Write data to host settings storage.
func (hg *HostGroup) SSWrite(ctx context.Context, params SectionParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.SS_Write", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

// SSGetNames Get section names from host settings storage.
//
// If product is empty then names will contain all product names.
// If product is not empty and version is empty then names will contain all versions for the specified product name.
// If product is not empty and version is not empty then names will contain all sections for the specified product and version.
func (hg *HostGroup) SSGetNames(ctx context.Context, params SectionParams) (*PxgValArrayOfString, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.SS_GetNames", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValArrayOfString := new(PxgValArrayOfString)
	raw, err := hg.client.Do(ctx, request, &pxgValArrayOfString)
	return pxgValArrayOfString, raw, err
}

// SSRead Read data from host settings storage.
func (hg *HostGroup) SSRead(ctx context.Context, params SectionParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.SS_Read", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

// UpdateGroupParam struct
type UpdateGroupParam struct {
	NGroup int64      `json:"nGroup,omitempty"`
	PInfo  PInfoGroup `json:"pInfo,omitempty"`
}

type PInfoGroup struct {
	Name                     string      `json:"name,omitempty"`
	ParentID                 int64       `json:"parentId,omitempty"`
	AutoRemovePeriod         bool        `json:"autoRemovePeriod,omitempty"`
	NotifyPeriod             int64       `json:"notifyPeriod,omitempty"`
	KLGRPHlfInherited        bool        `json:"KLGRP_HlfInherited,omitempty"`
	KLGRPHlfForceChildren    bool        `json:"KLGRP_HlfForceChildren,omitempty"`
	Level                    int64       `json:"level,omitempty"`
	KlsrvHststatCritical     interface{} `json:"KLSRV_HSTSTAT_CRITICAL,omitempty"`
	KlsrvHststatWarning      interface{} `json:"KLSRV_HSTSTAT_WARNING,omitempty"`
	CGrpAutoInstallPackageID []int64     `json:"c_grp_autoInstallPackageId"`
	GrpEnableFscan           bool        `json:"grp_enable_fscan,omitempty"`
}

// UpdateGroup Change attributes of existing administration group.
func (hg *HostGroup) UpdateGroup(ctx context.Context, params UpdateGroupParam) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.UpdateGroup", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

// UpdateHost Modify specified attributes for host.
func (hg *HostGroup) UpdateHost(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.UpdateHost", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

// UpdateHostsMultiple Update attributes of multiple computers.
func (hg *HostGroup) UpdateHostsMultiple(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.UpdateHostsMultiple", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

// UpdateIncidentParams struct using in HostGroup.UpdateIncident
type UpdateIncidentParams struct {
	NID   int64          `json:"nId,omitempty"`
	PData *PIncidentData `json:"pData,omitempty"`
}

type PIncidentData struct {
	KlhstWksHostname string `json:"KLHST_WKS_HOSTNAME"`
	//Incident body
	KlincdtBody string `json:"KLINCDT_BODY,omitempty"`
	//Incident severity
	KlincdtSeverity int64 `json:"KLINCDT_SEVERITY,omitempty"`
	//Time of incident entry creation
	KlincdtAdded *DateTime `json:"KLINCDT_ADDED,omitempty"`
	//"IsHandled" flag. True - if incident marked as "Handled", otherwise false
	KlincdtIsHandled bool `json:"KLINCDT_IS_HANDLED,omitempty"`
}

// UpdateIncident Modify properties of an existing incident.
func (hg *HostGroup) UpdateIncident(ctx context.Context, params UpdateIncidentParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.UpdateIncident", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hg.client.Do(ctx, request, nil)
	return raw, err
}

// ZeroVirusCountForGroup Zero virus count for hosts in group and all subgroups.
func (hg *HostGroup) ZeroVirusCountForGroup(ctx context.Context, nParent int64) (*WActionGUID, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nParent": %d}`, nParent))
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.ZeroVirusCountForGroup", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	wActionGUID := new(WActionGUID)
	raw, err := hg.client.Do(ctx, request, &wActionGUID)
	return wActionGUID, raw, err
}

// ZeroVirusCountForHosts Zero virus count for specified hosts.
func (hg *HostGroup) ZeroVirusCountForHosts(ctx context.Context, params interface{}) (*WActionGUID, []byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", hg.client.Server+"/api/v1.0/HostGroup.ZeroVirusCountForHosts", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	wActionGUID := new(WActionGUID)
	raw, err := hg.client.Do(ctx, request, &wActionGUID)
	return wActionGUID, raw, err
}
