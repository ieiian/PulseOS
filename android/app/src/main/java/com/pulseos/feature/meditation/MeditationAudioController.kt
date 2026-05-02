package com.pulseos.feature.meditation

class MeditationAudioController {
    private var playing: Boolean = false

    fun isPlaying(): Boolean = playing

    fun toggle() {
        playing = !playing
    }
}

