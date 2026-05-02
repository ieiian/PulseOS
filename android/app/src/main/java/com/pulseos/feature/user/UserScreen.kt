package com.pulseos.feature.user

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Divider
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import com.pulseos.domain.model.UserProfile

@Composable
fun UserScreen(profile: UserProfile) {
    Column(
        modifier = Modifier
            .fillMaxSize()
            .padding(16.dp),
        verticalArrangement = Arrangement.spacedBy(16.dp),
    ) {
        Text("用户", style = MaterialTheme.typography.headlineSmall)
        Section("个人信息", "${profile.name.ifBlank { "未设置姓名" }} · ${profile.age.ifBlank { "-" }} 岁")
        Section("健康目标", goalLabel(profile.primaryGoal))
        Section("身体指标", "${profile.heightCm.ifBlank { "-" }} cm / ${profile.weightKg.ifBlank { "-" }} kg")
        Section("统计数据", "连续记录 0 天，后续接后端统计接口")
        Divider()
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

