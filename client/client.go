// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateDeviceRequest struct {
	ChannelId         *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	City              *string                `json:"City,omitempty" xml:"City,omitempty"`
	DeviceModelNumber *string                `json:"DeviceModelNumber,omitempty" xml:"DeviceModelNumber,omitempty"`
	DeviceName        *string                `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceType        *string                `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	District          *string                `json:"District,omitempty" xml:"District,omitempty"`
	ExtraMap          map[string]interface{} `json:"ExtraMap,omitempty" xml:"ExtraMap,omitempty"`
	FirstScene        *string                `json:"FirstScene,omitempty" xml:"FirstScene,omitempty"`
	Floor             *string                `json:"Floor,omitempty" xml:"Floor,omitempty"`
	LocationName      *string                `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
	MediaId           *string                `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OuterCode         *string                `json:"OuterCode,omitempty" xml:"OuterCode,omitempty"`
	Province          *string                `json:"Province,omitempty" xml:"Province,omitempty"`
	SecondScene       *string                `json:"SecondScene,omitempty" xml:"SecondScene,omitempty"`
}

func (s CreateDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceRequest) GoString() string {
	return s.String()
}

func (s *CreateDeviceRequest) SetChannelId(v string) *CreateDeviceRequest {
	s.ChannelId = &v
	return s
}

func (s *CreateDeviceRequest) SetCity(v string) *CreateDeviceRequest {
	s.City = &v
	return s
}

func (s *CreateDeviceRequest) SetDeviceModelNumber(v string) *CreateDeviceRequest {
	s.DeviceModelNumber = &v
	return s
}

func (s *CreateDeviceRequest) SetDeviceName(v string) *CreateDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *CreateDeviceRequest) SetDeviceType(v string) *CreateDeviceRequest {
	s.DeviceType = &v
	return s
}

func (s *CreateDeviceRequest) SetDistrict(v string) *CreateDeviceRequest {
	s.District = &v
	return s
}

func (s *CreateDeviceRequest) SetExtraMap(v map[string]interface{}) *CreateDeviceRequest {
	s.ExtraMap = v
	return s
}

func (s *CreateDeviceRequest) SetFirstScene(v string) *CreateDeviceRequest {
	s.FirstScene = &v
	return s
}

func (s *CreateDeviceRequest) SetFloor(v string) *CreateDeviceRequest {
	s.Floor = &v
	return s
}

func (s *CreateDeviceRequest) SetLocationName(v string) *CreateDeviceRequest {
	s.LocationName = &v
	return s
}

func (s *CreateDeviceRequest) SetMediaId(v string) *CreateDeviceRequest {
	s.MediaId = &v
	return s
}

func (s *CreateDeviceRequest) SetOuterCode(v string) *CreateDeviceRequest {
	s.OuterCode = &v
	return s
}

func (s *CreateDeviceRequest) SetProvince(v string) *CreateDeviceRequest {
	s.Province = &v
	return s
}

func (s *CreateDeviceRequest) SetSecondScene(v string) *CreateDeviceRequest {
	s.SecondScene = &v
	return s
}

type CreateDeviceShrinkRequest struct {
	ChannelId         *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	City              *string `json:"City,omitempty" xml:"City,omitempty"`
	DeviceModelNumber *string `json:"DeviceModelNumber,omitempty" xml:"DeviceModelNumber,omitempty"`
	DeviceName        *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceType        *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	District          *string `json:"District,omitempty" xml:"District,omitempty"`
	ExtraMapShrink    *string `json:"ExtraMap,omitempty" xml:"ExtraMap,omitempty"`
	FirstScene        *string `json:"FirstScene,omitempty" xml:"FirstScene,omitempty"`
	Floor             *string `json:"Floor,omitempty" xml:"Floor,omitempty"`
	LocationName      *string `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
	MediaId           *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OuterCode         *string `json:"OuterCode,omitempty" xml:"OuterCode,omitempty"`
	Province          *string `json:"Province,omitempty" xml:"Province,omitempty"`
	SecondScene       *string `json:"SecondScene,omitempty" xml:"SecondScene,omitempty"`
}

func (s CreateDeviceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDeviceShrinkRequest) SetChannelId(v string) *CreateDeviceShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *CreateDeviceShrinkRequest) SetCity(v string) *CreateDeviceShrinkRequest {
	s.City = &v
	return s
}

func (s *CreateDeviceShrinkRequest) SetDeviceModelNumber(v string) *CreateDeviceShrinkRequest {
	s.DeviceModelNumber = &v
	return s
}

func (s *CreateDeviceShrinkRequest) SetDeviceName(v string) *CreateDeviceShrinkRequest {
	s.DeviceName = &v
	return s
}

func (s *CreateDeviceShrinkRequest) SetDeviceType(v string) *CreateDeviceShrinkRequest {
	s.DeviceType = &v
	return s
}

func (s *CreateDeviceShrinkRequest) SetDistrict(v string) *CreateDeviceShrinkRequest {
	s.District = &v
	return s
}

func (s *CreateDeviceShrinkRequest) SetExtraMapShrink(v string) *CreateDeviceShrinkRequest {
	s.ExtraMapShrink = &v
	return s
}

func (s *CreateDeviceShrinkRequest) SetFirstScene(v string) *CreateDeviceShrinkRequest {
	s.FirstScene = &v
	return s
}

func (s *CreateDeviceShrinkRequest) SetFloor(v string) *CreateDeviceShrinkRequest {
	s.Floor = &v
	return s
}

func (s *CreateDeviceShrinkRequest) SetLocationName(v string) *CreateDeviceShrinkRequest {
	s.LocationName = &v
	return s
}

func (s *CreateDeviceShrinkRequest) SetMediaId(v string) *CreateDeviceShrinkRequest {
	s.MediaId = &v
	return s
}

func (s *CreateDeviceShrinkRequest) SetOuterCode(v string) *CreateDeviceShrinkRequest {
	s.OuterCode = &v
	return s
}

func (s *CreateDeviceShrinkRequest) SetProvince(v string) *CreateDeviceShrinkRequest {
	s.Province = &v
	return s
}

func (s *CreateDeviceShrinkRequest) SetSecondScene(v string) *CreateDeviceShrinkRequest {
	s.SecondScene = &v
	return s
}

type CreateDeviceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Model     *string `json:"Model,omitempty" xml:"Model,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeviceResponseBody) SetCode(v string) *CreateDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDeviceResponseBody) SetMessage(v string) *CreateDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDeviceResponseBody) SetModel(v string) *CreateDeviceResponseBody {
	s.Model = &v
	return s
}

func (s *CreateDeviceResponseBody) SetRequestId(v string) *CreateDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeviceResponseBody) SetSuccess(v bool) *CreateDeviceResponseBody {
	s.Success = &v
	return s
}

type CreateDeviceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceResponse) GoString() string {
	return s.String()
}

func (s *CreateDeviceResponse) SetHeaders(v map[string]*string) *CreateDeviceResponse {
	s.Headers = v
	return s
}

func (s *CreateDeviceResponse) SetStatusCode(v int32) *CreateDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeviceResponse) SetBody(v *CreateDeviceResponseBody) *CreateDeviceResponse {
	s.Body = v
	return s
}

type DeleteCreativeInfoRequest struct {
	AccountNo  *string `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	MainId     *int64  `json:"MainId,omitempty" xml:"MainId,omitempty"`
	UpdateUser *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
}

func (s DeleteCreativeInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCreativeInfoRequest) GoString() string {
	return s.String()
}

func (s *DeleteCreativeInfoRequest) SetAccountNo(v string) *DeleteCreativeInfoRequest {
	s.AccountNo = &v
	return s
}

func (s *DeleteCreativeInfoRequest) SetBizId(v string) *DeleteCreativeInfoRequest {
	s.BizId = &v
	return s
}

func (s *DeleteCreativeInfoRequest) SetId(v int64) *DeleteCreativeInfoRequest {
	s.Id = &v
	return s
}

func (s *DeleteCreativeInfoRequest) SetMainId(v int64) *DeleteCreativeInfoRequest {
	s.MainId = &v
	return s
}

func (s *DeleteCreativeInfoRequest) SetUpdateUser(v string) *DeleteCreativeInfoRequest {
	s.UpdateUser = &v
	return s
}

type DeleteCreativeInfoResponseBody struct {
	Code         *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCreativeInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCreativeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCreativeInfoResponseBody) SetCode(v int32) *DeleteCreativeInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCreativeInfoResponseBody) SetData(v bool) *DeleteCreativeInfoResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCreativeInfoResponseBody) SetErrorMessage(v string) *DeleteCreativeInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteCreativeInfoResponseBody) SetRequestId(v string) *DeleteCreativeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCreativeInfoResponseBody) SetSuccess(v bool) *DeleteCreativeInfoResponseBody {
	s.Success = &v
	return s
}

type DeleteCreativeInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCreativeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCreativeInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCreativeInfoResponse) GoString() string {
	return s.String()
}

func (s *DeleteCreativeInfoResponse) SetHeaders(v map[string]*string) *DeleteCreativeInfoResponse {
	s.Headers = v
	return s
}

func (s *DeleteCreativeInfoResponse) SetStatusCode(v int32) *DeleteCreativeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCreativeInfoResponse) SetBody(v *DeleteCreativeInfoResponseBody) *DeleteCreativeInfoResponse {
	s.Body = v
	return s
}

type GetBrandPageRequest struct {
	AccountNo *string `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	MainId    *int64  `json:"MainId,omitempty" xml:"MainId,omitempty"`
	MainName  *string `json:"MainName,omitempty" xml:"MainName,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetBrandPageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBrandPageRequest) GoString() string {
	return s.String()
}

func (s *GetBrandPageRequest) SetAccountNo(v string) *GetBrandPageRequest {
	s.AccountNo = &v
	return s
}

func (s *GetBrandPageRequest) SetMainId(v int64) *GetBrandPageRequest {
	s.MainId = &v
	return s
}

func (s *GetBrandPageRequest) SetMainName(v string) *GetBrandPageRequest {
	s.MainName = &v
	return s
}

func (s *GetBrandPageRequest) SetPageIndex(v int32) *GetBrandPageRequest {
	s.PageIndex = &v
	return s
}

func (s *GetBrandPageRequest) SetPageSize(v int32) *GetBrandPageRequest {
	s.PageSize = &v
	return s
}

type GetBrandPageResponseBody struct {
	Code      *int64                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetBrandPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBrandPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBrandPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetBrandPageResponseBody) SetCode(v int64) *GetBrandPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetBrandPageResponseBody) SetData(v *GetBrandPageResponseBodyData) *GetBrandPageResponseBody {
	s.Data = v
	return s
}

func (s *GetBrandPageResponseBody) SetRequestId(v string) *GetBrandPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBrandPageResponseBody) SetSuccess(v bool) *GetBrandPageResponseBody {
	s.Success = &v
	return s
}

type GetBrandPageResponseBodyData struct {
	List     []*GetBrandPageResponseBodyDataList   `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageInfo *GetBrandPageResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
}

func (s GetBrandPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetBrandPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBrandPageResponseBodyData) SetList(v []*GetBrandPageResponseBodyDataList) *GetBrandPageResponseBodyData {
	s.List = v
	return s
}

func (s *GetBrandPageResponseBodyData) SetPageInfo(v *GetBrandPageResponseBodyDataPageInfo) *GetBrandPageResponseBodyData {
	s.PageInfo = v
	return s
}

type GetBrandPageResponseBodyDataList struct {
	AccountNo    *string `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	AccountType  *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	Company      *string `json:"Company,omitempty" xml:"Company,omitempty"`
	MainId       *int64  `json:"MainId,omitempty" xml:"MainId,omitempty"`
	MainName     *string `json:"MainName,omitempty" xml:"MainName,omitempty"`
	ParentMainId *int64  `json:"ParentMainId,omitempty" xml:"ParentMainId,omitempty"`
}

func (s GetBrandPageResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetBrandPageResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetBrandPageResponseBodyDataList) SetAccountNo(v string) *GetBrandPageResponseBodyDataList {
	s.AccountNo = &v
	return s
}

func (s *GetBrandPageResponseBodyDataList) SetAccountType(v string) *GetBrandPageResponseBodyDataList {
	s.AccountType = &v
	return s
}

func (s *GetBrandPageResponseBodyDataList) SetCompany(v string) *GetBrandPageResponseBodyDataList {
	s.Company = &v
	return s
}

func (s *GetBrandPageResponseBodyDataList) SetMainId(v int64) *GetBrandPageResponseBodyDataList {
	s.MainId = &v
	return s
}

func (s *GetBrandPageResponseBodyDataList) SetMainName(v string) *GetBrandPageResponseBodyDataList {
	s.MainName = &v
	return s
}

func (s *GetBrandPageResponseBodyDataList) SetParentMainId(v int64) *GetBrandPageResponseBodyDataList {
	s.ParentMainId = &v
	return s
}

type GetBrandPageResponseBodyDataPageInfo struct {
	Page        *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalNumber *int32 `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s GetBrandPageResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s GetBrandPageResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *GetBrandPageResponseBodyDataPageInfo) SetPage(v int32) *GetBrandPageResponseBodyDataPageInfo {
	s.Page = &v
	return s
}

