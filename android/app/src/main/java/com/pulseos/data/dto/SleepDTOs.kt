package com.pulseos.data.dto

import com.google.gson.annotations.SerializedName

data class SleepTodaySummaryDTO(
    val session: SleepSessionDTO,
    val events: List<SleepEventDTO>,
)

data class SleepSessionDTO(
    val id: String,
    val status: String,
    @SerializedName("started_at") val startedAt: String,
    @SerializedName("ended_at") val endedAt: String?,
    @SerializedName("duration_m") val durationM: Int,
    val score: Int,
    @SerializedName("audio_url") val audioURL: String,
    val advice: String,
)

data class SleepEventDTO(
    val type: String,
    val timestamp: String,
    val level: String,
)

data class SleepStartRequestDTO(
    @SerializedName("audio_url") val audioURL: String,
)

data class SleepEndRequestDTO(
    @SerializedName("session_id") val sessionID: String,
)
