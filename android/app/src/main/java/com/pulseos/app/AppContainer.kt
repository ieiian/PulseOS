package com.pulseos.app

import com.pulseos.core.config.AppConfig
import com.pulseos.core.network.ApiClient
import com.pulseos.data.api.ActivityService
import com.pulseos.data.api.DietService
import com.pulseos.data.api.HomeService
import com.pulseos.data.api.MeditationService
import com.pulseos.data.api.SleepService
import com.pulseos.data.api.UserService

class AppContainer(
    val config: AppConfig = AppConfig.default(),
) {
    val client: ApiClient = ApiClient(config.apiBaseUrl)
    val home: HomeService = HomeService(client)
    val diet: DietService = DietService(client)
    val activity: ActivityService = ActivityService(client)
    val meditation: MeditationService = MeditationService(client)
    val sleep: SleepService = SleepService(client)
    val user: UserService = UserService(client)
}
