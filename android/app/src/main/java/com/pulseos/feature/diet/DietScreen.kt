package com.pulseos.feature.diet

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
import com.pulseos.data.api.DietService
import com.pulseos.data.dto.DietRecordDTO
import com.pulseos.data.dto.TodayPlanDTO

@Composable
fun DietScreen(service: DietService) {
    var plan by remember { mutableStateOf<TodayPlanDTO?>(null) }
    var records by remember { mutableStateOf<List<DietRecordDTO>>(emptyList()) }
    var isLoading by remember { mutableStateOf(false) }
    var error by remember { mutableStateOf<String?>(null) }

    LaunchedEffect(Unit) {
        isLoading = true
        try {
            plan = service.getTodayPlan()
            records = service.listRecords()
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
        if (plan != null) {
            val p = plan!!
            DietCard("今日饮食目标", "目标热量 ${p.targetCalories} 千卡 · ${p.fastingPlan.name} ${p.fastingPlan.window}")
            Row(horizontalArrangement = Arrangement.spacedBy(8.dp)) {
                Button(onClick = {}) { Text("拍照分析") }
                Button(onClick = {}) { Text("快速记录") }
            }
            p.options.forEach { option ->
                DietCard(option.title, "${option.description}\n${option.items.joinToString(" / ")}")
            }
            if (p.commonMeals.isNotEmpty()) {
                DietCard("常用饮食", p.commonMeals.joinToString("\n") { "${it.title}: ${it.items.joinToString(" / ")}" })
            }
        } else if (isLoading) {
            CircularProgressIndicator()
        }

        if (records.isNotEmpty()) {
            records.forEach { record ->
                val foods = record.foods.joinToString("、") { "${it.name} ${it.calories}千卡" }
                DietCard("饮食记录", "${record.mealType ?: "未分类"} · ${record.totalCalories} 千卡\n$foods")
            }
        }

        if (error != null) {
            DietCard("加载失败", error!!)
        }
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
