import Foundation

actor SleepService {
    private let client: APIClient

    init(client: APIClient) {
        self.client = client
    }

    func startSession(audioURL: String = "") async throws -> SleepSessionDTO {
        try await client.post("/api/v1/sleep/sessions/start", body: SleepStartRequestDTO(audioURL: audioURL))
    }

    func endSession(sessionID: String = "") async throws -> SleepSessionDTO {
        try await client.post("/api/v1/sleep/sessions/end", body: SleepEndRequestDTO(sessionID: sessionID))
    }

    func getToday() async throws -> SleepTodaySummaryDTO {
        try await client.get("/api/v1/sleep/today", as: SleepTodaySummaryDTO.self)
    }
}
