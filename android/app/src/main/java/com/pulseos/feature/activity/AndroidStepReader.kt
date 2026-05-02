package com.pulseos.feature.activity

import android.content.Context
import android.hardware.Sensor
import android.hardware.SensorManager

class AndroidStepReader(context: Context) {
    private val sensorManager = context.getSystemService(Context.SENSOR_SERVICE) as? SensorManager

    fun isAvailable(): Boolean {
        return sensorManager?.getDefaultSensor(Sensor.TYPE_STEP_COUNTER) != null
    }

    fun getPreviewSteps(): Int {
        return if (isAvailable()) 5620 else 4200
    }
}

