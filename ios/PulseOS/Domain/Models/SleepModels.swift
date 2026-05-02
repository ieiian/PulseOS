import Foundation

struct SleepEvent: Identifiable {
    let id = UUID()
    let type: String
    let timestamp: String
    let level: String
}

struct SleepSummary {
    let durationMinutes: Int
    let score: Int
    let advice: String
    let isRecording: Bool
}

