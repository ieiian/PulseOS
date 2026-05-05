package com.pulseos.feature.home

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import com.pulseos.data.api.HomeService
import com.pulseos.data.dto.DashboardDTO
import com.pulseos.domain.model.UserProfile

@Composable
fun HomeScreen(profile: UserProfile, service: HomeService) {
    var dashboard by remember { mutableStateOf<DashboardDTO?>(null) }
    var isLoading by remember { mutableStateOf(false) }
    var error by remember { mutableStateOf<String?>(null) }

    LaunchedEffect(Unit) {
        isLoading = true
        try {
            dashboard = service.getDashboard()
        } catch (e: Exception) {
            error = e.message
        } finally {
            isLoading = false
        }
    }

    Column(
        modifier = Modifier
            .fillMaxSize()
            .verticalScroll(rememberScrollState())
            .padding(16.dp),
        verticalArrangement = Arrangement.spacedBy(12.dp),
    ) {
        if (dashboard != null) {
            val d = dashboard!!
            HomeCard(title = "今日概览", body = "${profile.name.ifBlank { "你" }}，${d.actionItem}")
            HomeCard(title = "今日总分", body = "${d.today.totalScore} 分")
            HomeCard(title = "饮食", body = d.dietSummary)
            HomeCard(title = "运动", body = d.activitySummary)
            HomeCard(title = "睡眠", body = d.sleepSummary)
            HomeCard(title = "冥想", body = d.meditationNote)
            HomeCard(title = "趋势", body = d.trends.joinToString(" / "))
        } else if (isLoading) {
            CircularProgressIndicator()
        }

        if (error != null) {
            HomeCard(title = "错误", body = error!!)
        }
    }
}

@Composable
private fun HomeCard(title: String, body: String) {
    Column(
        modifier = Modifier
            .fillMaxWidth()
            .background(
                color = MaterialTheme.colorScheme.surface,
                shape = RoundedCornerShape(20.dp),
            )
            .padding(20.dp),
        verticalArrangement = Arrangement.spacedBy(8.dp),
    ) {
        Text(title, style = MaterialTheme.typography.titleMedium)
        Text(body, style = MaterialTheme.typography.bodyMedium)
    }
}
