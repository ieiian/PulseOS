package com.pulseos.domain.model

data class ActivitySummary(
    val steps: Int,
    val cardioPoints: Int,
    val weeklyGoal: Int,
    val reminder: String,
)

data class ActivityTrend(
    val label: String,
    val points: Int,
)

