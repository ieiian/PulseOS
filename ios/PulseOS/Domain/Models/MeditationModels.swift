import Foundation

struct BreathMode: Identifiable {
    let id = UUID()
    let key: String
    let title: String
    let inhaleSec: Int
    let holdSec: Int
    let exhaleSec: Int
    let description: String
}

struct MeditationSummary {
    let totalMinutes: Int
    let sessionCount: Int
}

