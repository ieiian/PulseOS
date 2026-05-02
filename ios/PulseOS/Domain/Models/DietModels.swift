import Foundation

struct DietPlanOption: Identifiable {
    let id = UUID()
    let title: String
    let description: String
    let items: [String]
}

struct DietStatus {
    let recommendation: String
    let summary: String
    let explanation: String
}

