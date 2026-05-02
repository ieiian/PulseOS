package com.pulseos.domain.model

data class HomeDashboard(
    val totalScore: Int,
    val actionItem: String,
    val dietSummary: String,
    val activitySummary: String,
    val sleepSummary: String,
    val meditationNote: String,
    val trends: List<Int>,
)

