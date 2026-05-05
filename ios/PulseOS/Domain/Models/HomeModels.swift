import Foundation

struct HomeDashboard: Decodable {
    let totalScore: Int
    let actionItem: String
    let dietSummary: String
    let activitySummary: String
    let sleepSummary: String
    let meditationNote: String
    let trends: [Int]

    enum CodingKeys: String, CodingKey {
        case totalScore
        case actionItem = "action_item"
        case dietSummary = "diet_summary"
        case activitySummary = "activity_summary"
        case sleepSummary = "sleep_summary"
        case meditationNote = "meditation_note"
        case trends
    }

    init(totalScore: Int, actionItem: String, dietSummary: String, activitySummary: String, sleepSummary: String, meditationNote: String, trends: [Int]) {
        self.totalScore = totalScore
        self.actionItem = actionItem
        self.dietSummary = dietSummary
        self.activitySummary = activitySummary
        self.sleepSummary = sleepSummary
        self.meditationNote = meditationNote
        self.trends = trends
    }
}
