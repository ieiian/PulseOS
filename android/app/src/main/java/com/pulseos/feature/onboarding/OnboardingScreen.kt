package com.pulseos.feature.onboarding

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Button
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import com.pulseos.domain.model.UserProfile

@Composable
fun OnboardingScreen(
    initialProfile: UserProfile,
    onComplete: (UserProfile) -> Unit,
) {
    val profileState = remember { mutableStateOf(initialProfile) }

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
            onClick = { onComplete(profileState.value) },
            modifier = Modifier.fillMaxWidth(),
        ) {
            Text("完成设置")
        }
    }
}

