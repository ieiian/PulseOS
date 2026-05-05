import SwiftUI

@MainActor
final class MeditationViewModel: ObservableObject {
    @Published var summary: MeditationTodaySummaryDTO?
    @Published var isLoading = false
    @Published var errorMessage: String?

    private let service: MeditationService

    init(service: MeditationService) {
        self.service = service
    }

    func loadToday() async {
        isLoading = true
        defer { isLoading = false }
        do {
            summary = try await service.getTodaySummary()
        } catch {
            errorMessage = error.localizedDescription
        }
    }

    func recordSession(modeKey: String, durationS: Int, audioKey: String) async {
        isLoading = true
        defer { isLoading = false }
        do {
            let request = MeditationSessionRequestDTO(modeKey: modeKey, durationS: durationS, audioKey: audioKey)
            _ = try await service.recordSession(request)
            await loadToday()
        } catch {
            errorMessage = error.localizedDescription
        }
    }
}

struct MeditationView: View {
    @StateObject private var controller = MeditationAudioController()
    @StateObject private var viewModel: MeditationViewModel

    init(service: MeditationService = APIContainer.shared.meditation) {
        _viewModel = StateObject(wrappedValue: MeditationViewModel(service: service))
    }

    @State private var selectedModeKey: String?

    private var modes: [BreathModeDTO] {
        viewModel.summary?.modes ?? []
    }

    private var selectedMode: BreathModeDTO? {
        if let key = selectedModeKey {
            return modes.first { $0.key == key }
        }
        return modes.first
    }

    var body: some View {
        ScrollView {
            VStack(spacing: 12) {
                if let summary = viewModel.summary {
                    MeditationCard(title: "今日冥想", content: "已完成 \(summary.totalDurationS / 60) 分钟，\(summary.completedCount) 次训练")
                } else if viewModel.isLoading {
                    ProgressView("加载中...")
                }

                if let mode = selectedMode ?? modes.first {
                    MeditationCard(
                        title: mode.title,
                        content: "吸气 \(mode.inhaleSec)s · 停顿 \(mode.holdSec)s · 呼气 \(mode.exhaleSec)s\n\(mode.description)"
                    )
                }

                HStack(spacing: 8) {
                    Button(controller.isPlaying ? "暂停音频" : "播放音频") {
                        controller.toggle()
                    }
                    .buttonStyle(.borderedProminent)

                    if let mode = selectedMode ?? modes.first {
                        Button("开始呼吸") {
                            Task {
                                await viewModel.recordSession(modeKey: mode.key, durationS: mode.inhaleSec + mode.holdSec + mode.exhaleSec, audioKey: "")
                            }
                        }
                        .buttonStyle(.bordered)
                    }
                }
                .frame(maxWidth: .infinity, alignment: .leading)

                ForEach(modes, id: \.key) { mode in
                    Button(mode.title) {
                        selectedModeKey = mode.key
                    }
                    .buttonStyle(.bordered)
                    .tint(selectedModeKey == mode.key ? PulseTheme.brand : .gray.opacity(0.35))
                    .frame(maxWidth: .infinity, alignment: .leading)
                }

                if let error = viewModel.errorMessage {
                    MeditationCard(title: "错误", content: error)
                }
            }
            .padding(16)
        }
        .background(PulseTheme.background)
        .task {
            await viewModel.loadToday()
        }
    }
}

private struct MeditationCard: View {
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
