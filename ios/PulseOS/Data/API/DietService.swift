import Foundation

actor DietService {
    private let client: APIClient

    init(client: APIClient) {
        self.client = client
    }

    func getTodayPlan() async throws -> TodayPlanDTO {
        try await client.get("/api/v1/diet/plan/today", as: TodayPlanDTO.self)
    }

    func uploadPhoto(filename: String) async throws -> PhotoUploadDTO {
        try await client.postEmpty("/api/v1/diet/photo-upload?filename=\(filename.addingPercentEncoding(withAllowedCharacters: .urlQueryAllowed) ?? filename)")
    }

    func analyze(_ request: AnalyzeRequestDTO) async throws -> AnalyzeResultDTO {
        try await client.post("/api/v1/diet/analyze", body: request)
    }

    func listRecords() async throws -> [DietRecordDTO] {
        try await client.get("/api/v1/diet/records", as: [DietRecordDTO].self)
    }

    func quickRecord(_ request: AnalyzeRequestDTO) async throws -> DietRecordDTO {
        try await client.post("/api/v1/diet/records", body: request)
    }
}
