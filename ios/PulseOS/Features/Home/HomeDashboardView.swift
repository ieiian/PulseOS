import SwiftUI

struct HomeDashboardView: View {
    let profile: UserProfile
    private let dashboard = HomeDashboard(
        totalScore: 67,
        actionItem: "优先完成今天的运动积分和一次饮食记录。",
        dietSummary: "目标热量 1900 千卡，已配置轻断食窗口。",
        activitySummary: "今日 5620 步，心肺强化 75 分。",
        sleepSummary: "昨夜评分 55，时长 430 分钟。",
        meditationNote: "今日冥想 10 分钟。",
        trends: [58, 63, 61, 70, 68, 74, 67]
    )

    var body: some View {
        ScrollView {
            VStack(spacing: 12) {
                HomeCard(title: "今日概览", body: "\(profile.name.isEmpty ? "你" : profile.name)，\(dashboard.actionItem)")
                HomeCard(title: "今日总分", body: "\(dashboard.totalScore) 分")
                HomeCard(title: "饮食", body: dashboard.dietSummary)
                HomeCard(title: "运动", body: dashboard.activitySummary)
                HomeCard(title: "睡眠", body: dashboard.sleepSummary)
                HomeCard(title: "冥想", body: dashboard.meditationNote)
                HomeCard(title: "趋势", body: dashboard.trends.map(String.init).joined(separator: " / "))
            }
            .padding(16)
        }
        .background(PulseTheme.background)
    }
}


private struct HomeCard: View {
    let title: String
    let body: String

    var body: some View {
        VStack(alignment: .leading, spacing: 8) {
            Text(title)
                .font(.headline)
            Text(body)
                .foregroundStyle(.secondary)
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .padding(20)
        .background(PulseTheme.card)
        .clipShape(RoundedRectangle(cornerRadius: 20, style: .continuous))
    }
}
