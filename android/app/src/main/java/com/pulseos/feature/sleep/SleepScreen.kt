package com.pulseos.feature.sleep

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
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
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import com.pulseos.data.api.SleepService
import com.pulseos.data.dto.SleepSessionDTO
import com.pulseos.data.dto.SleepTodaySummaryDTO
import kotlinx.coroutines.launch

@Composable
fun SleepScreen(service: SleepService) {
    var today by remember { mutableStateOf<SleepTodaySummaryDTO?>(null) }
    var activeSession by remember { mutableStateOf<SleepSessionDTO?>(null) }
    var isLoading by remember { mutableStateOf(false) }
    var error by remember { mutableStateOf<String?>(null) }
    val scope = rememberCoroutineScope()

    LaunchedEffect(Unit) {
        isLoading = true
        try {
            val result = service.getToday()
            today = result
            if (result.session.status == "in_progress") {
                activeSession = result.session
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
        if (today != null) {
            val t = today!!
            val session = t.session

            if (session.status == "completed") {
                SleepCard("昨夜概览", "${session.durationM} 分钟 · 评分 ${session.score}")
                if (session.advice.isNotBlank()) {
                    SleepCard("睡前建议", session.advice)
                }
            }

            if (t.events.isNotEmpty()) {
                t.events.forEach { event ->
                    SleepCard("事件时间轴", "${event.timestamp} · ${event.type} · ${event.level}")
                }
            }
        }

        val isTracking = activeSession != null
        Button(
            onClick = {
                scope.launch {
                    try {
                        if (isTracking) {
                            val session = activeSession ?: return@launch
                            val result = service.endSession(session.id)
                            activeSession = null
                            today = today?.copy(session = result) ?: SleepTodaySummaryDTO(result, emptyList())
                        } else {
                            val result = service.startSession()
                            activeSession = result
                        }
                    } catch (e: Exception) {
                        error = e.message
                    }
                }
            },
            modifier = Modifier.fillMaxWidth(),
        ) {
            Text(if (isTracking) "结束监测" else "开始睡眠")
        }

        if (isLoading && today == null) {
            CircularProgressIndicator()
        }

        if (error != null) {
            SleepCard("错误", error!!)
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
