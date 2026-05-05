import Foundation

// MARK: - Dashboard (matches backend scoring.Dashboard JSON)

struct DashboardDTO: Codable {
    let today: DailyScoreDTO
    let actionItem: String
    let dietSummary: String
    let activitySummary: String
    let sleepSummary: String
    let meditationNote: String
    let trends: [Int]

    enum CodingKeys: String, CodingKey {
        case today
        case actionItem = "action_item"
        case dietSummary = "diet_summary"
        case activitySummary = "activity_summary"
        case sleepSummary = "sleep_summary"
        case meditationNote = "meditation_note"
        case trends
    }
    
    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        today = try container.decode(DailyScoreDTO.self, forKey: .today)
        actionItem = try container.decodeIfPresent(String.self, forKey: .actionItem) ?? ""
        dietSummary = try container.decodeIfPresent(String.self, forKey: .dietSummary) ?? ""
        activitySummary = try container.decodeIfPresent(String.self, forKey: .activitySummary) ?? ""
        sleepSummary = try container.decodeIfPresent(String.self, forKey: .sleepSummary) ?? ""
        meditationNote = try container.decodeIfPresent(String.self, forKey: .meditationNote) ?? ""
        trends = try container.decodeIfPresent([Int].self, forKey: .trends) ?? []
    }
}

struct DailyScoreDTO: Codable {
    let date: String
    let dietScore: Int
    let activityScore: Int
    let sleepScore: Int
    let totalScore: Int

    enum CodingKeys: String, CodingKey {
        case date
        case dietScore = "diet_score"
        case activityScore = "activity_score"
        case sleepScore = "sleep_score"
        case totalScore = "total_score"
    }
    
    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        date = try container.decodeIfPresent(String.self, forKey: .date) ?? ""
        dietScore = try container.decodeIfPresent(Int.self, forKey: .dietScore) ?? 0
        activityScore = try container.decodeIfPresent(Int.self, forKey: .activityScore) ?? 0
        sleepScore = try container.decodeIfPresent(Int.self, forKey: .sleepScore) ?? 0
        totalScore = try container.decodeIfPresent(Int.self, forKey: .totalScore) ?? 0
    }
}
