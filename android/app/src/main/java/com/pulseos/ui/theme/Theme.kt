package com.pulseos.ui.theme

import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.darkColorScheme
import androidx.compose.material3.lightColorScheme
import androidx.compose.runtime.Composable
import com.pulseos.core.designsystem.PulseColors

private val LightColors = lightColorScheme(
    primary = PulseColors.Brand,
    surface = PulseColors.Surface,
)

private val DarkColors = darkColorScheme(
    primary = PulseColors.Brand,
)

@Composable
fun PulseTheme(
    darkTheme: Boolean = false,
    content: @Composable () -> Unit,
) {
    MaterialTheme(
        colorScheme = if (darkTheme) DarkColors else LightColors,
        content = content,
    )
}

