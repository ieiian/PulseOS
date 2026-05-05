import Foundation

struct ActivitySummary: Decodable {
    let steps: Int
    let cardioPoints: Int
    let weeklyGoal: Int
    let reminder: String

    enum CodingKeys: String, CodingKey {
        case steps
        case cardioPoints = "cardio_points"
        case weeklyGoal = "weekly_goal"
        case reminder
    }

    init(steps: Int, cardioPoints: Int, weeklyGoal: Int, reminder: String) {
        self.steps = steps
        self.cardioPoints = cardioPoints
        self.weeklyGoal = weeklyGoal
        self.reminder = reminder
    }
}

struct ActivityTrend: Identifiable, Decodable {
    let id: UUID
    let label: String
    let points: Int

    enum CodingKeys: String, CodingKey {
        case label, points
    }

    init(id: UUID = UUID(), label: String, points: Int) {
        self.id = id
        self.label = label
        self.points = points
    }

    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        id = UUID()
        label = try container.decode(String.self, forKey: .label)
        points = try container.decode(Int.self, forKey: .points)
    }
}
