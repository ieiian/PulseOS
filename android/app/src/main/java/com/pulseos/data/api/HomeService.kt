package com.pulseos.data.api

import com.google.gson.reflect.TypeToken
import com.pulseos.core.network.ApiClient
import com.pulseos.data.dto.DashboardDTO

class HomeService(private val client: ApiClient) {

    private val dashboardType = object : TypeToken<DashboardDTO>() {}

    suspend fun getDashboard(): DashboardDTO =
        client.get("/api/v1/home/dashboard", dashboardType)
}
