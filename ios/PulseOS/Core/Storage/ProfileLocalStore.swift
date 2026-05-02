import Foundation

final class ProfileLocalStore {
    private let defaults = UserDefaults.standard

    func save(_ profile: UserProfile) {
        if let data = try? JSONEncoder().encode(profile) {
            defaults.set(data, forKey: "pulseos.profile")
        }
    }

    func load() -> UserProfile {
        guard
            let data = defaults.data(forKey: "pulseos.profile"),
            let profile = try? JSONDecoder().decode(UserProfile.self, from: data)
        else {
            return UserProfile()
        }

        return profile
    }
}

