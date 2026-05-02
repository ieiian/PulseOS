package com.pulseos.feature.sleep

class SleepMonitorController {
    private var recording = false

    fun isRecording(): Boolean = recording

    fun start() {
        recording = true
    }

    fun stop() {
        recording = false
    }
}

