package com.pulseos.data.dto

import com.google.gson.annotations.SerializedName

data class ProfileDTO(
    val id: Long? = null,
    val name: String,
    val age: Int,
    val gender: String,
    @SerializedName("height_cm") val heightCm: Int,
    @SerializedName("weight_kg") val weightKg: Int,
    @SerializedName("primary_goal") val primaryGoal: String,
    @SerializedName("secondary_goals") val secondaryGoals: List<String>,
    @SerializedName("health_flags") val healthFlags: List<String>,
)

data class SettingsDTO(
    @SerializedName("notifications_enabled") val notificationsEnabled: Boolean,
    @SerializedName("step_permission_granted") val stepPermissionGranted: Boolean,
    @SerializedName("microphone_permission_granted") val microphonePermissionGranted: Boolean,
    @SerializedName("sleep_reminder_enabled") val sleepReminderEnabled: Boolean,
)

data class StatsDTO(
    @SerializedName("current_streak") val currentStreak: Int,
    @SerializedName("days_tracked") val daysTracked: Int,
)
