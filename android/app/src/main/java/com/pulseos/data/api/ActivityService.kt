package com.pulseos.data.api

import com.google.gson.reflect.TypeToken
import com.pulseos.core.network.ApiClient
import com.pulseos.data.dto.ActivityRecordDTO
import com.pulseos.data.dto.ActivityTodaySummaryDTO
import com.pulseos.data.dto.ActivityWeekSummaryDTO
import com.pulseos.data.dto.ManualActivityRequestDTO

class ActivityService(private val client: ApiClient) {

    private val recordsListType = object : TypeToken<List<ActivityRecordDTO>>() {}
    private val recordType = object : TypeToken<ActivityRecordDTO>() {}
    private val todayType = object : TypeToken<ActivityTodaySummaryDTO>() {}
    private val weekType = object : TypeToken<ActivityWeekSummaryDTO>() {}
    private val requestType = object : TypeToken<ManualActivityRequestDTO>() {}

    suspend fun listRecords(): List<ActivityRecordDTO> =
        client.get("/api/v1/activity/records", recordsListType)

    suspend fun recordManual(request: ManualActivityRequestDTO): ActivityRecordDTO =
        client.post("/api/v1/activity/records", request, requestType, recordType)

    suspend fun getTodaySummary(): ActivityTodaySummaryDTO =
        client.get("/api/v1/activity/today", todayType)

    suspend fun getWeekSummary(): ActivityWeekSummaryDTO =
        client.get("/api/v1/activity/week", weekType)
}
