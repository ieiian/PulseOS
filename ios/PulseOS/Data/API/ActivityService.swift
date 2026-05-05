import Foundation

actor ActivityService {
    private let client: APIClient

    init(client: APIClient) {
        self.client = client
    }

    func listRecords() async throws -> [ActivityRecordDTO] {
        try await client.get("/api/v1/activity/records", as: [ActivityRecordDTO].self)
    }

    func recordManual(_ request: ManualActivityRequestDTO) async throws -> ActivityRecordDTO {
        try await client.post("/api/v1/activity/records", body: request)
    }

    func getTodaySummary() async throws -> ActivityTodaySummaryDTO {
        try await client.get("/api/v1/activity/today", as: ActivityTodaySummaryDTO.self)
    }

    func getWeekSummary() async throws -> ActivityWeekSummaryDTO {
        try await client.get("/api/v1/activity/week", as: ActivityWeekSummaryDTO.self)
    }
}
