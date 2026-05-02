import SwiftUI

struct MeditationView: View {
    @StateObject private var controller = MeditationAudioController()
    @State private var selectedMode = defaultModes().first!

    var body: some View {
        ScrollView {
            VStack(spacing: 12) {
                MeditationCard(title: "今日冥想", body: "已完成 10 分钟，1 次训练")
                MeditationCard(
                    title: selectedMode.title,
                    body: "吸气 \(selectedMode.inhaleSec)s · 停顿 \(selectedMode.holdSec)s · 呼气 \(selectedMode.exhaleSec)s\n\(selectedMode.description)"
                )
                HStack(spacing: 8) {
                    Button(controller.isPlaying ? "暂停音频" : "播放音频") {
                        controller.toggle()
                    }
                    .buttonStyle(.borderedProminent)

                    Button("开始呼吸") {}
                        .buttonStyle(.bordered)
                }
                .frame(maxWidth: .infinity, alignment: .leading)

                ForEach(defaultModes()) { mode in
                    Button(mode.title) {
                        selectedMode = mode
                    }
                    .buttonStyle(.bordered)
                    .frame(maxWidth: .infinity, alignment: .leading)
                }
            }
            .padding(16)
        }
        .background(PulseTheme.background)
    }
}

private func defaultModes() -> [BreathMode] {
    [
        BreathMode(key: "calm", title: "平静呼吸", inhaleSec: 4, holdSec: 2, exhaleSec: 6, description: "适合放慢节奏，降低紧绷感。"),
        BreathMode(key: "focus", title: "专注呼吸", inhaleSec: 4, holdSec: 4, exhaleSec: 4, description: "适合进入工作或学习状态前。"),
        BreathMode(key: "sleep", title: "睡前呼吸", inhaleSec: 4, holdSec: 7, exhaleSec: 8, description: "适合睡前稳定呼吸节律。"),
    ]
}

private struct MeditationCard: View {
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

