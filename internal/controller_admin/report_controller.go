/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package controller_admin

import (
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/report_admin"
	"github.com/answerdev/answer/pkg/converter"
	"github.com/gin-gonic/gin"
)

// ReportController report controller
type ReportController struct {
	reportService *report_admin.ReportAdminService
}

// NewReportController new controller
func NewReportController(reportService *report_admin.ReportAdminService) *ReportController {
	return &ReportController{reportService: reportService}
}

// ListReportPage godoc
// @Summary list report page
// @Description list report records
// @Security ApiKeyAuth
// @Tags admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param status query string true "status" Enums(pending, completed)
// @Param object_type query string true "object_type" Enums(all, question,answer,comment)
// @Param page query int false "page size"
// @Param page_size query int false "page size"
// @Success 200 {object} handler.RespBody
// @Router /answer/admin/api/reports/page [get]
func (rc *ReportController) ListReportPage(ctx *gin.Context) {
	var (
		objectType = ctx.Query("object_type")
		status     = ctx.Query("status")
		page       = converter.StringToInt(ctx.DefaultQuery("page", "1"))
		pageSize   = converter.StringToInt(ctx.DefaultQuery("page_size", "20"))
	)

	dto := schema.GetReportListPageDTO{
		ObjectType: objectType,
		Status:     status,
		Page:       page,
		PageSize:   pageSize,
	}

	resp, err := rc.reportService.ListReportPage(ctx, dto)
	if err != nil {
		handler.HandleResponse(ctx, err, schema.ErrTypeModal)
	} else {
		handler.HandleResponse(ctx, err, resp)
	}
}

// Handle godoc
// @Summary handle flag
// @Description handle flag
// @Security ApiKeyAuth
// @Tags admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.ReportHandleReq true "flag"
// @Success 200 {object} handler.RespBody
// @Router /answer/admin/api/report/ [put]
func (rc *ReportController) Handle(ctx *gin.Context) {
	req := schema.ReportHandleReq{}
	if handler.BindAndCheck(ctx, &req) {
		return
	}

	err := rc.reportService.HandleReported(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}
