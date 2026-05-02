package com.pulseos.feature.home

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import com.pulseos.domain.model.HomeDashboard
import com.pulseos.domain.model.UserProfile

@Composable
fun HomeScreen(profile: UserProfile) {
    val dashboard = HomeDashboard(
        totalScore = 67,
        actionItem = "优先完成今天的运动积分和一次饮食记录。",
        dietSummary = "目标热量 1900 千卡，已配置轻断食窗口。",
        activitySummary = "今日 5620 步，心肺强化 75 分。",
        sleepSummary = "昨夜评分 55，时长 430 分钟。",
        meditationNote = "今日冥想 10 分钟。",
        trends = listOf(58, 63, 61, 70, 68, 74, 67),
    )

    Column(
        modifier = Modifier
            .fillMaxSize()
            .padding(16.dp),
        verticalArrangement = Arrangement.spacedBy(12.dp),
    ) {
        HomeCard(title = "今日概览", body = "${profile.name.ifBlank { "你" }}，${dashboard.actionItem}")
        HomeCard(title = "今日总分", body = "${dashboard.totalScore} 分")
        HomeCard(title = "饮食", body = dashboard.dietSummary)
        HomeCard(title = "运动", body = dashboard.activitySummary)
        HomeCard(title = "睡眠", body = dashboard.sleepSummary)
        HomeCard(title = "冥想", body = dashboard.meditationNote)
        HomeCard(title = "趋势", body = dashboard.trends.joinToString(" / "))
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
