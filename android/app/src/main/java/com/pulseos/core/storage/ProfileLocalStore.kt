package com.pulseos.core.storage

import android.content.Context
import com.pulseos.domain.model.UserProfile

class ProfileLocalStore(context: Context) {
    private val prefs = context.getSharedPreferences("pulseos_profile", Context.MODE_PRIVATE)

    fun save(profile: UserProfile) {
        prefs.edit()
            .putString("name", profile.name)
            .putString("age", profile.age)
            .putString("gender", profile.gender)
            .putString("height_cm", profile.heightCm)
            .putString("weight_kg", profile.weightKg)
            .putString("primary_goal", profile.primaryGoal)
            .apply()
    }

    fun load(): UserProfile = UserProfile(
        name = prefs.getString("name", "") ?: "",
        age = prefs.getString("age", "") ?: "",
        gender = prefs.getString("gender", "") ?: "",
        heightCm = prefs.getString("height_cm", "") ?: "",
        weightKg = prefs.getString("weight_kg", "") ?: "",
        primaryGoal = prefs.getString("primary_goal", "maintain") ?: "maintain",
    )
}