func (s *GetBrandPageResponseBodyDataPageInfo) SetPageSize(v int32) *GetBrandPageResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetBrandPageResponseBodyDataPageInfo) SetTotalNumber(v int32) *GetBrandPageResponseBodyDataPageInfo {
	s.TotalNumber = &v
	return s
}

type GetBrandPageResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBrandPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBrandPageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBrandPageResponse) GoString() string {
	return s.String()
}

func (s *GetBrandPageResponse) SetHeaders(v map[string]*string) *GetBrandPageResponse {
	s.Headers = v
	return s
}

func (s *GetBrandPageResponse) SetStatusCode(v int32) *GetBrandPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBrandPageResponse) SetBody(v *GetBrandPageResponseBody) *GetBrandPageResponse {
	s.Body = v
	return s
}

type GetBusinessIdRequest struct {
	BusinessId *string `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
}

func (s GetBusinessIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBusinessIdRequest) GoString() string {
	return s.String()
}

func (s *GetBusinessIdRequest) SetBusinessId(v string) *GetBusinessIdRequest {
	s.BusinessId = &v
	return s
}

type GetBusinessIdResponseBody struct {
	Code         *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBusinessIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBusinessIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetBusinessIdResponseBody) SetCode(v int32) *GetBusinessIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetBusinessIdResponseBody) SetData(v string) *GetBusinessIdResponseBody {
	s.Data = &v
	return s
}

func (s *GetBusinessIdResponseBody) SetErrorMessage(v string) *GetBusinessIdResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetBusinessIdResponseBody) SetRequestId(v string) *GetBusinessIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBusinessIdResponseBody) SetSuccess(v bool) *GetBusinessIdResponseBody {
	s.Success = &v
	return s
}

type GetBusinessIdResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBusinessIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBusinessIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBusinessIdResponse) GoString() string {
	return s.String()
}

func (s *GetBusinessIdResponse) SetHeaders(v map[string]*string) *GetBusinessIdResponse {
	s.Headers = v
	return s
}

func (s *GetBusinessIdResponse) SetStatusCode(v int32) *GetBusinessIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBusinessIdResponse) SetBody(v *GetBusinessIdResponseBody) *GetBusinessIdResponse {
	s.Body = v
	return s
}

type GetCreativeInfoRequest struct {
	AccountNo *string `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	BizId     *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	MainId    *int64  `json:"MainId,omitempty" xml:"MainId,omitempty"`
}

func (s GetCreativeInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCreativeInfoRequest) GoString() string {
	return s.String()
}

func (s *GetCreativeInfoRequest) SetAccountNo(v string) *GetCreativeInfoRequest {
	s.AccountNo = &v
	return s
}

func (s *GetCreativeInfoRequest) SetBizId(v string) *GetCreativeInfoRequest {
	s.BizId = &v
	return s
}

func (s *GetCreativeInfoRequest) SetId(v int64) *GetCreativeInfoRequest {
	s.Id = &v
	return s
}

func (s *GetCreativeInfoRequest) SetMainId(v int64) *GetCreativeInfoRequest {
	s.MainId = &v
	return s
}

type GetCreativeInfoResponseBody struct {
	Code         *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *GetCreativeInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCreativeInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCreativeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetCreativeInfoResponseBody) SetCode(v int32) *GetCreativeInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetCreativeInfoResponseBody) SetData(v *GetCreativeInfoResponseBodyData) *GetCreativeInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetCreativeInfoResponseBody) SetErrorMessage(v string) *GetCreativeInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetCreativeInfoResponseBody) SetRequestId(v string) *GetCreativeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCreativeInfoResponseBody) SetSuccess(v bool) *GetCreativeInfoResponseBody {
	s.Success = &v
	return s
}

type GetCreativeInfoResponseBodyData struct {
	AccountNo       *string `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	ChainValue      *string `json:"ChainValue,omitempty" xml:"ChainValue,omitempty"`
	ComponentIdList *string `json:"ComponentIdList,omitempty" xml:"ComponentIdList,omitempty"`
	Id              *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	MainId          *int64  `json:"MainId,omitempty" xml:"MainId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageId          *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskCount       *int32  `json:"TaskCount,omitempty" xml:"TaskCount,omitempty"`
	Url             *string `json:"Url,omitempty" xml:"Url,omitempty"`
	UrlType         *string `json:"UrlType,omitempty" xml:"UrlType,omitempty"`
}

func (s GetCreativeInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCreativeInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCreativeInfoResponseBodyData) SetAccountNo(v string) *GetCreativeInfoResponseBodyData {
	s.AccountNo = &v
	return s
}

func (s *GetCreativeInfoResponseBodyData) SetChainValue(v string) *GetCreativeInfoResponseBodyData {
	s.ChainValue = &v
	return s
}

func (s *GetCreativeInfoResponseBodyData) SetComponentIdList(v string) *GetCreativeInfoResponseBodyData {
	s.ComponentIdList = &v
	return s
}

func (s *GetCreativeInfoResponseBodyData) SetId(v int32) *GetCreativeInfoResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetCreativeInfoResponseBodyData) SetMainId(v int64) *GetCreativeInfoResponseBodyData {
	s.MainId = &v
	return s
}

func (s *GetCreativeInfoResponseBodyData) SetName(v string) *GetCreativeInfoResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetCreativeInfoResponseBodyData) SetPageId(v string) *GetCreativeInfoResponseBodyData {
	s.PageId = &v
	return s
}

func (s *GetCreativeInfoResponseBodyData) SetStatus(v int32) *GetCreativeInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetCreativeInfoResponseBodyData) SetTaskCount(v int32) *GetCreativeInfoResponseBodyData {
	s.TaskCount = &v
	return s
}

func (s *GetCreativeInfoResponseBodyData) SetUrl(v string) *GetCreativeInfoResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetCreativeInfoResponseBodyData) SetUrlType(v string) *GetCreativeInfoResponseBodyData {
	s.UrlType = &v
	return s
}

type GetCreativeInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCreativeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCreativeInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCreativeInfoResponse) GoString() string {
	return s.String()
}

func (s *GetCreativeInfoResponse) SetHeaders(v map[string]*string) *GetCreativeInfoResponse {
	s.Headers = v
	return s
}

func (s *GetCreativeInfoResponse) SetStatusCode(v int32) *GetCreativeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCreativeInfoResponse) SetBody(v *GetCreativeInfoResponseBody) *GetCreativeInfoResponse {
	s.Body = v
	return s
}

type GetLeadsListPageRequest struct {
	ComponentId *int64 `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	ContentId   *int64 `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	CreativeId  *int64 `json:"CreativeId,omitempty" xml:"CreativeId,omitempty"`
	EndTime     *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MainId      *int64 `json:"MainId,omitempty" xml:"MainId,omitempty"`
	PageIndex   *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime   *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskId      *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetLeadsListPageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLeadsListPageRequest) GoString() string {
	return s.String()
}

func (s *GetLeadsListPageRequest) SetComponentId(v int64) *GetLeadsListPageRequest {
	s.ComponentId = &v
	return s
}

func (s *GetLeadsListPageRequest) SetContentId(v int64) *GetLeadsListPageRequest {
	s.ContentId = &v
	return s
}

func (s *GetLeadsListPageRequest) SetCreativeId(v int64) *GetLeadsListPageRequest {
	s.CreativeId = &v
	return s
}

func (s *GetLeadsListPageRequest) SetEndTime(v int64) *GetLeadsListPageRequest {
	s.EndTime = &v
	return s
}

func (s *GetLeadsListPageRequest) SetMainId(v int64) *GetLeadsListPageRequest {
	s.MainId = &v
	return s
}

func (s *GetLeadsListPageRequest) SetPageIndex(v int32) *GetLeadsListPageRequest {
	s.PageIndex = &v
	return s
}

func (s *GetLeadsListPageRequest) SetPageSize(v int32) *GetLeadsListPageRequest {
	s.PageSize = &v
	return s
}

func (s *GetLeadsListPageRequest) SetStartTime(v int64) *GetLeadsListPageRequest {
	s.StartTime = &v
	return s
}

func (s *GetLeadsListPageRequest) SetTaskId(v int64) *GetLeadsListPageRequest {
	s.TaskId = &v
	return s
}

type GetLeadsListPageResponseBody struct {
	Code         *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *GetLeadsListPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLeadsListPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLeadsListPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetLeadsListPageResponseBody) SetCode(v int32) *GetLeadsListPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetLeadsListPageResponseBody) SetData(v *GetLeadsListPageResponseBodyData) *GetLeadsListPageResponseBody {
	s.Data = v
	return s
}

func (s *GetLeadsListPageResponseBody) SetErrorMessage(v string) *GetLeadsListPageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetLeadsListPageResponseBody) SetRequestId(v string) *GetLeadsListPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLeadsListPageResponseBody) SetSuccess(v bool) *GetLeadsListPageResponseBody {
	s.Success = &v
	return s
}

type GetLeadsListPageResponseBodyData struct {
	List     []*GetLeadsListPageResponseBodyDataList   `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageInfo *GetLeadsListPageResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
}

func (s GetLeadsListPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetLeadsListPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLeadsListPageResponseBodyData) SetList(v []*GetLeadsListPageResponseBodyDataList) *GetLeadsListPageResponseBodyData {
	s.List = v
	return s
}

func (s *GetLeadsListPageResponseBodyData) SetPageInfo(v *GetLeadsListPageResponseBodyDataPageInfo) *GetLeadsListPageResponseBodyData {
	s.PageInfo = v
	return s
}

type GetLeadsListPageResponseBodyDataList struct {
	ComponentId  *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	ContentId    *int64  `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	CreativeId   *int32  `json:"CreativeId,omitempty" xml:"CreativeId,omitempty"`
	CreativeName *string `json:"CreativeName,omitempty" xml:"CreativeName,omitempty"`
	LeadsDetail  *string `json:"LeadsDetail,omitempty" xml:"LeadsDetail,omitempty"`
	SerialId     *int64  `json:"SerialId,omitempty" xml:"SerialId,omitempty"`
	TaskId       *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetLeadsListPageResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetLeadsListPageResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetLeadsListPageResponseBodyDataList) SetComponentId(v string) *GetLeadsListPageResponseBodyDataList {
	s.ComponentId = &v
	return s
}

func (s *GetLeadsListPageResponseBodyDataList) SetContentId(v int64) *GetLeadsListPageResponseBodyDataList {
	s.ContentId = &v
	return s
}

func (s *GetLeadsListPageResponseBodyDataList) SetCreativeId(v int32) *GetLeadsListPageResponseBodyDataList {
	s.CreativeId = &v
	return s
}

func (s *GetLeadsListPageResponseBodyDataList) SetCreativeName(v string) *GetLeadsListPageResponseBodyDataList {
	s.CreativeName = &v
	return s
}

func (s *GetLeadsListPageResponseBodyDataList) SetLeadsDetail(v string) *GetLeadsListPageResponseBodyDataList {
	s.LeadsDetail = &v
	return s
}

func (s *GetLeadsListPageResponseBodyDataList) SetSerialId(v int64) *GetLeadsListPageResponseBodyDataList {
	s.SerialId = &v
	return s
}

func (s *GetLeadsListPageResponseBodyDataList) SetTaskId(v int64) *GetLeadsListPageResponseBodyDataList {
	s.TaskId = &v
	return s
}

