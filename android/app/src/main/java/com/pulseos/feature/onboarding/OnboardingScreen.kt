package com.pulseos.feature.onboarding

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Button
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import com.pulseos.app.AppContainer
import com.pulseos.data.dto.ProfileDTO
import com.pulseos.domain.model.UserProfile
import kotlinx.coroutines.launch

@Composable
fun OnboardingScreen(
    initialProfile: UserProfile,
    container: AppContainer,
    onComplete: (UserProfile) -> Unit,
) {
    val profileState = remember { mutableStateOf(initialProfile) }
    var isSubmitting by remember { mutableStateOf(false) }
    var error by remember { mutableStateOf<String?>(null) }
    val scope = rememberCoroutineScope()

    Column(
        modifier = Modifier
            .fillMaxSize()
            .padding(24.dp),
        verticalArrangement = Arrangement.spacedBy(16.dp),
    ) {
        Text("欢迎来到 PulseOS")
        Text("先完成基础资料和健康目标设置。")

        OutlinedTextField(
            modifier = Modifier.fillMaxWidth(),
            value = profileState.value.name,
            onValueChange = { profileState.value = profileState.value.copy(name = it) },
            label = { Text("姓名") },
        )
        OutlinedTextField(
            modifier = Modifier.fillMaxWidth(),
            value = profileState.value.age,
            onValueChange = { profileState.value = profileState.value.copy(age = it) },
            label = { Text("年龄") },
        )
        OutlinedTextField(
            modifier = Modifier.fillMaxWidth(),
            value = profileState.value.gender,
            onValueChange = { profileState.value = profileState.value.copy(gender = it) },
            label = { Text("性别") },
        )
        OutlinedTextField(
            modifier = Modifier.fillMaxWidth(),
            value = profileState.value.heightCm,
            onValueChange = { profileState.value = profileState.value.copy(heightCm = it) },
            label = { Text("身高（cm）") },
        )
        OutlinedTextField(
            modifier = Modifier.fillMaxWidth(),
            value = profileState.value.weightKg,
            onValueChange = { profileState.value = profileState.value.copy(weightKg = it) },
            label = { Text("体重（kg）") },
        )
        GoalPicker(
            selected = profileState.value.primaryGoal,
            onSelected = { profileState.value = profileState.value.copy(primaryGoal = it) },
        )
        Button(
            onClick = {
                val p = profileState.value
                val dto = ProfileDTO(
                    name = p.name,
                    age = p.age.toIntOrNull() ?: 0,
                    gender = p.gender,
                    heightCm = p.heightCm.toIntOrNull() ?: 0,
                    weightKg = p.weightKg.toIntOrNull() ?: 0,
                    primaryGoal = p.primaryGoal,
                    secondaryGoals = emptyList(),
                    healthFlags = emptyList(),
                )
                isSubmitting = true
                error = null
                scope.launch {
                    try {
                        container.user.onboard(dto)
                        onComplete(p)
                    } catch (e: Exception) {
                        error = e.message
                    } finally {
                        isSubmitting = false
                    }
                }
            },
            modifier = Modifier.fillMaxWidth(),
            enabled = !isSubmitting,
        ) {
            Text("完成设置")
        }

        if (isSubmitting) {
            CircularProgressIndicator()
        }
        if (error != null) {
            Text("提交失败: ${error!!}", color = androidx.compose.ui.graphics.Color.Red)
        }
    }
}
