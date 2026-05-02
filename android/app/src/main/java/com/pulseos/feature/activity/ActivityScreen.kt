package com.pulseos.feature.activity

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Button
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import com.pulseos.domain.model.ActivitySummary
import com.pulseos.domain.model.ActivityTrend

@Composable
fun ActivityScreen(stepPreview: Int) {
    val summary = ActivitySummary(
        steps = stepPreview,
        cardioPoints = 75,
        weeklyGoal = 150,
        reminder = "本周还差 75 分，优先安排 30 分钟快走或 20 分钟骑行。",
    )
    val trends = listOf(
        ActivityTrend("Mon", 12),
        ActivityTrend("Tue", 18),
        ActivityTrend("Wed", 0),
        ActivityTrend("Thu", 35),
        ActivityTrend("Fri", 10),
        ActivityTrend("Sat", 0),
        ActivityTrend("Sun", 75),
    )

    Column(
        modifier = Modifier.padding(16.dp),
        verticalArrangement = Arrangement.spacedBy(12.dp),
    ) {
        ActivityCard("今日步数与积分", "${summary.steps} 步 · ${summary.cardioPoints} 心肺强化分")
        ActivityCard("周目标进度", "目标 ${summary.weeklyGoal} 分，当前完成 ${summary.cardioPoints} 分")
        ActivityCard("今日提醒", summary.reminder)
        Row(horizontalArrangement = Arrangement.spacedBy(8.dp)) {
            Button(onClick = {}) { Text("手动补录") }
            Button(onClick = {}) { Text("查看周趋势") }
        }
        trends.forEach { trend ->
            ActivityCard("${trend.label} 趋势", "${trend.points} 分")
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