type GetLeadsListPageResponseBodyDataPageInfo struct {
	Page        *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalNumber *int32 `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s GetLeadsListPageResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s GetLeadsListPageResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *GetLeadsListPageResponseBodyDataPageInfo) SetPage(v int32) *GetLeadsListPageResponseBodyDataPageInfo {
	s.Page = &v
	return s
}

func (s *GetLeadsListPageResponseBodyDataPageInfo) SetPageSize(v int32) *GetLeadsListPageResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetLeadsListPageResponseBodyDataPageInfo) SetTotalNumber(v int32) *GetLeadsListPageResponseBodyDataPageInfo {
	s.TotalNumber = &v
	return s
}

type GetLeadsListPageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLeadsListPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLeadsListPageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLeadsListPageResponse) GoString() string {
	return s.String()
}

func (s *GetLeadsListPageResponse) SetHeaders(v map[string]*string) *GetLeadsListPageResponse {
	s.Headers = v
	return s
}

func (s *GetLeadsListPageResponse) SetStatusCode(v int32) *GetLeadsListPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLeadsListPageResponse) SetBody(v *GetLeadsListPageResponseBody) *GetLeadsListPageResponse {
	s.Body = v
	return s
}

type GetMainPartPageRequest struct {
	MainId    *int64  `json:"MainId,omitempty" xml:"MainId,omitempty"`
	MainName  *string `json:"MainName,omitempty" xml:"MainName,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetMainPartPageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMainPartPageRequest) GoString() string {
	return s.String()
}

func (s *GetMainPartPageRequest) SetMainId(v int64) *GetMainPartPageRequest {
	s.MainId = &v
	return s
}

func (s *GetMainPartPageRequest) SetMainName(v string) *GetMainPartPageRequest {
	s.MainName = &v
	return s
}

func (s *GetMainPartPageRequest) SetPageIndex(v int32) *GetMainPartPageRequest {
	s.PageIndex = &v
	return s
}

func (s *GetMainPartPageRequest) SetPageSize(v int32) *GetMainPartPageRequest {
	s.PageSize = &v
	return s
}

type GetMainPartPageResponseBody struct {
	Code      *int64                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetMainPartPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMainPartPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMainPartPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetMainPartPageResponseBody) SetCode(v int64) *GetMainPartPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetMainPartPageResponseBody) SetData(v *GetMainPartPageResponseBodyData) *GetMainPartPageResponseBody {
	s.Data = v
	return s
}

func (s *GetMainPartPageResponseBody) SetRequestId(v string) *GetMainPartPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMainPartPageResponseBody) SetSuccess(v bool) *GetMainPartPageResponseBody {
	s.Success = &v
	return s
}

type GetMainPartPageResponseBodyData struct {
	List     []*GetMainPartPageResponseBodyDataList   `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageInfo *GetMainPartPageResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
}

func (s GetMainPartPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMainPartPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMainPartPageResponseBodyData) SetList(v []*GetMainPartPageResponseBodyDataList) *GetMainPartPageResponseBodyData {
	s.List = v
	return s
}

func (s *GetMainPartPageResponseBodyData) SetPageInfo(v *GetMainPartPageResponseBodyDataPageInfo) *GetMainPartPageResponseBodyData {
	s.PageInfo = v
	return s
}

type GetMainPartPageResponseBodyDataList struct {
	AccountTypeList []*GetMainPartPageResponseBodyDataListAccountTypeList `json:"AccountTypeList,omitempty" xml:"AccountTypeList,omitempty" type:"Repeated"`
	Company         *string                                               `json:"Company,omitempty" xml:"Company,omitempty"`
	MainId          *int64                                                `json:"MainId,omitempty" xml:"MainId,omitempty"`
	MainName        *string                                               `json:"MainName,omitempty" xml:"MainName,omitempty"`
}

func (s GetMainPartPageResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetMainPartPageResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetMainPartPageResponseBodyDataList) SetAccountTypeList(v []*GetMainPartPageResponseBodyDataListAccountTypeList) *GetMainPartPageResponseBodyDataList {
	s.AccountTypeList = v
	return s
}

func (s *GetMainPartPageResponseBodyDataList) SetCompany(v string) *GetMainPartPageResponseBodyDataList {
	s.Company = &v
	return s
}

func (s *GetMainPartPageResponseBodyDataList) SetMainId(v int64) *GetMainPartPageResponseBodyDataList {
	s.MainId = &v
	return s
}

func (s *GetMainPartPageResponseBodyDataList) SetMainName(v string) *GetMainPartPageResponseBodyDataList {
	s.MainName = &v
	return s
}

type GetMainPartPageResponseBodyDataListAccountTypeList struct {
	AccountType     *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AccountTypeDesc *string `json:"AccountTypeDesc,omitempty" xml:"AccountTypeDesc,omitempty"`
}

func (s GetMainPartPageResponseBodyDataListAccountTypeList) String() string {
	return tea.Prettify(s)
}

func (s GetMainPartPageResponseBodyDataListAccountTypeList) GoString() string {
	return s.String()
}

func (s *GetMainPartPageResponseBodyDataListAccountTypeList) SetAccountType(v string) *GetMainPartPageResponseBodyDataListAccountTypeList {
	s.AccountType = &v
	return s
}

func (s *GetMainPartPageResponseBodyDataListAccountTypeList) SetAccountTypeDesc(v string) *GetMainPartPageResponseBodyDataListAccountTypeList {
	s.AccountTypeDesc = &v
	return s
}

type GetMainPartPageResponseBodyDataPageInfo struct {
	Page        *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalNumber *int32 `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s GetMainPartPageResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s GetMainPartPageResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *GetMainPartPageResponseBodyDataPageInfo) SetPage(v int32) *GetMainPartPageResponseBodyDataPageInfo {
	s.Page = &v
	return s
}

func (s *GetMainPartPageResponseBodyDataPageInfo) SetPageSize(v int32) *GetMainPartPageResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetMainPartPageResponseBodyDataPageInfo) SetTotalNumber(v int32) *GetMainPartPageResponseBodyDataPageInfo {
	s.TotalNumber = &v
	return s
}

type GetMainPartPageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMainPartPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMainPartPageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMainPartPageResponse) GoString() string {
	return s.String()
}

func (s *GetMainPartPageResponse) SetHeaders(v map[string]*string) *GetMainPartPageResponse {
	s.Headers = v
	return s
}

func (s *GetMainPartPageResponse) SetStatusCode(v int32) *GetMainPartPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMainPartPageResponse) SetBody(v *GetMainPartPageResponseBody) *GetMainPartPageResponse {
	s.Body = v
	return s
}

type GetOssUploadSignatureRequest struct {
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
}

func (s GetOssUploadSignatureRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOssUploadSignatureRequest) GoString() string {
	return s.String()
}

func (s *GetOssUploadSignatureRequest) SetFileName(v string) *GetOssUploadSignatureRequest {
	s.FileName = &v
	return s
}

func (s *GetOssUploadSignatureRequest) SetFileType(v string) *GetOssUploadSignatureRequest {
	s.FileType = &v
	return s
}

type GetOssUploadSignatureResponseBody struct {
	Data      *GetOssUploadSignatureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *int32                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode  *int32                                 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOssUploadSignatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOssUploadSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssUploadSignatureResponseBody) SetData(v *GetOssUploadSignatureResponseBodyData) *GetOssUploadSignatureResponseBody {
	s.Data = v
	return s
}

func (s *GetOssUploadSignatureResponseBody) SetErrorCode(v int32) *GetOssUploadSignatureResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetOssUploadSignatureResponseBody) SetHttpCode(v int32) *GetOssUploadSignatureResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetOssUploadSignatureResponseBody) SetRequestId(v string) *GetOssUploadSignatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssUploadSignatureResponseBody) SetSuccess(v bool) *GetOssUploadSignatureResponseBody {
	s.Success = &v
	return s
}

type GetOssUploadSignatureResponseBodyData struct {
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	Expire      *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host        *string `json:"Host,omitempty" xml:"Host,omitempty"`
	OssKey      *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy      *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature   *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GetOssUploadSignatureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOssUploadSignatureResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOssUploadSignatureResponseBodyData) SetAccessKeyId(v string) *GetOssUploadSignatureResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetOssUploadSignatureResponseBodyData) SetExpire(v string) *GetOssUploadSignatureResponseBodyData {
	s.Expire = &v
	return s
}

func (s *GetOssUploadSignatureResponseBodyData) SetHost(v string) *GetOssUploadSignatureResponseBodyData {
	s.Host = &v
	return s
}

func (s *GetOssUploadSignatureResponseBodyData) SetOssKey(v string) *GetOssUploadSignatureResponseBodyData {
	s.OssKey = &v
	return s
}

func (s *GetOssUploadSignatureResponseBodyData) SetPolicy(v string) *GetOssUploadSignatureResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetOssUploadSignatureResponseBodyData) SetSignature(v string) *GetOssUploadSignatureResponseBodyData {
	s.Signature = &v
	return s
}

type GetOssUploadSignatureResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOssUploadSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOssUploadSignatureResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOssUploadSignatureResponse) GoString() string {
	return s.String()
}

func (s *GetOssUploadSignatureResponse) SetHeaders(v map[string]*string) *GetOssUploadSignatureResponse {
	s.Headers = v
	return s
}

func (s *GetOssUploadSignatureResponse) SetStatusCode(v int32) *GetOssUploadSignatureResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssUploadSignatureResponse) SetBody(v *GetOssUploadSignatureResponseBody) *GetOssUploadSignatureResponse {
	s.Body = v
	return s
}

type GetRelatedByCreativeIdRequest struct {
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetRelatedByCreativeIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRelatedByCreativeIdRequest) GoString() string {
	return s.String()
}

func (s *GetRelatedByCreativeIdRequest) SetId(v int32) *GetRelatedByCreativeIdRequest {
	s.Id = &v
	return s
}

type GetRelatedByCreativeIdResponseBody struct {
	Code         *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         []*GetRelatedByCreativeIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMessage *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRelatedByCreativeIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRelatedByCreativeIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetRelatedByCreativeIdResponseBody) SetCode(v int32) *GetRelatedByCreativeIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetRelatedByCreativeIdResponseBody) SetData(v []*GetRelatedByCreativeIdResponseBodyData) *GetRelatedByCreativeIdResponseBody {
	s.Data = v
	return s
}

func (s *GetRelatedByCreativeIdResponseBody) SetErrorMessage(v string) *GetRelatedByCreativeIdResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetRelatedByCreativeIdResponseBody) SetRequestId(v string) *GetRelatedByCreativeIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRelatedByCreativeIdResponseBody) SetSuccess(v bool) *GetRelatedByCreativeIdResponseBody {
	s.Success = &v
	return s
}

type GetRelatedByCreativeIdResponseBodyData struct {
	ContentId   *int64  `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	ContentName *string `json:"ContentName,omitempty" xml:"ContentName,omitempty"`
	TaskId      *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName    *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s GetRelatedByCreativeIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRelatedByCreativeIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRelatedByCreativeIdResponseBodyData) SetContentId(v int64) *GetRelatedByCreativeIdResponseBodyData {
	s.ContentId = &v
	return s
}

func (s *GetRelatedByCreativeIdResponseBodyData) SetContentName(v string) *GetRelatedByCreativeIdResponseBodyData {
	s.ContentName = &v
	return s
}

func (s *GetRelatedByCreativeIdResponseBodyData) SetTaskId(v int64) *GetRelatedByCreativeIdResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetRelatedByCreativeIdResponseBodyData) SetTaskName(v string) *GetRelatedByCreativeIdResponseBodyData {
	s.TaskName = &v
	return s
}

type GetRelatedByCreativeIdResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRelatedByCreativeIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRelatedByCreativeIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRelatedByCreativeIdResponse) GoString() string {
	return s.String()
}

func (s *GetRelatedByCreativeIdResponse) SetHeaders(v map[string]*string) *GetRelatedByCreativeIdResponse {
	s.Headers = v
	return s
}

func (s *GetRelatedByCreativeIdResponse) SetStatusCode(v int32) *GetRelatedByCreativeIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRelatedByCreativeIdResponse) SetBody(v *GetRelatedByCreativeIdResponseBody) *GetRelatedByCreativeIdResponse {
	s.Body = v
	return s
}

type GetUserFinishedAdRequest struct {
	Adid      *int64  `json:"Adid,omitempty" xml:"Adid,omitempty"`
	Clicklink *string `json:"Clicklink,omitempty" xml:"Clicklink,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Mediaid   *string `json:"Mediaid,omitempty" xml:"Mediaid,omitempty"`
	Tagid     *string `json:"Tagid,omitempty" xml:"Tagid,omitempty"`
	Uid       *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetUserFinishedAdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserFinishedAdRequest) GoString() string {
	return s.String()
}

func (s *GetUserFinishedAdRequest) SetAdid(v int64) *GetUserFinishedAdRequest {
	s.Adid = &v
	return s
}

