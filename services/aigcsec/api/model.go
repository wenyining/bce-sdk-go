/*
 * Copyright 2017 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

// model.go - definitions of the request arguments and results data structure model for VCR

package api

import "github.com/wenyining/bce-sdk-go/bce"

const (
	URI_PREFIX     = bce.URI_PREFIX
	INPUT_URI      = "/rcs/llm/input/analyze"
	OUTPUT_URI     = "/rcs/llm/output/analyze"
	MULTIMODAL_URI = "/rcs/llm/input/multimodal"
)

// InputAnalyzeRequest 定义输入分析请求的结构
type InputAnalyzeRequest struct {
	Query      string              `json:"query"`
	HistoryQA  []map[string]string `json:"historyQA,omitempty"`
	AppID      string              `json:"appid,omitempty"`
	Stream     bool                `json:"stream,omitempty"`
	TemplateID string              `json:"templateId,omitempty"`
	UserID     string              `json:"userId,omitempty"`
}

// InputAnalyzeResponseData 定义返回数据结构
type InputAnalyzeResponseData struct {
	IsSafe        int                    `json:"isSafe"`
	Action        int                    `json:"action"`
	HitType       string                 `json:"hitType,omitempty"`
	SubHitType    string                 `json:"subHitType,omitempty"`
	LangType      string                 `json:"langType,omitempty"`
	Score         float64                `json:"score,omitempty"`
	HitWord       string                 `json:"hitWord,omitempty"`
	Redline       map[string]interface{} `json:"redline,omitempty"`
	SafeChat      string                 `json:"safeChat,omitempty"`
	DefaultAnswer string                 `json:"defaultAnswer,omitempty"`
}

// InputAnalyzeResponse 定义输入分析响应的结构
type InputAnalyzeResponse struct {
	RequestID string                   `json:"request_id"`
	RetCode   string                   `json:"ret_code"`
	RetMsg    string                   `json:"ret_msg"`
	RetData   InputAnalyzeResponseData `json:"ret_data"`
}
