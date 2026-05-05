import SwiftUI

@MainActor
final class ActivityViewModel: ObservableObject {
    @Published var todaySummary: ActivityTodaySummaryDTO?
    @Published var weekSummary: ActivityWeekSummaryDTO?
    @Published var isLoading = false
    @Published var errorMessage: String?

    private let service: ActivityService

    init(service: ActivityService) {
        self.service = service
    }

    func loadToday() async {
        isLoading = true
        defer { isLoading = false }
        do {
            todaySummary = try await service.getTodaySummary()
        } catch {
            errorMessage = "加载今日数据失败: \(error.localizedDescription)"
        }
    }

    func loadWeek() async {
        do {
            weekSummary = try await service.getWeekSummary()
        } catch {
            errorMessage = "加载周数据失败: \(error.localizedDescription)"
        }
    }

    func recordManual(activityType: String, minutes: Int, intensity: String, steps: Int) async {
        isLoading = true
        defer { isLoading = false }
        do {
            let request = ManualActivityRequestDTO(activityType: activityType, minutes: minutes, intensity: intensity, steps: steps)
            _ = try await service.recordManual(request)
            await loadToday()
        } catch {
            errorMessage = error.localizedDescription
        }
    }
}

struct ActivityView: View {
    @StateObject private var viewModel: ActivityViewModel

    init(service: ActivityService = APIContainer.shared.activity) {
        _viewModel = StateObject(wrappedValue: ActivityViewModel(service: service))
    }

    var body: some View {
        ScrollView {
            VStack(spacing: 12) {
                if let summary = viewModel.todaySummary {
                    ActivityCard(title: "今日步数与积分", content: "\(summary.steps) 步 · \(summary.cardioPoints) 心肺强化分")
                    ActivityCard(title: "周目标进度", content: "目标 \(summary.weeklyGoal) 分，当前完成 \(summary.cardioPoints) 分")
                    ActivityCard(title: "今日提醒", content: summary.reminder)
                } else if viewModel.isLoading {
                    ProgressView("加载中...")
                }

                if let week = viewModel.weekSummary {
                    ActivityCard(title: "本周状态", content: "\(week.status) · 总计 \(week.totalPoints) 分")
                    ForEach(Array(zip(week.dailyPoints.indices, week.dailyPoints)), id: \.0) { index, points in
                        let labels = ["周一", "周二", "周三", "周四", "周五", "周六", "周日"]
                        if index < labels.count {
                            ActivityCard(title: "\(labels[index]) 趋势", content: "\(points) 分")
                        }
                    }
                }

                if let error = viewModel.errorMessage {
                    ActivityCard(title: "错误", content: error)
                }
            }
            .padding(16)
        }
        .background(PulseTheme.background)
        .task {
            await viewModel.loadToday()
            await viewModel.loadWeek()
        }
    }
}

private struct ActivityCard: View {
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