func (s *GetUserFinishedAdRequest) SetClicklink(v string) *GetUserFinishedAdRequest {
	s.Clicklink = &v
	return s
}

func (s *GetUserFinishedAdRequest) SetId(v string) *GetUserFinishedAdRequest {
	s.Id = &v
	return s
}

func (s *GetUserFinishedAdRequest) SetMediaid(v string) *GetUserFinishedAdRequest {
	s.Mediaid = &v
	return s
}

func (s *GetUserFinishedAdRequest) SetTagid(v string) *GetUserFinishedAdRequest {
	s.Tagid = &v
	return s
}

func (s *GetUserFinishedAdRequest) SetUid(v string) *GetUserFinishedAdRequest {
	s.Uid = &v
	return s
}

type GetUserFinishedAdResponseBody struct {
	ErrorCode *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string                              `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Ext       map[string]*string                   `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Header    *GetUserFinishedAdResponseBodyHeader `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetUserFinishedAdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserFinishedAdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserFinishedAdResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserFinishedAdResponseBody) SetErrorCode(v string) *GetUserFinishedAdResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUserFinishedAdResponseBody) SetErrorMsg(v string) *GetUserFinishedAdResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetUserFinishedAdResponseBody) SetExt(v map[string]*string) *GetUserFinishedAdResponseBody {
	s.Ext = v
	return s
}

func (s *GetUserFinishedAdResponseBody) SetHeader(v *GetUserFinishedAdResponseBodyHeader) *GetUserFinishedAdResponseBody {
	s.Header = v
	return s
}

func (s *GetUserFinishedAdResponseBody) SetRequestId(v string) *GetUserFinishedAdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserFinishedAdResponseBody) SetResult(v *GetUserFinishedAdResponseBodyResult) *GetUserFinishedAdResponseBody {
	s.Result = v
	return s
}

func (s *GetUserFinishedAdResponseBody) SetSuccess(v bool) *GetUserFinishedAdResponseBody {
	s.Success = &v
	return s
}

type GetUserFinishedAdResponseBodyHeader struct {
	CostTime *int64  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	RpcId    *string `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	TraceId  *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	Version  *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetUserFinishedAdResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s GetUserFinishedAdResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *GetUserFinishedAdResponseBodyHeader) SetCostTime(v int64) *GetUserFinishedAdResponseBodyHeader {
	s.CostTime = &v
	return s
}

func (s *GetUserFinishedAdResponseBodyHeader) SetRpcId(v string) *GetUserFinishedAdResponseBodyHeader {
	s.RpcId = &v
	return s
}

func (s *GetUserFinishedAdResponseBodyHeader) SetTraceId(v string) *GetUserFinishedAdResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *GetUserFinishedAdResponseBodyHeader) SetVersion(v string) *GetUserFinishedAdResponseBodyHeader {
	s.Version = &v
	return s
}

type GetUserFinishedAdResponseBodyResult struct {
	Commission    *string `json:"Commission,omitempty" xml:"Commission,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	MarketingType *string `json:"MarketingType,omitempty" xml:"MarketingType,omitempty"`
	Objective     *string `json:"Objective,omitempty" xml:"Objective,omitempty"`
	Success       *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserFinishedAdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetUserFinishedAdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetUserFinishedAdResponseBodyResult) SetCommission(v string) *GetUserFinishedAdResponseBodyResult {
	s.Commission = &v
	return s
}

func (s *GetUserFinishedAdResponseBodyResult) SetId(v string) *GetUserFinishedAdResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetUserFinishedAdResponseBodyResult) SetMarketingType(v string) *GetUserFinishedAdResponseBodyResult {
	s.MarketingType = &v
	return s
}

func (s *GetUserFinishedAdResponseBodyResult) SetObjective(v string) *GetUserFinishedAdResponseBodyResult {
	s.Objective = &v
	return s
}

func (s *GetUserFinishedAdResponseBodyResult) SetSuccess(v bool) *GetUserFinishedAdResponseBodyResult {
	s.Success = &v
	return s
}

type GetUserFinishedAdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserFinishedAdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserFinishedAdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserFinishedAdResponse) GoString() string {
	return s.String()
}

func (s *GetUserFinishedAdResponse) SetHeaders(v map[string]*string) *GetUserFinishedAdResponse {
	s.Headers = v
	return s
}

func (s *GetUserFinishedAdResponse) SetStatusCode(v int32) *GetUserFinishedAdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserFinishedAdResponse) SetBody(v *GetUserFinishedAdResponseBody) *GetUserFinishedAdResponse {
	s.Body = v
	return s
}

type ListAdvertisingRequest struct {
	App      *ListAdvertisingRequestApp    `json:"App,omitempty" xml:"App,omitempty" type:"Struct"`
	Dealtype *int32                        `json:"Dealtype,omitempty" xml:"Dealtype,omitempty"`
	Device   *ListAdvertisingRequestDevice `json:"Device,omitempty" xml:"Device,omitempty" type:"Struct"`
	Ext      map[string]interface{}        `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Id       *string                       `json:"Id,omitempty" xml:"Id,omitempty"`
	Imp      []*ListAdvertisingRequestImp  `json:"Imp,omitempty" xml:"Imp,omitempty" type:"Repeated"`
	Test     *int32                        `json:"Test,omitempty" xml:"Test,omitempty"`
	User     *ListAdvertisingRequestUser   `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s ListAdvertisingRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAdvertisingRequest) GoString() string {
	return s.String()
}

func (s *ListAdvertisingRequest) SetApp(v *ListAdvertisingRequestApp) *ListAdvertisingRequest {
	s.App = v
	return s
}

func (s *ListAdvertisingRequest) SetDealtype(v int32) *ListAdvertisingRequest {
	s.Dealtype = &v
	return s
}

func (s *ListAdvertisingRequest) SetDevice(v *ListAdvertisingRequestDevice) *ListAdvertisingRequest {
	s.Device = v
	return s
}

func (s *ListAdvertisingRequest) SetExt(v map[string]interface{}) *ListAdvertisingRequest {
	s.Ext = v
	return s
}

func (s *ListAdvertisingRequest) SetId(v string) *ListAdvertisingRequest {
	s.Id = &v
	return s
}

func (s *ListAdvertisingRequest) SetImp(v []*ListAdvertisingRequestImp) *ListAdvertisingRequest {
	s.Imp = v
	return s
}

func (s *ListAdvertisingRequest) SetTest(v int32) *ListAdvertisingRequest {
	s.Test = &v
	return s
}

func (s *ListAdvertisingRequest) SetUser(v *ListAdvertisingRequestUser) *ListAdvertisingRequest {
	s.User = v
	return s
}

type ListAdvertisingRequestApp struct {
	Ext     map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Mediaid *string                `json:"Mediaid,omitempty" xml:"Mediaid,omitempty"`
	Sn      *string                `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s ListAdvertisingRequestApp) String() string {
	return tea.Prettify(s)
}

func (s ListAdvertisingRequestApp) GoString() string {
	return s.String()
}

func (s *ListAdvertisingRequestApp) SetExt(v map[string]interface{}) *ListAdvertisingRequestApp {
	s.Ext = v
	return s
}

func (s *ListAdvertisingRequestApp) SetMediaid(v string) *ListAdvertisingRequestApp {
	s.Mediaid = &v
	return s
}

func (s *ListAdvertisingRequestApp) SetSn(v string) *ListAdvertisingRequestApp {
	s.Sn = &v
	return s
}

type ListAdvertisingRequestDevice struct {
	Androidid      *string                          `json:"Androidid,omitempty" xml:"Androidid,omitempty"`
	Androidmd5     *string                          `json:"Androidmd5,omitempty" xml:"Androidmd5,omitempty"`
	Caid           *string                          `json:"Caid,omitempty" xml:"Caid,omitempty"`
	Carrier        *string                          `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	Connectiontype *int32                           `json:"Connectiontype,omitempty" xml:"Connectiontype,omitempty"`
	Devicetype     *int32                           `json:"Devicetype,omitempty" xml:"Devicetype,omitempty"`
	Geo            *ListAdvertisingRequestDeviceGeo `json:"Geo,omitempty" xml:"Geo,omitempty" type:"Struct"`
	Idfa           *string                          `json:"Idfa,omitempty" xml:"Idfa,omitempty"`
	Imei           *string                          `json:"Imei,omitempty" xml:"Imei,omitempty"`
	Imeimd5        *string                          `json:"Imeimd5,omitempty" xml:"Imeimd5,omitempty"`
	Ip             *string                          `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Language       *string                          `json:"Language,omitempty" xml:"Language,omitempty"`
	Mac            *string                          `json:"Mac,omitempty" xml:"Mac,omitempty"`
	Macmd5         *string                          `json:"Macmd5,omitempty" xml:"Macmd5,omitempty"`
	Make           *string                          `json:"Make,omitempty" xml:"Make,omitempty"`
	Model          *string                          `json:"Model,omitempty" xml:"Model,omitempty"`
	Oaid           *string                          `json:"Oaid,omitempty" xml:"Oaid,omitempty"`
	Os             *string                          `json:"Os,omitempty" xml:"Os,omitempty"`
	Osv            *string                          `json:"Osv,omitempty" xml:"Osv,omitempty"`
	Ua             *string                          `json:"Ua,omitempty" xml:"Ua,omitempty"`
	Utdid          *string                          `json:"Utdid,omitempty" xml:"Utdid,omitempty"`
}

func (s ListAdvertisingRequestDevice) String() string {
	return tea.Prettify(s)
}

func (s ListAdvertisingRequestDevice) GoString() string {
	return s.String()
}

func (s *ListAdvertisingRequestDevice) SetAndroidid(v string) *ListAdvertisingRequestDevice {
	s.Androidid = &v
	return s
}

func (s *ListAdvertisingRequestDevice) SetAndroidmd5(v string) *ListAdvertisingRequestDevice {
	s.Androidmd5 = &v
	return s
}

func (s *ListAdvertisingRequestDevice) SetCaid(v string) *ListAdvertisingRequestDevice {
	s.Caid = &v
	return s
}

func (s *ListAdvertisingRequestDevice) SetCarrier(v string) *ListAdvertisingRequestDevice {
	s.Carrier = &v
	return s
}

func (s *ListAdvertisingRequestDevice) SetConnectiontype(v int32) *ListAdvertisingRequestDevice {
	s.Connectiontype = &v
	return s
}

func (s *ListAdvertisingRequestDevice) SetDevicetype(v int32) *ListAdvertisingRequestDevice {
	s.Devicetype = &v
	return s
}

func (s *ListAdvertisingRequestDevice) SetGeo(v *ListAdvertisingRequestDeviceGeo) *ListAdvertisingRequestDevice {
	s.Geo = v
	return s
}

func (s *ListAdvertisingRequestDevice) SetIdfa(v string) *ListAdvertisingRequestDevice {
	s.Idfa = &v
	return s
}

func (s *ListAdvertisingRequestDevice) SetImei(v string) *ListAdvertisingRequestDevice {
	s.Imei = &v
	return s
}

func (s *ListAdvertisingRequestDevice) SetImeimd5(v string) *ListAdvertisingRequestDevice {
	s.Imeimd5 = &v
	return s
}

func (s *ListAdvertisingRequestDevice) SetIp(v string) *ListAdvertisingRequestDevice {
	s.Ip = &v
	return s
}

func (s *ListAdvertisingRequestDevice) SetLanguage(v string) *ListAdvertisingRequestDevice {
	s.Language = &v
	return s
}

func (s *ListAdvertisingRequestDevice) SetMac(v string) *ListAdvertisingRequestDevice {
	s.Mac = &v
	return s
}

func (s *ListAdvertisingRequestDevice) SetMacmd5(v string) *ListAdvertisingRequestDevice {
	s.Macmd5 = &v
	return s
}

func (s *ListAdvertisingRequestDevice) SetMake(v string) *ListAdvertisingRequestDevice {
	s.Make = &v
	return s
}

func (s *ListAdvertisingRequestDevice) SetModel(v string) *ListAdvertisingRequestDevice {
	s.Model = &v
	return s
}

func (s *ListAdvertisingRequestDevice) SetOaid(v string) *ListAdvertisingRequestDevice {
	s.Oaid = &v
	return s
}

func (s *ListAdvertisingRequestDevice) SetOs(v string) *ListAdvertisingRequestDevice {
	s.Os = &v
	return s
}

func (s *ListAdvertisingRequestDevice) SetOsv(v string) *ListAdvertisingRequestDevice {
	s.Osv = &v
	return s
}

