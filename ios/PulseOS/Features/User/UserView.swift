import SwiftUI

@MainActor
final class UserViewModel: ObservableObject {
    @Published var profile: ProfileDTO?
    @Published var stats: StatsDTO?
    @Published var isLoading = false
    @Published var errorMessage: String?

    private let service: UserService

    init(service: UserService) {
        self.service = service
    }

    func loadProfile() async {
        isLoading = true
        defer { isLoading = false }
        do {
            profile = try await service.getProfile()
            stats = try await service.getStats()
        } catch {
            errorMessage = error.localizedDescription
        }
    }
}

struct UserView: View {
    let profile: UserProfile
    @StateObject private var viewModel: UserViewModel

    init(profile: UserProfile, service: UserService = APIContainer.shared.user) {
        self.profile = profile
        _viewModel = StateObject(wrappedValue: UserViewModel(service: service))
    }

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 16) {
                Text("用户")
                    .font(.title2.weight(.semibold))

                if let apiProfile = viewModel.profile {
                    UserSection(title: "个人信息", content: "\(apiProfile.name.isEmpty ? "未设置姓名" : apiProfile.name) · \(apiProfile.age == 0 ? "-" : "\(apiProfile.age)") 岁")
                    UserSection(title: "健康目标", content: goalLabel(apiProfile.primaryGoal))
                    UserSection(title: "身体指标", content: "\(apiProfile.heightCm == 0 ? "-" : "\(apiProfile.heightCm)") cm / \(apiProfile.weightKg == 0 ? "-" : "\(apiProfile.weightKg)") kg")
                } else {
                    UserSection(title: "个人信息", content: "\(profile.name.isEmpty ? "未设置姓名" : profile.name) · \(profile.age.isEmpty ? "-" : profile.age) 岁")
                    UserSection(title: "健康目标", content: goalLabel(profile.primaryGoal))
                    UserSection(title: "身体指标", content: "\(profile.heightCM.isEmpty ? "-" : profile.heightCM) cm / \(profile.weightKG.isEmpty ? "-" : profile.weightKG) kg")
                }

                if let stats = viewModel.stats {
                    UserSection(title: "统计数据", content: "连续记录 \(stats.currentStreak) 天，累计 \(stats.daysTracked) 天")
                } else {
                    UserSection(title: "统计数据", content: "加载中...")
                }

                Divider()
                UserSection(title: "设置", content: "通知、权限、设备管理结构已预留")

                if let error = viewModel.errorMessage {
                    UserSection(title: "错误", content: error)
                }
            }
            .padding(16)
        }
        .background(PulseTheme.background)
        .task {
            await viewModel.loadProfile()
        }
    }

    private func goalLabel(_ goal: String) -> String {
        switch goal {
        case "fat_loss":
            return "减脂"
        case "muscle_gain":
            return "增肌"
        case "sleep_recovery":
            return "作息恢复"
        case "stress_relief":
            return "减压放松"
        default:
            return "维持健康"
        }
    }
}

private struct UserSection: View {
    let title: String
    let content: String

    var body: some View {
        VStack(alignment: .leading, spacing: 6) {
            Text(title)
                .font(.headline)
            Text(content)
                .foregroundStyle(.secondary)
        }
    }
}
