package com.pulseos.feature.sleep

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Button
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import com.pulseos.domain.model.SleepEvent
import com.pulseos.domain.model.SleepSummary

@Composable
fun SleepScreen() {
    val controller = remember { SleepMonitorController() }
    var isRecording by remember { mutableStateOf(controller.isRecording()) }
    val summary = SleepSummary(
        durationMinutes = 430,
        score = 55,
        advice = "昨夜呼噜/梦话事件偏多，建议连续观察几晚并避免睡前饮酒。",
        isRecording = isRecording,
    )
    val events = listOf(
        SleepEvent("snore", "01:12", "medium"),
        SleepEvent("talk", "03:46", "low"),
        SleepEvent("snore", "05:21", "high"),
    )

    Column(
        modifier = Modifier.padding(16.dp),
        verticalArrangement = Arrangement.spacedBy(12.dp),
    ) {
        SleepCard("昨夜概览", "${summary.durationMinutes} 分钟 · 评分 ${summary.score}")
        SleepCard("睡前建议", summary.advice)
        Button(
            onClick = {
                if (isRecording) controller.stop() else controller.start()
                isRecording = controller.isRecording()
            },
            modifier = Modifier.fillMaxWidth(),
        ) {
            Text(if (isRecording) "结束监测" else "开始睡眠")
        }
        events.forEach { event ->
            SleepCard("事件时间轴", "${event.timestamp} · ${event.type} · ${event.level}")
        }
    }
}

@Composable
private fun SleepCard(title: String, body: String) {
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

