package com.pulseos

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import com.pulseos.app.PulseApplication
import com.pulseos.navigation.PulseNavHost
import com.pulseos.ui.theme.PulseTheme

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        val container = (application as PulseApplication).container
        setContent {
            PulseTheme {
                PulseNavHost(container)
            }
        }
    }
}

