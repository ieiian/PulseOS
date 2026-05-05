package com.pulseos.data.dto

import com.google.gson.annotations.SerializedName

data class DashboardDTO(
    val today: DailyScoreDTO,
    @SerializedName("action_item") val actionItem: String,
    @SerializedName("diet_summary") val dietSummary: String,
    @SerializedName("activity_summary") val activitySummary: String,
    @SerializedName("sleep_summary") val sleepSummary: String,
    @SerializedName("meditation_note") val meditationNote: String,
    val trends: List<Int>,
)

data class DailyScoreDTO(
    val date: String,
    @SerializedName("diet_score") val dietScore: Int,
    @SerializedName("activity_score") val activityScore: Int,
    @SerializedName("sleep_score") val sleepScore: Int,
    @SerializedName("total_score") val totalScore: Int,
)
