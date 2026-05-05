package com.pulseos.feature.meditation

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.Button
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
import com.pulseos.data.api.MeditationService
import com.pulseos.data.dto.BreathModeDTO
import com.pulseos.data.dto.MeditationTodaySummaryDTO

@Composable
fun MeditationScreen(service: MeditationService) {
    var summary by remember { mutableStateOf<MeditationTodaySummaryDTO?>(null) }
    var isLoading by remember { mutableStateOf(false) }
    var error by remember { mutableStateOf<String?>(null) }
    var selectedMode by remember { mutableStateOf<BreathModeDTO?>(null) }
    var isPlaying by remember { mutableStateOf(false) }

    LaunchedEffect(Unit) {
        isLoading = true
        try {
            val result = service.getTodaySummary()
            summary = result
            if (result.modes.isNotEmpty() && selectedMode == null) {
                selectedMode = result.modes.first()
            }
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
        if (summary != null) {
            val s = summary!!
            MeditationCard("今日冥想", "已完成 ${s.totalDurationS / 60} 分钟，${s.completedCount} 次训练")

            if (s.recentSessions.isNotEmpty()) {
                s.recentSessions.forEach { session ->
                    MeditationCard("冥想记录", "${session.modeKey} · ${session.durationS / 60} 分钟")
                }
            }
        }

        if (selectedMode != null) {
            val mode = selectedMode!!
            MeditationCard(
                mode.title,
                "吸气 ${mode.inhaleSec}s · 停顿 ${mode.holdSec}s · 呼气 ${mode.exhaleSec}s\n${mode.description}",
            )
        }

        Row(horizontalArrangement = Arrangement.spacedBy(8.dp)) {
            Button(onClick = { isPlaying = !isPlaying }) {
                Text(if (isPlaying) "暂停音频" else "播放音频")
            }
            Button(onClick = {}) {
                Text("开始呼吸")
            }
        }

        summary?.modes?.forEach { mode ->
            Button(
                onClick = { selectedMode = mode },
                modifier = Modifier.fillMaxWidth(),
            ) {
                Text(mode.title)
            }
        }

        if (isLoading && summary == null) {
            CircularProgressIndicator()
        }

        if (error != null) {
            MeditationCard("加载失败", error!!)
        }
    }
}

@Composable
private fun MeditationCard(title: String, body: String) {
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