func (s *ListAdvertisingRequestDevice) SetUa(v string) *ListAdvertisingRequestDevice {
	s.Ua = &v
	return s
}

func (s *ListAdvertisingRequestDevice) SetUtdid(v string) *ListAdvertisingRequestDevice {
	s.Utdid = &v
	return s
}

type ListAdvertisingRequestDeviceGeo struct {
	City     *string  `json:"City,omitempty" xml:"City,omitempty"`
	District *string  `json:"District,omitempty" xml:"District,omitempty"`
	Lat      *float64 `json:"Lat,omitempty" xml:"Lat,omitempty"`
	Lon      *float64 `json:"Lon,omitempty" xml:"Lon,omitempty"`
	Province *string  `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s ListAdvertisingRequestDeviceGeo) String() string {
	return tea.Prettify(s)
}

func (s ListAdvertisingRequestDeviceGeo) GoString() string {
	return s.String()
}

func (s *ListAdvertisingRequestDeviceGeo) SetCity(v string) *ListAdvertisingRequestDeviceGeo {
	s.City = &v
	return s
}

func (s *ListAdvertisingRequestDeviceGeo) SetDistrict(v string) *ListAdvertisingRequestDeviceGeo {
	s.District = &v
	return s
}

func (s *ListAdvertisingRequestDeviceGeo) SetLat(v float64) *ListAdvertisingRequestDeviceGeo {
	s.Lat = &v
	return s
}

func (s *ListAdvertisingRequestDeviceGeo) SetLon(v float64) *ListAdvertisingRequestDeviceGeo {
	s.Lon = &v
	return s
}

func (s *ListAdvertisingRequestDeviceGeo) SetProvince(v string) *ListAdvertisingRequestDeviceGeo {
	s.Province = &v
	return s
}

type ListAdvertisingRequestImp struct {
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Tagid *string `json:"Tagid,omitempty" xml:"Tagid,omitempty"`
}

func (s ListAdvertisingRequestImp) String() string {
	return tea.Prettify(s)
}

func (s ListAdvertisingRequestImp) GoString() string {
	return s.String()
}

func (s *ListAdvertisingRequestImp) SetId(v string) *ListAdvertisingRequestImp {
	s.Id = &v
	return s
}

func (s *ListAdvertisingRequestImp) SetTagid(v string) *ListAdvertisingRequestImp {
	s.Tagid = &v
	return s
}

type ListAdvertisingRequestUser struct {
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Usertype *string `json:"Usertype,omitempty" xml:"Usertype,omitempty"`
}

func (s ListAdvertisingRequestUser) String() string {
	return tea.Prettify(s)
}

func (s ListAdvertisingRequestUser) GoString() string {
	return s.String()
}

func (s *ListAdvertisingRequestUser) SetId(v string) *ListAdvertisingRequestUser {
	s.Id = &v
	return s
}

func (s *ListAdvertisingRequestUser) SetUsertype(v string) *ListAdvertisingRequestUser {
	s.Usertype = &v
	return s
}

type ListAdvertisingShrinkRequest struct {
	AppShrink    *string `json:"App,omitempty" xml:"App,omitempty"`
	Dealtype     *int32  `json:"Dealtype,omitempty" xml:"Dealtype,omitempty"`
	DeviceShrink *string `json:"Device,omitempty" xml:"Device,omitempty"`
	ExtShrink    *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ImpShrink    *string `json:"Imp,omitempty" xml:"Imp,omitempty"`
	Test         *int32  `json:"Test,omitempty" xml:"Test,omitempty"`
	UserShrink   *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s ListAdvertisingShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAdvertisingShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAdvertisingShrinkRequest) SetAppShrink(v string) *ListAdvertisingShrinkRequest {
	s.AppShrink = &v
	return s
}

func (s *ListAdvertisingShrinkRequest) SetDealtype(v int32) *ListAdvertisingShrinkRequest {
	s.Dealtype = &v
	return s
}

func (s *ListAdvertisingShrinkRequest) SetDeviceShrink(v string) *ListAdvertisingShrinkRequest {
	s.DeviceShrink = &v
	return s
}

func (s *ListAdvertisingShrinkRequest) SetExtShrink(v string) *ListAdvertisingShrinkRequest {
	s.ExtShrink = &v
	return s
}

func (s *ListAdvertisingShrinkRequest) SetId(v string) *ListAdvertisingShrinkRequest {
	s.Id = &v
	return s
}

func (s *ListAdvertisingShrinkRequest) SetImpShrink(v string) *ListAdvertisingShrinkRequest {
	s.ImpShrink = &v
	return s
}

func (s *ListAdvertisingShrinkRequest) SetTest(v int32) *ListAdvertisingShrinkRequest {
	s.Test = &v
	return s
}

func (s *ListAdvertisingShrinkRequest) SetUserShrink(v string) *ListAdvertisingShrinkRequest {
	s.UserShrink = &v
	return s
}

type ListAdvertisingResponseBody struct {
	Errorcode *string                            `json:"Errorcode,omitempty" xml:"Errorcode,omitempty"`
	Errormsg  *string                            `json:"Errormsg,omitempty" xml:"Errormsg,omitempty"`
	Ext       map[string]*string                 `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Header    *ListAdvertisingResponseBodyHeader `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListAdvertisingResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAdvertisingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAdvertisingResponseBody) GoString() string {
	return s.String()
}

func (s *ListAdvertisingResponseBody) SetErrorcode(v string) *ListAdvertisingResponseBody {
	s.Errorcode = &v
	return s
}

func (s *ListAdvertisingResponseBody) SetErrormsg(v string) *ListAdvertisingResponseBody {
	s.Errormsg = &v
	return s
}

func (s *ListAdvertisingResponseBody) SetExt(v map[string]*string) *ListAdvertisingResponseBody {
	s.Ext = v
	return s
}

func (s *ListAdvertisingResponseBody) SetHeader(v *ListAdvertisingResponseBodyHeader) *ListAdvertisingResponseBody {
	s.Header = v
	return s
}

func (s *ListAdvertisingResponseBody) SetRequestId(v string) *ListAdvertisingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAdvertisingResponseBody) SetResult(v *ListAdvertisingResponseBodyResult) *ListAdvertisingResponseBody {
	s.Result = v
	return s
}

func (s *ListAdvertisingResponseBody) SetSuccess(v bool) *ListAdvertisingResponseBody {
	s.Success = &v
	return s
}

type ListAdvertisingResponseBodyHeader struct {
	CostTime *int64  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	RpcId    *string `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	TraceId  *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	Version  *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAdvertisingResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s ListAdvertisingResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *ListAdvertisingResponseBodyHeader) SetCostTime(v int64) *ListAdvertisingResponseBodyHeader {
	s.CostTime = &v
	return s
}

func (s *ListAdvertisingResponseBodyHeader) SetRpcId(v string) *ListAdvertisingResponseBodyHeader {
	s.RpcId = &v
	return s
}

func (s *ListAdvertisingResponseBodyHeader) SetTraceId(v string) *ListAdvertisingResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *ListAdvertisingResponseBodyHeader) SetVersion(v string) *ListAdvertisingResponseBodyHeader {
	s.Version = &v
	return s
}

type ListAdvertisingResponseBodyResult struct {
	Bidid   *string                                     `json:"Bidid,omitempty" xml:"Bidid,omitempty"`
	Id      *string                                     `json:"Id,omitempty" xml:"Id,omitempty"`
	Seatbid []*ListAdvertisingResponseBodyResultSeatbid `json:"Seatbid,omitempty" xml:"Seatbid,omitempty" type:"Repeated"`
}

func (s ListAdvertisingResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAdvertisingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAdvertisingResponseBodyResult) SetBidid(v string) *ListAdvertisingResponseBodyResult {
	s.Bidid = &v
	return s
}

func (s *ListAdvertisingResponseBodyResult) SetId(v string) *ListAdvertisingResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListAdvertisingResponseBodyResult) SetSeatbid(v []*ListAdvertisingResponseBodyResultSeatbid) *ListAdvertisingResponseBodyResult {
	s.Seatbid = v
	return s
}

type ListAdvertisingResponseBodyResultSeatbid struct {
	Bid []*ListAdvertisingResponseBodyResultSeatbidBid `json:"Bid,omitempty" xml:"Bid,omitempty" type:"Repeated"`
}

func (s ListAdvertisingResponseBodyResultSeatbid) String() string {
	return tea.Prettify(s)
}

func (s ListAdvertisingResponseBodyResultSeatbid) GoString() string {
	return s.String()
}

func (s *ListAdvertisingResponseBodyResultSeatbid) SetBid(v []*ListAdvertisingResponseBodyResultSeatbidBid) *ListAdvertisingResponseBodyResultSeatbid {
	s.Bid = v
	return s
}

type ListAdvertisingResponseBodyResultSeatbidBid struct {
	Ads   []*ListAdvertisingResponseBodyResultSeatbidBidAds `json:"Ads,omitempty" xml:"Ads,omitempty" type:"Repeated"`
	Impid *string                                           `json:"Impid,omitempty" xml:"Impid,omitempty"`
}

func (s ListAdvertisingResponseBodyResultSeatbidBid) String() string {
	return tea.Prettify(s)
}

func (s ListAdvertisingResponseBodyResultSeatbidBid) GoString() string {
	return s.String()
}

func (s *ListAdvertisingResponseBodyResultSeatbidBid) SetAds(v []*ListAdvertisingResponseBodyResultSeatbidBidAds) *ListAdvertisingResponseBodyResultSeatbidBid {
	s.Ads = v
	return s
}

func (s *ListAdvertisingResponseBodyResultSeatbidBid) SetImpid(v string) *ListAdvertisingResponseBodyResultSeatbidBid {
	s.Impid = &v
	return s
}

type ListAdvertisingResponseBodyResultSeatbidBidAds struct {
	Crid          *string                                                 `json:"Crid,omitempty" xml:"Crid,omitempty"`
	Crurl         *string                                                 `json:"Crurl,omitempty" xml:"Crurl,omitempty"`
	Icon          *ListAdvertisingResponseBodyResultSeatbidBidAdsIcon     `json:"Icon,omitempty" xml:"Icon,omitempty" type:"Struct"`
	Id            *string                                                 `json:"Id,omitempty" xml:"Id,omitempty"`
	Images        []*ListAdvertisingResponseBodyResultSeatbidBidAdsImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	Interacttype  *int32                                                  `json:"Interacttype,omitempty" xml:"Interacttype,omitempty"`
	Labeltype     *string                                                 `json:"Labeltype,omitempty" xml:"Labeltype,omitempty"`
	Landingurls   []*string                                               `json:"Landingurls,omitempty" xml:"Landingurls,omitempty" type:"Repeated"`
	Marketingtype *string                                                 `json:"Marketingtype,omitempty" xml:"Marketingtype,omitempty"`
	Objective     *string                                                 `json:"Objective,omitempty" xml:"Objective,omitempty"`
	Price         *string                                                 `json:"Price,omitempty" xml:"Price,omitempty"`
	Seat          *string                                                 `json:"Seat,omitempty" xml:"Seat,omitempty"`
	Style         *string                                                 `json:"Style,omitempty" xml:"Style,omitempty"`
	Title         *string                                                 `json:"Title,omitempty" xml:"Title,omitempty"`
	Trackers      *ListAdvertisingResponseBodyResultSeatbidBidAdsTrackers `json:"Trackers,omitempty" xml:"Trackers,omitempty" type:"Struct"`
	Type          *string                                                 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAdvertisingResponseBodyResultSeatbidBidAds) String() string {
	return tea.Prettify(s)
}

func (s ListAdvertisingResponseBodyResultSeatbidBidAds) GoString() string {
	return s.String()
}

func (s *ListAdvertisingResponseBodyResultSeatbidBidAds) SetCrid(v string) *ListAdvertisingResponseBodyResultSeatbidBidAds {
	s.Crid = &v
	return s
}

func (s *ListAdvertisingResponseBodyResultSeatbidBidAds) SetCrurl(v string) *ListAdvertisingResponseBodyResultSeatbidBidAds {
	s.Crurl = &v
	return s
}

func (s *ListAdvertisingResponseBodyResultSeatbidBidAds) SetIcon(v *ListAdvertisingResponseBodyResultSeatbidBidAdsIcon) *ListAdvertisingResponseBodyResultSeatbidBidAds {
	s.Icon = v
	return s
}

