package com.pulseos.domain.model

data class DietPlanOption(
    val title: String,
    val description: String,
    val items: List<String>,
)

data class DietStatus(
    val recommendation: String,
    val summary: String,
    val explanation: String,
)

