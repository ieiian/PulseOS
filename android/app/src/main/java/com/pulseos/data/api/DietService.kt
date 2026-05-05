package com.pulseos.data.api

import com.google.gson.reflect.TypeToken
import com.pulseos.core.network.ApiClient
import com.pulseos.data.dto.AnalyzeRequestDTO
import com.pulseos.data.dto.AnalyzeResultDTO
import com.pulseos.data.dto.DietRecordDTO
import com.pulseos.data.dto.PhotoUploadDTO
import com.pulseos.data.dto.TodayPlanDTO

class DietService(private val client: ApiClient) {

    private val planType = object : TypeToken<TodayPlanDTO>() {}
    private val uploadType = object : TypeToken<PhotoUploadDTO>() {}
    private val analyzeType = object : TypeToken<AnalyzeResultDTO>() {}
    private val recordsListType = object : TypeToken<List<DietRecordDTO>>() {}
    private val requestType = object : TypeToken<AnalyzeRequestDTO>() {}
    private val recordType = object : TypeToken<DietRecordDTO>() {}

    suspend fun getTodayPlan(): TodayPlanDTO =
        client.get("/api/v1/diet/plan/today", planType)

    suspend fun uploadPhoto(filename: String): PhotoUploadDTO =
        client.postEmpty("/api/v1/diet/photo-upload?filename=${filename}", uploadType)

    suspend fun analyze(request: AnalyzeRequestDTO): AnalyzeResultDTO =
        client.post("/api/v1/diet/analyze", request, requestType, analyzeType)

    suspend fun listRecords(): List<DietRecordDTO> =
        client.get("/api/v1/diet/records", recordsListType)

    suspend fun quickRecord(request: AnalyzeRequestDTO): DietRecordDTO =
        client.post("/api/v1/diet/records", request, requestType, recordType)
}