func (s *ListAdvertisingResponseBodyResultSeatbidBidAds) SetId(v string) *ListAdvertisingResponseBodyResultSeatbidBidAds {
	s.Id = &v
	return s
}

func (s *ListAdvertisingResponseBodyResultSeatbidBidAds) SetImages(v []*ListAdvertisingResponseBodyResultSeatbidBidAdsImages) *ListAdvertisingResponseBodyResultSeatbidBidAds {
	s.Images = v
	return s
}

func (s *ListAdvertisingResponseBodyResultSeatbidBidAds) SetInteracttype(v int32) *ListAdvertisingResponseBodyResultSeatbidBidAds {
	s.Interacttype = &v
	return s
}

func (s *ListAdvertisingResponseBodyResultSeatbidBidAds) SetLabeltype(v string) *ListAdvertisingResponseBodyResultSeatbidBidAds {
	s.Labeltype = &v
	return s
}

func (s *ListAdvertisingResponseBodyResultSeatbidBidAds) SetLandingurls(v []*string) *ListAdvertisingResponseBodyResultSeatbidBidAds {
	s.Landingurls = v
	return s
}

func (s *ListAdvertisingResponseBodyResultSeatbidBidAds) SetMarketingtype(v string) *ListAdvertisingResponseBodyResultSeatbidBidAds {
	s.Marketingtype = &v
	return s
}

func (s *ListAdvertisingResponseBodyResultSeatbidBidAds) SetObjective(v string) *ListAdvertisingResponseBodyResultSeatbidBidAds {
	s.Objective = &v
	return s
}

func (s *ListAdvertisingResponseBodyResultSeatbidBidAds) SetPrice(v string) *ListAdvertisingResponseBodyResultSeatbidBidAds {
	s.Price = &v
	return s
}

func (s *ListAdvertisingResponseBodyResultSeatbidBidAds) SetSeat(v string) *ListAdvertisingResponseBodyResultSeatbidBidAds {
	s.Seat = &v
	return s
}

func (s *ListAdvertisingResponseBodyResultSeatbidBidAds) SetStyle(v string) *ListAdvertisingResponseBodyResultSeatbidBidAds {
	s.Style = &v
	return s
}

func (s *ListAdvertisingResponseBodyResultSeatbidBidAds) SetTitle(v string) *ListAdvertisingResponseBodyResultSeatbidBidAds {
	s.Title = &v
	return s
}

func (s *ListAdvertisingResponseBodyResultSeatbidBidAds) SetTrackers(v *ListAdvertisingResponseBodyResultSeatbidBidAdsTrackers) *ListAdvertisingResponseBodyResultSeatbidBidAds {
	s.Trackers = v
	return s
}

func (s *ListAdvertisingResponseBodyResultSeatbidBidAds) SetType(v string) *ListAdvertisingResponseBodyResultSeatbidBidAds {
	s.Type = &v
	return s
}

type ListAdvertisingResponseBodyResultSeatbidBidAdsIcon struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListAdvertisingResponseBodyResultSeatbidBidAdsIcon) String() string {
	return tea.Prettify(s)
}

func (s ListAdvertisingResponseBodyResultSeatbidBidAdsIcon) GoString() string {
	return s.String()
}

func (s *ListAdvertisingResponseBodyResultSeatbidBidAdsIcon) SetUrl(v string) *ListAdvertisingResponseBodyResultSeatbidBidAdsIcon {
	s.Url = &v
	return s
}

type ListAdvertisingResponseBodyResultSeatbidBidAdsImages struct {
	Desc   *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	Url    *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListAdvertisingResponseBodyResultSeatbidBidAdsImages) String() string {
	return tea.Prettify(s)
}

func (s ListAdvertisingResponseBodyResultSeatbidBidAdsImages) GoString() string {
	return s.String()
}

func (s *ListAdvertisingResponseBodyResultSeatbidBidAdsImages) SetDesc(v string) *ListAdvertisingResponseBodyResultSeatbidBidAdsImages {
	s.Desc = &v
	return s
}

func (s *ListAdvertisingResponseBodyResultSeatbidBidAdsImages) SetFormat(v string) *ListAdvertisingResponseBodyResultSeatbidBidAdsImages {
	s.Format = &v
	return s
}

func (s *ListAdvertisingResponseBodyResultSeatbidBidAdsImages) SetUrl(v string) *ListAdvertisingResponseBodyResultSeatbidBidAdsImages {
	s.Url = &v
	return s
}

type ListAdvertisingResponseBodyResultSeatbidBidAdsTrackers struct {
	Imps []*string `json:"Imps,omitempty" xml:"Imps,omitempty" type:"Repeated"`
}

func (s ListAdvertisingResponseBodyResultSeatbidBidAdsTrackers) String() string {
	return tea.Prettify(s)
}

func (s ListAdvertisingResponseBodyResultSeatbidBidAdsTrackers) GoString() string {
	return s.String()
}

func (s *ListAdvertisingResponseBodyResultSeatbidBidAdsTrackers) SetImps(v []*string) *ListAdvertisingResponseBodyResultSeatbidBidAdsTrackers {
	s.Imps = v
	return s
}

type ListAdvertisingResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAdvertisingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAdvertisingResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAdvertisingResponse) GoString() string {
	return s.String()
}

func (s *ListAdvertisingResponse) SetHeaders(v map[string]*string) *ListAdvertisingResponse {
	s.Headers = v
	return s
}

func (s *ListAdvertisingResponse) SetStatusCode(v int32) *ListAdvertisingResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAdvertisingResponse) SetBody(v *ListAdvertisingResponseBody) *ListAdvertisingResponse {
	s.Body = v
	return s
}

type QueryAuditResultRequest struct {
	DspId *string   `json:"DspId,omitempty" xml:"DspId,omitempty"`
	Ids   []*string `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s QueryAuditResultRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAuditResultRequest) GoString() string {
	return s.String()
}

func (s *QueryAuditResultRequest) SetDspId(v string) *QueryAuditResultRequest {
	s.DspId = &v
	return s
}

func (s *QueryAuditResultRequest) SetIds(v []*string) *QueryAuditResultRequest {
	s.Ids = v
	return s
}

type QueryAuditResultResponseBody struct {
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	Records   []*QueryAuditResultResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *int32                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Total     *int32                                 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryAuditResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAuditResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAuditResultResponseBody) SetMessage(v string) *QueryAuditResultResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAuditResultResponseBody) SetRecords(v []*QueryAuditResultResponseBodyRecords) *QueryAuditResultResponseBody {
	s.Records = v
	return s
}

func (s *QueryAuditResultResponseBody) SetRequestId(v string) *QueryAuditResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAuditResultResponseBody) SetStatus(v int32) *QueryAuditResultResponseBody {
	s.Status = &v
	return s
}

func (s *QueryAuditResultResponseBody) SetTotal(v int32) *QueryAuditResultResponseBody {
	s.Total = &v
	return s
}

type QueryAuditResultResponseBodyRecords struct {
	Crid   *string `json:"Crid,omitempty" xml:"Crid,omitempty"`
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	State  *int32  `json:"State,omitempty" xml:"State,omitempty"`
}

func (s QueryAuditResultResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s QueryAuditResultResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *QueryAuditResultResponseBodyRecords) SetCrid(v string) *QueryAuditResultResponseBodyRecords {
	s.Crid = &v
	return s
}

func (s *QueryAuditResultResponseBodyRecords) SetReason(v string) *QueryAuditResultResponseBodyRecords {
	s.Reason = &v
	return s
}

func (s *QueryAuditResultResponseBodyRecords) SetState(v int32) *QueryAuditResultResponseBodyRecords {
	s.State = &v
	return s
}

type QueryAuditResultResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryAuditResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAuditResultResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAuditResultResponse) GoString() string {
	return s.String()
}

func (s *QueryAuditResultResponse) SetHeaders(v map[string]*string) *QueryAuditResultResponse {
	s.Headers = v
	return s
}

func (s *QueryAuditResultResponse) SetStatusCode(v int32) *QueryAuditResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAuditResultResponse) SetBody(v *QueryAuditResultResponseBody) *QueryAuditResultResponse {
	s.Body = v
	return s
}

type ReportImpressionRequest struct {
	Impressionlink *string `json:"Impressionlink,omitempty" xml:"Impressionlink,omitempty"`
}

func (s ReportImpressionRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportImpressionRequest) GoString() string {
	return s.String()
}

func (s *ReportImpressionRequest) SetImpressionlink(v string) *ReportImpressionRequest {
	s.Impressionlink = &v
	return s
}

type ReportImpressionResponseBody struct {
	ErrorCode *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string                             `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Ext       map[string]*string                  `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Header    *ReportImpressionResponseBodyHeader `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ReportImpressionResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReportImpressionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportImpressionResponseBody) GoString() string {
	return s.String()
}

func (s *ReportImpressionResponseBody) SetErrorCode(v string) *ReportImpressionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ReportImpressionResponseBody) SetErrorMsg(v string) *ReportImpressionResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ReportImpressionResponseBody) SetExt(v map[string]*string) *ReportImpressionResponseBody {
	s.Ext = v
	return s
}

func (s *ReportImpressionResponseBody) SetHeader(v *ReportImpressionResponseBodyHeader) *ReportImpressionResponseBody {
	s.Header = v
	return s
}

func (s *ReportImpressionResponseBody) SetRequestId(v string) *ReportImpressionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportImpressionResponseBody) SetResult(v *ReportImpressionResponseBodyResult) *ReportImpressionResponseBody {
	s.Result = v
	return s
}

func (s *ReportImpressionResponseBody) SetSuccess(v bool) *ReportImpressionResponseBody {
	s.Success = &v
	return s
}

type ReportImpressionResponseBodyHeader struct {
	CostTime *int64  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	RpcId    *string `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	TraceId  *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	Version  *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ReportImpressionResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s ReportImpressionResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *ReportImpressionResponseBodyHeader) SetCostTime(v int64) *ReportImpressionResponseBodyHeader {
	s.CostTime = &v
	return s
}

func (s *ReportImpressionResponseBodyHeader) SetRpcId(v string) *ReportImpressionResponseBodyHeader {
	s.RpcId = &v
	return s
}

func (s *ReportImpressionResponseBodyHeader) SetTraceId(v string) *ReportImpressionResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *ReportImpressionResponseBodyHeader) SetVersion(v string) *ReportImpressionResponseBodyHeader {
	s.Version = &v
	return s
}

type ReportImpressionResponseBodyResult struct {
	Bidid   *string `json:"Bidid,omitempty" xml:"Bidid,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReportImpressionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ReportImpressionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ReportImpressionResponseBodyResult) SetBidid(v string) *ReportImpressionResponseBodyResult {
	s.Bidid = &v
	return s
}

func (s *ReportImpressionResponseBodyResult) SetSuccess(v bool) *ReportImpressionResponseBodyResult {
	s.Success = &v
	return s
}

type ReportImpressionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReportImpressionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReportImpressionResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportImpressionResponse) GoString() string {
	return s.String()
}

func (s *ReportImpressionResponse) SetHeaders(v map[string]*string) *ReportImpressionResponse {
	s.Headers = v
	return s
}

func (s *ReportImpressionResponse) SetStatusCode(v int32) *ReportImpressionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportImpressionResponse) SetBody(v *ReportImpressionResponseBody) *ReportImpressionResponse {
	s.Body = v
	return s
}

type SendSmsRequest struct {
	NowStamp     *int64  `json:"NowStamp,omitempty" xml:"NowStamp,omitempty"`
	PhoneNumbers *string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty"`
	SignKey      *string `json:"SignKey,omitempty" xml:"SignKey,omitempty"`
}

func (s SendSmsRequest) String() string {
	return tea.Prettify(s)
}

func (s SendSmsRequest) GoString() string {
	return s.String()
}

func (s *SendSmsRequest) SetNowStamp(v int64) *SendSmsRequest {
	s.NowStamp = &v
	return s
}

func (s *SendSmsRequest) SetPhoneNumbers(v string) *SendSmsRequest {
	s.PhoneNumbers = &v
	return s
}

func (s *SendSmsRequest) SetSignKey(v string) *SendSmsRequest {
	s.SignKey = &v
	return s
}

type SendSmsResponseBody struct {
	Data         *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpCode     *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendSmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendSmsResponseBody) GoString() string {
	return s.String()
}

func (s *SendSmsResponseBody) SetData(v bool) *SendSmsResponseBody {
	s.Data = &v
	return s
}

