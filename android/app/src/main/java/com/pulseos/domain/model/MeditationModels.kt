package com.pulseos.domain.model

data class BreathMode(
    val key: String,
    val title: String,
    val inhaleSec: Int,
    val holdSec: Int,
    val exhaleSec: Int,
    val description: String,
)

data class MeditationSummary(
    val totalMinutes: Int,
    val sessionCount: Int,
)

