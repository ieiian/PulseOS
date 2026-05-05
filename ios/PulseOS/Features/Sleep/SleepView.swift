import SwiftUI

@MainActor
final class SleepViewModel: ObservableObject {
    @Published var summary: SleepTodaySummaryDTO?
    @Published var isLoading = false
    @Published var errorMessage: String?

    private let service: SleepService

    init(service: SleepService) {
        self.service = service
    }

    func loadToday() async {
        isLoading = true
        defer { isLoading = false }
        do {
            summary = try await service.getToday()
        } catch {
            errorMessage = error.localizedDescription
        }
    }

    func startSession(audioURL: String = "") async {
        isLoading = true
        defer { isLoading = false }
        do {
            _ = try await service.startSession(audioURL: audioURL)
            await loadToday()
        } catch {
            errorMessage = error.localizedDescription
        }
    }

    func endSession() async {
        guard let id = summary?.session.id, !id.isEmpty else {
            errorMessage = "没有进行中的睡眠会话"
            return
        }
        isLoading = true
        defer { isLoading = false }
        do {
            _ = try await service.endSession(sessionID: id)
            await loadToday()
        } catch {
            errorMessage = error.localizedDescription
        }
    }
}

struct SleepView: View {
    @StateObject private var monitor = SleepMonitorController()
    @StateObject private var viewModel: SleepViewModel

    init(service: SleepService = APIContainer.shared.sleep) {
        _viewModel = StateObject(wrappedValue: SleepViewModel(service: service))
    }

    var body: some View {
        ScrollView {
            VStack(spacing: 12) {
                if let summary = viewModel.summary {
                    let session = summary.session
                    SleepCard(title: "昨夜概览", content: "\(session.durationM) 分钟 · 评分 \(session.score)")
                    if !session.advice.isEmpty {
                        SleepCard(title: "睡前建议", content: session.advice)
                    }

                    ForEach(summary.events, id: \.timestamp) { event in
                        SleepCard(title: "事件时间轴", content: "\(event.timestamp) · \(event.type) · \(event.level)")
                    }
                } else if viewModel.isLoading {
                    ProgressView("加载中...")
                }

                Button(monitor.isRecording ? "结束监测" : "开始睡眠") {
                    Task {
                        if monitor.isRecording {
                            monitor.stop()
                            await viewModel.endSession()
                        } else {
                            monitor.start()
                            await viewModel.startSession()
                        }
                    }
                }
                .buttonStyle(.borderedProminent)
                .frame(maxWidth: .infinity, alignment: .leading)

                if let error = viewModel.errorMessage {
                    SleepCard(title: "错误", content: error)
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

private struct SleepCard: View {
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
