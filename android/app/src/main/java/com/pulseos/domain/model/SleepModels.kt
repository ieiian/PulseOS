package com.pulseos.domain.model

data class SleepEvent(
    val type: String,
    val timestamp: String,
    val level: String,
)

data class SleepSummary(
    val durationMinutes: Int,
    val score: Int,
    val advice: String,
    val isRecording: Boolean,
)

