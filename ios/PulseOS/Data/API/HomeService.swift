import Foundation

actor HomeService {
    private let client: APIClient

    init(client: APIClient) {
        self.client = client
    }

    func getDashboard() async throws -> DashboardDTO {
        try await client.get("/api/v1/home/dashboard", as: DashboardDTO.self)
    }
}
