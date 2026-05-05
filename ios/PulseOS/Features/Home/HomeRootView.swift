import SwiftUI

struct HomeRootView: View {
    @State private var profile = ProfileLocalStore().load()

    var body: some View {
        if profile.name.isEmpty {
            OnboardingView(profile: profile) { updatedProfile in
                profile = updatedProfile
                ProfileLocalStore().save(updatedProfile)
            }
        } else {
            TabView {
                HomeDashboardView(profile: profile)
                    .tabItem {
                        Label("首页", systemImage: "house")
                    }

                DietView()
                    .tabItem {
                        Label("饮食", systemImage: "fork.knife")
                    }

                ActivityView()
                    .tabItem {
                        Label("运动", systemImage: "figure.walk")
                    }

                MeditationView()
                    .tabItem {
                        Label("冥想", systemImage: "wind")
                    }

                SleepView()
                    .tabItem {
                        Label("睡眠", systemImage: "moon.zzz")
                    }

                UserView(profile: profile)
                    .tabItem {
                        Label("用户", systemImage: "person")
                    }
            }
        }
    }
}

#if DEBUG
struct HomeRootView_Previews: PreviewProvider {
    static var previews: some View {
        HomeRootView()
    }
}
#endif
