package com.pulseos.navigation

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.NavigationBar
import androidx.compose.material3.NavigationBarItem
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import com.pulseos.core.storage.ProfileLocalStore
import com.pulseos.feature.activity.ActivityScreen
import com.pulseos.feature.activity.AndroidStepReader
import com.pulseos.feature.diet.DietScreen
import com.pulseos.feature.home.HomeScreen
import com.pulseos.feature.meditation.MeditationScreen
import com.pulseos.feature.onboarding.OnboardingScreen
import com.pulseos.feature.sleep.SleepScreen
import com.pulseos.feature.user.UserScreen

@Composable
fun PulseNavHost() {
    val context = LocalContext.current
    val store = remember { ProfileLocalStore(context) }
    val stepReader = remember { AndroidStepReader(context) }
    var profile by remember { mutableStateOf(store.load()) }
    var tab by remember { mutableStateOf("home") }
    val onboardingComplete = profile.name.isNotBlank()

    if (!onboardingComplete) {
        OnboardingScreen(
            initialProfile = profile,
            onComplete = {
                profile = it
                store.save(it)
            },
        )
        return
    }

    Scaffold(
        modifier = Modifier.fillMaxSize(),
        bottomBar = {
            NavigationBar {
                listOf(
                    "home" to "首页",
                    "diet" to "饮食",
                    "activity" to "运动",
                    "meditation" to "冥想",
                    "sleep" to "睡眠",
                    "user" to "用户",
                ).forEach { (value, label) ->
                    NavigationBarItem(
                        selected = tab == value,
                        onClick = { tab = value },
                        icon = { Text("•") },
                        label = { Text(label) },
                    )
                }
            }
        },
    ) { paddingValues ->
        Box(modifier = Modifier.padding(paddingValues)) {
            when (tab) {
                "activity" -> ActivityScreen(stepPreview = stepReader.getPreviewSteps())
                "diet" -> DietScreen()
                "meditation" -> MeditationScreen()
                "sleep" -> SleepScreen()
                "user" -> UserScreen(profile)
                else -> HomeScreen(profile)
            }
        }
    }
}
