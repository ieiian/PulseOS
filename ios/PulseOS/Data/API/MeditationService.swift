import Foundation

actor MeditationService {
    private let client: APIClient

    init(client: APIClient) {
        self.client = client
    }

    func recordSession(_ request: MeditationSessionRequestDTO) async throws -> MeditationSessionDTO {
        try await client.post("/api/v1/meditation/sessions", body: request)
    }

    func getTodaySummary() async throws -> MeditationTodaySummaryDTO {
        try await client.get("/api/v1/meditation/today", as: MeditationTodaySummaryDTO.self)
    }
}