func (s *SendSmsResponseBody) SetErrorCode(v int32) *SendSmsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SendSmsResponseBody) SetErrorMessage(v string) *SendSmsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SendSmsResponseBody) SetHttpCode(v int32) *SendSmsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *SendSmsResponseBody) SetRequestId(v string) *SendSmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendSmsResponseBody) SetSuccess(v bool) *SendSmsResponseBody {
	s.Success = &v
	return s
}

type SendSmsResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendSmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendSmsResponse) String() string {
	return tea.Prettify(s)
}

func (s SendSmsResponse) GoString() string {
	return s.String()
}

func (s *SendSmsResponse) SetHeaders(v map[string]*string) *SendSmsResponse {
	s.Headers = v
	return s
}

func (s *SendSmsResponse) SetStatusCode(v int32) *SendSmsResponse {
	s.StatusCode = &v
	return s
}

func (s *SendSmsResponse) SetBody(v *SendSmsResponseBody) *SendSmsResponse {
	s.Body = v
	return s
}

type SyncInfoRequest struct {
	AccountNo       *string `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	BizId           *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ChainValue      *string `json:"ChainValue,omitempty" xml:"ChainValue,omitempty"`
	ComponentIdList *string `json:"ComponentIdList,omitempty" xml:"ComponentIdList,omitempty"`
	CreateUser      *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	MainId          *int64  `json:"MainId,omitempty" xml:"MainId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NextChainValue  *string `json:"NextChainValue,omitempty" xml:"NextChainValue,omitempty"`
	OssFileUrl      *string `json:"OssFileUrl,omitempty" xml:"OssFileUrl,omitempty"`
	PageId          *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateUser      *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
	Url             *string `json:"Url,omitempty" xml:"Url,omitempty"`
	UrlType         *int32  `json:"UrlType,omitempty" xml:"UrlType,omitempty"`
}

func (s SyncInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncInfoRequest) GoString() string {
	return s.String()
}

func (s *SyncInfoRequest) SetAccountNo(v string) *SyncInfoRequest {
	s.AccountNo = &v
	return s
}

func (s *SyncInfoRequest) SetBizId(v string) *SyncInfoRequest {
	s.BizId = &v
	return s
}

func (s *SyncInfoRequest) SetChainValue(v string) *SyncInfoRequest {
	s.ChainValue = &v
	return s
}

func (s *SyncInfoRequest) SetComponentIdList(v string) *SyncInfoRequest {
	s.ComponentIdList = &v
	return s
}

func (s *SyncInfoRequest) SetCreateUser(v string) *SyncInfoRequest {
	s.CreateUser = &v
	return s
}

func (s *SyncInfoRequest) SetId(v int64) *SyncInfoRequest {
	s.Id = &v
	return s
}

func (s *SyncInfoRequest) SetMainId(v int64) *SyncInfoRequest {
	s.MainId = &v
	return s
}

func (s *SyncInfoRequest) SetName(v string) *SyncInfoRequest {
	s.Name = &v
	return s
}

func (s *SyncInfoRequest) SetNextChainValue(v string) *SyncInfoRequest {
	s.NextChainValue = &v
	return s
}

func (s *SyncInfoRequest) SetOssFileUrl(v string) *SyncInfoRequest {
	s.OssFileUrl = &v
	return s
}

func (s *SyncInfoRequest) SetPageId(v string) *SyncInfoRequest {
	s.PageId = &v
	return s
}

func (s *SyncInfoRequest) SetStatus(v int32) *SyncInfoRequest {
	s.Status = &v
	return s
}

func (s *SyncInfoRequest) SetUpdateUser(v string) *SyncInfoRequest {
	s.UpdateUser = &v
	return s
}

func (s *SyncInfoRequest) SetUrl(v string) *SyncInfoRequest {
	s.Url = &v
	return s
}

func (s *SyncInfoRequest) SetUrlType(v int32) *SyncInfoRequest {
	s.UrlType = &v
	return s
}

type SyncInfoResponseBody struct {
	Code         *int32                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *SyncInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SyncInfoResponseBody) SetCode(v int32) *SyncInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SyncInfoResponseBody) SetData(v *SyncInfoResponseBodyData) *SyncInfoResponseBody {
	s.Data = v
	return s
}

func (s *SyncInfoResponseBody) SetErrorMessage(v string) *SyncInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SyncInfoResponseBody) SetRequestId(v string) *SyncInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncInfoResponseBody) SetSuccess(v bool) *SyncInfoResponseBody {
	s.Success = &v
	return s
}

type SyncInfoResponseBodyData struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SyncInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SyncInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *SyncInfoResponseBodyData) SetId(v int64) *SyncInfoResponseBodyData {
	s.Id = &v
	return s
}

type SyncInfoResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SyncInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncInfoResponse) GoString() string {
	return s.String()
}

func (s *SyncInfoResponse) SetHeaders(v map[string]*string) *SyncInfoResponse {
	s.Headers = v
	return s
}

func (s *SyncInfoResponse) SetStatusCode(v int32) *SyncInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncInfoResponse) SetBody(v *SyncInfoResponseBody) *SyncInfoResponse {
	s.Body = v
	return s
}

type UpdateAdxCreativeContentRequest struct {
	Ad    []*UpdateAdxCreativeContentRequestAd `json:"Ad,omitempty" xml:"Ad,omitempty" type:"Repeated"`
	DspId *string                              `json:"DspId,omitempty" xml:"DspId,omitempty"`
}

func (s UpdateAdxCreativeContentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAdxCreativeContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateAdxCreativeContentRequest) SetAd(v []*UpdateAdxCreativeContentRequestAd) *UpdateAdxCreativeContentRequest {
	s.Ad = v
	return s
}

func (s *UpdateAdxCreativeContentRequest) SetDspId(v string) *UpdateAdxCreativeContentRequest {
	s.DspId = &v
	return s
}

type UpdateAdxCreativeContentRequestAd struct {
	Clicks       []*string                                    `json:"Clicks,omitempty" xml:"Clicks,omitempty" type:"Repeated"`
	Crid         *string                                      `json:"Crid,omitempty" xml:"Crid,omitempty"`
	Enddate      *string                                      `json:"Enddate,omitempty" xml:"Enddate,omitempty"`
	Imps         []*string                                    `json:"Imps,omitempty" xml:"Imps,omitempty" type:"Repeated"`
	Interacttype *int32                                       `json:"Interacttype,omitempty" xml:"Interacttype,omitempty"`
	Nativead     []*UpdateAdxCreativeContentRequestAdNativead `json:"Nativead,omitempty" xml:"Nativead,omitempty" type:"Repeated"`
	Op           *int32                                       `json:"Op,omitempty" xml:"Op,omitempty"`
	Ostype       *string                                      `json:"Ostype,omitempty" xml:"Ostype,omitempty"`
	Startdate    *string                                      `json:"Startdate,omitempty" xml:"Startdate,omitempty"`
	Template     *int32                                       `json:"Template,omitempty" xml:"Template,omitempty"`
	Type         *int32                                       `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateAdxCreativeContentRequestAd) String() string {
	return tea.Prettify(s)
}

func (s UpdateAdxCreativeContentRequestAd) GoString() string {
	return s.String()
}

func (s *UpdateAdxCreativeContentRequestAd) SetClicks(v []*string) *UpdateAdxCreativeContentRequestAd {
	s.Clicks = v
	return s
}

func (s *UpdateAdxCreativeContentRequestAd) SetCrid(v string) *UpdateAdxCreativeContentRequestAd {
	s.Crid = &v
	return s
}

func (s *UpdateAdxCreativeContentRequestAd) SetEnddate(v string) *UpdateAdxCreativeContentRequestAd {
	s.Enddate = &v
	return s
}

func (s *UpdateAdxCreativeContentRequestAd) SetImps(v []*string) *UpdateAdxCreativeContentRequestAd {
	s.Imps = v
	return s
}

func (s *UpdateAdxCreativeContentRequestAd) SetInteracttype(v int32) *UpdateAdxCreativeContentRequestAd {
	s.Interacttype = &v
	return s
}

func (s *UpdateAdxCreativeContentRequestAd) SetNativead(v []*UpdateAdxCreativeContentRequestAdNativead) *UpdateAdxCreativeContentRequestAd {
	s.Nativead = v
	return s
}

func (s *UpdateAdxCreativeContentRequestAd) SetOp(v int32) *UpdateAdxCreativeContentRequestAd {
	s.Op = &v
	return s
}

func (s *UpdateAdxCreativeContentRequestAd) SetOstype(v string) *UpdateAdxCreativeContentRequestAd {
	s.Ostype = &v
	return s
}

func (s *UpdateAdxCreativeContentRequestAd) SetStartdate(v string) *UpdateAdxCreativeContentRequestAd {
	s.Startdate = &v
	return s
}

func (s *UpdateAdxCreativeContentRequestAd) SetTemplate(v int32) *UpdateAdxCreativeContentRequestAd {
	s.Template = &v
	return s
}

func (s *UpdateAdxCreativeContentRequestAd) SetType(v int32) *UpdateAdxCreativeContentRequestAd {
	s.Type = &v
	return s
}

type UpdateAdxCreativeContentRequestAdNativead struct {
	Attrname  *string `json:"Attrname,omitempty" xml:"Attrname,omitempty"`
	Attrvalue *string `json:"Attrvalue,omitempty" xml:"Attrvalue,omitempty"`
	H         *int32  `json:"H,omitempty" xml:"H,omitempty"`
	Mime      *string `json:"Mime,omitempty" xml:"Mime,omitempty"`
	W         *int32  `json:"W,omitempty" xml:"W,omitempty"`
}

func (s UpdateAdxCreativeContentRequestAdNativead) String() string {
	return tea.Prettify(s)
}

func (s UpdateAdxCreativeContentRequestAdNativead) GoString() string {
	return s.String()
}

func (s *UpdateAdxCreativeContentRequestAdNativead) SetAttrname(v string) *UpdateAdxCreativeContentRequestAdNativead {
	s.Attrname = &v
	return s
}

func (s *UpdateAdxCreativeContentRequestAdNativead) SetAttrvalue(v string) *UpdateAdxCreativeContentRequestAdNativead {
	s.Attrvalue = &v
	return s
}

func (s *UpdateAdxCreativeContentRequestAdNativead) SetH(v int32) *UpdateAdxCreativeContentRequestAdNativead {
	s.H = &v
	return s
}

func (s *UpdateAdxCreativeContentRequestAdNativead) SetMime(v string) *UpdateAdxCreativeContentRequestAdNativead {
	s.Mime = &v
	return s
}

func (s *UpdateAdxCreativeContentRequestAdNativead) SetW(v int32) *UpdateAdxCreativeContentRequestAdNativead {
	s.W = &v
	return s
}

type UpdateAdxCreativeContentResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateAdxCreativeContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAdxCreativeContentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAdxCreativeContentResponseBody) SetMessage(v string) *UpdateAdxCreativeContentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAdxCreativeContentResponseBody) SetRequestId(v string) *UpdateAdxCreativeContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAdxCreativeContentResponseBody) SetStatus(v int32) *UpdateAdxCreativeContentResponseBody {
	s.Status = &v
	return s
}

type UpdateAdxCreativeContentResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAdxCreativeContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAdxCreativeContentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAdxCreativeContentResponse) GoString() string {
	return s.String()
}

func (s *UpdateAdxCreativeContentResponse) SetHeaders(v map[string]*string) *UpdateAdxCreativeContentResponse {
	s.Headers = v
	return s
}

func (s *UpdateAdxCreativeContentResponse) SetStatusCode(v int32) *UpdateAdxCreativeContentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAdxCreativeContentResponse) SetBody(v *UpdateAdxCreativeContentResponseBody) *UpdateAdxCreativeContentResponse {
	s.Body = v
	return s
}

type VerifySmsCodeRequest struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	NowStamp     *int64  `json:"NowStamp,omitempty" xml:"NowStamp,omitempty"`
	PhoneNumbers *string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty"`
	SignKey      *string `json:"SignKey,omitempty" xml:"SignKey,omitempty"`
}

func (s VerifySmsCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifySmsCodeRequest) GoString() string {
	return s.String()
}

func (s *VerifySmsCodeRequest) SetCode(v string) *VerifySmsCodeRequest {
	s.Code = &v
	return s
}

func (s *VerifySmsCodeRequest) SetNowStamp(v int64) *VerifySmsCodeRequest {
	s.NowStamp = &v
	return s
}

func (s *VerifySmsCodeRequest) SetPhoneNumbers(v string) *VerifySmsCodeRequest {
	s.PhoneNumbers = &v
	return s
}

