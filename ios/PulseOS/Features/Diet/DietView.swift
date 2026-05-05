import SwiftUI

@MainActor
final class DietViewModel: ObservableObject {
    @Published var plan: TodayPlanDTO?
    @Published var records: [DietRecordDTO] = []
    @Published var analyzeResult: AnalyzeResultDTO?
    @Published var isLoading = false
    @Published var errorMessage: String?

    private let service: DietService

    init(service: DietService) {
        self.service = service
    }

    func loadPlan() async {
        isLoading = true
        defer { isLoading = false }
        do {
            plan = try await service.getTodayPlan()
        } catch {
            errorMessage = "加载饮食计划失败: \(error.localizedDescription)"
        }
    }

    func loadRecords() async {
        do {
            records = try await service.listRecords()
        } catch {
            errorMessage = "加载饮食记录失败: \(error.localizedDescription)"
        }
    }

    func quickRecord(items: [String], mealType: String) async {
        isLoading = true
        defer { isLoading = false }
        do {
            let request = AnalyzeRequestDTO(imageURL: nil, mealType: mealType, manualItems: items)
            let record = try await service.quickRecord(request)
            records.insert(record, at: 0)
        } catch {
            errorMessage = error.localizedDescription
        }
    }
}

struct DietView: View {
    @StateObject private var viewModel: DietViewModel

    init(service: DietService = APIContainer.shared.diet) {
        _viewModel = StateObject(wrappedValue: DietViewModel(service: service))
    }

    var body: some View {
        ScrollView {
            VStack(spacing: 12) {
                if let plan = viewModel.plan {
                    DietCard(title: "今日饮食目标", content: "目标热量 \(plan.targetCalories) 千卡，\(plan.fastingPlan.name) \(plan.fastingPlan.window)")

                    ForEach(plan.options, id: \.title) { option in
                        DietCard(
                            title: option.title,
                            content: "\(option.description)\n\(option.items.joined(separator: " / "))"
                        )
                    }
                } else if viewModel.isLoading {
                    ProgressView("加载中...")
                }

                ForEach(viewModel.records, id: \.id) { record in
                    let recColor: Color = {
                        switch record.recommendation {
                        case "recommended": return Color(red: 0.87, green: 0.95, blue: 0.89)
                        case "not_recommended": return Color(red: 0.98, green: 0.89, blue: 0.84)
                        case "forbidden": return Color(red: 0.96, green: 0.83, blue: 0.85)
                        default: return Color(red: 0.98, green: 0.94, blue: 0.78)
                        }
                    }()
                    DietCard(
                        title: "\(record.mealType ?? "记录") · \(record.totalCalories) 千卡",
                        content: record.recommendation
                    )
                    .background(recColor)
                }

                if let error = viewModel.errorMessage {
                    DietCard(title: "错误", content: error)
                }
            }
            .padding(16)
        }
        .background(PulseTheme.background)
        .task {
            await viewModel.loadPlan()
            await viewModel.loadRecords()
        }
    }
}

private struct DietCard: View {
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
