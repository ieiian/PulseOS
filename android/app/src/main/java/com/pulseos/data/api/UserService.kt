package com.pulseos.data.api

import com.google.gson.reflect.TypeToken
import com.pulseos.core.network.ApiClient
import com.pulseos.data.dto.ProfileDTO
import com.pulseos.data.dto.SettingsDTO
import com.pulseos.data.dto.StatsDTO

class UserService(private val client: ApiClient) {

    private val profileType = object : TypeToken<ProfileDTO>() {}
    private val settingsType = object : TypeToken<SettingsDTO>() {}
    private val statsType = object : TypeToken<StatsDTO>() {}

    suspend fun onboard(profile: ProfileDTO): ProfileDTO =
        client.post("/api/v1/users/onboarding", profile, profileType, profileType)

    suspend fun getProfile(): ProfileDTO =
        client.get("/api/v1/users/profile", profileType)

    suspend fun updateProfile(profile: ProfileDTO): ProfileDTO =
        client.put("/api/v1/users/profile", profile, profileType, profileType)

    suspend fun getSettings(): SettingsDTO =
        client.get("/api/v1/users/settings", settingsType)

    suspend fun updateSettings(settings: SettingsDTO): SettingsDTO =
        client.put("/api/v1/users/settings", settings, settingsType, settingsType)

    suspend fun getStats(): StatsDTO =
        client.get("/api/v1/users/stats", statsType)
}
