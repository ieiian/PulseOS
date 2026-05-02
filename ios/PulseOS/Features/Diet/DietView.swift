import SwiftUI

struct DietView: View {
    private let status = DietStatus(
        recommendation: "caution",
        summary: "识别到鸡胸肉沙拉和糙米饭，总热量约 480 千卡。",
        explanation: "适合正餐，但要注意当天总热量和晚间加餐。"
    )

    private let options = [
        DietPlanOption(title: "高蛋白午餐", description: "控制主食份量，优先蛋白质和蔬菜。", items: ["鸡胸肉", "西兰花", "糙米饭"]),
        DietPlanOption(title: "常用早餐", description: "快速记录模板", items: ["无糖酸奶", "鸡蛋", "香蕉"]),
    ]

    var body: some View {
        ScrollView {
            VStack(spacing: 12) {
                DietCard(title: "今日饮食目标", body: "目标热量 1900 千卡，轻断食窗口 12:00 - 20:00")
                RecommendationCard(status: status)
                HStack(spacing: 8) {
                    Button("拍照分析") {}
                        .buttonStyle(.borderedProminent)
                    Button("快速记录") {}
                        .buttonStyle(.bordered)
                }
                .frame(maxWidth: .infinity, alignment: .leading)

                ForEach(options) { option in
                    DietCard(
                        title: option.title,
                        body: "\(option.description)\n\(option.items.joined(separator: " / "))"
                    )
                }
            }
            .padding(16)
        }
        .background(PulseTheme.background)
    }
}

private struct RecommendationCard: View {
    let status: DietStatus

    var body: some View {
        let pair = labelAndColor()

        VStack(alignment: .leading, spacing: 8) {
            Text(pair.0)
                .font(.caption.weight(.semibold))
                .padding(.horizontal, 10)
                .padding(.vertical, 6)
                .background(.white.opacity(0.8))
                .clipShape(Capsule())
            Text(status.summary)
                .font(.headline)
            Text(status.explanation)
                .foregroundStyle(.secondary)
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .padding(20)
        .background(pair.1)
        .clipShape(RoundedRectangle(cornerRadius: 20, style: .continuous))
    }

    private func labelAndColor() -> (String, Color) {
        switch status.recommendation {
        case "recommended":
            return ("建议食用", Color(red: 0.87, green: 0.95, blue: 0.89))
        case "not_recommended":
            return ("不建议食用", Color(red: 0.98, green: 0.89, blue: 0.84))
        case "forbidden":
            return ("禁止食用", Color(red: 0.96, green: 0.83, blue: 0.85))
        default:
            return ("注意食用", Color(red: 0.98, green: 0.94, blue: 0.78))
        }
    }
}

private struct DietCard: View {
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

