package com.pulseos.feature.diet

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.AssistChip
import androidx.compose.material3.Button
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.unit.dp
import com.pulseos.domain.model.DietPlanOption
import com.pulseos.domain.model.DietStatus

@Composable
fun DietScreen() {
    val status = DietStatus(
        recommendation = "caution",
        summary = "识别到鸡胸肉沙拉和糙米饭，总热量约 480 千卡。",
        explanation = "适合正餐，但要注意当天总热量和晚间加餐。",
    )

    val options = listOf(
        DietPlanOption("高蛋白午餐", "控制主食份量，优先蛋白质和蔬菜。", listOf("鸡胸肉", "西兰花", "糙米饭")),
        DietPlanOption("常用早餐", "快速记录模板", listOf("无糖酸奶", "鸡蛋", "香蕉")),
    )

    Column(
        modifier = Modifier.padding(16.dp),
        verticalArrangement = Arrangement.spacedBy(12.dp),
    ) {
        DietCard("今日饮食目标", "目标热量 1900 千卡，轻断食窗口 12:00 - 20:00")
        RecommendationCard(status)
        Row(horizontalArrangement = Arrangement.spacedBy(8.dp)) {
            Button(onClick = {}) { Text("拍照分析") }
            Button(onClick = {}) { Text("快速记录") }
        }
        options.forEach { option ->
            DietCard(option.title, "${option.description}\n${option.items.joinToString(" / ")}")
        }
    }
}

@Composable
private fun RecommendationCard(status: DietStatus) {
    val (label, bg) = when (status.recommendation) {
        "recommended" -> "建议食用" to Color(0xFFDFF3E3)
        "not_recommended" -> "不建议食用" to Color(0xFFFBE3D8)
        "forbidden" -> "禁止食用" to Color(0xFFF7D4D8)
        else -> "注意食用" to Color(0xFFFCEFC7)
    }

    Column(
        modifier = Modifier
            .fillMaxWidth()
            .background(bg, RoundedCornerShape(20.dp))
            .padding(20.dp),
        verticalArrangement = Arrangement.spacedBy(8.dp),
    ) {
        AssistChip(onClick = {}, label = { Text(label) })
        Text(status.summary, style = MaterialTheme.typography.titleMedium)
        Text(status.explanation, style = MaterialTheme.typography.bodyMedium)
    }
}

@Composable
private fun DietCard(title: String, body: String) {
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

