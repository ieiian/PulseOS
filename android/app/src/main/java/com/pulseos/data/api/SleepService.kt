package com.pulseos.data.api

import com.google.gson.reflect.TypeToken
import com.pulseos.core.network.ApiClient
import com.pulseos.data.dto.SleepEndRequestDTO
import com.pulseos.data.dto.SleepSessionDTO
import com.pulseos.data.dto.SleepStartRequestDTO
import com.pulseos.data.dto.SleepTodaySummaryDTO

class SleepService(private val client: ApiClient) {

    private val sessionType = object : TypeToken<SleepSessionDTO>() {}
    private val startRequestType = object : TypeToken<SleepStartRequestDTO>() {}
    private val endRequestType = object : TypeToken<SleepEndRequestDTO>() {}
    private val todayType = object : TypeToken<SleepTodaySummaryDTO>() {}

    suspend fun startSession(audioURL: String = ""): SleepSessionDTO =
        client.post("/api/v1/sleep/sessions/start", SleepStartRequestDTO(audioURL), startRequestType, sessionType)

    suspend fun endSession(sessionID: String): SleepSessionDTO =
        client.post("/api/v1/sleep/sessions/end", SleepEndRequestDTO(sessionID), endRequestType, sessionType)

    suspend fun getToday(): SleepTodaySummaryDTO =
        client.get("/api/v1/sleep/today", todayType)
}
