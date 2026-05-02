package com.pulseos.domain.model

data class UserProfile(
    val name: String = "",
    val age: String = "",
    val gender: String = "",
    val heightCm: String = "",
    val weightKg: String = "",
    val primaryGoal: String = "maintain",
)

