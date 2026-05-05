package com.pulseos.feature.user

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.HorizontalDivider
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
import com.pulseos.data.api.UserService
import com.pulseos.data.dto.StatsDTO
import com.pulseos.domain.model.UserProfile

@Composable
fun UserScreen(profile: UserProfile, service: UserService) {
    var stats by remember { mutableStateOf<StatsDTO?>(null) }
    var isLoading by remember { mutableStateOf(false) }
    var error by remember { mutableStateOf<String?>(null) }

    LaunchedEffect(Unit) {
        isLoading = true
        try {
            stats = service.getStats()
        } catch (e: Exception) {
            error = e.message
        } finally {
            isLoading = false
        }
    }

    Column(
        modifier = Modifier
            .fillMaxSize()
            .padding(16.dp)
            .verticalScroll(rememberScrollState()),
        verticalArrangement = Arrangement.spacedBy(16.dp),
    ) {
        Text("用户", style = MaterialTheme.typography.headlineSmall)
        Section("个人信息", "${profile.name.ifBlank { "未设置姓名" }} · ${profile.age.ifBlank { "-" }} 岁")
        Section("健康目标", goalLabel(profile.primaryGoal))
        Section("身体指标", "${profile.heightCm.ifBlank { "-" }} cm / ${profile.weightKg.ifBlank { "-" }} kg")

        if (stats != null) {
            val s = stats!!
            Section("统计数据", "连续记录 ${s.currentStreak} 天 · 共记录 ${s.daysTracked} 天")
        } else if (isLoading) {
            CircularProgressIndicator()
        }

        if (error != null) {
            Section("加载统计失败", error!!)
        }

        HorizontalDivider()
        Section("设置", "通知、权限、设备管理结构已预留")
    }
}

@Composable
private fun Section(title: String, body: String) {
    Column(
        modifier = Modifier.fillMaxWidth(),
        verticalArrangement = Arrangement.spacedBy(6.dp),
    ) {
        Text(title, style = MaterialTheme.typography.titleMedium)
        Text(body, style = MaterialTheme.typography.bodyMedium)
    }
}

private fun goalLabel(goal: String): String = when (goal) {
    "fat_loss" -> "减脂"
    "muscle_gain" -> "增肌"
    "sleep_recovery" -> "作息恢复"
    "stress_relief" -> "减压放松"
    else -> "维持健康"
}
