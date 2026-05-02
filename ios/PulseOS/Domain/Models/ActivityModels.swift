import Foundation

struct ActivitySummary {
    let steps: Int
    let cardioPoints: Int
    let weeklyGoal: Int
    let reminder: String
}

struct ActivityTrend: Identifiable {
    let id = UUID()
    let label: String
    let points: Int
}

