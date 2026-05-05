import SwiftUI

@MainActor
final class HomeDashboardViewModel: ObservableObject {
    @Published var dashboard: DashboardDTO?
    @Published var isLoading = false
    @Published var errorMessage: String?

    private let service: HomeService

    init(service: HomeService) {
        self.service = service
    }

    func loadDashboard() async {
        isLoading = true
        errorMessage = nil
        do {
            dashboard = try await service.getDashboard()
        } catch {
            errorMessage = error.localizedDescription
        }
        isLoading = false
    }
}

struct HomeDashboardView: View {
    let profile: UserProfile
    @StateObject private var viewModel: HomeDashboardViewModel

    init(profile: UserProfile, service: HomeService = APIContainer.shared.home) {
        self.profile = profile
        _viewModel = StateObject(wrappedValue: HomeDashboardViewModel(service: service))
    }

    var body: some View {
        ScrollView {
            VStack(spacing: 12) {
                if let dashboard = viewModel.dashboard {
                    HomeCard(title: "今日概览", content: "\(profile.name.isEmpty ? "你" : profile.name)，\(dashboard.actionItem)")
                    HomeCard(title: "今日总分", content: "\(dashboard.today.totalScore) 分")
                    HomeCard(title: "饮食", content: dashboard.dietSummary)
                    HomeCard(title: "运动", content: dashboard.activitySummary)
                    HomeCard(title: "睡眠", content: dashboard.sleepSummary)
                    HomeCard(title: "冥想", content: dashboard.meditationNote)
                    HomeCard(title: "趋势", content: dashboard.trends.map(String.init).joined(separator: " / "))
                } else if viewModel.isLoading {
                    ProgressView("加载中...")
                } else if let error = viewModel.errorMessage {
                    HomeCard(title: "错误", content: error)
                }
            }
            .padding(16)
        }
        .background(PulseTheme.background)
        .task {
            await viewModel.loadDashboard()
        }
    }
}


private struct HomeCard: View {
    let title: String
    let content: String

    var body: some View {
        VStack(alignment: .leading, spacing: 8) {
            Text(title)
                .font(.headline)
            Text(content)
                .foregroundStyle(.secondary)
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .padding(20)
        .background(PulseTheme.card)
        .clipShape(RoundedRectangle(cornerRadius: 20, style: .continuous))
    }
}
