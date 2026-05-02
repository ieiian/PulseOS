import SwiftUI

struct ActivityView: View {
    private let summary: ActivitySummary
    private let trends: [ActivityTrend]

    init(stepPreview: Int) {
        self.summary = ActivitySummary(
            steps: stepPreview,
            cardioPoints: 75,
            weeklyGoal: 150,
            reminder: "本周还差 75 分，优先安排 30 分钟快走或 20 分钟骑行。"
        )
        self.trends = [
            ActivityTrend(label: "Mon", points: 12),
            ActivityTrend(label: "Tue", points: 18),
            ActivityTrend(label: "Wed", points: 0),
            ActivityTrend(label: "Thu", points: 35),
            ActivityTrend(label: "Fri", points: 10),
            ActivityTrend(label: "Sat", points: 0),
            ActivityTrend(label: "Sun", points: 75),
        ]
    }

    var body: some View {
        ScrollView {
            VStack(spacing: 12) {
                ActivityCard(title: "今日步数与积分", body: "\(summary.steps) 步 · \(summary.cardioPoints) 心肺强化分")
                ActivityCard(title: "周目标进度", body: "目标 \(summary.weeklyGoal) 分，当前完成 \(summary.cardioPoints) 分")
                ActivityCard(title: "今日提醒", body: summary.reminder)
                HStack(spacing: 8) {
                    Button("手动补录") {}
                        .buttonStyle(.borderedProminent)
                    Button("查看周趋势") {}
                        .buttonStyle(.bordered)
                }
                .frame(maxWidth: .infinity, alignment: .leading)

                ForEach(trends) { trend in
                    ActivityCard(title: "\(trend.label) 趋势", body: "\(trend.points) 分")
                }
            }
            .padding(16)
        }
        .background(PulseTheme.background)
    }
}

private struct ActivityCard: View {
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

