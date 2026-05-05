import Foundation

// MARK: - Today Summary (matches backend activity.TodaySummary JSON)

struct ActivityTodaySummaryDTO: Codable {
    let steps: Int
    let cardioPoints: Int
    let stepGoal: Int
    let weeklyGoal: Int
    let remainingPoints: Int
    let reminder: String
    let recentActivities: [ActivityRecordDTO]

    enum CodingKeys: String, CodingKey {
        case steps
        case cardioPoints = "cardio_points"
        case stepGoal = "step_goal"
        case weeklyGoal = "weekly_goal"
        case remainingPoints = "remaining_points"
        case reminder
        case recentActivities = "recent_activities"
    }
    
    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        steps = try container.decodeIfPresent(Int.self, forKey: .steps) ?? 0
        cardioPoints = try container.decodeIfPresent(Int.self, forKey: .cardioPoints) ?? 0
        stepGoal = try container.decodeIfPresent(Int.self, forKey: .stepGoal) ?? 0
        weeklyGoal = try container.decodeIfPresent(Int.self, forKey: .weeklyGoal) ?? 0
        remainingPoints = try container.decodeIfPresent(Int.self, forKey: .remainingPoints) ?? 0
        reminder = try container.decodeIfPresent(String.self, forKey: .reminder) ?? ""
        recentActivities = try container.decodeIfPresent([ActivityRecordDTO].self, forKey: .recentActivities) ?? []
    }
}

// MARK: - Week Summary (matches backend activity.WeekSummary JSON)

struct ActivityWeekSummaryDTO: Codable {
    let dailyPoints: [Int]
    let totalPoints: Int
    let weeklyGoal: Int
    let status: String
    let tips: [String]

    enum CodingKeys: String, CodingKey {
        case dailyPoints = "daily_points"
        case totalPoints = "total_points"
        case weeklyGoal = "weekly_goal"
        case status, tips
    }
    
    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        dailyPoints = try container.decodeIfPresent([Int].self, forKey: .dailyPoints) ?? []
        totalPoints = try container.decodeIfPresent(Int.self, forKey: .totalPoints) ?? 0
        weeklyGoal = try container.decodeIfPresent(Int.self, forKey: .weeklyGoal) ?? 0
        status = try container.decodeIfPresent(String.self, forKey: .status) ?? ""
        tips = try container.decodeIfPresent([String].self, forKey: .tips) ?? []
    }
}

// MARK: - Record (matches backend activity.Record JSON)

struct ActivityRecordDTO: Codable {
    let id: String
    let source: String
    let activityType: String
    let steps: Int
    let minutes: Int
    let intensity: String
    let cardioPoints: Int

    enum CodingKeys: String, CodingKey {
        case id, source
        case activityType = "activity_type"
        case steps, minutes, intensity
        case cardioPoints = "cardio_points"
    }
    
    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        id = try container.decodeIfPresent(String.self, forKey: .id) ?? ""
        source = try container.decodeIfPresent(String.self, forKey: .source) ?? ""
        activityType = try container.decodeIfPresent(String.self, forKey: .activityType) ?? ""
        steps = try container.decodeIfPresent(Int.self, forKey: .steps) ?? 0
        minutes = try container.decodeIfPresent(Int.self, forKey: .minutes) ?? 0
        intensity = try container.decodeIfPresent(String.self, forKey: .intensity) ?? ""
        cardioPoints = try container.decodeIfPresent(Int.self, forKey: .cardioPoints) ?? 0
    }
}

// MARK: - Manual Record Request

struct ManualActivityRequestDTO: Codable {
    let activityType: String
    let minutes: Int
    let intensity: String
    let steps: Int

    enum CodingKeys: String, CodingKey {
        case activityType = "activity_type"
        case minutes, intensity, steps
    }
}
