package com.pulseos.data.dto

import com.google.gson.annotations.SerializedName

data class MeditationTodaySummaryDTO(
    @SerializedName("total_duration_s") val totalDurationS: Int,
    @SerializedName("completed_count") val completedCount: Int,
    @SerializedName("recent_sessions") val recentSessions: List<MeditationSessionDTO>,
    val modes: List<BreathModeDTO>,
)

data class MeditationSessionDTO(
    val id: String,
    @SerializedName("mode_key") val modeKey: String,
    @SerializedName("duration_s") val durationS: Int,
    @SerializedName("audio_key") val audioKey: String,
)

data class BreathModeDTO(
    val key: String,
    val title: String,
    @SerializedName("inhale_sec") val inhaleSec: Int,
    @SerializedName("hold_sec") val holdSec: Int,
    @SerializedName("exhale_sec") val exhaleSec: Int,
    val description: String,
)

data class MeditationSessionRequestDTO(
    @SerializedName("mode_key") val modeKey: String,
    @SerializedName("duration_s") val durationS: Int,
    @SerializedName("audio_key") val audioKey: String,
)