func (s *VerifySmsCodeRequest) SetSignKey(v string) *VerifySmsCodeRequest {
	s.SignKey = &v
	return s
}

type VerifySmsCodeResponseBody struct {
	Code         *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VerifySmsCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifySmsCodeResponseBody) GoString() string {
	return s.String()
}

func (s *VerifySmsCodeResponseBody) SetCode(v int32) *VerifySmsCodeResponseBody {
	s.Code = &v
	return s
}

func (s *VerifySmsCodeResponseBody) SetData(v bool) *VerifySmsCodeResponseBody {
	s.Data = &v
	return s
}

func (s *VerifySmsCodeResponseBody) SetErrorMessage(v string) *VerifySmsCodeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *VerifySmsCodeResponseBody) SetRequestId(v string) *VerifySmsCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifySmsCodeResponseBody) SetSuccess(v bool) *VerifySmsCodeResponseBody {
	s.Success = &v
	return s
}

type VerifySmsCodeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VerifySmsCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifySmsCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifySmsCodeResponse) GoString() string {
	return s.String()
}

func (s *VerifySmsCodeResponse) SetHeaders(v map[string]*string) *VerifySmsCodeResponse {
	s.Headers = v
	return s
}

func (s *VerifySmsCodeResponse) SetStatusCode(v int32) *VerifySmsCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifySmsCodeResponse) SetBody(v *VerifySmsCodeResponseBody) *VerifySmsCodeResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("imarketing"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDeviceWithOptions(tmpReq *CreateDeviceRequest, runtime *util.RuntimeOptions) (_result *CreateDeviceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateDeviceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ExtraMap)) {
		request.ExtraMapShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtraMap, tea.String("ExtraMap"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		body["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.City)) {
		body["City"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceModelNumber)) {
		body["DeviceModelNumber"] = request.DeviceModelNumber
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		body["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		body["DeviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.District)) {
		body["District"] = request.District
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraMapShrink)) {
		body["ExtraMap"] = request.ExtraMapShrink
	}

	if !tea.BoolValue(util.IsUnset(request.FirstScene)) {
		body["FirstScene"] = request.FirstScene
	}

	if !tea.BoolValue(util.IsUnset(request.Floor)) {
		body["Floor"] = request.Floor
	}

	if !tea.BoolValue(util.IsUnset(request.LocationName)) {
		body["LocationName"] = request.LocationName
	}

	if !tea.BoolValue(util.IsUnset(request.MediaId)) {
		body["MediaId"] = request.MediaId
	}

	if !tea.BoolValue(util.IsUnset(request.OuterCode)) {
		body["OuterCode"] = request.OuterCode
	}

	if !tea.BoolValue(util.IsUnset(request.Province)) {
		body["Province"] = request.Province
	}

	if !tea.BoolValue(util.IsUnset(request.SecondScene)) {
		body["SecondScene"] = request.SecondScene
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDevice"),
		Version:     tea.String("2022-07-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDevice(request *CreateDeviceRequest) (_result *CreateDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDeviceResponse{}
	_body, _err := client.CreateDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCreativeInfoWithOptions(request *DeleteCreativeInfoRequest, runtime *util.RuntimeOptions) (_result *DeleteCreativeInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountNo)) {
		query["AccountNo"] = request.AccountNo
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.MainId)) {
		query["MainId"] = request.MainId
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateUser)) {
		query["UpdateUser"] = request.UpdateUser
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCreativeInfo"),
		Version:     tea.String("2022-07-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCreativeInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCreativeInfo(request *DeleteCreativeInfoRequest) (_result *DeleteCreativeInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCreativeInfoResponse{}
	_body, _err := client.DeleteCreativeInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBrandPageWithOptions(request *GetBrandPageRequest, runtime *util.RuntimeOptions) (_result *GetBrandPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountNo)) {
		query["AccountNo"] = request.AccountNo
	}

	if !tea.BoolValue(util.IsUnset(request.MainId)) {
		query["MainId"] = request.MainId
	}

	if !tea.BoolValue(util.IsUnset(request.MainName)) {
		query["MainName"] = request.MainName
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBrandPage"),
		Version:     tea.String("2022-07-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBrandPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBrandPage(request *GetBrandPageRequest) (_result *GetBrandPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBrandPageResponse{}
	_body, _err := client.GetBrandPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBusinessIdWithOptions(request *GetBusinessIdRequest, runtime *util.RuntimeOptions) (_result *GetBusinessIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessId)) {
		query["BusinessId"] = request.BusinessId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBusinessId"),
		Version:     tea.String("2022-07-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBusinessIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBusinessId(request *GetBusinessIdRequest) (_result *GetBusinessIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBusinessIdResponse{}
	_body, _err := client.GetBusinessIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCreativeInfoWithOptions(request *GetCreativeInfoRequest, runtime *util.RuntimeOptions) (_result *GetCreativeInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountNo)) {
		query["AccountNo"] = request.AccountNo
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.MainId)) {
		query["MainId"] = request.MainId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCreativeInfo"),
		Version:     tea.String("2022-07-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCreativeInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCreativeInfo(request *GetCreativeInfoRequest) (_result *GetCreativeInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCreativeInfoResponse{}
	_body, _err := client.GetCreativeInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLeadsListPageWithOptions(request *GetLeadsListPageRequest, runtime *util.RuntimeOptions) (_result *GetLeadsListPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComponentId)) {
		query["ComponentId"] = request.ComponentId
	}

	if !tea.BoolValue(util.IsUnset(request.ContentId)) {
		query["ContentId"] = request.ContentId
	}

	if !tea.BoolValue(util.IsUnset(request.CreativeId)) {
		query["CreativeId"] = request.CreativeId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MainId)) {
		query["MainId"] = request.MainId
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLeadsListPage"),
		Version:     tea.String("2022-07-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLeadsListPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLeadsListPage(request *GetLeadsListPageRequest) (_result *GetLeadsListPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLeadsListPageResponse{}
	_body, _err := client.GetLeadsListPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMainPartPageWithOptions(request *GetMainPartPageRequest, runtime *util.RuntimeOptions) (_result *GetMainPartPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MainId)) {
		query["MainId"] = request.MainId
	}

	if !tea.BoolValue(util.IsUnset(request.MainName)) {
		query["MainName"] = request.MainName
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMainPartPage"),
		Version:     tea.String("2022-07-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMainPartPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMainPartPage(request *GetMainPartPageRequest) (_result *GetMainPartPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMainPartPageResponse{}
	_body, _err := client.GetMainPartPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOssUploadSignatureWithOptions(request *GetOssUploadSignatureRequest, runtime *util.RuntimeOptions) (_result *GetOssUploadSignatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOssUploadSignature"),
		Version:     tea.String("2022-07-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOssUploadSignatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOssUploadSignature(request *GetOssUploadSignatureRequest) (_result *GetOssUploadSignatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOssUploadSignatureResponse{}
	_body, _err := client.GetOssUploadSignatureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRelatedByCreativeIdWithOptions(request *GetRelatedByCreativeIdRequest, runtime *util.RuntimeOptions) (_result *GetRelatedByCreativeIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRelatedByCreativeId"),
		Version:     tea.String("2022-07-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRelatedByCreativeIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRelatedByCreativeId(request *GetRelatedByCreativeIdRequest) (_result *GetRelatedByCreativeIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRelatedByCreativeIdResponse{}
	_body, _err := client.GetRelatedByCreativeIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserFinishedAdWithOptions(request *GetUserFinishedAdRequest, runtime *util.RuntimeOptions) (_result *GetUserFinishedAdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserFinishedAd"),
		Version:     tea.String("2022-07-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserFinishedAdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserFinishedAd(request *GetUserFinishedAdRequest) (_result *GetUserFinishedAdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserFinishedAdResponse{}
	_body, _err := client.GetUserFinishedAdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAdvertisingWithOptions(tmpReq *ListAdvertisingRequest, runtime *util.RuntimeOptions) (_result *ListAdvertisingResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListAdvertisingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.App))) {
		request.AppShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.App), tea.String("App"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Device))) {
		request.DeviceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Device), tea.String("Device"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Ext)) {
		request.ExtShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ext, tea.String("Ext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Imp)) {
		request.ImpShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Imp, tea.String("Imp"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.User))) {
		request.UserShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.User), tea.String("User"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAdvertising"),
		Version:     tea.String("2022-07-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAdvertisingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAdvertising(request *ListAdvertisingRequest) (_result *ListAdvertisingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAdvertisingResponse{}
	_body, _err := client.ListAdvertisingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAuditResultWithOptions(request *QueryAuditResultRequest, runtime *util.RuntimeOptions) (_result *QueryAuditResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DspId)) {
		query["DspId"] = request.DspId
	}

	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		query["Ids"] = request.Ids
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAuditResult"),
		Version:     tea.String("2022-07-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAuditResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAuditResult(request *QueryAuditResultRequest) (_result *QueryAuditResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAuditResultResponse{}
	_body, _err := client.QueryAuditResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportImpressionWithOptions(request *ReportImpressionRequest, runtime *util.RuntimeOptions) (_result *ReportImpressionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Impressionlink)) {
		query["Impressionlink"] = request.Impressionlink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportImpression"),
		Version:     tea.String("2022-07-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportImpressionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportImpression(request *ReportImpressionRequest) (_result *ReportImpressionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportImpressionResponse{}
	_body, _err := client.ReportImpressionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendSmsWithOptions(request *SendSmsRequest, runtime *util.RuntimeOptions) (_result *SendSmsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NowStamp)) {
		query["NowStamp"] = request.NowStamp
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumbers)) {
		query["PhoneNumbers"] = request.PhoneNumbers
	}

	if !tea.BoolValue(util.IsUnset(request.SignKey)) {
		query["SignKey"] = request.SignKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendSms"),
		Version:     tea.String("2022-07-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendSmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendSms(request *SendSmsRequest) (_result *SendSmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendSmsResponse{}
	_body, _err := client.SendSmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncInfoWithOptions(request *SyncInfoRequest, runtime *util.RuntimeOptions) (_result *SyncInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountNo)) {
		query["AccountNo"] = request.AccountNo
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.ChainValue)) {
		query["ChainValue"] = request.ChainValue
	}

	if !tea.BoolValue(util.IsUnset(request.ComponentIdList)) {
		query["ComponentIdList"] = request.ComponentIdList
	}

	if !tea.BoolValue(util.IsUnset(request.CreateUser)) {
		query["CreateUser"] = request.CreateUser
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.MainId)) {
		query["MainId"] = request.MainId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NextChainValue)) {
		query["NextChainValue"] = request.NextChainValue
	}

	if !tea.BoolValue(util.IsUnset(request.OssFileUrl)) {
		query["OssFileUrl"] = request.OssFileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.PageId)) {
		query["PageId"] = request.PageId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateUser)) {
		query["UpdateUser"] = request.UpdateUser
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	if !tea.BoolValue(util.IsUnset(request.UrlType)) {
		query["UrlType"] = request.UrlType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncInfo"),
		Version:     tea.String("2022-07-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncInfo(request *SyncInfoRequest) (_result *SyncInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncInfoResponse{}
	_body, _err := client.SyncInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAdxCreativeContentWithOptions(request *UpdateAdxCreativeContentRequest, runtime *util.RuntimeOptions) (_result *UpdateAdxCreativeContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ad)) {
		query["Ad"] = request.Ad
	}

	if !tea.BoolValue(util.IsUnset(request.DspId)) {
		query["DspId"] = request.DspId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAdxCreativeContent"),
		Version:     tea.String("2022-07-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAdxCreativeContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAdxCreativeContent(request *UpdateAdxCreativeContentRequest) (_result *UpdateAdxCreativeContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAdxCreativeContentResponse{}
	_body, _err := client.UpdateAdxCreativeContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifySmsCodeWithOptions(request *VerifySmsCodeRequest, runtime *util.RuntimeOptions) (_result *VerifySmsCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.NowStamp)) {
		query["NowStamp"] = request.NowStamp
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumbers)) {
		query["PhoneNumbers"] = request.PhoneNumbers
	}

	if !tea.BoolValue(util.IsUnset(request.SignKey)) {
		query["SignKey"] = request.SignKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifySmsCode"),
		Version:     tea.String("2022-07-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifySmsCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifySmsCode(request *VerifySmsCodeRequest) (_result *VerifySmsCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifySmsCodeResponse{}
	_body, _err := client.VerifySmsCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
