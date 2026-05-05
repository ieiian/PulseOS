package com.pulseos.feature.activity

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
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
import com.pulseos.data.api.ActivityService
import com.pulseos.data.dto.ActivityTodaySummaryDTO
import com.pulseos.data.dto.ActivityWeekSummaryDTO

@Composable
fun ActivityScreen(service: ActivityService) {
    var today by remember { mutableStateOf<ActivityTodaySummaryDTO?>(null) }
    var week by remember { mutableStateOf<ActivityWeekSummaryDTO?>(null) }
    var isLoading by remember { mutableStateOf(false) }
    var error by remember { mutableStateOf<String?>(null) }

    LaunchedEffect(Unit) {
        isLoading = true
        try {
            today = service.getTodaySummary()
            week = service.getWeekSummary()
        } catch (e: Exception) {
            error = e.message
        } finally {
            isLoading = false
        }
    }

    Column(
        modifier = Modifier
            .padding(16.dp)
            .verticalScroll(rememberScrollState()),
        verticalArrangement = Arrangement.spacedBy(12.dp),
    ) {
        if (today != null) {
            val t = today!!
            ActivityCard("今日步数与积分", "${t.steps} 步 · ${t.cardioPoints} 心肺强化分")
            ActivityCard("周目标进度", "目标 ${t.weeklyGoal} 分，当前完成 ${t.cardioPoints} 分，还差 ${t.remainingPoints} 分")
            if (t.reminder.isNotBlank()) {
                ActivityCard("今日提醒", t.reminder)
            }
            if (t.recentActivities.isNotEmpty()) {
                t.recentActivities.forEach { r ->
                    ActivityCard("运动记录", "${r.activityType} · ${r.minutes} 分钟 · ${r.intensity} · ${r.cardioPoints} 分")
                }
            }
        }

        if (week != null) {
            val w = week!!
            val dayLabels = listOf("周一", "周二", "周三", "周四", "周五", "周六", "周日")
            val trendLines = w.dailyPoints.mapIndexed { i, pts ->
                "${dayLabels.getOrElse(i) { "第${i + 1}天" }}: $pts 分"
            }
            ActivityCard("周趋势", "${trendLines.joinToString("\n")}\n总计 ${w.totalPoints} / ${w.weeklyGoal} 分 · ${w.status}")
            if (w.tips.isNotEmpty()) {
                ActivityCard("本周建议", w.tips.joinToString("\n"))
            }
        }

        if (isLoading && today == null) {
            CircularProgressIndicator()
        }

        if (error != null) {
            ActivityCard("加载失败", error!!)
        }
    }
}

@Composable
private fun ActivityCard(title: String, body: String) {
    Column(
        modifier = Modifier
            .fillMaxWidth()
            .background(MaterialTheme.colorScheme.surface, RoundedCornerShape(20.dp))
            .padding(20.dp),
        verticalArrangement = Arrangement.spacedBy(8.dp),
    ) {
        Text(title, style = MaterialTheme.typography.titleMedium)
        Text(body, style = MaterialTheme.typography.bodyMedium)
    }
}
