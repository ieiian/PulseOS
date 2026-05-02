package com.pulseos.core.config

data class AppConfig(
    val appName: String,
    val apiBaseUrl: String,
) {
    companion object {
        fun default(): AppConfig = AppConfig(
            appName = "PulseOS",
            apiBaseUrl = "http://10.0.2.2:8080",
        )
    }
}

