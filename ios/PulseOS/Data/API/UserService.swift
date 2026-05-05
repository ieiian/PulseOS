import Foundation

actor UserService {
    private let client: APIClient

    init(client: APIClient) {
        self.client = client
    }

    func onboard(profile: ProfileDTO) async throws -> ProfileDTO {
        try await client.post("/api/v1/users/onboarding", body: profile)
    }

    func getProfile() async throws -> ProfileDTO {
        try await client.get("/api/v1/users/profile", as: ProfileDTO.self)
    }

    func updateProfile(_ profile: ProfileDTO) async throws -> ProfileDTO {
        try await client.put("/api/v1/users/profile", body: profile)
    }

    func getSettings() async throws -> SettingsDTO {
        try await client.get("/api/v1/users/settings", as: SettingsDTO.self)
    }

    func updateSettings(_ settings: SettingsDTO) async throws -> SettingsDTO {
        try await client.put("/api/v1/users/settings", body: settings)
    }

    func getStats() async throws -> StatsDTO {
        try await client.get("/api/v1/users/stats", as: StatsDTO.self)
    }
}
