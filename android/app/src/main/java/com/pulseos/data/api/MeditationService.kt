package com.pulseos.data.api

import com.google.gson.reflect.TypeToken
import com.pulseos.core.network.ApiClient
import com.pulseos.data.dto.MeditationSessionDTO
import com.pulseos.data.dto.MeditationSessionRequestDTO
import com.pulseos.data.dto.MeditationTodaySummaryDTO

class MeditationService(private val client: ApiClient) {

    private val sessionType = object : TypeToken<MeditationSessionDTO>() {}
    private val requestType = object : TypeToken<MeditationSessionRequestDTO>() {}
    private val todayType = object : TypeToken<MeditationTodaySummaryDTO>() {}

    suspend fun recordSession(request: MeditationSessionRequestDTO): MeditationSessionDTO =
        client.post("/api/v1/meditation/sessions", request, requestType, sessionType)

    suspend fun getTodaySummary(): MeditationTodaySummaryDTO =
        client.get("/api/v1/meditation/today", todayType)
}
