package com.pulseos.feature.meditation

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
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.getValue
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import com.pulseos.domain.model.BreathMode

@Composable
fun MeditationScreen() {
    val controller = remember { MeditationAudioController() }
    var selectedMode by remember { mutableStateOf(defaultModes().first()) }
    var isPlaying by remember { mutableStateOf(controller.isPlaying()) }

    Column(
        modifier = Modifier.padding(16.dp),
        verticalArrangement = Arrangement.spacedBy(12.dp),
    ) {
        MeditationCard("今日冥想", "已完成 10 分钟，1 次训练")
        MeditationCard(
            selectedMode.title,
            "吸气 ${selectedMode.inhaleSec}s · 停顿 ${selectedMode.holdSec}s · 呼气 ${selectedMode.exhaleSec}s\n${selectedMode.description}",
        )
        Row(horizontalArrangement = Arrangement.spacedBy(8.dp)) {
            Button(onClick = {
                controller.toggle()
                isPlaying = controller.isPlaying()
            }) {
                Text(if (isPlaying) "暂停音频" else "播放音频")
            }
            Button(onClick = {}) {
                Text("开始呼吸")
            }
        }
        defaultModes().forEach { mode ->
            Button(onClick = { selectedMode = mode }, modifier = Modifier.fillMaxWidth()) {
                Text(mode.title)
            }
        }
    }
}

private fun defaultModes(): List<BreathMode> = listOf(
    BreathMode("calm", "平静呼吸", 4, 2, 6, "适合放慢节奏，降低紧绷感。"),
    BreathMode("focus", "专注呼吸", 4, 4, 4, "适合进入工作或学习状态前。"),
    BreathMode("sleep", "睡前呼吸", 4, 7, 8, "适合睡前稳定呼吸节律。"),
)

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

