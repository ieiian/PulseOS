import SwiftUI

struct SleepView: View {
    @StateObject private var controller = SleepMonitorController()

    var body: some View {
        let summary = SleepSummary(
            durationMinutes: 430,
            score: 55,
            advice: "昨夜呼噜/梦话事件偏多，建议连续观察几晚并避免睡前饮酒。",
            isRecording: controller.isRecording
        )
        let events = [
            SleepEvent(type: "snore", timestamp: "01:12", level: "medium"),
            SleepEvent(type: "talk", timestamp: "03:46", level: "low"),
            SleepEvent(type: "snore", timestamp: "05:21", level: "high"),
        ]

        return ScrollView {
            VStack(spacing: 12) {
                SleepCard(title: "昨夜概览", body: "\(summary.durationMinutes) 分钟 · 评分 \(summary.score)")
                SleepCard(title: "睡前建议", body: summary.advice)
                Button(controller.isRecording ? "结束监测" : "开始睡眠") {
                    if controller.isRecording {
                        controller.stop()
                    } else {
                        controller.start()
                    }
                }
                .buttonStyle(.borderedProminent)
                .frame(maxWidth: .infinity, alignment: .leading)

                ForEach(events) { event in
                    SleepCard(title: "事件时间轴", body: "\(event.timestamp) · \(event.type) · \(event.level)")
                }
            }
            .padding(16)
        }
        .background(PulseTheme.background)
    }
}

private struct SleepCard: View {
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

